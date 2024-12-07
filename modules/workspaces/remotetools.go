package workspaces

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"
	"github.com/hashicorp/go-retryablehttp"
)

type FakeClientResponse struct {
	Response       *http.Response
	Err            error
	BodyContent    string
	ResolveContent func(url *url.URL) string
}

func (f *FakeClientResponse) RoundTrip(req *http.Request) (*http.Response, error) {
	res := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBufferString(f.ResolveContent(req.URL))),
		Header:     make(http.Header),
	}
	return res, nil
}

func CreateMockHTTPClient(contentResolver func(url *url.URL) string) *retryablehttp.Client {
	client := retryablehttp.NewClient()

	response := &http.Response{
		StatusCode: 200,
		Header:     make(http.Header),
	}

	client.HTTPClient.Transport = &FakeClientResponse{
		Response:       response,
		ResolveContent: contentResolver,
	}

	return client
}

type HTTPRequestOptions struct {
	Method  string      // HTTP method (GET, POST, PUT, DELETE, etc.)
	Headers http.Header // HTTP headers
	Body    interface{} // Request body
}

func MakeHTTPRequest(
	client *retryablehttp.Client,
	url string,
	queryParams any,
	options HTTPRequestOptions,
) ([]byte, *http.Response, *IError) {

	if queryParams != nil {
		v, err := query.Values(queryParams)
		if err != nil {
			return nil, nil, CastToIError(err)
		}

		if !strings.HasSuffix(url, "?") {
			url += "?"
		}

		url += v.Encode()
	}

	// Set the HTTP method, URL, and request body for the retryable request
	req, err := retryablehttp.NewRequest(strings.ToUpper(options.Method), url, options.Body)

	if err != nil {
		return nil, nil, GormErrorToIError(err)
	}

	req.Header = options.Headers

	fmt.Println(1, req.Header)

	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, CastToIError(err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp, CastToIError(err)
	}

	return body, resp, nil
}
