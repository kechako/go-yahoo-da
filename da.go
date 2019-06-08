package da

import (
	"context"
	"net/http"
	"net/url"

	"golang.org/x/xerrors"
)

// APIEndpoint is a endpoint url of an API for japanese dependency parsing.
var APIEndpoint = "https://jlp.yahooapis.jp/DAService/V1/parse"

// A Client represents a client of an API for japanese dependency parsing.
type Client struct {
	AppID      string
	HTTPClient *http.Client
}

// NewClient returns a new *Client that has appID.
func NewClient(appID string) *Client {
	return &Client{
		AppID:      appID,
		HTTPClient: http.DefaultClient,
	}
}

// Parse returns a ResultSet parsed japanese dependency.
func (c *Client) Parse(ctx context.Context, text string) (ResultSet, error) {
	var result ResultSet

	req, err := http.NewRequest(http.MethodGet, c.requestURL(text), nil)
	if err != nil {
		return result, err
	}
	req = req.WithContext(ctx)

	res, err := c.httpClient().Do(req)
	if err != nil {
		return result, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return result, xerrors.Errorf("Got http status %s", res.Status)
	}

	result, err = decodeResultSet(res.Body)

	return result, err
}

func (c *Client) requestURL(text string) string {
	_url, err := url.Parse(APIEndpoint)
	if err != nil {
		panic(err)
	}
	query := url.Values{}
	query.Add("appid", c.AppID)
	query.Add("sentence", text)
	_url.RawQuery = query.Encode()

	return _url.String()
}

func (c *Client) httpClient() *http.Client {
	if c.HTTPClient == nil {
		return http.DefaultClient
	}

	return c.HTTPClient
}
