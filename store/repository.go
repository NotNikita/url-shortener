package store

import (
	"url-shortener/model"

	"github.com/gofiber/fiber/v2"
)

// Urls store is store for urls
type UrlsRepo interface {
	GetUrl(*fiber.Ctx, string) *model.DBUrlData
	CreateUrl(*fiber.Ctx, *model.DBUrlData) (*model.DBUrlData, error)
	// TODO:
	UpdateUrl(*fiber.Ctx, *model.DBUrlData) (*model.DBUrlData, error)
	DeleteUrl(*fiber.Ctx, string) error
}
