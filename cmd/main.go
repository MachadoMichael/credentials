package main

import (
	"github.com/MachadoMichael/GoAPI/infra/database/repository"
	"github.com/MachadoMichael/GoAPI/pkg/route"
)

func main() {
	repository.InitDB()
	route.Init()
}
