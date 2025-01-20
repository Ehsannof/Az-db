package models

import (
	"errors"

	"example.com/db"
	"example.com/utils"
)

type User struct{
	ID int64
	Email string `binding:"required"` 
	Password string `binding:"required"` 
	Name string
	PhoneNumber string
	Address string  
	BirthYear int64
}

func (user User) Save() error{
	query := `INSERT INTO users(email, password, name, PhoneNumber, Address, BirthYear) VALUES (?, ?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil{
		return err
	}

	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil{
		return err
	}
	result , err := stmt.Exec(user.Email, hashedPassword, user.Name, user.PhoneNumber, user.Address, user.BirthYear)

	if err != nil {
		return err	
	}
	user.ID, err = result.LastInsertId()
	return err
}

func (user *User) ValidateCredentials() error {
   query := "SELECT id, password FROM users WHERE email=?"
   row := db.DB.QueryRow(query, user.Email)     
   
   var retrievedPassword string
   err := row.Scan(&user.ID, &retrievedPassword)

   if err != nil{
       return err
   }

   passwordIsValid := utils.CheckPasswordHash(user.Password, retrievedPassword)
   if !passwordIsValid{
        return errors.New("Credentials invalid")
   }

   return nil
}

func GetAllUsers() ([]User, error){
	query := "SELECT * FROM users"
	result,err := db.DB.Query(query)
	if err != nil{
		return nil,err
	}

	defer result.Close()
	var users []User
	for result.Next(){
		var user User
		var err error
		err = result.Scan(&user.ID, &user.Email, &user.Password, &user.Name, &user.PhoneNumber, &user.Address, &user.BirthYear)
		if err != nil{
			return nil,err
		}
		users = append(users, user)
	}
	return users,nil
}

    // "email": "loqeb@mailinator.com",
    // "password": "123"
