package group

import (
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/group"
)

func ProtoToGroupResponse(response *group.Group) *ResponseGroup {
	return &ResponseGroup{
		ID:         response.ID().String(),
		CreatedAt:  response.CreatedAt(),
		ModifiedAt: response.ModifiedAt(),
		Group: Group{
			ShortGroup: ShortGroup{
				Name:        response.Name().Value(),
				Description: response.Description().Value(),
			},
			ContactsAmount: response.ContactCount(),
		},
	}
}
