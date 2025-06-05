package http

import (
	"encoding/json"
	"net/http"
	"mime"
	"fmt"
	"github.com/spyzhov/ajson"

)


type JSONHTTPResponse struct {
	// bodyBytes is the raw response body. It's not JSON-unmarshalled.
	// We keep it around so that we can unmarshal it into a struct later,
	// if needed (via the UnmarshalJSON function).
	bodyBytes []byte

	// Code is the HTTP status code of the response.
	Code int

	// Headers are the HTTP headers of the response.
	Headers http.Header

	// body is the JSON-unmarshalled response body. Aside from the fact
	// that it's JSON-unmarshalled, it's identical to bodyBytes.
	// If there were no bytes this will be nil.
	body *ajson.Node
}

// Body returns JSON node. If it is empty the flag will indicate so.
// Empty response body is a special case and should be handled explicitly.
func (j *JSONHTTPResponse) Body() (*ajson.Node, bool) {
	if j.body == nil {
		return nil, false
	}

	return j.body, true
}


// parseJSONResponse parses the given HTTP response and returns a JSONHTTPResponse.
func parseResponse(res *http.Response, body []byte) (*JSONHTTPResponse, error) {
	// empty response body should not be parsed as JSON since it will cause ajson to err
	if len(body) == 0 {
		return nil, nil //nolint:nilnil
	}
	// Ensure the response is JSON
	ct := res.Header.Get("Content-Type")
	if len(ct) > 0 {
		mimeType, _, err := mime.ParseMediaType(ct)
		if err != nil {
			return nil, fmt.Errorf("failed to parse content type: %w", err)
		}

		// Providers implementing JSONAPISpeicifcations returns application/vnd.api+json
		if mimeType != "application/json" && mimeType != "application/vnd.api+json" {
			return nil, fmt.Errorf("%w: expected content type to be application/json or application/vnd.api+json , got %s",
				ErrNotJSON, mimeType,
			)
		}
	}

	// Unmarshall the response body into JSON
	jsonBody, err := ajson.Unmarshal(body)
	if err != nil {
		return nil, NewHTTPStatusError(res.StatusCode, fmt.Errorf("failed to unmarshall response body into JSON: %w", err))
	}

	return &JSONHTTPResponse{
		bodyBytes: body,
		Code:      res.StatusCode,
		Headers:   res.Header,
		body:      jsonBody,
	}, nil
}

// UnmarshalJSON deserializes the response body into the given type.
func UnmarshalResponse[T any]( body []byte) (*T, error) {
	var data T
	if len(body) == 0 {
		return nil, nil //nolint:nilnil
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body into JSON: %w", err)
	}

	return &data, nil
}