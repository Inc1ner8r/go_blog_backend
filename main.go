package main

import (
	"database/sql"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/inciner8r/blog_backend_go/routes"
)

func main() {

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	r := gin.Default()
	r.Use(cors.Default())
	routes.Routes(r)
	r.Run()
}
