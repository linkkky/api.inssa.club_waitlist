package controllers

import (
	"inssa_club_waitlist_backend/cmd/server/errors"
	"inssa_club_waitlist_backend/cmd/server/forms"
	"inssa_club_waitlist_backend/cmd/server/models"
	"inssa_club_waitlist_backend/cmd/server/utils"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func (ctrler *Controller) AddInterest(c *gin.Context) {
	var form forms.AddInterest
	var count int64
	db := utils.GetDB().Instance

	if err := c.ShouldBindJSON(&form); err != nil {
		utils.AbortWithErrorResponse(c, http.StatusBadRequest, errors.ValidationError, err.Error())
		return
	}
	if ok, err := govalidator.ValidateStruct(form); !ok {
		utils.AbortWithErrorResponse(c, http.StatusBadRequest, errors.ValidationError, err.Error())
		return
	}
	// validators

	db.Model(&models.Interest{}).Where("email = ?", form.Email).Count(&count)
	if count > 0 {
		utils.AbortWithErrorResponse(c, http.StatusBadRequest, errors.DuplicateEmailError, "")
		return
	}
	// check unique

	interest := models.Interest{ClubhouseUserID: form.ClubhouseUserID, Email: form.Email}
	if err := interest.Create(db); err != nil {
		utils.AbortWithErrorResponse(c, http.StatusBadRequest, errors.DuplicateEmailError, err.Error())
	}
	response := forms.AddInterestResponse{ID: interest.ID, ClubhouseUserID: interest.ClubhouseUserID, Email: interest.Email}

	c.JSON(http.StatusCreated, response)
}

func (ctrl *Controller) DeleteInterest(c *gin.Context) {
	var form forms.DeleteInterest
	var count int64

	db := utils.GetDB().Instance
	if err := c.ShouldBindJSON(&form); err != nil {
		utils.AbortWithErrorResponse(c, http.StatusBadRequest, errors.ValidationError, err.Error())
		return
	}

	db.Where("email = ?", form.Email).Delete(&models.Interest{})

	c.Status(200)
}
