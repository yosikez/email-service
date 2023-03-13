package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yosikez/email-notification-services/database"
	"github.com/yosikez/email-notification-services/model"
)

type MailController struct{}

func NewMailController() *MailController {
	return &MailController{}
}

func (m *MailController) FindAll(c *gin.Context) {
	var mails []model.Mail

	result := database.DB.Find(&mails)

	if err := result.Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to find mails",
			"error":   err.Error(),
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "failed to find mails",
			"error":   "record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": mails,
	})
}

func (m *MailController) FindById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid mail id",
			"error":   "id must be a number",
		})
		return
	}

	var mail model.Mail
	if err := database.DB.First(&mail, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "failed to find mail",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": mail,
	})
}

func (m *MailController) FindByAction(c *gin.Context) {
	action := c.Param("action")

	if action != "update" && action != "delete" && action != "create" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid action",
			"error":   "allowed action just update, create or delete",
		})
		return
	}

	var mails []model.Mail
	result := database.DB.Where("action = ?", action).Find(&mails)

	if err := result.Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to find mails",
			"error":   err.Error(),
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "failed to find mails",
			"error":   "record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data" : mails,
	})
}
