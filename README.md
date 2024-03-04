# server-play

trying things with server

Implement Redis, understand what problem it solves, what are alternatives of it

#### DB

Just use SQLite file. It can be just persisted.

#### Schema//data modeling

This is many-to-many relationship. Each user can have multiple fav beans and each bean can have multiple users that like them. So we need a associative table.

<users>
id: int primary key
name: string
email: string

<coffee>
id: int
name: string
origin: string
profile: string

<join_users_coffee>
user_id: int
coffee_id: int
