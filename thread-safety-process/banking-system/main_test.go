package bankingsystem

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
	"testing"
	"time"

	pgxtxpool "github.com/rasatmaja/pgx-txpool"
	"github.com/rasatmaja/phoenix-med/banking-system/model"
	"github.com/rasatmaja/phoenix-med/banking-system/repository"
	"github.com/rasatmaja/phoenix-med/banking-system/service"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type TestSuite struct {
	repo *repository.Repository
	srv  *service.Service
}

func TestMain(t *testing.T) {
	ctx := context.Background()

	suite := &TestSuite{}
	// setup test containers
	assert.NotPanics(t, func() { suite.Setup(ctx) })

	// run tests
	t.Run("TestMigration", suite.Migration)
	t.Run("TestCreateUser", suite.CreateUser)
	t.Run("TestCreateTransaction", suite.CreateTransaction)
}

func (ts *TestSuite) Setup(ctx context.Context) {
	var pgUsername, pgPassword, pgDatabase, pgHost, pgPort string
	pgUsername = "postgres-user"
	pgPassword = "postgres-password"
	pgDatabase = "postgres-db"

	postgres, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "postgres:16-alpine",
			ExposedPorts: []string{"5432/tcp"},
			AutoRemove:   true,
			Env: map[string]string{
				"POSTGRES_USER":     pgUsername,
				"POSTGRES_PASSWORD": pgPassword,
				"POSTGRES_DB":       pgDatabase,
			},
			WaitingFor: wait.ForListeningPort("5432/tcp").WithStartupTimeout(10 * time.Second),
		},
		Started: true,
	})
	if err != nil {
		panic(err)
	}

	// get hostname from generated test container
	pgHost, err = postgres.Host(ctx)
	if err != nil {
		panic(err)
	}

	// get exposed port from generated test container
	exposedPort, err := postgres.MappedPort(ctx, "5432")
	if err != nil {
		panic(err)
	}

	pgPort = exposedPort.Port()

	// setup database
	db := pgxtxpool.New(
		pgxtxpool.SetHost(pgHost, pgPort),
		pgxtxpool.SetCredential(pgUsername, pgPassword),
		pgxtxpool.SetDatabase(pgDatabase),
		pgxtxpool.WithSSLMode("disable"),
		pgxtxpool.WithMaxConns(20),
		pgxtxpool.WithMaxIdleConns("30s"),
		pgxtxpool.WithMaxConnLifetime("5m"),
	)

	if err := db.Ping(ctx); err != nil {
		panic(err)
	}

	ts.repo = repository.NewRepository(db)
	ts.srv = service.NewService(ts.repo)
}

// Migration tests repository Migration method
func (ts *TestSuite) Migration(t *testing.T) {
	ctx := context.Background()
	err := ts.repo.Migration(ctx)
	assert.NoError(t, err, "failed execute migration")

	columnsUsers, err := ts.repo.ShowColomns(ctx, "users")
	assert.NoError(t, err, "failed to get columns users")
	assert.ElementsMatch(t, []string{"id", "name", "balance"}, columnsUsers)

	columnsTransactions, err := ts.repo.ShowColomns(ctx, "transactions")
	assert.NoError(t, err, "failed to get columns transactions")
	assert.ElementsMatch(t, []string{"id", "user_id", "type", "amount"}, columnsTransactions)
}

// CreateUser tests service CreateUser method
func (ts *TestSuite) CreateUser(t *testing.T) {

	ctx := context.Background()
	cases := []struct {
		name string
		user model.User
		trx  []model.Transaction
	}{
		{
			name: "USR001: should success",
			user: model.User{
				ID:      "USR001",
				Name:    "John Doe",
				Balance: 1000,
			},
			trx: []model.Transaction{
				{
					ID:     "TRX001_INIT",
					UserID: "USR001",
					Type:   model.TransactionTypeInitialBalance,
					Amount: 1000,
				},
			},
		},
	}

	wg := new(sync.WaitGroup)

	for _, c := range cases {
		wg.Add(1)

		go func() {
			t.Run(c.name, func(t *testing.T) {
				defer wg.Done()

				err := ts.srv.CreateUser(ctx, c.user, c.trx...)
				assert.NoError(t, err, fmt.Sprintf("failed to create user %s", c.user.ID))
			})
		}()
	}

	wg.Wait()

	usr, err := ts.repo.GetUserByID(ctx, "USR001")
	assert.NoError(t, err, "failed to get user")
	assert.Equal(t, "USR001", usr.ID)
	assert.Equal(t, "John Doe", usr.Name)
	assert.Equal(t, 1000.0, usr.Balance)
}

type users struct {
	model.User
	mu sync.Mutex
}

func (u *users) UpdateBalance(amount float64) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.Balance += amount
}

func (ts *TestSuite) CreateTransaction(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name string
		trx  model.Transaction
	}{
		{
			name: "TRX001: should success",
			trx: model.Transaction{
				ID:     "TRX001",
				UserID: "USR001",
				Type:   model.TransactionTypeDeposit,
				Amount: 1000,
			},
		},
		{
			name: "TRX002: should success",
			trx: model.Transaction{
				ID:     "TRX002",
				UserID: "USR001",
				Type:   model.TransactionTypeWithdraw,
				Amount: 500,
			},
		},
		{
			name: "TRX003: balance not enough",
			trx: model.Transaction{
				ID:     "TRX003",
				UserID: "USR001",
				Type:   model.TransactionTypeWithdraw,
				Amount: 99000,
			},
		},
	}

	wg := new(sync.WaitGroup)

	//collect balance that successfully execute
	//and use it as initial balance for next test case
	InitialUsr, err := ts.repo.GetUserByID(ctx, "USR001")
	assert.NoError(t, err, "failed to get user")

	usr := &users{User: InitialUsr}

	for _, c := range cases {
		wg.Add(1)
		go func() {
			t.Run(c.name, func(t *testing.T) {
				defer wg.Done()

				slog.InfoContext(ctx, fmt.Sprintf("%s - Started", c.trx.ID), "type", c.trx.Type, "amount", c.trx.Amount, "user_id", c.trx.UserID)
				defer slog.InfoContext(ctx, fmt.Sprintf("%s - Done", c.trx.ID), "type", c.trx.Type, "amount", c.trx.Amount, "user_id", c.trx.UserID)

				err := ts.srv.CreateTransaction(ctx, c.trx)
				if err != nil {
					slog.Error("failed to create transaction", "error", err, "id", c.trx.ID)
					return
				}

				amount := c.trx.Amount
				if c.trx.Type == model.TransactionTypeWithdraw {
					amount = c.trx.Amount * -1
				}
				usr.UpdateBalance(amount)
			})
		}()
	}
	wg.Wait()

	updatedUsr, err := ts.repo.GetUserByID(ctx, "USR001")
	assert.NoError(t, err, "failed to get user")
	assert.Equal(t, usr.Balance, updatedUsr.Balance)
	slog.Info("final balance", "balance", usr.Balance)
}
