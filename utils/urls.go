package utils

import (
	"net/url"
)

func IsValidUrl(urlToValidate string) bool {
	u, err := url.Parse(urlToValidate)
	return err == nil && u.Scheme != "" && u.Host != ""
}
