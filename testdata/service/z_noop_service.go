package service

import (
	"context"
	"errors"

	"github.com/bongnv/kitgen/testdata/model"
)

// NoopService ...
type NoopService struct{}

// PostProfile ...
func (_ NoopService) PostProfile(_ context.Context, _ model.Profile) error {
	return errors.New("method PostProfile is unimplemented")
}

// GetProfile ...
func (_ NoopService) GetProfile(_ context.Context, _ string) (model.Profile, error) {
	return model.Profile{}, errors.New("method GetProfile is unimplemented")
}

// PutProfile ...
func (_ NoopService) PutProfile(_ context.Context, _ string, _ model.Profile) error {
	return errors.New("method PutProfile is unimplemented")
}

// PatchProfile ...
func (_ NoopService) PatchProfile(_ context.Context, _ string, _ model.Profile) error {
	return errors.New("method PatchProfile is unimplemented")
}

// DeleteProfile ...
func (_ NoopService) DeleteProfile(_ context.Context, _ string) error {
	return errors.New("method DeleteProfile is unimplemented")
}

// GetAddresses ...
func (_ NoopService) GetAddresses(_ context.Context, _ string) ([]model.Address, error) {
	return nil, errors.New("method GetAddresses is unimplemented")
}

// GetAddress ...
func (_ NoopService) GetAddress(_ context.Context, _ string, _ string) (model.Address, error) {
	return model.Address{}, errors.New("method GetAddress is unimplemented")
}

// PostAddress ...
func (_ NoopService) PostAddress(_ context.Context, _ string, _ model.Address) (int, error) {
	return 0, errors.New("method PostAddress is unimplemented")
}

// DeleteAddress ...
func (_ NoopService) DeleteAddress(_ context.Context, _ string, _ string) (string, error) {
	return "", errors.New("method DeleteAddress is unimplemented")
}
