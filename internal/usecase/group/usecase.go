package group

import (
	"github.com/evgeniy-dammer/clean-architecture/internal/usecase/adapters/storage"
)

type UseCase struct {
	options        Options
	adapterStorage storage.Group
}

type Options struct{}

func New(storage storage.Group, options Options) *UseCase {
	uc := &UseCase{
		adapterStorage: storage,
	}

	uc.SetOptions(options)

	return uc
}

func (uc *UseCase) SetOptions(options Options) {
	if uc.options != options {
		uc.options = options
	}
}
