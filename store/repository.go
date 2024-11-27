package store

import (
	"url-shortener/model"

	"github.com/gofiber/fiber/v2"
)

// Urls store is store for urls
type UrlsRepo interface {
	GetUrl(*fiber.Ctx, string) *model.UrlData
	CreateUrl(*fiber.Ctx, *model.UrlData) (*model.UrlData, error)
	// TODO:
	UpdateUrl(*fiber.Ctx, *model.UrlData) (*model.UrlData, error)
	DeleteUrl(*fiber.Ctx, string) error
}
