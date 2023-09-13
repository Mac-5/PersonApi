package services

import "github.com/Mac-5/models"

// Here i am creating the contracts that will be implemented in the person.serviceimpl.go

type PersonService interface{
	CreatePerson(createRequest *models.CreatePersonModel)(*models.Person, error)
	GetOnePerson(id int)(*models.Person, error)
	GetAllPerson()([]*models.Person, error)
	UpdatePerson(id int, updatePerson *models.UpdatePersonModel)(*models.Person,error)
	DeletePerson(id int)error
}