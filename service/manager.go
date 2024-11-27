package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"

	"url-shortener/store"
)

type ServiceManager struct {
	UrlsService *UrlsWebService
}

// NewServiceManager creates new service manager
func NewServiceManager(ctx *fiber.Ctx, store *store.Store) (*ServiceManager, error) {
	if store == nil {
		return nil, errors.New("No store provided")
	}

	return &ServiceManager{
		UrlsService: NewUrlsWebService(ctx, store),
	}, nil
}
