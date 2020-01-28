package entity

var UserMock = User{
	FullName: "Yared Girma",
	Email:    "yaredgirmadb123@gmail.com",
	Password: "$2a$12$tGd/ljs96H7/nOwErVEq0uzPyydFtXZzvcSYNdiKy5MHvg93J28Py",
	RoleID:   1,
}

var UserMock2 = User{
	FullName: "Yared Girma",
	Email:    "yaredgirmadb1234@gmail.com",
	Password: "$2a$12$tGd/ljs96H7/nOwErVEq0uzPyydFtXZzvcSYNdiKy5MHvg93J28Py",
	RoleID:   2,
}

var UserMock3 = User{
	FullName: "Girma Deyaso",
	Email:    "yaredgirmadb123@outlook.com",
	Password: "$2a$12$tGd/ljs96H7/nOwErVEq0uzPyydFtXZzvcSYNdiKy5MHvg93J28Py",
	RoleID:   1,
}

var UserMock4 = User{
	FullName: "Deyaso Balcha",
	Email:    "yaredgirmadb12@gmail.com",
	Password: "$2a$12$tGd/ljs96H7/nOwErVEq0uzPyydFtXZzvcSYNdiKy5MHvg93J28Py",
	RoleID:   1,
}

var UserMock5 = User{
	FullName: "Nathan Getachew",
	Email:    "NathanGetachew@gmail.com",
	Password: "$2a$12$tGd/ljs96H7/nOwErVEq0uzPyydFtXZzvcSYNdiKy5MHvg93J28Py",
	RoleID:   1,
}
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
	CategoryID:  1,
	Answers:     nil,
}

var QuestionMock2 = Question{
	Title:       "About our Testing",
	Description: "Hello Gophers, Good Afternoon. I've noticed that we've not implemented our testing as of now so i was wondering when we were gonna do that? ",
	Image:       "",
	UserID:      1,
	CategoryID:  1,
	Answers:     nil,
}
var QuestionMock3 = Question{
	Title:       "About the BasketBall tournament",
	Description: "Hey guys. I am a fellow student who loves watching basketball matches and i was happy when the school anounced there will be one. but eh schedule for the matches are not posted. will there be a tournament? if so, when?",
	Image:       "",
	UserID:      1,
	CategoryID:  2,
	Answers:     nil,
}

var QuestionMock4 = Question{
	Title:       "About ETRSS-1 ",
	Description: "hey guys, I heard information about ETRSS-1 regarding how the information it collects will be disclosed to the people. will us average joes be able to view the incredible pictures the satelite will take or is it disclosed to authoried personel only?",
	Image:       "",
	UserID:      1,
	CategoryID:  3,
	Answers:     nil,
}

var QuestionMock5 = Question{
	Title:       "About Section-3's Soccer team",
	Description: "Hey guys, i heard that our section (Section 3 - SE 3rd year) is forming a soccer team for the competition to be held a week from now. are there any empty spaces left for me to join?",
	Image:       "",
	UserID:      5,
	CategoryID:  2,
	Answers:     nil,
}

var QuestionMock6 = Question{
	Title:       "About Computers",
	Description: "Hey guys, I'm a student thats very fascinated with computers and want to know how they work under the hood but I dont know where to acquire the resources to fulfill my curiosity. Any recomendations on where I should get some books or videos?",
	Image:       "",
	UserID:      5,
	CategoryID:  3,
	Answers:     nil,
}

var QuestionMock7 = Question{
	Title:       "About the OOP assignment",
	Description: "Hey guys, what day is the submission date of the OOP assignment for 3rd year students ",
	Image:       "",
	UserID:      3,
	CategoryID:  1,
	Answers:     nil,
}

var QuestionMock8 = Question{
	Title:       "About Github",
	Description: "Hey guys, I was having difficulties with using github's Web GUI. should I learn to use it or just continue with bash?",
	Image:       "",
	UserID:      3,
	CategoryID:  1,
	Answers:     nil,
}

var QuestionMock9 = Question{
	Title:       "About our HPE classes",
	Description: "Hey guys, I am a first year and i couldn't help but wonder if my classes will have physical education classes like they were saying in the news. will we have PE classes and if so then how will it be graded?",
	Image:       "",
	UserID:      4,
	CategoryID:  2,
	Answers:     nil,
}

var QuestionMock10 = Question{
	Title:       "About Hard drives",
	Description: "Hey guys, I wanted to ask regaring how hard drives work. Why do they call them mechanical drives and do they have a limitted lifespan?",
	Image:       "",
	UserID:      5,
	CategoryID:  3,
	Answers:     nil,
}

var QuestionMock11 = Question{
	Title:       "About the Championship prizes",
	Description: "Hey guys, I am a competitor in the soccer championship bieng held at 5k. what will be the prizes we will get if we happen to win the competition?",
	Image:       "",
	UserID:      5,
	CategoryID:  2,
	Answers:     nil,
}

var AnswerMock = Answer{
	UserID:     3,
	QuestionID: 1,
	Message:    "There are ways in which many to many relationships and other forms of relations can be shown in gorm. quit easily might I add. visit the native documentation for gorm for details",
}

var AnswerMock2 = Answer{
	UserID:     1,
	QuestionID: 2,
	Message:    "On Jan28th If all goes well :)",
}

var AnswerMock3 = Answer{
	UserID:     1,
	QuestionID: 2,
	Message:    "Today",
}

var AnswerMock4 = Answer{
	UserID:     1,
	QuestionID: 4,
	Message:    "Some, if not all, the info the satelite collects will be disclosed for public eyes. Don't expect a lot tho. Baby steps :)",
}

var AnswerMock5 = Answer{
	UserID:     4,
	QuestionID: 5,
	Message:    "Hey, Im sorry to tell you buut the emtpy spaces are filled as of now. I will notify you if things change.",
}

var AnswerMock6 = Answer{
	UserID:     1,
	QuestionID: 6,
	Message:    "try using linus's youtube channel (Linus Tech Tips) for a daily fix of tech news and if that's not enough then visit thier website",
}

var AnswerMock7 = Answer{
	UserID:     1,
	QuestionID: 7,
	Message:    "February 1st",
}

var AnswerMock8 = Answer{
	UserID:     3,
	QuestionID: 7,
	Message:    "January 31st",
}

var AnswerMock9 = Answer{
	UserID:     1,
	QuestionID: 9,
	Message:    "LOL. yes you will be taking PE classes. Don't worry though the classes you will be taking will be graded on the basis of attendance only",
}

var AnswerMock10 = Answer{
	UserID:     1,
	QuestionID: 10,
	Message:    "hey. Hard drives are a form of storage medium that are called mechanical cause they have spinning magnetic platers and moving pins that read and write to them. Hard drives theoretically can live forever tho they usually fail due to mechanical issues. read more on the web for better info. ",
}

var AnswerMock11 = Answer{
	UserID:     5,
	QuestionID: 3,
	Message:    "You're not gonna like this but the matches were cancelled :(",
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
