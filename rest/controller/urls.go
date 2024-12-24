package controller

import (
	"context"
	"fmt"
	"log"
	"time"
	"url-shortener/logger"
	"url-shortener/model"
	"url-shortener/rest/middleware"
	"url-shortener/service"
	"url-shortener/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/patrickmn/go-cache"
)

// Urls Controller
type UrlsController struct {
	ctx      context.Context
	services *service.ServiceManager
	logger   *logger.Logger
	cache    *cache.Cache
}

// Creates a new Urls Controller
func NewUrlsController(ctx context.Context, services *service.ServiceManager, logger *logger.Logger, cache *cache.Cache) *UrlsController {
	return &UrlsController{
		ctx:      ctx,
		services: services,
		logger:   logger,
		cache:    cache,
	}
}

// Interface of UrlsController: shorten a given url
func (cnt *UrlsController) ShortenUrl(c *fiber.Ctx) error {
	var newUrlData model.ViewUrlData
	form, err := c.MultipartForm()
	if err != nil {
		log.Fatalf("Error parsing provided data for shortenURL: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse form: " + err.Error(),
		})
	}

	long_url := form.Value["long_url"]
	if len(long_url) == 0 || !utils.IsValidUrl(long_url[0]) {
		log.Fatalf("Error when validating provided url to shorten")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to validate url: " + err.Error(),
		})
	}

	now := time.Now().Format(time.RFC3339)
	nowPlusMonth := time.Now().AddDate(0, 1, 0).Format(time.RFC3339)
	newUrlData = model.ViewUrlData{
		ShortUrl:    "",
		OriginalUrl: long_url[0],
		ExpiresAt:   nowPlusMonth,
		CreatedAt:   now,
	}
	fmt.Println("urlData", newUrlData)
	createdUrl, err := cnt.services.UrlsService.CreateUrl(c.Context(), &newUrlData)
	if err != nil {
		log.Printf("500: Failed to shorten url %s", long_url[0])
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to shorten your url: " + err.Error(),
		})
	}

	middleware.PutOriginalUrlInCache(cnt.cache, createdUrl.ShortUrl, long_url[0])

	cnt.logger.Debugln("Created short url: ", createdUrl.ShortUrl)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"shortenURL": createdUrl.ShortUrl,
	})
}

// Interface of UrlsController: get original url for provided short and redirect
func (cnt *UrlsController) GetOriginalUrl(c *fiber.Ctx) error {
	var newUrlData *model.ViewUrlData
	param := c.Params("shortCode")
	if len(param) < 1 { // TODO: HASH_LENGTH {
		log.Printf("Error when retrieving shortCode from request url")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to retrieve shortCode in getURL function",
		})
	}

	newUrlData, err := cnt.services.UrlsService.GetUrl(c.Context(), param)
	if err != nil {
		log.Printf("Failed to retrieve shortcode from database")
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": "Error while retrieving or finding provided shortcode" + err.Error(),
			})
	}

	middleware.PutOriginalUrlInCache(cnt.cache, param, newUrlData.OriginalUrl)

	return c.Redirect(newUrlData.OriginalUrl, fiber.StatusFound)
}
