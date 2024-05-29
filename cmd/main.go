package main

import (
	"github.com/MachadoMichael/credentials/infra/database"
	"github.com/MachadoMichael/credentials/pkg/route"
)

func main() {
	database.Init()
	defer database.CloseDb()
	route.Init()
}
