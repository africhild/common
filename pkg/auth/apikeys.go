// auth/apikey.go
package auth

import (
    "fmt"
    "net/http"
)

// APIKeysAuth holds the API key and the header or query parameter name
type APIKeysAuth struct {
    KeyArray []APIKeyAuth
    InHeader bool // If true, use headers; if false, use query parameters
}

func (a *APIKeysAuth) ApplyAuth(req *http.Request) error {
    for _, keyAuth := range a.KeyArray {
        if keyAuth.Key == "" || keyAuth.Field == "" {
            return fmt.Errorf("missing API key or field name")
        }

        if a.InHeader {
            req.Header.Set(keyAuth.Field, keyAuth.Key)
        } else {
            q := req.URL.Query()
            q.Set(keyAuth.Field, keyAuth.Key)
            req.URL.RawQuery = q.Encode()
        }
    }

    return nil
}