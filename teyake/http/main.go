package main

import (
	"html/template"
	"net/http"


	"teyake/entity"

	catRepoImp "teyake/category/repository"
	ansRepoImp "teyake/answer/repository"
	catServiceImp "teyake/category/service"
	quesRepoImp "teyake/question/repository"
	quesServiceImp "teyake/question/service"
	ansServiceImp "teyake/answer/service"
	"teyake/teyake/http/handler"
	userRepoImp "teyake/user/repository"
	userServiceImp "teyake/user/service"

	"teyake/util"
	"teyake/util/token"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func createTables(dbconn *gorm.DB) []error {
	errs := dbconn.CreateTable(&entity.User{}, &entity.Session{}, &entity.Role{}, &entity.Question{}, &entity.Answer{},&entity.Category{}).GetErrors()
	if errs != nil {
		return errs
	}
	return nil
}

func main() {
	dbconn, err := gorm.Open("postgres", util.DBConnectString)
	defer dbconn.Close()
	templ:=template.Must(template.New("main").Funcs(util.AvailableFuncMaps).ParseGlob("../../ui/templates/*"))
	fs := http.FileServer(http.Dir("../../ui/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	if err != nil {
		panic(err)
	}

	//Create a new csrf signing key for forms
	csrfSignKey := []byte(token.GenerateRandomID(32))


	userRepo := userRepoImp.NewUserGormRepo(dbconn)
	userService := userServiceImp.NewUserService(userRepo)

	sessionRepo := userRepoImp.NewSessionGormRepo(dbconn)
	sessionService := userServiceImp.NewSessionService(sessionRepo)

	roleRepo := userRepoImp.NewRoleGormRepo(dbconn)
	roleServ := userServiceImp.NewRoleService(roleRepo)

	questionRepo := quesRepoImp.NewQuestionGormRepo(dbconn)
	questionService := quesServiceImp.NewQuestionService(questionRepo)

	categoryRepo:=catRepoImp.NewCategoryGormRepo(dbconn)
	categoryService:=catServiceImp.NewCategoryService(categoryRepo)

	answerRepo:=ansRepoImp.NewAnswerGormRepo(dbconn)
	answerService:=ansServiceImp.NewAnswerService(answerRepo)

	//Uncomment the following lines after you created a fresh teyake db
	//createTables(dbconn)
	//roleServ.StoreRole(&entity.UserRoleMock)
	//roleServ.StoreRole(&entity.AdminRoleMock)
	//questionService.StoreQuestion(&entity.QuestionMock)
	//categoryService.StoreCategory(&entity.CategoryMock1)
	//categoryService.StoreCategory(&entity.CategoryMock2)
	//categoryService.StoreCategory(&entity.CategoryMock3)


	userHandler := handler.NewUserHandler(templ, userService, sessionService, roleServ, csrfSignKey)
	indexHandler := handler.NewIndexHandler(templ,questionService,categoryService)
	questionHandler :=handler.NewQuestionHandler(templ,questionService,answerService,categoryService,csrfSignKey)
	http.Handle("/",userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(indexHandler.Index))))
	http.Handle("/question",userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(questionHandler.QuestionHandler))))
	http.Handle("/question/new",userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(questionHandler.NewQuestion))))
	http.HandleFunc("/login", userHandler.Login)
	http.HandleFunc("/signup", userHandler.SignUp)
	http.Handle("/logout", userHandler.Authenticated(http.HandlerFunc(userHandler.Logout)))
	http.ListenAndServe(":8181", nil)
}
