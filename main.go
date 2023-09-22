package main

import (
	"github.com/gin-gonic/gin"
	"github.com/santaniello/gin-api-rest/controllers/aluno"
	"github.com/santaniello/gin-api-rest/database"
)

func main() {
	database.ConectaComBancoDeDados()
	r := gin.Default()
	aluno.RegisterRoutes(r)
	r.Run(":5000")
}
