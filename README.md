# Go CRUD
Go CRUD App for practice

## DB Setup

### Migrations

#### Create migrations
Use: https://github.com/golang-migrate/migrate

`go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`

**Generate migration**

```sh
migrate create -ext sql -dir db/migrations -seq create_users_table
```

**Run migrations**

```sh
export POSTGRESQL_URL='postgres://postgres:password@localhost:5432/example?sslmode=disable'

migrate -database ${POSTGRESQL_URL} -path db/migrations up
```
