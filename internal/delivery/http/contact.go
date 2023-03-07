package http

import (
	"fmt"
	"net/http"
	"time"

	jsonContact "github.com/evgeniy-dammer/clean-architecture/internal/delivery/http/contact"
	domainContact "github.com/evgeniy-dammer/clean-architecture/internal/domain/contact"
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/contact/age"
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/contact/name"
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/contact/patronymic"
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/contact/surname"
	"github.com/evgeniy-dammer/clean-architecture/pkg/tools/converter"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/pagination"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/phone"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/query"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/queryparameter"
	"github.com/gin-gonic/gin"
)

var mappingSortsContact = query.SortsOptions{
	"name":        {},
	"surname":     {},
	"patronymic":  {},
	"phoneNumber": {},
	"email":       {},
	"gender":      {},
	"age":         {},
}

// CreateContact
// @Summary Create contact method.
// @Description Create contact method.
// @Tags contacts
// @Accept  json
// @Produce json
// @Param   contact 	body 		jsonContact.ShortContact 		    true  "Contact data"
// @Success 201			{object}  	jsonContact.ResponseContact 		true  "Contact structure"
// @Success 200
// @Failure 400 		{object}    ErrorResponse
// @Failure 403	 		"Forbidden"
// @Failure 404 	    {object} 	ErrorResponse			"404 Not Found"
// @Router /contacts/ [post].
func (d *Delivery) CreateContact(ctx *gin.Context) {
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
		*phone.New(contact.PhoneNumber), contact.Email, *contactName, *contactSurname, *contactPatronymic, *contactAge, contact.Gender, //nolint:lll
	)
	if err != nil {
		SetError(ctx, http.StatusBadRequest, err)

		return
	}

	response, err := d.ucContact.CreateContact(dContact)
	if err != nil {
		SetError(ctx, http.StatusInternalServerError, err)

		return
	}

	if len(response) > 0 {
		ctx.JSON(http.StatusCreated, jsonContact.ToContactResponse(response[0]))
	} else {
		ctx.Status(http.StatusOK)
	}
}

// UpdateContact
// @Summary Update contact method.
// @Description Update contact method.
// @Tags contacts
// @Accept  json
// @Produce json
// @Param   id 			path 		string 						true  "Contact ID"
// @Param   contact 	body 		jsonContact.ShortContact	true  "Contact data"
// @Success 200			{object}  	jsonContact.ResponseContact true  "Contact structure"
// @Failure 400 		{object}    ErrorResponse
// @Failure 403	 		"Forbidden"
// @Failure 404 	    {object} 	ErrorResponse			  		  "404 Not Found"
// @Router /contacts/{id} [put].
func (d *Delivery) UpdateContact(ctx *gin.Context) {
	var contactID jsonContact.ID
	if err := ctx.ShouldBindUri(&contactID); err != nil {
		SetError(ctx, http.StatusBadRequest, err)

		return
	}

	contact := jsonContact.ShortContact{}
	if err := ctx.ShouldBindJSON(&contact); err != nil {
		SetError(ctx, http.StatusInternalServerError, err)

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

	dContact, _ := domainContact.NewWithID(
		converter.StringToUUID(contactID.Value), time.Now().UTC(), time.Now().UTC(), *phone.New(contact.PhoneNumber),
		contact.Email, *contactName, *contactSurname, *contactPatronymic, *contactAge, contact.Gender,
	)

	response, err := d.ucContact.UpdateContact(dContact)
	if err != nil {
		SetError(ctx, http.StatusInternalServerError, err)

		return
	}

	ctx.JSON(http.StatusOK, jsonContact.ToContactResponse(response))
}

// DeleteContact
// @Summary Delete contact method.
// @Description Delete contact method.
// @Tags contacts
// @Accept  json
// @Produce json
// @Param   id 			path 		string 			true 	"Contact ID"
// @Failure 400 		{object}    ErrorResponse
// @Failure 403	 		"Forbidden"
// @Failure 404 	    {object} 	ErrorResponse			"404 Not Found"
// @Router /contacts/{id} [delete].
func (d *Delivery) DeleteContact(ctx *gin.Context) {
	var contactID jsonContact.ID

	if err := ctx.ShouldBindUri(&contactID); err != nil {
		SetError(ctx, http.StatusBadRequest, err)

		return
	}

	if err := d.ucContact.DeleteContact(converter.StringToUUID(contactID.Value)); err != nil {
		SetError(ctx, http.StatusInternalServerError, err)

		return
	}

	ctx.Status(http.StatusOK)
}

// ListContact
// @Summary Get contact list method.
// @Description Get contact list method.
// @Tags contacts
// @Accept  json
// @Produce json
// @Param 	limit 		query 		int 					false "Item count" default(10) min(0) max(100)
// @Param 	offset 		query 		int 					false "item offset" default(0) min(0)
// @Param 	sort 		query 		string 					false "Field sort" default(name)
// @Success 200			{object}  	jsonContact.ListContact true  "Contact list"
// @Failure 400 		{object}    ErrorResponse
// @Failure 403	 		"Forbidden"
// @Failure 404 	    {object} 	ErrorResponse			"404 Not Found"
// @Router /contacts/ [get].
func (d *Delivery) ListContact(ctx *gin.Context) {
	params, err := query.ParseQuery(ctx, query.Options{
		Sorts: mappingSortsContact,
	})
	if err != nil {
		SetError(ctx, http.StatusBadRequest, err)

		return
	}

	param := queryparameter.QueryParameter{
		Sorts: params.Sorts,
		Pagination: pagination.Pagination{
			Limit:  params.Limit,
			Offset: params.Offset,
		},
	}

	contacts, err := d.ucContact.GetListContact(param)
	if err != nil {
		SetError(ctx, http.StatusInternalServerError, err)

		return
	}

	count, err := d.ucContact.CountContact(param)
	if err != nil {
		SetError(ctx, http.StatusInternalServerError, err)

		return
	}

	result := jsonContact.ListContact{
		Total:  count,
		Limit:  params.Limit,
		Offset: params.Offset,
		List:   []*jsonContact.ResponseContact{},
	}

	for _, value := range contacts {
		result.List = append(result.List, jsonContact.ToContactResponse(value))
	}

	ctx.JSON(http.StatusOK, result)
}

// ReadContactByID
// @Summary Get contact by ID method.
// @Description Get contact by ID method.
// @Tags contacts
// @Accept  json
// @Produce json
// @Param   id 			path 		string 						true "Contact ID"
// @Success 200			{object}  	jsonContact.ResponseContact true "Contact structure"
// @Failure 400 		{object}    ErrorResponse
// @Failure 403	 		"Forbidden"
// @Failure 404 	    {object} 	ErrorResponse					  "404 Not Found"
// @Router /contacts/{id} [get].
func (d *Delivery) ReadContactByID(ctx *gin.Context) {
	var contactID jsonContact.ID

	if err := ctx.ShouldBindUri(&contactID); err != nil {
		SetError(ctx, http.StatusBadRequest, err)

		return
	}

	response, err := d.ucContact.GetContactByID(converter.StringToUUID(contactID.Value))
	if err != nil {
		SetError(ctx, http.StatusInternalServerError, err)

		return
	}

	ctx.JSON(http.StatusOK, jsonContact.ToContactResponse(response))
}
