package models

import (
	"io"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func (p *Page) parse(body io.Reader) {
	doc := html.NewTokenizer(body)
	for {
		tokenType := doc.Next()
		switch tokenType {
		case html.ErrorToken:
			return
		case html.StartTagToken:
			p.parseToken(doc)
		}
	}
}

func (p *Page) parseToken(doc *html.Tokenizer) {
	token := doc.Token()

	switch token.DataAtom {
	case atom.Title:
		doc.Next()
		token = doc.Token()
		p.Title = token.Data
	case atom.Meta:
		p.fetchMeta(token)
	}
}

func (p *Page) fetchMeta(token html.Token) {
	name, content := fetchMetaAttributes(token)
	switch name {
	case "description":
		p.Description = content
	case "keywords":
		p.Keywords = content
	case "og:image":
		p.OGImage = content
	}
}

func fetchMetaAttributes(token html.Token) (name, content string) {
	for _, attr := range token.Attr {
		key := strings.ToLower(attr.Key)
		val := strings.ToLower(attr.Val)
		switch key {
		case "name":
			name = val
		case "property":
			name = val
		case "content":
			content = attr.Val
		}
	}
	return
}
