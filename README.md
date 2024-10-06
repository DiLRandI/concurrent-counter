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

## Solutions - Failures

### Solution-01

[link.](https://github.com/DiLRandI/concurrent-counter/blob/main/cmd/solution-01/main.go#L34-L48) To run the code

```bash
make run APP=solution-01
```

This is the core of the problem: it's just independently reading the value, incrementing it, and writing it back. The end result is always incorrect in this case.

### Solution-02

[link.](https://github.com/DiLRandI/concurrent-counter/blob/main/cmd/solution-02/main.go#L34-L48) To run the code

```bash
make run APP=solution-02
```

This is mostly same as [Solution-01](#solution-01). Only different is table now have where statement on primary key. But still the result is wrong.

### Solution-03

[link.](https://github.com/DiLRandI/concurrent-counter/blob/main/cmd/solution-03/main.go#L35-64) To run the code

```bash
make run APP=solution-03
```

This is mostly same as [Solution-01](#solution-01) but using database transaction with default level. But still the result is wrong.

### Solution-04

[link.](https://github.com/DiLRandI/concurrent-counter/blob/main/cmd/solution-04/main.go#L35-64) To run the code

```bash
make run APP=solution-04
```

This is mostly same as [Solution-02](#solution-02). but using database transaction with default level. But still the result is wrong.

### Solution-05

[link.](https://github.com/DiLRandI/concurrent-counter/blob/main/cmd/solution-03/main.go#L35-64) To run the code

```bash
make run APP=solution-05
```

This is mostly same as [Solution-03](#solution-03), this time we use database transaction with isolation level [serializable](https://en.wikipedia.org/wiki/Isolation_(database_systems)#Serializable).
Initially, each application instance can successfully increment the counter value. However, due to a deadlock situation, the application will eventually fail. Ultimately, only one instance will remain operational, continuing to increment the counter value.

### Solution-06

[link.](https://github.com/DiLRandI/concurrent-counter/blob/main/cmd/solution-04/main.go#L35-64) To run the code

```bash
make run APP=solution-06
```

This is mostly same as [Solution-04](#solution-04), this time we use database transaction with isolation level [serializable](https://en.wikipedia.org/wiki/Isolation_(database_systems)#Serializable).
Initially, each application instance can successfully increment the counter value. However, due to a deadlock situation, the application will eventually fail. Ultimately, only one instance will remain operational, continuing to increment the counter value.
