// Package json provides utilities for working with JSON in HTTP request & responses.
package json

import (
	jsonEncoder "encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const (
	// contentTypeKey is the content type header key
	contentTypeKey = "Content-Type"

	// contentTypeVal is the json content type header value
	contentTypeVal = "application/json"
)

// Service provides utilities for working with JSON in HTTP request & responses.
type Service struct{}

func NewService() *Service {
	return &Service{}
}

// Write writes a JSON response to the provided http.ResponseWriter.
func (s *Service) Write(w http.ResponseWriter, data Response) {
	if data.StatusCode < 100 {
		data.StatusCode = http.StatusOK
	}

	if data.Status == "" {
		if data.StatusCode >= 500 {
			data.Status = StatusError
		} else {
			data.Status = StatusSuccess
		}
	}

	if data.Error != nil {
		data.Status = StatusError

		// if utils.IsProd() {
		// 	data.Error.Stack = ""
		// 	data.Error.Cause = ""
		// }
	}

	if data.Message == "" {
		data.Message = http.StatusText(data.StatusCode)
	}

	w.Header().Add(contentTypeKey, fmt.Sprintf("%s; charset=utf-8", contentTypeVal))
	w.WriteHeader(data.StatusCode)
	jsonEncoder.NewEncoder(w).Encode(data)
}

// Unmarshal attempts to parse the JSON request body into v
func (s *Service) UnmarshalBody(r *http.Request, v any) (err error) {
	t := r.Header.Get(contentTypeKey)

	defer func() {
		if err != nil {
			err = fmt.Errorf("json unmarshal error: %w", err)
		}
	}()

	if strings.Split(t, ";")[0] != contentTypeVal {
		err = fmt.Errorf("unexpected content type: \"%s\"", t)
		return err
	}

	if err = jsonEncoder.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}

	return nil
}
