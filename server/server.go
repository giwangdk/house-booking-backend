package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {

	r := gin.Default()
	return r
}

func Init() {
	r := initRouter()
	err := r.Run()
	if err != nil {
		fmt.Println("error while running server", err)
		return
	}

}
