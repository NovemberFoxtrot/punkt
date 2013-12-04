package server

import (
	"fmt"
	"io/ioutil"
	"log"
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

	fmt.Println("uri:", uri)

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
		} else if !strings.Contains(string(p), "404") {
			t.Errorf("header response doen't match:\n%s", p)
		}
	}
}

func TestPunkt(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(Index))
	defer ts.Close()

	res, err := http.Get(ts.URL)

	if err != nil {
		log.Fatal(err)
	}

	p, err := ioutil.ReadAll(res.Body)

	res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	if !strings.Contains(string(p), "punkt") {
		t.Errorf("header response does not match:\n%s", p)
	}
}

func TestDude(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		log.Fatal(err)
	}

	w := httptest.NewRecorder()

	Index(w, req)

	fmt.Printf("%d - %s", w.Code, w.Body.String())

	if w.Body.String() != "punkt" {
		t.Errorf("header response doen't match:\n")
	}
}
