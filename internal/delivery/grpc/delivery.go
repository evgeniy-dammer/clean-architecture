package grpc

import (
	contact "github.com/evgeniy-dammer/clean-architecture/internal/delivery/grpc/interface"
	"github.com/evgeniy-dammer/clean-architecture/internal/usecase"
)

type Delivery struct {
	contact.UnimplementedContactServiceServer
	options   Options
	ucContact usecase.Contact
	ucGroup   usecase.Group
}

type Options struct{}

func New(ucContact usecase.Contact, ucGroup usecase.Group, options Options) *Delivery {
	d := &Delivery{
		ucContact: ucContact,
		ucGroup:   ucGroup,
	}

	d.SetOptions(options)

	return d
}

func (d *Delivery) SetOptions(options Options) {
	if d.options != options {
		d.options = options
	}
}
