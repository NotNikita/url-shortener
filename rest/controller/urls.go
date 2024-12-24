package controller

import (
	"context"
	"time"

	"url-shortener/logger"
	"url-shortener/model"
	"url-shortener/rest/middleware"
	"url-shortener/service"
	"url-shortener/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/patrickmn/go-cache"
	"go.uber.org/zap"
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
// @Summary Shorten a URL
// @Description Shortens a given URL and returns the shortened URL
// @Tags URLs
// @Accept multipart/form-data
// @Produce json
// @Param long_url formData string true "Long URL"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /shorten [post]
func (cnt *UrlsController) ShortenUrl(c *fiber.Ctx) error {
	var newUrlData model.ViewUrlData
	form, err := c.MultipartForm()
	if err != nil {
		cnt.logger.Debugln("Error parsing provided data for shortenURL: %v", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse form: " + err.Error(),
		})
	}

	longUrl := form.Value["long_url"]
	if len(longUrl) == 0 || !utils.IsValidUrl(longUrl[0]) {
		cnt.logger.Debugln("Error when validating provided url to shorten")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to validate url",
		})
	}

	now := time.Now().Format(time.RFC3339)
	nowPlusMonth := time.Now().AddDate(0, 1, 0).Format(time.RFC3339)
	newUrlData = model.ViewUrlData{
		ShortUrl:    "",
		OriginalUrl: longUrl[0],
		ExpiresAt:   nowPlusMonth,
		CreatedAt:   now,
	}

	createdUrl, err := cnt.services.UrlsService.CreateUrl(c.Context(), &newUrlData)
	if err != nil {
		cnt.logger.Debugln("Failed to shorten URL", zap.String("long_url", longUrl[0]), zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to shorten your url: " + err.Error(),
		})
	}

	middleware.PutOriginalUrlInCache(cnt.cache, createdUrl.ShortUrl, longUrl[0])

	cnt.logger.Debugln("Created short url: ", zap.String("short_url", createdUrl.ShortUrl))
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"shortenURL": createdUrl.ShortUrl,
	})
}

// Interface of UrlsController: get original url for provided short and redirect
// @Summary Get original URL
// @Description Retrieves the original URL for a given short code and redirects
// @Tags URLs
// @Produce json
// @Param shortCode path string true "Short Code"
// @Success 302
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /{shortCode} [get]
func (cnt *UrlsController) GetOriginalUrl(c *fiber.Ctx) error {
	var newUrlData *model.ViewUrlData
	param := c.Params("shortCode")
	if len(param) < 1 {
		cnt.logger.Debugln("Error when retrieving shortCode from request url")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to retrieve shortCode in getURL function",
		})
	}

	newUrlData, err := cnt.services.UrlsService.GetUrl(c.Context(), param)
	if err != nil {
		cnt.logger.Debugln("Failed to retrieve shortcode from database")
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": "Error while retrieving or finding provided shortcode" + err.Error(),
			})
	}

	middleware.PutOriginalUrlInCache(cnt.cache, param, newUrlData.OriginalUrl)

	return c.Redirect(newUrlData.OriginalUrl, fiber.StatusFound)
}

// Interface of UrlsController: update existing short url with new origin value
// @Summary Update original URL for short URL
// @Description Update existing short url with new origin value
// @Tags URLs
// @Param shortCode path string true "Short Code"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /{shortCode} [put]
func (cnt *UrlsController) UpdateUrl(c *fiber.Ctx) error {
	var newUrlData model.ViewUrlData
	shortCodeParam := c.Params("shortCode")
	if len(shortCodeParam) < 1 {
		cnt.logger.Debugln("Error when retrieving shortCode from request url")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to retrieve shortCode in getURL function",
		})
	}

	form, err := c.MultipartForm()
	if err != nil {
		cnt.logger.Debugf("Error parsing provided data for shortenURL: %v", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse form: " + err.Error(),
		})
	}

	longUrl := form.Value["long_url"]
	if len(longUrl) == 0 || !utils.IsValidUrl(longUrl[0]) {
		cnt.logger.Debugln("Error when validating provided origin url to update")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to validate origin url",
		})
	}
	newUrlData = model.ViewUrlData{
		ShortUrl:    shortCodeParam,
		OriginalUrl: longUrl[0],
		ExpiresAt:   time.Now().AddDate(0, 1, 0).Format(time.RFC3339),
	}

	updatedUrlData, err := cnt.services.UrlsService.UpdateUrl(c.Context(), &newUrlData)
	if err != nil {
		cnt.logger.Debugf("Error while updating url of %v", newUrlData)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update url",
		})
	}

	middleware.PutOriginalUrlInCache(cnt.cache, shortCodeParam, updatedUrlData.OriginalUrl)

	return c.SendStatus(fiber.StatusNoContent)
}

// Interface of UrlsController: delete existing short url
// @Summary Delete short URL
// @Description Delete existing short url
// @Tags URLs
// @Param shortCode path string true "Short Code"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /{shortCode} [delete]
func (cnt *UrlsController) DeleteUrl(c *fiber.Ctx) error {
	shortCodeParam := c.Params("shortCode")
	if len(shortCodeParam) < 1 {
		cnt.logger.Debugln("Error when retrieving shortCode from request url")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to retrieve shortCode in getURL function",
		})
	}

	err := cnt.services.UrlsService.DeleteUrl(c.Context(), shortCodeParam)
	if err != nil {
		cnt.logger.Debugln("Failed to delete shortcode from database")
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": "Error while deleting provided shortcode" + err.Error(),
			})
	}

	middleware.EjectOriginalUrlFromCache(cnt.cache, shortCodeParam)

	return c.SendStatus(fiber.StatusNoContent)
}
