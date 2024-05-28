package main

import (
	"github.com/MachadoMichael/GoAPI/infra/database"
	"github.com/MachadoMichael/GoAPI/pkg/route"
)

func main() {
	database.Init()
	defer database.CloseDb()
	route.Init()
}
