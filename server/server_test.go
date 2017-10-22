package server

import (
	"crawler/logger"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"testing"

	"github.com/valyala/fasthttp"
)

const PORT = "8080"

func init() {
	logger.Mute()
}

func startServerOnPort(t *testing.T, h fasthttp.RequestHandler) io.Closer {
	ln, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", PORT))
	if err != nil {
		t.Fatalf("cannot start tcp server on port %s: %s", PORT, err)
	}
	go fasthttp.Serve(ln, h)
	return ln
}

func TestSaveHandler(t *testing.T) {
	defer startServerOnPort(t, NewRouter()).Close()
	url := "http://localhost:" + PORT + "/tasks/"
	contentType := "text/plain"

	for _, tc := range testCasesSaveHandler {
		resp, err := http.Post(url, contentType, strings.NewReader(tc.ReqBody))
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != tc.Status {
			t.Errorf("HTTP status code: '%d', expected: '%d'\n", resp.StatusCode, tc.Status)
		}
		actual, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(actual) != tc.RespBody {
			t.Errorf("Response body: '%s', expected: '%s'\n", actual, tc.RespBody)
		}
	}
}

func TestViewHandler(t *testing.T) {
	defer startServerOnPort(t, NewRouter()).Close()
	url := "http://localhost:" + PORT + "/tasks/"

	for _, tc := range testCasesViewHandler {
		tc.Setup()

		resp, err := http.Get(url + tc.Id)
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != tc.Status {
			t.Errorf("HTTP status code: '%d', expected: '%d'\n", resp.StatusCode, tc.Status)
		}
		actual, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(actual) != tc.RespBody {
			t.Errorf("Response body: '%s', expected: '%s'\n", actual, tc.RespBody)
		}
	}
}

func TestDeleteHandler(t *testing.T) {
	defer startServerOnPort(t, NewRouter()).Close()
	url := "http://localhost:" + PORT + "/tasks/"

	for _, tc := range testCasesDeleteHandler {
		tc.Setup()

		resp, err := http.Get(url + tc.Id + "?delete=1")
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != tc.Status {
			t.Errorf("HTTP status code: '%d', expected: '%d'\n", resp.StatusCode, tc.Status)
		}
		actual, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(actual) != tc.RespBody {
			t.Errorf("Response body: '%s', expected: '%s'\n", actual, tc.RespBody)
		}
	}
}
