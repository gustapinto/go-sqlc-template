//go:generate go run github.com/sqlc-dev/sqlc/cmd/sqlc@v1.28.0 generate
//go:generate go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.18.2 -database=postgresql://<db-user>:<bd-password>@localhost:5432/<db-name>?sslmode=disable -path=database/migrations up
package main
