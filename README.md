# bstack-go

A new web aapp that will serve as a hub for Counter Strike players to network,
create teams, and recreate players.

This is an application that I've been creating and recreating for a very 
long time. I've been switching the programming language and framework mid 
creation for a long time. I hope that this will be the final implementation.

# Development

## Environment Variables
Fields that need to be populated for the application to run


```bash
export DB_USER=some_username
export DB_PASSWORD=some_password
export DB_HOST_POST=xxx.xxx.xxx.xxx:xxxx
export DB_NAME=name_of_database
```

## Golang Migrate

### Create SQL migration files
```bash
migrate create -ext sql -dir migrations/ -seq create_team_table.sql
```

### Migrate
```bash
migrate -path migrations/ -database="mysql://bstack_user:password@tcp(localhost:3306)/bstack_db" -verbose up
```
