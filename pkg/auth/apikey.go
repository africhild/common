// auth/apikey.go
package auth

import (
    "fmt"
    "net/http"
)

// // APIKeyAuth holds the API key and the header or query parameter name
type APIKeyAuth struct {
    Key   string
    Field string // e.g., "X-API-Key" or "api_key"
    InHeader bool
}

// // ApplyAuth adds the API key to the request (header or query parameter)
// func (a *APIKeyAuth) ApplyAuth(req *http.Request) error {
//     if a.Key == "" || a.Field == "" {
//         return fmt.Errorf("missing API key or field name")
//     }

//     if a.InHeader {
//         // Put API key in a custom header
//         req.Header.Set(a.Field, a.Key)
//     } else {
//         // Put API key as a query param
//         q := req.URL.Query()
//         q.Set(a.Field, a.Key)
//         req.URL.RawQuery = q.Encode()
//     }

//     return nil
// }

func (a *APIKeyAuth) ApplyAuth(req *http.Request) error {    
    if a.Key == "" || a.Field == "" {
        return fmt.Errorf("missing API key or field name")
    }
    if a.InHeader {
        req.Header.Set(a.Field, a.Key)
    } else {
        q := req.URL.Query()
        q.Set(a.Field, a.Key)
        req.URL.RawQuery = q.Encode()
    }
    return nil
}