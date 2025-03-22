## SQL query Optimization
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

## Solution
### Optimized Query
```sql
SELECT 
    customer_id,
    SUM(amount) as total_spent
FROM orders
WHERE order_date >= DATE_TRUNC('month', CURRENT_DATE - INTERVAL '1 month')
    AND order_date < DATE_TRUNC('month', CURRENT_DATE)
GROUP BY customer_id
ORDER BY total_spent DESC
LIMIT 5;
```

### Performance Improvement in a Production Environment
1. **Indexing**: Add indexes on `order_date` and `amount` columns to improve the performance of the query.
2. **Partitioning**: Partition the `orders` table by `order_date` to improve the performance of the query.