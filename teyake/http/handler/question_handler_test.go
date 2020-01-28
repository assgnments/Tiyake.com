package handler

import (
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"teyake/entity"
	"teyake/util/token"

	catRepoImp "teyake/category/repository"
	catServiceImp "teyake/category/service"
	quesRepoImp "teyake/question/repository"
	quesServiceImp "teyake/question/service"

	userRepoImp "teyake/user/repository"
	userServiceImp "teyake/user/service"
)

var client *http.Client
var server *httptest.Server
var csrfSignKey []byte

func init() {
	//templ := template.Must(template.New("main").Funcs(util.AvailableFuncMaps).ParseGlob("ui/templates/*"))

	mux := http.NewServeMux()
	userRepo := userRepoImp.NewMockUserRepo(
		map[uint]*entity.User{
			0: {
				Model: gorm.Model{
					ID: 0,
				},
				FullName: "Nathan",
				RoleID:   0,
			},
			1: {
				Model: gorm.Model{
					ID: 1,
				},
				FullName: "Habib",
				RoleID:   1,
			},
		},
		)
	userService := userServiceImp.NewUserService(userRepo)


	roleRepo := userRepoImp.NewMockRoleRepo(	map[string]*entity.Role{
		"USER": {
			Model: gorm.Model{
				ID: 0,
			},
			Name: "USER",
		},
		"ADMIN": {
			Model: gorm.Model{
				ID: 1,
			},
			Name: "ADMIN",
		},
	})
	roleServ := userServiceImp.NewRoleService(roleRepo)


	sessionRepo := userRepoImp.NewMockSessionRepo(map[uint]*entity.Session{
		0: {
			Model: gorm.Model{
				ID: 0,
			},
			UUID:       0,
			SessionId:  "0",
			SigningKey: []byte("demo_key"),
		},
		1: {
			Model: gorm.Model{
				ID: 1,
			},
			UUID:       1,
			SessionId:  "1",
			SigningKey: []byte("demo_key"),
		},
	}, )
	sessionService := userServiceImp.NewSessionService(sessionRepo)

	questionRepo := quesRepoImp.NewMockQuestionRepo(map[uint]*entity.Question{
		0: {
			Model:       gorm.Model{
				ID:1,
			},
			Title:       "This is a question",
			Description: "Demo Question",
			UserID:      0,
			CategoryID:  0,
			Answers:     nil,
		},
	})
	questionService := quesServiceImp.NewQuestionService(questionRepo)

	categoryRepo := catRepoImp.NewMockCategoryRepo(map[uint]*entity.Category{
		0: {
			Model: gorm.Model{
				ID: 0,
			},
			Name: "Science",
		},
		1: {
			Model: gorm.Model{
				ID: 0,
			},
			Name: "Programming",
		}},
	)
	categoryService := catServiceImp.NewCategoryService(categoryRepo)

	userHandler := NewUserHandler(nil, userService, sessionService, roleServ, csrfSignKey)
	questionHandler := NewQuestionHandler(nil, questionService, nil, categoryService, nil, csrfSignKey)
	indexHandler := NewIndexHandler(nil, questionService, categoryService)
	http.Handle("/question", userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(questionHandler.QuestionHandler))))
	http.Handle("/question/new", userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(questionHandler.NewQuestion))))
	http.Handle("/", userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(indexHandler.Index))))
	server = httptest.NewTLSServer(mux)
	client = server.Client()
	csrfSignKey = []byte(token.GenerateRandomID(32))

}

func TestIndexHandler(t *testing.T) {
	defer server.Close()
	resp, err := client.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, resp.StatusCode)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	println(body)

}
