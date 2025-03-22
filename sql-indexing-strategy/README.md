## Indexing strategy
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

## Solution
Bassed on the given scenario, the following indexing strategy can be implemented on all the columns with individual indexes:
1. Index on `username`: This index can be used for fast point on certain queries like for authentication.
2. Index on `email`: This index can be used for fast point on certain queries like for authentication or email existance validation.
3. Index on `created_at`: This index can be used for query 2, as it is a timestamp that can be used to filter the results (but its `OPTIONAL`)

```sql
CREATE INDEX idx_username ON users (username);
CREATE INDEX idx_email ON users (email);
CREATE INDEX idx_created_at ON users (created_at);
```
***why not composite indexes?***
Bassed on the given scenario, composite indexes are not necessary, as the queries target a single column independently (username OR email OR created_at).

***what the trade-offs?***
The trade-offs using individual indexes are:
- Indexes need more storage space, because each indexs requires space to store copies of the indexed columns, pointers to the original rows, and additional metadata.
- Write performance: Each index requires additional processing during INSERT, UPDATE, and DELETE operations as the database needs to maintain the index structure.

***why using index on `created_at` optional?***
Using an index on `created_at` depends on the use case and query frequency. If the query is rarely used, then an index on `created_at` is not recommended. This scenario is typically used for internal tooling, marketing, or admin panels where slower query performance is tolerable.