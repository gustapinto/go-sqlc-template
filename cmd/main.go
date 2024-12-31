package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gustapinto/go-sqlc-template/internal/user"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	userService := user.NewService(db)

	user, err := userService.Create("sample.login", "sample.password", "sample@email")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", user)
}
