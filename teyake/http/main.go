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

	upvoteRepoImp "teyake/upvote/repository"
	upvoteServiceImp "teyake/upvote/service"
	userRepoImp "teyake/user/repository"
	userServiceImp "teyake/user/service"
	"teyake/teyake/http/handler"
	"teyake/util"
	"teyake/util/token"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func createTables(dbconn *gorm.DB) []error {
	errs := dbconn.CreateTable(&entity.Role{},&entity.User{},&entity.Session{},&entity.Question{},&entity.Category{},entity.Answer{},entity.UpVote{}).GetErrors()
	if errs != nil {
		return errs
	}
	return nil
}

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

	upvoteRepo := upvoteRepoImp.NewUpVoteGormRepo(dbconn)
	upvoteService := upvoteServiceImp.NewUpVoteService(upvoteRepo)

	//Uncomment the following lines after you created a fresh teyake db
	
	//createTables(dbconn)
	//
	//roleServ.StoreRole(&entity.UserRoleMock)
	//roleServ.StoreRole(&entity.AdminRoleMock)
	//
	//userService.StoreUser(&entity.UserMock)
	//userService.StoreUser(&entity.UserMock2)
	//userService.StoreUser(&entity.UserMock3)
	//userService.StoreUser(&entity.UserMock4)
	//userService.StoreUser(&entity.UserMock5)
	//
	//questionService.StoreQuestion(&entity.QuestionMock)
	//questionService.StoreQuestion(&entity.QuestionMock2)
	//questionService.StoreQuestion(&entity.QuestionMock3)
	//questionService.StoreQuestion(&entity.QuestionMock4)
	//questionService.StoreQuestion(&entity.QuestionMock5)
	//questionService.StoreQuestion(&entity.QuestionMock6)
	//questionService.StoreQuestion(&entity.QuestionMock7)
	//questionService.StoreQuestion(&entity.QuestionMock8)
	//questionService.StoreQuestion(&entity.QuestionMock9)
	//questionService.StoreQuestion(&entity.QuestionMock10)
	//questionService.StoreQuestion(&entity.QuestionMock11)
	//
	//answerService.StoreAnswer(&entity.AnswerMock)
	//answerService.StoreAnswer(&entity.AnswerMock2)
	//answerService.StoreAnswer(&entity.AnswerMock3)
	//answerService.StoreAnswer(&entity.AnswerMock4)
	//answerService.StoreAnswer(&entity.AnswerMock5)
	//answerService.StoreAnswer(&entity.AnswerMock6)
	//answerService.StoreAnswer(&entity.AnswerMock7)
	//answerService.StoreAnswer(&entity.AnswerMock8)
	//answerService.StoreAnswer(&entity.AnswerMock9)
	//answerService.StoreAnswer(&entity.AnswerMock10)
	//answerService.StoreAnswer(&entity.AnswerMock11)
	//
	//upvoteService.StoreUpVote(&entity.UpVoteMock)
	//
	//categoryService.StoreCategory(&entity.CategoryMock1)
	//categoryService.StoreCategory(&entity.CategoryMock2)
	//categoryService.StoreCategory(&entity.CategoryMock3)
	//

	userHandler := handler.NewUserHandler(templ, userService, sessionService, roleServ, csrfSignKey)
	indexHandler := handler.NewIndexHandler(templ, questionService, categoryService)
	questionHandler := handler.NewQuestionHandler(templ, questionService, answerService, categoryService, upvoteService, csrfSignKey)
	adminHandler := handler.NewAdminUsersHandler(templ, userService, roleServ, csrfSignKey)
	adminquestionHandler := handler.NewAdminQuestionHandler(templ, questionService, categoryService, csrfSignKey)
	adminanswerHandler := handler.NewAdminAnswerHandler(templ, answerService, questionService, csrfSignKey)
	admincategoryHandler := handler.NewAdminCategoryHandler(templ, categoryService, csrfSignKey)

	http.Handle("/admin", userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(userHandler.Admin))))
	http.Handle("/admin/users", userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(adminHandler.AdminUsers))))
	http.Handle("/admin/users/update", userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(adminHandler.AdminUsersUpdate))))
	http.Handle("/admin/users/delete", userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(adminHandler.AdminUsersDelete))))
	http.Handle("/admin/users/new", userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(adminHandler.AdminUsersNew))))

	http.Handle("/admin/categories", userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(admincategoryHandler.AdminCategoriesHandler))))
	http.Handle("/admin/categories/delete", userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(admincategoryHandler.AdminCategoriesDelete))))
	http.Handle("/admin/categories/update", userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(admincategoryHandler.AdminCategoriesUpdate))))
	http.Handle("/admin/categories/new", userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(admincategoryHandler.AdminCategoriesNew))))

	http.Handle("/admin/questions", userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(adminquestionHandler.AdminQuestionHandler))))
	http.Handle("/admin/questions/delete", userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(adminquestionHandler.AdminQuestionsDelete))))

	http.Handle("/admin/answers", userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(adminanswerHandler.AdminAnswerHandler))))
	http.Handle("/admin/answers/delete", userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(adminanswerHandler.AdminAnswersDelete))))

	http.Handle("/", userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(indexHandler.Index))))
	http.Handle("/question", userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(questionHandler.QuestionHandler))))
	http.Handle("/question/new", userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(questionHandler.NewQuestion))))
	http.Handle("/question/upvote",userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(questionHandler.UpvoteHandler))))
	http.HandleFunc("/login", userHandler.Login)
	http.HandleFunc("/signup", userHandler.SignUp)
	http.HandleFunc("/question/search", indexHandler.SearchQuestions)
	http.Handle("/logout", userHandler.Authenticated(http.HandlerFunc(userHandler.Logout)))
	http.ListenAndServe(":8181", nil)
}
