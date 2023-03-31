package main

import (
	"fmt"
	"personnel-management-backend/dao"
	"personnel-management-backend/router"
)

func main() {
	fmt.Println("Hello go")
	dao.DBinit()
	defer dao.CloseDB()
	r := router.SetupRouter()
	r.Run(":8080")

}
