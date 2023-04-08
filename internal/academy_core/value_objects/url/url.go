package models

import (
	"fmt"
	"net/url"
)

type Url string

func (str *Url) IsUrl() bool {
	u, err := url.Parse(string(*str))
	return err == nil && u.Scheme != "" && u.Host != ""
}

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
