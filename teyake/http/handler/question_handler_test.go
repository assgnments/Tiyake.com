package handler

import (
	"bytes"
	"fmt"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/html"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	catRepoImp "teyake/category/repository"
	catServiceImp "teyake/category/service"
	"teyake/entity"
	quesRepoImp "teyake/question/repository"
	quesServiceImp "teyake/question/service"
	"teyake/util"
	"teyake/util/token"
	"time"

	userRepoImp "teyake/user/repository"
	userServiceImp "teyake/user/service"
)

var client *http.Client
var server *httptest.Server
var csrfSignKey []byte
var signingKey []byte

func init() {
	templ := template.Must(template.New("main").Funcs(util.AvailableFuncMaps).ParseGlob("../../../ui/templates/*"))
	signingKey=[]byte("demo_key")
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

	roleRepo := userRepoImp.NewMockRoleRepo(map[uint]*entity.Role{
		0: {
			Model: gorm.Model{
				ID: 0,
			},
			Name: "USER",
		},
		1: {
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
			SigningKey: signingKey,
		},
	}, )
	sessionService := userServiceImp.NewSessionService(sessionRepo)

	questionRepo := quesRepoImp.NewMockQuestionRepo(map[uint]*entity.Question{
		0: {
			Model: gorm.Model{
				ID: 1,
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

	userHandler := NewUserHandler(templ, userService, sessionService, roleServ, csrfSignKey)
	questionHandler := NewQuestionHandler(templ, questionService, nil, categoryService, nil, csrfSignKey)
	indexHandler := NewIndexHandler(templ, questionService, categoryService)
	mux.Handle("/question", userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(questionHandler.QuestionHandler))))
	mux.Handle("/question/new", userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(questionHandler.NewQuestion))))
	mux.Handle("/", userHandler.Authenticated(userHandler.Authorized(http.HandlerFunc(indexHandler.Index))))
	server = httptest.NewTLSServer(mux)
	client = server.Client()
	csrfSignKey = []byte(token.GenerateRandomID(32))

}

func TestIndexHandler(t *testing.T) {
	defer server.Close()
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/", server.URL), nil)
	if err != nil {
		t.Fatal(err)
	}

	req.AddCookie(getCookie())
	resp,err:=client.Do(req)
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
	displayHtml(body)

}

func displayHtml(body []byte)  {
		doc, _ := html.Parse(strings.NewReader(string(body)))

	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, doc)
	fmt.Println(buf.String())
}

//func crawlAndFindString(key string, body []byte) bool{
//	doc, _ := html.Parse(strings.NewReader(string(body)))
//	//displayHtml(doc)
//	var crawler func(*html.Node) bool
//	crawler = func(node *html.Node) bool{
//		println("Checking "+node.Data)
//		if node.Data == key {
//			println("Found "+node.Data)
//			return  true
//		}
//		for child := node.FirstChild; child != nil; child = child.NextSibling {
//			crawler(child)
//		}
//		return false
//	}
//	return  crawler(doc)
//}

func getCookie() *http.Cookie  {
	claims := token.NewClaims("0", time.Now().AddDate(1,0,0).Unix())
	signedString, _ := token.Generate(signingKey, claims)

	return &http.Cookie{
		Name:       "session_key",
		Value:      signedString,
		Expires:    time.Now().AddDate(1,0,0),
		RawExpires: "",
		MaxAge:     500000,
		HttpOnly:   false,
	}
}
