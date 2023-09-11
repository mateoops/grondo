package controllers

import (
	"grondo/db/sql"
	"grondo/models"

	"github.com/gin-gonic/gin"
)

// UserPassCreate godoc
// @Summary      Create User with Password (UserPass) object
// @Tags         cronjob credentials
// @Accept 		 json
// @Produce      json
// @Success      200  {array}  models.UserWithPassword
// @Failure      500  {object}  string
// @Param        password body models.UserWithPassword true "UserWithPassword object"
// @Router       /user/password [post]
func UserPassCreate(c *gin.Context) {
	var body struct {
		Username string
		Password string
	}

	c.Bind(&body)

	credential := models.UserWithPassword{Username: body.Username, Password: body.Password}

	result := sql.DB.Create(&credential)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"result": "success",
	})
}

// UserPassIndex godoc
// @Summary      List all UserWithPassword objects
// @Tags         cronjob credentials
// @Produce      json
// @Success      200  {object}  models.UserWithPassword
// @Failure      500  {object}  string
// @Router       /user/password [get]
func UserPassIndex(c *gin.Context) {
	var credentials []models.UserWithPassword

	result := sql.DB.Find(&credentials)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"result": credentials,
	})
}

// UserPassShow godoc
// @Summary      Get UserWithPassword object by providing ID
// @Tags         cronjob credentials
// @Produce      json
// @Success      200  {array}  models.UserWithPassword
// @Failure      500  {object}  string
// @Param        id  path  int  true  "ID"
// @Router       /user/password/{id} [get]
func UserPassShow(c *gin.Context) {
	id := c.Param("id")

	var credential models.UserWithPassword

	result := sql.DB.First(&credential, id)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"result": credential,
	})
}

// UserPassUpdate godoc
// @Summary      Update UserWithPassword object by providing ID
// @Tags         cronjob credentials
// @Accept 		 json
// @Produce      json
// @Success      200  {array}  models.UserWithPassword
// @Failure      500  {object}  string
// @Param        password body models.UserWithPassword true "UserWithPassword object"
// @Param        id  path  int  true  "ID"
// @Router       /user/password/{id} [put]
func UserPassUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Username string
		Password string
	}

	c.Bind(&body)

	credential := models.UserWithPassword{Username: body.Username, Password: body.Password}

	result := sql.DB.Model(&credential).Where("id = ?", id).Updates(&credential)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"result": "success",
	})
}

// UserPassDelete godoc
// @Summary      Delete UserWithPassword object by providing ID
// @Tags         cronjob credentials
// @Produce      json
// @Success      200  {object}  string
// @Failure      500  {object}  string
// @Param        id  path  int  true  "ID"
// @Router       /user/password/{id} [delete]
func UserPassDelete(c *gin.Context) {
	id := c.Param("id")

	var credential models.UserWithPassword

	result := sql.DB.Delete(&credential, id)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"result": "success",
	})
}

// UserSSHCreate godoc
// @Summary      Create User with SSHKey (UserWithSSHKey) object
// @Tags         cronjob credentials
// @Accept 		 json
// @Produce      json
// @Success      200  {array}  models.UserWithSSHKey
// @Failure      500  {object}  string
// @Param        sshkey body models.UserWithSSHKey true "UserWithSSHKey object"
// @Router       /user/sshkey [post]
func UserSSHCreate(c *gin.Context) {
	var body struct {
		Username   string
		SSHKey     string
		Passphrase string
	}

	c.Bind(&body)

	credential := models.UserWithSSHKey{Username: body.Username, SSHKey: body.SSHKey, Passphrase: body.Passphrase}

	result := sql.DB.Create(&credential)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"result": "success",
	})
}

// UserSSHIndex godoc
// @Summary      List all UserWithSSHKey objects
// @Tags         cronjob credentials
// @Produce      json
// @Success      200  {object}  models.UserWithSSHKey
// @Failure      500  {object}  string
// @Router       /user/sshkey [get]
func UserSSHIndex(c *gin.Context) {
	var credentials []models.UserWithSSHKey

	result := sql.DB.Find(&credentials)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"result": credentials,
	})
}

// UserSSHShow godoc
// @Summary      Get UserWithSSHKey object by providing ID
// @Tags         cronjob credentials
// @Produce      json
// @Success      200  {array}  models.UserWithSSHKey
// @Failure      500  {object}  string
// @Param        id  path  int  true  "ID"
// @Router       /user/sshkey/{id} [get]
func UserSSHShow(c *gin.Context) {
	id := c.Param("id")

	var credential models.UserWithSSHKey

	result := sql.DB.First(&credential, id)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"result": credential,
	})
}

// UserSSHUpdate godoc
// @Summary      Update UserWithSSHKey object by providing ID
// @Tags         cronjob credentials
// @Accept 		 json
// @Produce      json
// @Success      200  {array}  models.UserWithSSHKey
// @Failure      500  {object}  string
// @Param        sshkey body models.UserWithSSHKey true "UserWithSSHKey object"
// @Router       /user/sshkey/{id} [put]
func UserSSHUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Username   string
		SSHKey     string
		Passphrase string
	}

	c.Bind(&body)

	credential := models.UserWithSSHKey{Username: body.Username, SSHKey: body.SSHKey, Passphrase: body.Passphrase}

	result := sql.DB.Model(&credential).Where("id = ?", id).Updates(&credential)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"result": "success",
	})
}

// UserSSHDelete godoc
// @Summary      Delete UserWithSSHKey object by providing ID
// @Tags         cronjob credentials
// @Produce      json
// @Success      200  {object}  string
// @Failure      500  {object}  string
// @Param        id  path  int  true  "ID"
// @Router       /user/sshkey/{id} [delete]
func UserSSHDelete(c *gin.Context) {
	id := c.Param("id")

	var credential models.UserWithSSHKey

	result := sql.DB.Delete(&credential, id)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"result": "success",
	})
}
