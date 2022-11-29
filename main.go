package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/Riku0617/myapi/api"
	"github.com/Riku0617/myapi/controllers"
	"github.com/Riku0617/myapi/services"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbUser     = "docker"
	dbPassword = "docker"
	dbDatabase = "sampledb"
	dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
)

func main() {
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("fail to connect db")
		return
	}

	service := services.NewMyAppService(db)
	con := controllers.NewArticleController(service)
	con2 := controllers.NewCommentController(service)

	r := api.NewRouter(con, con2)

	log.Println("server start at port 8000")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
