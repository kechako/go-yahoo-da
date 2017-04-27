package da

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

const testAppID = "test_application_id"

func TestNewClient(t *testing.T) {
	client := NewClient(testAppID)
	if client == nil {
		t.Fatal("shoud not be return nil.")
	}

	if client.AppID != testAppID {
		t.Errorf("want %s but %s", testAppID, client.AppID)
	}

	if client.HTTPClient != http.DefaultClient {
		t.Error("HTTPClient is not http.DefaultClient.")
	}
}

func Test_httpClient(t *testing.T) {
	client := NewClient("dummy")
	if client.httpClient() == nil {
		t.Error("Shoud not be return nil.")
	}

	httpClient := &http.Client{}
	client.HTTPClient = httpClient
	if client.httpClient() == nil {
		t.Error("Shoud not be return nil.")
	}
	if client.httpClient() != httpClient {
		t.Error("httpClient() returns invalid http.Client.")
	}

	client.HTTPClient = nil
	if client.httpClient() == nil {
		t.Error("Shoud not be return nil.")
	}
	if client.httpClient() != http.DefaultClient {
		t.Error("HTTPClient is not http.DefaultClient.")
	}
}

const (
	testText   = "うちの庭には二羽鶏がいます。"
	testReqURL = `https://jlp.yahooapis.jp/DAService/V1/parse?appid=test_application_id&sentence=%E3%81%86%E3%81%A1%E3%81%AE%E5%BA%AD%E3%81%AB%E3%81%AF%E4%BA%8C%E7%BE%BD%E9%B6%8F%E3%81%8C%E3%81%84%E3%81%BE%E3%81%99%E3%80%82`
)

func Test_requestURL(t *testing.T) {
	client := NewClient(testAppID)

	reqURL := client.requestURL(testText)
	if reqURL != testReqURL {
		t.Errorf("want %s but %s", testReqURL, reqURL)
	}
}

func Test_Parse(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.FormValue("appid") != testAppID || r.FormValue("sentence") != testText {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		fmt.Fprint(w, testXML)
	}))

	APIEndpoint = ts.URL
	client := NewClient(testAppID)
	res, err := client.Parse(context.Background(), testText)
	if err != nil {
		t.Fatalf("shoud not be fail: %v", err)
	}

	if !reflect.DeepEqual(res, testResultSet) {
		t.Errorf("want %#v\nbut %#v", testResultSet, res)
	}
}
