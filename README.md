## Technical Test Nexmedis

### [01. API for an e-commerce platform ↗](./ecommers-api-design/README.md)
You are tasked with designing an API for an e-commerce platform. The system must support the following features:
- User registration and authentication
- Viewing and searching products
- Adding items to a shopping cart
- Completing a purchase

Design the RESTful endpoints for the above features. Describe your choice of HTTP methods (GET, POST, PUT, DELETE), URL structure, and the expected response formats. Assume that users need to authenticate before performing certain actions (e.g., adding items to the cart).

### [02. Indexing strategy ↗](./sql-indexing-strategy/README.md)
Consider a database table Users with the following columns:
```
id (Primary Key)
username
email
created_at
```
Your task is to design an indexing strategy to optimize the following queries:
1. Fetch a user by username.
2. Fetch users who signed up after a certain date (created_at > "2023-01-01").
3. Fetch a user by email.

Explain which columns you would index, and whether composite indexes or individual indexes would be appropriate for each query. Discuss trade-offs in terms of read and write performance.

### [03. Thread safety process ↗](./thread-safety-process/README.md)
You need to implement a function that simulates a bank account system. Multiple users can simultaneously access and update their account balance. Your system must ensure that concurrent access does not result in race conditions.
Implement the function that:
- Deposits money into an account.
- Withdraws money from an account (ensuring there’s enough balance).
- Ensures thread-safety while handling concurrent deposits and withdrawals.

### [04. SQL query Optimization ↗](./sql-query-optimization/README.md)
Given a database table orders with the following schema:
```sql
CREATE TABLE orders (
    id INT PRIMARY KEY,
    customer_id INT,
    product_id INT,
    order_date TIMESTAMP,
    amount DECIMAL(10, 2)
);
```
> Assume that customer_id is indexed, but amount and order_date are not indexed.

Write an optimized SQL query to find the top 5 customers who spent the most money in the past month.
How would you improve the performance of this query in a production environment?

### [05. Refactoring Monolithic Service ↗](./refactoring-monolithic-service/README.md)
You are tasked with refactoring a monolithic service that handles multiple responsibilities such as authentication, file uploads, and user data processing. The system has become slow and hard to maintain. How would you approach refactoring the service?
- What steps would you take to decompose the service into smaller, more manageable services?
- How would you ensure that the new system is backward compatible with the old one during the transition?
