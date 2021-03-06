package users

import (
	"fmt"
	"github.com/DevAgani/bookstore_users-api/datasources/mysql/users_db"
	"github.com/DevAgani/bookstore_users-api/logger"
	"github.com/DevAgani/bookstore_users-api/utils/errors"
	"github.com/DevAgani/bookstore_users-api/utils/mysql_utils"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	errorNoRows = "no rows in result set"
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created,status,password) VALUES(?, ?, ?, ?,?,?);"
	queryGetUser = "SELECT id, first_name, last_name, email, date_created,status FROM users where id=?;"
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser = "DELETE FROM users where id=?;"
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, date_created,status from users where status=?;"
)


func (user *User) Get()  *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil{
		logger.Error("error when trying to prepare get user statement",err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName,&user.LastName,&user.Email,&user.DateCreated,&user.Status);err != nil{
		logger.Error("error when trying to  get a user ",err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func (user *User) Save() *errors.RestErr  {
	stmt, err  := users_db.Client.Prepare(queryInsertUser)
	if err != nil{
		logger.Error("error when trying to prepare save user statement",err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	insertResult, insertErr := stmt.Exec(user.FirstName,user.LastName,user.Email,user.DateCreated,user.Status,user.Password)
	if insertErr != nil{
		logger.Error("error when trying to insert",err)
		return errors.NewInternalServerError("database error")
	}
	userId, err := insertResult.LastInsertId()
	if err != nil{
		logger.Error("error when trying to return last insert id",err)
		return errors.NewInternalServerError("database error")
	}
	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestErr  {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil{
		logger.Error("error when trying to prepare update user statement",err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.FirstName,user.LastName,user.Email,user.Id)
	if err != nil{
		logger.Error("error when trying to update a user",err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func (user *User) Delete() *errors.RestErr{
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil{
		logger.Error("error when trying to prepare delete user statement",err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	if _,err = stmt.Exec(user.Id);err != nil{
		logger.Error("error when trying to delete a user",err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func (user *User)FindByStatus(status string) ([]User, *errors.RestErr)  {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil{
		logger.Error("error when trying to prepare find by status user statement",err)
		return nil,errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	rows,err := stmt.Query(status)
	if err != nil{
		logger.Error("error when trying to find a user by status",err)
		return nil,errors.NewInternalServerError("database error")
	}
	defer rows.Close()
	results := make([]User,0)
	for rows.Next(){
		var user User
		if err := rows.Scan(&user.Id,&user.FirstName,&user.LastName,&user.Email,&user.DateCreated,&user.Status);err != nil{
			return nil,mysql_utils.ParseError(err)
		}
		results = append(results, user)
	}
	if len(results)== 0{
		return nil,errors.NewNotFoundError(fmt.Sprintf("no users matching status %s",status))
	}
	return results,nil
}
