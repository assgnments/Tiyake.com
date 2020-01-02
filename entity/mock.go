package entity

var UserMock = User{
	ID    : 1,
	FullName :"Habib",
	Email  :  "habib@gmail.com",
	Password :"123123123",
	RoleID  : 1,
}

var AdminRoleMock= Role{
	ID:   1,
	Name: "ADMIN",
}
var UserRoleMock= Role{
	ID:   2,
	Name: "USer",
}

