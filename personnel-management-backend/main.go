package main

import (
	"fmt"
	// "net/http"
	// "personnel-admin/controller"
	"personnel-admin/router"
	// "personnel-admin/dao"

	// "github.com/gin-gonic/gin"
)



func main() {
	fmt.Println("Hello go")
	// router.Run()
	r := router.SetupRouter()
	r.Run(":8080")

}
