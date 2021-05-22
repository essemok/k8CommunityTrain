package main

import (
	"github.com/takama/router"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	r := router.New()
	r.GET("/", home)

	ts := httptest.NewServer(r)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/")
	if err != nil {
		t.Fatal(err)
	}

	greeting, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		t.Fatal(err)
	}

	expectedGreeting := "Processing URL /..."
	testingGreeting := strings.Trim(string(greeting), "\n")

	if testingGreeting != expectedGreeting {
		t.Fatalf("Wrong greeting '%s', expected '%s'",
			testingGreeting, expectedGreeting,
		)
	}
}
