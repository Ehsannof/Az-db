package models

import (
	"time"

	"example.com/db"
)

type Event struct{
	ID int64
	Name string `binding:"required"`
	Description string `binding:"required"`
	Price int64 `binding:"required"`
	DateTime time.Time `binding:"required"` 
} 

var events = []Event{}

func (e *Event) Save() error {
	query := `
	INSERT INTO events(name, description, price, dateTime) 
	VALUES (?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil{
		return err
	}
	defer stmt.Close()
	result,err := stmt.Exec(e.Name, e.Description, e.Price, e.DateTime)
	if err != nil{
		return err
	}
	id,err := result.LastInsertId()
	e.ID = id
	return err
}

func GetAllEvents() ([]Event, error){
	query := "SELECT * FROM EVENTS"
	result,err := db.DB.Query(query)
	if err != nil{
		return nil,err
	}

	defer result.Close()
	var events []Event

	for result.Next(){
		var event Event
		var err error
		err = result.Scan(&event.ID, &event.Name, &event.Description, &event.Price, &event.DateTime)
		if err != nil{
			return nil,err
		}
		events = append(events, event)
	}
	return events,nil
}

func GetEventByID(id int64) (*Event, error){
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event
	var err error

	err = row.Scan(&event.ID, &event.Name, &event.Description, &event.Price, &event.DateTime)
		if err != nil{
			return nil, err
		}

	return &event, nil
} 

func (event Event) Update()  error {
	query := `
		UPDATE events
		SET name = ?, Description = ?, Price = ?, dateTime = ?
		WHERE id = ?
	`
	stmt,err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Name,event.Description,event.Price, event.DateTime, event.ID)
	return err
}

func (event Event) Delete() error {
    query := `DELETE FROM events
    WHERE ID = ?
    `
    stmt, err := db.DB.Prepare(query)

    if err != nil {
        return err
    }

    defer stmt.Close()
    _,err = stmt.Exec(event.ID)

    return err
}

// func (e Event) Register(userId int64) error{
// 	query := `INSERT INTO registrations(event_id, user_id) VALUES (?, ?)`
// 	stmt,err := db.DB.Prepare(query)
// 	if err != nil {
// 		return err
// 	}

// 	defer stmt.Close()

// 	_,err = stmt.Exec(e.ID, userId)

// 	return err
// }

// func (e Event)CancelRegistration(userId int64)error{
// 	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?"
// 	stmt,err := db.DB.Prepare(query)
// 	if err != nil {
// 		return err
// 	}

// 	defer stmt.Close()

// 	_,err = stmt.Exec(e.ID, userId)

// 	return err
// }
