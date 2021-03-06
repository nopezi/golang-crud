package main

import (
	"clean-gin-template/bootstrap"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
	_ "github.com/go-sql-driver/mysql"

)

func main() {
	godotenv.Load()
	fx.New(bootstrap.Module).Run()
}
