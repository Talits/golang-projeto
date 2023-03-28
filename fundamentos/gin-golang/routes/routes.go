package routes

import (
	"fundamentos/gin-golang/controller"

	"github.com/gin-gonic/gin"
)

func HandlerRequest() {

	r := gin.Default()
	r.GET("/alunos", controller.ExibeAlunos)
	r.Run(":8087")

}
