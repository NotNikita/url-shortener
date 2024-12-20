package service

import (
	"errors"
	"log"
	"math/rand"

	"github.com/gofiber/fiber/v2"
	"github.com/zeebo/xxh3"
)

const (
	BASE_62_CHARS    = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	SHORT_URL_LENGTH = 6
)

type HashingService struct {
	ctx    *fiber.Ctx
	hasher *xxh3.Hasher
}

func NewHashingService(ctx *fiber.Ctx) *HashingService {
	return &HashingService{
		ctx:    ctx,
		hasher: xxh3.New(),
	}
}

// TODO: make this as standalone routine
// that will start running after server will be up
// and it will be reading tasks from queue and return shortUrls
func (service *HashingService) GenerateXXHash3BasedOnOriginURL(ctx *fiber.Ctx, originUrl string) (string, error) {
	// generating salt
	salt := generateSalt(10)
	saltedInput := salt + originUrl

	// hashing salt+url
	_, err := service.hasher.WriteString(saltedInput)
	if err != nil {
		return "", err
	}
	hashValue := service.hasher.Sum64()

	// encoding hash into base62
	base62Hash := base62Encode(hashValue)

	// get shorturl of desired length
	if len(base62Hash) > SHORT_URL_LENGTH {
		result := base62Hash[:SHORT_URL_LENGTH]
		return result, errors.New("Url <" + originUrl + "> was hashed into <" + result + ">")
	}

	log.Printf("Error when hashing following url: %s", originUrl)
	return "", errors.New("Failed to hash following url, see logs " + originUrl)
}

// TODO: generate salt once for 10-30 urls
// GenerateSalt generates a random salt string of the given length
func generateSalt(length int) (salt string) {
	for i := 0; i < length; i++ {
		salt += string(BASE_62_CHARS[rand.Intn(len(BASE_62_CHARS))])
	}
	return salt
}

// Base62Encode encodes a number into a Base62 string
func base62Encode(num uint64) string {
	if num == 0 {
		return string(BASE_62_CHARS[0])
	}
	var result string
	for num > 0 {
		result = string(BASE_62_CHARS[num%62]) + result
		num /= 62
	}
	return result
}
