A very simple guestbook that works with a database

<img width="400" height="auto" alt="image" src="https://github.com/user-attachments/assets/3b252ba1-5e6f-4d49-9b79-ce10e7457f84" />

# Commands for creating in database
```sql
CREATE DATABASE Guest_book
USE Guest_book
CREATE TABLE IF NOT EXISTS entries (
name TEXT,
email TEXT,
message TEXT,
timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
)
```
It will be quite easy to copy and paste into any other project.
