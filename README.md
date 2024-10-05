# Concurrent Counter

## Problem Statement

The goal is to increment a numerical value stored in a database concurrently from multiple instances of an application. This must be done in a way that prevents conflicts and ensures accurate updates.

### Approach

To address this, we'll explore different solutions and organize them into separate subdirectories.

### Application Details

- Language: Golang
- Database: MySQL

Database Table: Counter (with a single column count_value)
Scenario

- Initial Value: The counter starts at 0.
- Increment Loop: A loop iterates 10,000 times, reading the current value, incrementing it by 1, and updating the database.
- Single Instance: With a single instance, the final value is expected to be 10,000.
- Multiple Instances: With multiple instances, the expected final value is 200,000 (10,000 increments per instance * 20 instances). However, the actual value is often lower due to read-related issues.

### Potential Solutions

- Pessimistic Locking: Use transactions to acquire exclusive locks on the Counter row before reading and updating. This can be inefficient for high concurrency.
- Optimistic Locking: Use versioning to detect conflicts. Each row has a version number. Before updating, the version is checked. If it has changed since the read, the update is retried.
- Distributed Locks: Employ a distributed locking mechanism like Redis or ZooKeeper to coordinate access across multiple instances.
- Database-Specific Features: Explore database-specific features like sequence numbers or atomic operations that can handle concurrent updates.

Note: The choice of solution depends on factors like performance requirements, consistency needs, and the specific characteristics of the database and application.
