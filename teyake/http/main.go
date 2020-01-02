package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"html/template"
	"net/http"
	"teyake/entity"
	"teyake/teyake/http/handler"
	userRepoImp "teyake/user/repository"
	userServiceImp "teyake/user/service"
)

func createTables(dbconn *gorm.DB)[]error{
	errs:= dbconn.CreateTable(&entity.User{},&entity.Role{},&entity.Session{}).GetErrors()
	if errs !=nil{
		return errs
	}
	return nil
}

func main(){
	dbconn, err:= gorm.Open("postgres","postgres://postgres:toor@localhost/teyake?sslmode=disable")
	defer dbconn.Close()
	templ := template.Must(template.ParseGlob("ui/templates/*"))
	fs := http.FileServer(http.Dir("ui/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	if err !=nil{
		panic(err)
	}


	//createTables(dbconn)
	userRepo := userRepoImp.NewUserGormRepo(dbconn)
	userService := userServiceImp.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(templ,userService)


	http.HandleFunc("/", userHandler.SignUp)
	http.HandleFunc("/login", userHandler.Login)
	http.HandleFunc("/signup", userHandler.SignUp)
	http.ListenAndServe(":8181", nil)
}