package http

import (
	"net/http"
	"time"

	jsonGroup "github.com/evgeniy-dammer/clean-architecture/internal/delivery/http/group"
	domainGroup "github.com/evgeniy-dammer/clean-architecture/internal/domain/group"
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/group/description"
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/group/name"
	"github.com/evgeniy-dammer/clean-architecture/pkg/tools/converter"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/pagination"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/query"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/queryparameter"
	"github.com/gin-gonic/gin"
)

var mappingSortsGroup = query.SortsOptions{
	"id":           {},
	"name":         {},
	"description":  {},
	"contactCount": {},
}

// CreateGroup
// @Summary Create group method.
// @Description Create group method.
// @Tags 	groups
// @Accept  json
// @Produce json
// @Param   group 		body 		jsonGroup.ShortGroup 	true	"Group data"
// @Success 200			{object}  	jsonGroup.ResponseGroup	true
// @Failure 400 		{object}    ErrorResponse
// @Failure 403	 		"Forbidden"
// @Failure 404 	    {object} 	ErrorResponse					"404 Not Found"
// @Router /groups/ [post].
func (d *Delivery) CreateGroup(ctx *gin.Context) {
	group := &jsonGroup.ShortGroup{}
	if err := ctx.ShouldBindJSON(&group); err != nil {
		SetError(ctx, http.StatusBadRequest, err)

		return
	}

	groupName, err := name.New(group.Name)
	if err != nil {
		SetError(ctx, http.StatusBadRequest, err)

		return
	}

	groupDescription, err := description.New(group.Description)
	if err != nil {
		SetError(ctx, http.StatusBadRequest, err)

		return
	}

	newGroup, err := d.ucGroup.CreateGroup(domainGroup.New(groupName, groupDescription))
	if err != nil {
		SetError(ctx, http.StatusInternalServerError, err)

		return
	}

	ctx.JSON(http.StatusOK, jsonGroup.ResponseGroup{
		ID:         newGroup.ID().String(),
		CreatedAt:  newGroup.CreatedAt(),
		ModifiedAt: newGroup.ModifiedAt(),
		Group: jsonGroup.Group{
			ShortGroup: jsonGroup.ShortGroup{
				Name:        newGroup.Name().Value(),
				Description: newGroup.Description().Value(),
			},
			ContactsAmount: newGroup.ContactCount(),
		},
	})
}

// UpdateGroup
// @Summary Update group method.
// @Description Update group method.
// @Tags 	groups
// @Accept  json
// @Produce json
// @Param   id 			path 		string 					true	"Group ID"
// @Param   group 		body 		jsonGroup.ShortGroup 	true	"Group data"
// @Success 200			{object}  	jsonGroup.ResponseGroup
// @Failure 400 		{object}    ErrorResponse
// @Failure 403	 		"Forbidden"
// @Failure 404 	    {object} 	ErrorResponse					"404 Not Found"
// @Router /groups/{id} [put].
func (d *Delivery) UpdateGroup(ctx *gin.Context) {
	var groupID jsonGroup.ID
	if err := ctx.ShouldBindUri(&groupID); err != nil {
		SetError(ctx, http.StatusBadRequest, err)

		return
	}

	group := jsonGroup.ShortGroup{}
	if err := ctx.ShouldBindJSON(&group); err != nil {
		SetError(ctx, http.StatusBadRequest, err)

		return
	}

	groupName, err := name.New(group.Name)
	if err != nil {
		SetError(ctx, http.StatusBadRequest, err)

		return
	}

	groupDescription, err := description.New(group.Description)
	if err != nil {
		SetError(ctx, http.StatusBadRequest, err)

		return
	}

	grp, err := domainGroup.NewWithID(
		converter.StringToUUID(groupID.Value), time.Now().UTC(), time.Now().UTC(), groupName, groupDescription, 0,
	)
	if err != nil {
		SetError(ctx, http.StatusInternalServerError, err)

		return
	}

	response, err := d.ucGroup.UpdateGroup(grp)
	if err != nil {
		SetError(ctx, http.StatusInternalServerError, err)

		return
	}

	ctx.JSON(http.StatusOK, jsonGroup.ProtoToGroupResponse(response))
}

// DeleteGroup
// @Summary Delete group method.
// @Description Delete group method.
// @Tags 	groups
// @Accept  json
// @Produce json
// @Param   id 			path 		string 			true 	"Group ID"
// @Success 200			{object}  	string
// @Failure 400 		{object}    ErrorResponse
// @Failure 403	 		"Forbidden"
// @Failure 404 	    {object} 	ErrorResponse			"404 Not Found"
// @Router /groups/{id} [delete].
func (d *Delivery) DeleteGroup(ctx *gin.Context) {
	var groupID jsonGroup.ID

	if err := ctx.ShouldBindUri(&groupID); err != nil {
		SetError(ctx, http.StatusBadRequest, err)

		return
	}

	if err := d.ucGroup.DeleteGroup(converter.StringToUUID(groupID.Value)); err != nil {
		SetError(ctx, http.StatusInternalServerError, err)

		return
	}

	ctx.Status(http.StatusOK)
}

// ListGroup
// @Summary Get group list method.
// @Description Get group list method.
// @Tags 	groups
// @Accept  json
// @Produce json
// @Param 	limit 		query 		int 					false "Item count" default(10) min(0) max(100)
// @Param 	offset 		query 		int 					false "Item offset" default(0) min(0)
// @Param 	sort 		query 		string 					false "Field sort" default(name)
// @Success 200			{object}  	jsonGroup.ListGroup
// @Failure 400 		{object}    ErrorResponse
// @Failure 403	 		"Forbidden"
// @Failure 404 	    {object} 	ErrorResponse			"404 Not Found"
// @Router /groups/ [get].
func (d *Delivery) ListGroup(ctx *gin.Context) {
	params, err := query.ParseQuery(ctx, query.Options{Sorts: mappingSortsGroup})
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

	groups, err := d.ucGroup.GetListGroup(param)
	if err != nil {
		SetError(ctx, http.StatusInternalServerError, err)

		return
	}

	count, err := d.ucContact.CountContact(param)
	if err != nil {
		SetError(ctx, http.StatusInternalServerError, err)

		return
	}

	list := make([]*jsonGroup.ResponseGroup, len(groups))

	for i, elem := range groups {
		list[i] = jsonGroup.ProtoToGroupResponse(elem)
	}

	ctx.JSON(http.StatusOK, jsonGroup.ListGroup{
		Total:  count,
		Limit:  params.Limit,
		Offset: params.Offset,
		List:   list,
	})
}

// ReadGroupByID
// @Summary Get group by ID method.
// @Description Get group by ID method.
// @Tags 	groups
// @Accept  json
// @Produce json
// @Param   id 			path 		string 					true 	"Group ID"
// @Success 200			{object}  	jsonGroup.ResponseGroup
// @Failure 400 		{object}    ErrorResponse
// @Failure 403	 		"Forbidden"
// @Failure 404 	    {object} 	ErrorResponse					"404 Not Found"
// @Router /groups/{id} [get].
func (d *Delivery) ReadGroupByID(ctx *gin.Context) {
	var groupID jsonGroup.ID
	if err := ctx.ShouldBindUri(&groupID); err != nil {
		SetError(ctx, http.StatusBadRequest, err)

		return
	}

	response, err := d.ucGroup.GetGroupByID(converter.StringToUUID(groupID.Value))
	if err != nil {
		SetError(ctx, http.StatusInternalServerError, err)

		return
	}

	ctx.JSON(http.StatusOK, jsonGroup.ProtoToGroupResponse(response))
}
