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

// @Summary Add to interest list (= waitlist)
// @Description add to interest list (= waitlist)
// @Accept json
// @Produce json
// @Param body body forms.AddInterestRequestDocument true "json body"
// @Success 201 {object} forms.AddInterestResponseDocument
// @Failure 400 {object} forms.ErrorResponse
// @Router /interest [post]
func (ctrler *Controller) AddInterest(c *gin.Context) {
	var form forms.AddInterestRequest
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

// @Summary Delete from interest list (= waitlist)
// @Description Soft delete from interest list (= waitlist)
// @Accept json
// @Produce json
// @Param body body forms.DeleteInterest true "json body"
// @Success 200
// @Failure 400 {object} forms.ErrorResponse
// @Failure 404
// @Router /interest [delete]
func (ctrl *Controller) DeleteInterest(c *gin.Context) {
	var form forms.DeleteInterest
	var count int64

	db := utils.GetDB().Instance
	if err := c.ShouldBindJSON(&form); err != nil {
		utils.AbortWithErrorResponse(c, http.StatusBadRequest, errors.ValidationError, err.Error())
		return
	}

	db.Model(&models.Interest{}).Where("email = ?", form.Email).Count(&count)
	if count == 0 {
		c.AbortWithStatus(404)
		return
	}
	// check unique

	db.Where("email = ?", form.Email).Delete(&models.Interest{})

	c.Status(200)
}
