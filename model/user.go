package model

type User struct {
	Id       int64
	Password string `json:"password" form:"password" uri:"password" `
	Username string `json:"username" form:"username" uri:"username"`
}
