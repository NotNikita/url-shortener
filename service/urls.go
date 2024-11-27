package service

import (
	"github.com/gofiber/fiber/v2"

	"url-shortener/model"
	"url-shortener/store"
)

type UrlsWebService struct {
	ctx   *fiber.Ctx
	store *store.Store
}

func NewUrlsWebService(ctx *fiber.Ctx, store *store.Store) *UrlsWebService {
	return &UrlsWebService{
		ctx,
		store,
	}
}

// GetUrl
func (service *UrlsWebService) GetUrl(ctx *fiber.Ctx, shortUrl string) *model.UrlData {
	// TODO: move logic from controller
	panic("Not implemented")
}

func (service *UrlsWebService) CreateUrl(ctx *fiber.Ctx, obj *model.UrlData) (*model.UrlData, error) {
	// TODO: move logic from controller
	panic("Not implemented")
}

// TODO:
func (service *UrlsWebService) UpdateUrl(ctx *fiber.Ctx, obj *model.UrlData) (*model.UrlData, error) {
	panic("Not implemented")
}

func (service *UrlsWebService) DeleteUrl(ctx *fiber.Ctx, shortUrl string) error {
	panic("Not implemented")
}
