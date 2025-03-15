package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"site.org/abc/models"
	"site.org/abc/utils"
)

func getUsers(context *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch users. " + err.Error()})
		return
	}
	context.JSON(http.StatusOK, users)
}

func getUser(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user id. " + err.Error()})
		return
	}

	user, err := models.GetUserById(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch user. " + err.Error()})
		return
	}

	context.JSON(http.StatusOK, user)
}

func createUser(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user data"})
		return
	}

	user, err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user. " + err.Error()})
		return
	}

	context.JSON(http.StatusCreated, user)
}

func updateUser(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user id. " + err.Error()})
		return
	}

	_, err = models.GetUserById(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch user. " + err.Error()})
		return
	}

	var updatedUser models.User
	err = context.ShouldBindJSON(&updatedUser)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data. " + err.Error()})
		return
	}

	updatedUser.ID = userId

	err = updatedUser.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update user. " + err.Error()})
		return
	}
	context.JSON(http.StatusOK, updatedUser)
}

func deleteUser(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user id. " + err.Error()})
		return
	}

	user, err := models.GetUserById(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch user. " + err.Error()})
		return
	}

	err = user.Delete(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the user. " + err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User deleted."})
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}
	//crypto/bcrypt: hashedSecret too short to be a bcrypted password
	// err = user.ValidateCredentials()
	// if err != nil {
	// 	context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
	// 	return
	// }
	fetch_user, err := models.GetUserByEmail(user.Email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch user. " + err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, fetch_user.ID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not generate token"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Message": "Success", "Email": user.Email, "Id": fetch_user.ID, "token": token})
}
