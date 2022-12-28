package model

import (
	"app/database"
	"log"
)

type UserModel struct {
	Id       int    `form:"id"`
	Account  string `form:"account"`
	Password string `form:"password"`
}

func (user *UserModel) QueryByAccount() UserModel {
	u := UserModel{}
	row := database.Db.QueryRow("select * from user where account = ?", user.Account)
	e := row.Scan(&u.Id, &u.Account, &u.Password)
	if e != nil {
		log.Panicln(e)
	}
	return u
}
