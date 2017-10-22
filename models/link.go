package models

import (
	"crawler/errors"
	"regexp"
)

var HOSTNAME_REGEXP = regexp.MustCompile(`^(?:https?:\/\/)((?:[\w-]+\.)+\w{2,11}?)(?:\/.*)?$`)

type Link struct {
	URL      string
	Hostname string
	Error    error
}

func NewLink(url string) *Link {
	hostname, err := parseURL(url)

	link := &Link{
		URL:      url,
		Hostname: hostname,
		Error:    err,
	}

	return link
}

func parseURL(url string) (string, error) {
	match := HOSTNAME_REGEXP.FindStringSubmatch(url)
	if len(match) != 2 {
		return "", errors.INVALID_URL
	}
	return match[1], nil
}
