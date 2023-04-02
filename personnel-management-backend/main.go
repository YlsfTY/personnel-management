package main

import (
	"fmt"
	"personnel-management-backend/dao"
	"personnel-management-backend/server"
)

func main() {
	fmt.Println("Hello go")
	dao.DBinit()
	server.SetupServer()
	defer dao.CloseDB()
}
