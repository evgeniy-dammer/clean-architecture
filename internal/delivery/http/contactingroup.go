package http

import (
	"fmt"
	"net/http"

	jsonContact "github.com/evgeniy-dammer/clean-architecture/internal/delivery/http/contact"
	jsonGroup "github.com/evgeniy-dammer/clean-architecture/internal/delivery/http/group"
	domainContact "github.com/evgeniy-dammer/clean-architecture/internal/domain/contact"
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/contact/age"
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/contact/name"
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/contact/patronymic"
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/contact/surname"
	"github.com/evgeniy-dammer/clean-architecture/pkg/tools/converter"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/phone"
	"github.com/gin-gonic/gin"
)

// CreateContactIntoGroup
// @Summary Create contact and add it into group.
// @Description Create contact and add it into group.
// @Security Cookies
// @Tags 	groups
// @Accept  json
// @Produce json
// @Param   id 			path 		string 						true	"Group ID"
// @Param   contact 	body 		jsonContact.ShortContact 	true	"Contact data"
// @Success 200
// @Failure 400 		{object}    ErrorResponse
// @Failure 403	 		"Forbidden"
// @Failure 404 	    {object} 	ErrorResponse						"404 Not Found"
// @Router /groups/{id}/contacts/ [post].
func (d *Delivery) CreateContactIntoGroup(ctx *gin.Context) {
	var groupID jsonGroup.ID
	if err := ctx.ShouldBindUri(&groupID); err != nil {
		SetError(ctx, http.StatusBadRequest, err)

		return
	}

	contact := jsonContact.ShortContact{}
	if err := ctx.ShouldBindJSON(&contact); err != nil {
		SetError(ctx, http.StatusBadRequest, fmt.Errorf("payload is not correct, Error: %w", err))

		return
	}

	contactAge, err := age.New(uint64(contact.Age))
	if err != nil {
		SetError(ctx, http.StatusBadRequest, err)

		return
	}

	contactName, err := name.New(contact.Name)
	if err != nil {
		SetError(ctx, http.StatusBadRequest, err)

		return
	}

	contactSurname, err := surname.New(contact.Surname)
	if err != nil {
		SetError(ctx, http.StatusBadRequest, err)

		return
	}

	contactPatronymic, err := patronymic.New(contact.Patronymic)
	if err != nil {
		SetError(ctx, http.StatusBadRequest, err)

		return
	}

	dContact, err := domainContact.New(
		*phone.New(contact.PhoneNumber),
		contact.Email,
		*contactName,
		*contactSurname,
		*contactPatronymic,
		*contactAge,
		contact.Gender,
	)
	if err != nil {
		SetError(ctx, http.StatusBadRequest, err)

		return
	}

	contacts, err := d.ucGroup.CreateContactIntoGroup(converter.StringToUUID(groupID.Value), dContact)
	if err != nil {
		SetError(ctx, http.StatusInternalServerError, err)

		return
	}

	list := []*jsonContact.ResponseContact{}

	for _, value := range contacts {
		list = append(list, jsonContact.ToContactResponse(value))
	}

	ctx.JSON(http.StatusOK, list)
}

// AddContactToGroup
// @Summary Add contact into group.
// @Description Add contact into group.
// @Tags 	groups
// @Accept  json
// @Produce json
// @Param   id 			path 		string 			true 	"Group ID"
// @Param   contactId 	path 		string 			true 	"Contact ID"
// @Success 200
// @Failure 400 		{object}    ErrorResponse
// @Failure 403	 		"Forbidden"
// @Failure 404 	    {object} 	ErrorResponse				"404 Not Found"
// @Router /groups/{id}/contacts/{contactId} [post].
func (d *Delivery) AddContactToGroup(ctx *gin.Context) {
	var groupID jsonGroup.ID
	if err := ctx.ShouldBindUri(&groupID); err != nil {
		SetError(ctx, http.StatusBadRequest, err)

		return
	}

	var contactID jsonGroup.ContactID
	if err := ctx.ShouldBindUri(&contactID); err != nil {
		SetError(ctx, http.StatusBadRequest, err)

		return
	}

	if err := d.ucGroup.AddContactToGroup(converter.StringToUUID(groupID.Value), converter.StringToUUID(contactID.Value)); err != nil { //nolint:lll
		SetError(ctx, http.StatusInternalServerError, err)

		return
	}

	ctx.Status(http.StatusOK)
}

// DeleteContactFromGroup
// @Summary Remove contact from group.
// @Description Remove contact from group.
// @Tags 	groups
// @Accept  json
// @Produce json
// @Param   id 			path 		string 			true 	"Group ID"
// @Param   contactId 	path 		string 			true 	"Contact ID"
// @Success 200
// @Failure 400 		{object}    ErrorResponse
// @Failure 403	 		"Forbidden"
// @Failure 404 	    {object} 	ErrorResponse			"404 Not Found"
// @Router /groups/{id}/contacts/{contactId} [delete].
func (d *Delivery) DeleteContactFromGroup(ctx *gin.Context) {
	var groupID jsonGroup.ID
	if err := ctx.ShouldBindUri(&groupID); err != nil {
		SetError(ctx, http.StatusBadRequest, err)

		return
	}

	var contactID jsonGroup.ContactID
	if err := ctx.ShouldBindUri(&contactID); err != nil {
		SetError(ctx, http.StatusBadRequest, err)

		return
	}

	if err := d.ucGroup.DeleteContactFromGroup(converter.StringToUUID(groupID.Value), converter.StringToUUID(contactID.Value)); err != nil { //nolint:lll
		SetError(ctx, http.StatusInternalServerError, err)

		return
	}

	ctx.Status(http.StatusOK)
}
