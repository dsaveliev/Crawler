package models

import (
	"crawler/logger"
	"fmt"
	"net/http"
	"testing"

	"gopkg.in/h2non/gock.v1"
)

var client = http.DefaultClient

func init() {
	logger.Mute()
}

func TestGetPage(t *testing.T) {
	defer gock.Off()
	gock.New("http://ya.ru").Get("/").Reply(http.StatusOK).BodyString("")
	gock.New("http://ya.ru").Get("/error").ReplyError(fmt.Errorf("Error"))

	for _, tc := range testCasesGetPage {
		page := GetPage(client, tc.Link)
		if page.URL != tc.URL {
			t.Fatalf("page.URL = %#v. Want: %#v",
				page.URL, tc.URL)
		}
		if page.StatusCode != tc.StatusCode {
			t.Fatalf("page.URL = %#v, page.StatusCode = %d. Want: %d",
				page.URL, page.StatusCode, tc.StatusCode)
		}
	}
}
