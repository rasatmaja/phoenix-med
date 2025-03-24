## Thread safety process
You need to implement a function that simulates a bank account system. Multiple users can simultaneously access and update their account balance. Your system must ensure that concurrent access does not result in race conditions.
Implement the function that:
- Deposits money into an account.
- Withdraws money from an account (ensuring there's enough balance).
- Ensures thread-safety while handling concurrent deposits and withdrawals.

## Solution
To ensure thread safety in banking transactions, we can implement two different approaches for handling concurrent access:

### Approach 1: Non-blocking (Fail Fast)
- During a deposit transaction, the account is temporarily locked
- If another user attempts to withdraw during this lock period, the system immediately returns an error
- Benefits: Low latency, immediate feedback
- Drawbacks: Higher failure rate during high concurrency

### Approach 2: Blocking (Wait and Retry)
- During a deposit transaction, the account is temporarily locked
- If another user attempts to withdraw during this lock period, the request waits in a queue
- Benefits: Higher success rate, better user experience
- Drawbacks: Potential increased latency

### Implementation Options

1. **In-Memory Synchronization**
   - Uses Golang's built-in `sync.Map` 
   - Simple implementation, suitable for single-instance deployments
   - Limited to single server scenarios

2. **Distributed Locking (Redis)**
   - Handles distributed system scenarios
   - Scales well with multiple service instances
   - Ideal for high-load production environments

### Implementation Details
> ***For this implementation, I'll use the Non-Blocking approach with  Golang's built-in `sync.Map` as locker.*** 

#### Architecture
The system is implemented using a clean architecture pattern with the following components:
- **Service Layer**: Handles business logic and transaction management
- **Repository Layer**: Manages data persistence and implements locking mechanism
- **Model Layer**: Defines data structures and error types

#### Testing
The implementation includes integration tests that:
- Verify concurrent transaction handling
- Test multiple scenarios (deposit, withdrawal, insufficient funds)
- Use testcontainers for PostgreSQL database setup