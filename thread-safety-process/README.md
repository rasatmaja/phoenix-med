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

> ***For this implementation, we'll use the Non-Blocking approach with  Golang's built-in `sync.Map` as locker.*** 