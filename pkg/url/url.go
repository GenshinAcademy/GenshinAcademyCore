package url

import (
	"fmt"
	"net/url"
)

type Url string

func CreateUrl(urlParts ...string) (Url, error) {
	if len(urlParts) == 0 {
		return "", fmt.Errorf("empty url")
	}

	baseUrl, err := url.Parse(urlParts[0])
	if err != nil {
		return "", err
	}

	for _, path := range urlParts[1:] {
		urlPath, err := url.Parse(path)
		if err != nil {
			return "", err
		}

		baseUrl = baseUrl.ResolveReference(urlPath)
	}

	return Url(baseUrl.String()), nil
}

func IsUrl(input string) bool {
	u, err := url.Parse(input)
	return err == nil && u.Scheme != ""
}
