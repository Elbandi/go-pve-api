package pve

import (
	"net/http/httptest"
	"net/url"
	"net/http"
	//"github.com/stretchr/testify/assert"
	"fmt"
	"testing"
	"reflect"
)

// Use the client to make requests on the server.
// Register handlers on mux to handle requests.
// Caller must close test server.
func testServer(t *testing.T, withTicketMux bool) (*http.Client, *http.ServeMux, *httptest.Server) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	transport := &RewriteTransport{&http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}}

	if withTicketMux {
		mux.HandleFunc("/api2/json/access/ticket", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, "{\"data\":{\"username\":\"root@pam\",\"CSRFPreventionToken\":\"5876ABF4:token\",\"ticket\":\"PVE:root@pam:5876ABF4::ticket\",\"cap\":{}}}")
		})
	}
	client := &http.Client{Transport: transport}
	return client, mux, server
}

type RewriteTransport struct {
	Transport http.RoundTripper
}

func (t *RewriteTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.URL.Scheme = "http"
	if t.Transport == nil {
		return http.DefaultTransport.RoundTrip(req)
	}
	return t.Transport.RoundTrip(req)
}

func assertMethod(t *testing.T, expectedMethod string, req *http.Request) {
	if actualMethod := req.Method; actualMethod != expectedMethod {
		t.Errorf("expected method %s, got %s", expectedMethod, actualMethod)
	}
}

// assertQuery tests that the Request has the expected url query key/val pairs
func assertQuery(t *testing.T, expected map[string]string, req *http.Request) {
	queryValues := req.URL.Query() // net/url Values is a map[string][]string
	expectedValues := url.Values{}
	for key, value := range expected {
		expectedValues.Add(key, value)
	}
	if !reflect.DeepEqual(expectedValues, queryValues) {
		t.Errorf("expected parameters %v, got %v", expected, req.URL.RawQuery)
	}
}

// assertPostForm tests that the Request has the expected key values pairs url
// encoded in its Body
func assertPostForm(t *testing.T, expected map[string]string, req *http.Request) {
	req.ParseForm() // parses request Body to put url.Values in r.Form/r.PostForm
	expectedValues := url.Values{}
	for key, value := range expected {
		expectedValues.Add(key, value)
	}
	if !reflect.DeepEqual(expectedValues, req.PostForm) {
		t.Errorf("expected parameters %v, got %v", expected, req.PostForm)
	}
}