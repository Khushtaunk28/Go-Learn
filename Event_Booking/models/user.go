package models

import (
	"Event_Booking/db"
	"Event_Booking/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required`
	Password string `binding:"required`
}

// save new user to the db
func (u User) Save() error {
	query := "INSERT INTO users(email,password) VALUES(?,?)"
	stmt,err:=db.DB.Prepare(query)
	if err!=nil{
		return err
	}
	defer stmt.Close()
	hashpsw,err:=utils.HashPassword(u.Password)
	if err !=nil{
		return err
	}
	res,err:=stmt.Exec(u.Email,hashpsw)
	if err!=nil{
		return err
	}

	userId,err:=res.LastInsertId()
	u.ID=userId
	return err
}