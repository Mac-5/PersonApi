package controllers

import (
	"net/http"
	"strconv"

	"github.com/Mac-5/models"
	"github.com/Mac-5/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PersonController struct {
	PersonService services.PersonService
}

func NewPersonController(personService services.PersonService) *PersonController {
	return &PersonController{
		PersonService: personService,
	}
}


func (pc *PersonController) GetOnePerson (c *gin.Context){
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})

		return 
	}
	person, err := pc.PersonService.GetOnePerson(id)
	if err != nil{
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, person)
}


func (pc *PersonController) CreatePerson(c *gin.Context){
	var person models.CreatePersonModel

	if err := c.ShouldBindJSON(&person); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if person.Name == ""{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Name is Required",
		})
		return
	}
	validate := validator.New()
	if err := validate.Struct(person); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	createdPerson, err := pc.PersonService.CreatePerson(&person)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create person",
		})
		return
	}
	c.JSON(http.StatusCreated, createdPerson)
}


func (pc *PersonController) GetAllPerson(c *gin.Context){
	persons, err:= pc.PersonService.GetAllPerson()
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve persons",
		})
		return
	}
	c.JSON(http.StatusOK, persons)
}


func (pc *PersonController) UpdatePerson (c *gin.Context){
	idParam := c.Param("id")
	id, err :=strconv.Atoi(idParam)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatePerson models.UpdatePersonModel
	if err := c.ShouldBindJSON(&updatePerson); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return 
	}

	person, err := pc.PersonService.UpdatePerson(id, &updatePerson)
	if err != nil{
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return 
	}

	c.JSON(http.StatusOK, person)

}

func (pc *PersonController) DeletePerson(c *gin.Context){
	idParam := c.Param("id")
	id, err :=strconv.Atoi(idParam)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = pc.PersonService.DeletePerson(id)
	if err != nil{
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return 
	}

	c.JSON(http.StatusNoContent, gin.H{
		"message": "The Person Has been deleted",
	})


}