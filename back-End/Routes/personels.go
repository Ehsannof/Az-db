package routes

import (
	"net/http"
	"strconv"

	"example.com/models"
	"github.com/gin-gonic/gin"
)



func getPersonels(context *gin.Context) {
	var err error
    personel, err := models.GetAllPersonel()
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"message": "Error to get Personel"})
        return
    }
    context.JSON(http.StatusOK, personel)
}

func getPersonel(context *gin.Context){
	personelId,err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error to get Personel"})
        return
	}
	personel, err := models.GetPersonelByID(personelId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse personel id."})
		return
	}
	context.JSON(http.StatusOK, personel)
}

func AddPersonel(context *gin.Context){
    var personel models.Personel
    err := context.ShouldBindJSON(&personel)
    if err != nil{
        context.JSON(http.StatusBadRequest, gin.H{"message":"Cant parse Data!"})
        return
    }
    err = personel.Save()
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error to add data"})
		return
	}

    context.JSON(http.StatusCreated, gin.H{"message": "personel added!", "person": personel})
}

func updatePersonel(context *gin.Context){ 
	var err error
	personelId,err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error to get personels"})
        return
	}

	var updatedPersonel models.Personel

	err = context.ShouldBindJSON(&updatedPersonel)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error to get personels"})
        return
	}

	updatedPersonel.ID = personelId
	err = updatedPersonel.Update2()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error to Update personel"})
        return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Personel data updated"})
}

func deletePersonel(context *gin.Context){
	var err error
	personelId,err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error to get personels"})
        return
	}

	personel, err := models.GetPersonelByID(personelId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete personel"})
        return
	} 

	err = personel.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the personel"})
        return
	}
	context.JSON(http.StatusOK, gin.H{"message": "personel deleted"})
}
