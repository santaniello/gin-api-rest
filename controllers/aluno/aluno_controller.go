package aluno

import (
	"github.com/gin-gonic/gin"
	"github.com/santaniello/gin-api-rest/database"
	"github.com/santaniello/gin-api-rest/models"
	"net/http"
	"strconv"
)

func exibeTodosAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

func saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "E ai " + nome + ", tudo beleza?",
	})
}

func criaNovovAluno(c *gin.Context) {
	var aluno models.Aluno
	err := c.ShouldBindJSON(&aluno)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func exibeAlunoPorId(c *gin.Context) {
	id := c.Params.ByName("id")
	var aluno models.Aluno
	database.DB.Find(&aluno, id)
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado",
		})
		return
	}
	c.JSON(http.StatusOK, aluno)
}

func deleteAlunoPorId(c *gin.Context) {
	id := c.Params.ByName("id")
	var aluno models.Aluno
	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{
		"msg": "Aluno " + id + " removido",
	})
}

func updateAluno(c *gin.Context) {
	id := c.Params.ByName("id")
	var alunoRequest models.Aluno
	var alunoDb models.Aluno
	err := c.ShouldBindJSON(&alunoRequest)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Find(&alunoDb, id)
	idInt, _ := strconv.ParseUint(id, 10, 32)
	alunoRequest.ID = uint(idInt)
	database.DB.Save(&alunoRequest)
	c.JSON(http.StatusOK, alunoRequest)
}

func exibeAlunoPorCpf(c *gin.Context) {
	cpf := c.Params.ByName("cpf")
	var aluno models.Aluno
	database.DB.Where(models.Aluno{CPF: cpf}).First(&aluno)
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado",
		})
		return
	}
	c.JSON(http.StatusOK, aluno)
}

func RegisterRoutes(r *gin.Engine) {
	r.GET("/alunos", exibeTodosAlunos)
	r.GET("/:nome", saudacao)
	r.GET("/alunos/:id", exibeAlunoPorId)
	r.GET("/alunos/cpf/:cpf", exibeAlunoPorCpf)
	r.PATCH("/alunos/:id", updateAluno)
	r.DELETE("/alunos/:id", deleteAlunoPorId)
	r.POST("/alunos", criaNovovAluno)
}
