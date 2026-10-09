
# Commands for creating in database
```sql
CREATE DATABASE Guest_book
USE Guest_book
CREATE TABLE IF NOT EXISTS entries (
name TEXT,
message TEXT,
timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
)
```
