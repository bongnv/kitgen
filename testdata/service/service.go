package service

import (
	"context"

	"github.com/bongnv/kitgen/testdata/model"
)

// Service ...
type Service interface {
	PostProfile(ctx context.Context, p model.Profile) error
	GetProfile(ctx context.Context, id string) (model.Profile, error)
	PutProfile(ctx context.Context, id string, p model.Profile) error
	PatchProfile(ctx context.Context, id string, p model.Profile) error
	DeleteProfile(ctx context.Context, id string) error
	GetAddresses(ctx context.Context, profileID string) ([]model.Address, error)
	GetAddress(ctx context.Context, profileID string, addressID string) (model.Address, error)
	PostAddress(ctx context.Context, profileID string, a model.Address) (int, error)
	DeleteAddress(ctx context.Context, profileID string, addressID string) (string, error)
}
