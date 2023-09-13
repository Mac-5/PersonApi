package services

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Mac-5/models"
	_ "github.com/lib/pq"
	
)

type PersonServiceImpl struct{
	db *sql.DB
	ctx context.Context
}

func NewPersonSeviceImpl(db *sql.DB, ctx context.Context) PersonService{
	return &PersonServiceImpl{
		db: db,
		ctx: ctx,
		
	}
}

// Create-Person Implementation
func (p *PersonServiceImpl) CreatePerson(Createperson *models.CreatePersonModel)(*models.Person, error){
	if Createperson == nil {
			return nil, errors.New("person cant be nil")
	}
	// Prepare the SQL statement for inserting a new person
	query := "INSERT INTO persons (name, email) VALUES ($1, $2) RETURNING id"
	stmt, err := p.db.PrepareContext(p.ctx, query)
	if err != nil{
		return nil, err
	}
	defer stmt.Close()
	var personID int
	err = stmt.QueryRowContext(p.ctx, Createperson.Name,Createperson.Email).Scan(&personID)
	if err != nil{
		return nil, err
	}

	// Create a Preson struct with the generated ID

	person := &models.Person{
		ID: personID,
		Name: Createperson.Name,
		Email: Createperson.Email,
	}
	
	return person,  nil

}


//Get-Person-By-Id implementation

func (p *PersonServiceImpl) GetOnePerson(id int)(*models.Person, error){
	if id <= 0{
		return nil, errors.New("invalid ID")
	}

	//Sql statement for fetching by id
	query := "SELECT id, name, email FROM persons WHERE id = $1"
	row := p.db.QueryRowContext(p.ctx, query, id)

	// Create a person struct to store the result 
	var person models.Person
	err := row.Scan(&person.ID, &person.Name, &person.Email)
	if err != nil{
		if err == sql.ErrNoRows{
			return nil, errors.New("person not found")
		}

		return nil, err
	}
	return &person, nil
}

// Get All Persons


func (p *PersonServiceImpl) GetAllPerson()([]*models.Person, error){
	query := "SELECT id, name, email FROM persons"
	rows, err := p.db.QueryContext(p.ctx, query)
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	var persons []*models.Person

	for rows.Next(){
		var person models.Person
		if err := rows.Scan(&person.ID, &person.Name, &person.Email); err != nil{
			return nil, err
		}
		persons  = append(persons, &person)

	}
	if err := rows.Err(); err != nil{
		return nil, err
	}
	return persons, nil
}


// Update-Person function

func (p *PersonServiceImpl) UpdatePerson (id int, updatePerson *models.UpdatePersonModel)(*models.Person, error){
	if id <= 0{
		return nil, errors.New("invalid id")
	}
	if updatePerson == nil{
		return nil, errors.New("updated persons data doesnot exist")
	}

	query := "UPDATE persons SET name = $1, email = $2 WHERE id = $3 RETURNING id"
	stmt, err := p.db.PrepareContext(p.ctx, query)
	if err != nil{
		return nil, err
	}
	defer stmt.Close()

	var personID int
	err = stmt.QueryRowContext(p.ctx, updatePerson.Name, updatePerson.Email, id).Scan(&personID)
	if err != nil{
		if err == sql.ErrNoRows{
			return nil, errors.New("person not found")
		}
		return nil, err
	}
	if personID == 0 {
		return nil, errors.New("person not found")
	}

	updated := &models.Person{
		ID: personID,
		Name: updatePerson.Name,
		Email: updatePerson.Email,
	}
	return updated, nil
}


// Delete-Person Function
func (p *PersonServiceImpl) DeletePerson (id int)error{
	if id <= 0 {
		return errors.New("invalid id")
	}
	query := "DELETE FROM persons WHERE id = $1"
	stmt, err := p.db.PrepareContext(p.ctx, query)
	if err != nil{
		return err
	}
	result, err := stmt.ExecContext(p.ctx, id)
	if err != nil{
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil{
		return err
	}
	if rowsAffected == 0{
		return errors.New("person not found")
	}
	return nil
}