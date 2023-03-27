package models

import "net/url"

type Url string

func (str *Url) IsUrl() bool {
	u, err := url.Parse(string(*str))
	return err == nil && u.Scheme != "" && u.Host != ""
}
