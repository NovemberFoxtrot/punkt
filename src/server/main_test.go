package server

import (
	_ "bufio"
	"io/ioutil"
	_ "net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestHeader(t *testing.T) {
	resp := httptest.NewRecorder()

	uri := "/"
	path := "/home/test"
	unlno := "997225821"

	param := make(url.Values)

	param["param1"] = []string{path}
	param["param2"] = []string{unlno}

	// req, err := http.NewRequest("GET", uri+param.Encode(), nil)

	req, err := http.NewRequest("GET", uri, nil)

	if err != nil {
		t.Fatal(err)
	}

	http.DefaultServeMux.ServeHTTP(resp, req)

	if p, err := ioutil.ReadAll(resp.Body); err != nil {
		t.Fail()
	} else {
		if strings.Contains(string(p), "Error") {
			t.Errorf("header response shouldn't return error: %s", p)
		} else if !strings.Contains(string(p), "Dude") {
			t.Errorf("header response doen't match:\n%s", p)
		}
	}
}
