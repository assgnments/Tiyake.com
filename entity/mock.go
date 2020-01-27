package entity

//var UserMock = User{
//	FullName :"Yared",
//	Email  :  "yaredgirmadb123@gmail.com",
//	Password :"123123123",
//	RoleID  : 1,
//}
//
var AdminRoleMock = Role{
	Name: "ADMIN",
}
var UserRoleMock = Role{
	Name: "USER",
}

var QuestionMock = Question{
	Title:       "Important Question",
	Description: "Hello Gophers, Good Afternoon. Recently I come across an issue in GORM github.com/jinzhu/gorm 18. In the associations like One2One, One2Many and Many2Many this GORM is not adding the Foreign Key and so we need to add this foreign key relation manually. So anybody knows this how to do without manual work, let us know.",
	Image:       "",
	UserID:      1,
	CategoryID:  0,
	Answers:     nil,
}

var AnswerMock = Answer{
	UserID:     1,
	QuestionID: 1,
	Message:    "Yo chalaw",
}

var UpVoteMock = UpVote{
	UserID:   1,
	AnswerID: 1,
}

var CategoryMock1 = Category{

	Name: "Programming",
}
var CategoryMock2 = Category{

	Name: "Sport",
}
var CategoryMock3 = Category{

	Name: "Science",
}
