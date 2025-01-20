package models

import (
	"example.com/db"
)

type Personel struct{
	ID int64
	Email string `binding:"required"` 
	Code int64 
	Description string  
	BirthYear int64
}

func (c *Personel) Save() error {
	query := `
	INSERT INTO personel(email, code, description, BirthYear) 
	VALUES (?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil{
		return err
	}
	defer stmt.Close()
	result,err := stmt.Exec(c.Email, c.Code, c.Description, c.BirthYear)
	if err != nil{
		return err
	}
	id,err := result.LastInsertId()
	c.ID = id
	return err
}

func GetAllPersonel() ([]Personel, error){
	query := "SELECT * FROM personel"
	result,err := db.DB.Query(query)
	if err != nil{
		return nil,err
	}

	defer result.Close()
	var personels []Personel

	for result.Next(){
		var personel Personel
		var err error
		err = result.Scan(&personel.ID, &personel.Email, &personel.Code, &personel.Description, &personel.BirthYear)
		if err != nil{
			return nil,err
		}
		personels = append(personels, personel)
	}
	return personels,nil
}

func GetPersonelByID(id int64) (*Personel, error){
	query := "SELECT * FROM personel WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var personel Personel
	var err error

	err = row.Scan(&personel.ID, &personel.Email, &personel.Code, &personel.Description, &personel.BirthYear)
		if err != nil{
			return nil, err
		}

	return &personel, nil
} 

func (c Personel) Update2()  error {
	query := `
		UPDATE personel
		SET email = ?, code = ?, Description = ?, birthYear = ?
		WHERE id = ?
	`
	stmt,err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(c.Email,c.Code,c.Description,c.BirthYear,c.ID)
	return err
}

func (c Personel) Delete() error {
    query := `DELETE FROM personel
    WHERE ID = ?
    `
    stmt, err := db.DB.Prepare(query)

    if err != nil {
        return err
    }

    defer stmt.Close()
    _,err = stmt.Exec(c.ID)

    return err
}
