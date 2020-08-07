package main

import (
	"github.com/at8109/go-echo-mysql/db"
	"github.com/at8109/go-echo-mysql/routes"
)

func main() {

	db.Init()

	e := routes.Init()
	e.Logger.Fatal(e.Start(":1234"))
}
