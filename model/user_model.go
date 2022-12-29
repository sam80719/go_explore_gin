package model

import (
	"app/database"
	"log"
)

type UserModel struct {
	Id       int    `form:"id"`
	Account  string `form:"account" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func (user *UserModel) Store() int64 {
	result, err := database.Db.Exec("insert into user (account, password) value (?,?)", user.Account, user.Password)
	if err != nil {
		log.Panicln(err.Error())
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Panicln(err.Error())
	}
	return id
}

func (user *UserModel) QueryByAccount() UserModel {
	u := UserModel{}
	row := database.Db.QueryRow("select * from user where account = ?", user.Account)
	err := row.Scan(&u.Id, &u.Account, &u.Password)
	if err != nil {
		log.Panicln(err.Error())
	}
	return u
}
