package models

import (
	"net/http"
	"strconv"
)

const (
	STATUS_INVALID_URL   = 0
	STATUS_NETWORK_ERROR = -1
)

type Page struct {
	StatusCode  int
	URL         string
	Title       string
	Description string
	Keywords    string
	OGImage     string
}

func GetPage(client *http.Client, link *Link) (page *Page) {
	page = &Page{URL: link.URL}

	err := page.validateLink(link)
	if err != nil {
		return
	}

	page.makeRequest(client)

	return
}

func (p *Page) ToSlice() []string {
	return []string{
		strconv.Itoa(p.StatusCode),
		p.URL,
		p.Title,
		p.Description,
		p.Keywords,
		p.OGImage,
	}
}

func (p *Page) validateLink(link *Link) error {
	if link.Error != nil {
		p.StatusCode = STATUS_INVALID_URL
	}
	return link.Error
}

func (p *Page) makeRequest(client *http.Client) {
	resp, err := client.Get(p.URL)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		p.StatusCode = STATUS_NETWORK_ERROR
		return
	}

	p.StatusCode = resp.StatusCode
	p.parse(resp.Body)
}
