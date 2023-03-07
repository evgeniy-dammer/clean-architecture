package contact

import (
	domainContact "github.com/evgeniy-dammer/clean-architecture/internal/domain/contact"
)

func ToContactResponse(response *domainContact.Contact) *ResponseContact {
	return &ResponseContact{
		ID:         response.ID().String(),
		CreatedAt:  response.CreatedAt(),
		ModifiedAt: response.ModifiedAt(),
		ShortContact: ShortContact{
			PhoneNumber: response.PhoneNumber().String(),
			Email:       response.Email(),
			Gender:      response.Gender(),
			Age:         uint8(response.Age()),
			Name:        response.Name().String(),
			Surname:     response.Surname().String(),
			Patronymic:  response.Patronymic().String(),
		},
	}
}
