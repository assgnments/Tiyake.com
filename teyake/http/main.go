package main

import (
	"html/template"
	"net/http"

	"teyake/entity"

	ansRepoImp "teyake/answer/repository"
	ansServiceImp "teyake/answer/service"
	catRepoImp "teyake/category/repository"
	catServiceImp "teyake/category/service"
	quesRepoImp "teyake/question/repository"
	quesServiceImp "teyake/question/service"
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
	errs := dbconn.CreateTable(&entity.User{}, &entity.Session{}, &entity.Role{}, &entity.Question{}, &entity.Answer{}, &entity.Category{}).GetErrors()
	if errs != nil {
		return errs
	}
	return nil
}

//type justFilesFilesystem struct {
//	fs http.FileSystem
//}
//func (fs justFilesFilesystem) Open(name string)(http.File,error){
//	f,err := fs.fs.Open(name)
//	if err!=nil{
//		return nil,err
//	}
//	return newreaddir{f},nil
//}
func main() {
	dbconn, err := gorm.Open("postgres", util.DBConnectString)
	defer dbconn.Close()
	templ := template.Must(template.New("main").Funcs(util.AvailableFuncMaps).ParseGlob("ui/templates/*"))
	fs := http.FileServer(http.Dir("ui/assets"))
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

	categoryRepo := catRepoImp.NewCategoryGormRepo(dbconn)
	categoryService := catServiceImp.NewCategoryService(categoryRepo)

	answerRepo := ansRepoImp.NewAnswerGormRepo(dbconn)
	answerService := ansServiceImp.NewAnswerService(answerRepo)

	//userService.StoreUser(&entity.UserMock)
	//Uncomment the following lines after you created a fresh teyake db
	//createTables(dbconn)
	//roleServ.StoreRole(&entity.UserRoleMock)
	//roleServ.StoreRole(&entity.AdminRoleMock)
	//questionService.StoreQuestion(&entity.QuestionMock)
	//categoryService.StoreCategory(&entity.CategoryMock1)
	//categoryService.StoreCategory(&entity.CategoryMock2)
	//categoryService.StoreCategory(&entity.CategoryMock3)

	userHandler := handler.NewUserHandler(templ, userService, sessionService, roleServ, csrfSignKey)
	indexHandler := handler.NewIndexHandler(templ, questionService, categoryService)
	questionHandler := handler.NewQuestionHandler(templ, questionService, answerService, categoryService, csrfSignKey)
	adminHandler := handler.NewAdminUsersHandler(templ, userService, roleServ, csrfSignKey)
	adminQuestionHandler := handler.NewAdminQuestionHandler(templ, questionService, categoryService, csrfSignKey)
	adminAnswerHandler := handler.NewAdminAnswerHandler(templ, answerService, questionService, csrfSignKey)
	adminCategoriesHandler := handler.NewAdminCategoryHandler(templ, categoryService, csrfSignKey)

	http.Handle("/admin", userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(userHandler.Admin))))
	http.Handle("/admin/users", userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(adminHandler.AdminUsers))))
	http.Handle("/admin/users/update", userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(adminHandler.AdminUsersUpdate))))
	http.Handle("/admin/users/new", userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(adminHandler.AdminUsersNew))))
	http.Handle("/admin/users/delete", userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(adminHandler.AdminUsersDelete))))

	http.Handle("/admin/questions", userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(adminQuestionHandler.AdminQuestionHandler))))
	http.Handle("/admin/questions/delete", userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(adminQuestionHandler.AdminQuestionsDelete))))
	http.Handle("/admin/answers", userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(adminAnswerHandler.AdminAnswerHandler))))
	http.Handle("/admin/answers/delete", userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(adminAnswerHandler.AdminAnswersDelete))))
	http.Handle("/admin/categories", userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(adminCategoriesHandler.AdminCategoriesHandler))))
	http.Handle("/admin/categories/delete", userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(adminCategoriesHandler.AdminCategoriesDelete))))
	http.Handle("/admin/categories/new", userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(adminCategoriesHandler.AdminCategoriesNew))))
	http.Handle("/admin/categories/update", userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(adminCategoriesHandler.AdminCategoriesUpdate))))

	http.Handle("/", userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(indexHandler.Index))))
	http.Handle("/question", userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(questionHandler.QuestionHandler))))
	http.Handle("/question/new", userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(questionHandler.NewQuestion))))
	http.HandleFunc("/login", userHandler.Login)
	http.HandleFunc("/signup", userHandler.SignUp)
	http.Handle("/logout", userHandler.Authenticated(http.HandlerFunc(userHandler.Logout)))
	http.ListenAndServe(":8181", nil)
}
