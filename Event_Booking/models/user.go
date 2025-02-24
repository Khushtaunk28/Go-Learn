package models

import "Event_Booking/db"

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
	res,err:=stmt.Exec(u.Email,u.Password)
	if err!=nil{
		return err
	}

	userId,err:=res.LastInsertId()
	u.ID=userId
	return err
}