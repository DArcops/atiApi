package v1

import (
	"net/http"
	"strconv"

	b64 "encoding/base64"

	"github.com/darcops/atiApi/models"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users := models.GetAllUsers()

	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusCreated, users)
}

func Register(c *gin.Context) {
	var user models.User

	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	if err := models.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusCreated, gin.H{})
}

func Login(c *gin.Context) {
	var user models.User

	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	if models.First(&user, "email = ? and password = ?", user.Email, user.Password).RecordNotFound() {
		c.JSON(http.StatusForbidden, gin.H{})
		return
	}

	token, err := models.GenerateToken([]byte(user.Email + "+" + strconv.FormatUint(uint64(user.ID), 10)))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, gin.H{
		"token":      b64.StdEncoding.EncodeToString(token),
		"user_name":  user.Name,
		"user_email": user.Email,
	})
	return

}

func GetProfile(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	user.Password = ""
	c.JSON(http.StatusOK, user)
	return
}

func AddAdmin(c *gin.Context) {
	var newAdmin models.NewAdmin

	if err := c.Bind(&newAdmin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	user := c.MustGet("user").(models.User)
	if !user.CanWrite {
		c.JSON(http.StatusForbidden, gin.H{})
		return
	}

	if err := user.AddNewAdmin(newAdmin.Email); err != nil {
		c.JSON(Err[err], gin.H{})
		return
	}

	c.JSON(http.StatusCreated, user)
}
