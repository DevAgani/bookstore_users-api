package users

import (
	"github.com/DevAgani/bookstore_users-api/datasources/mysql/users_db"
	"github.com/DevAgani/bookstore_users-api/utils/date_utils"
	"github.com/DevAgani/bookstore_users-api/utils/errors"
	"github.com/DevAgani/bookstore_users-api/utils/mysql_utils"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	errorNoRows = "no rows in result set"
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser = "SELECT id, first_name, last_name, email, date_created FROM users where id=?;"
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser = "DELETE FROM users where id=?;"
)


func (user *User) Get()  *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil{
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName,&user.LastName,&user.Email,&user.DateCreated);err != nil{
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Save() *errors.RestErr  {
	stmt, err  := users_db.Client.Prepare(queryInsertUser)
	if err != nil{
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	insertResult, insertErr := stmt.Exec(user.FirstName,user.LastName,user.Email,user.DateCreated)
	if insertErr != nil{
		return mysql_utils.ParseError(insertErr)
	}
	userId, err := insertResult.LastInsertId()
	if err != nil{
		return mysql_utils.ParseError(insertErr)
	}
	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestErr  {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil{
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.FirstName,user.LastName,user.Email,user.Id)
	if err != nil{
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Delete() *errors.RestErr{
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil{
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	if _,err = stmt.Exec(user.Id);err != nil{
		return mysql_utils.ParseError(err)
	}
	return nil
}
