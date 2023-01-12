package model

import (
	"app/database"
	"log"
)

type UserModel struct {
	Id       int    `form:"id"`
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func (user UserModel) Store() int64 {
	result, err := database.Db.Exec("insert into user (email, password) value (?,?)", user.Email, user.Password)
	if err != nil {
		log.Panicln(err.Error())
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Panicln(err.Error())
	}
	return id
}

func (user UserModel) QueryByAccount() UserModel {
	u := UserModel{}
	row := database.Db.QueryRow("select * from user where email = ?", user.Email)
	err := row.Scan(&u.Id, &u.Email, &u.Password)
	if err != nil {
		log.Println(err.Error())
	}
	return u
}

func (user UserModel) FindAll() []UserModel {
	rows, err := database.Db.Query("select * from user")
	if err != nil {
		log.Println(err.Error())
	}

	var users []UserModel

	for rows.Next() {
		var u UserModel
		if e := rows.Scan(&u.Id, &u.Email, &u.Password); e == nil {
			users = append(users, u)
		}
	}
	return users
}

func (user UserModel) FindById() UserModel {
	row := database.Db.QueryRow("select * from user where id=?", user.Id)

	if err := row.Scan(&user.Id, &user.Email, &user.Password); err != nil {
		log.Println("binding error", err.Error())
	}
	return user
}

func (user UserModel) DeleteById() {
	if _, err := database.Db.Exec("delete  from user where id =?", user.Id); err != nil {
		log.Println("delete error", err.Error())
	}
}
