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
    fmt.Printf("[DEBUG] ApplyAuth - Key: %s, Field: %s, InHeader: %v\n", a.Key, a.Field, a.InHeader)
    
    if a.Key == "" || a.Field == "" {
        return fmt.Errorf("missing API key or field name")
    }

    if a.InHeader {
        req.Header.Set(a.Field, a.Key)
        fmt.Printf("[DEBUG] Set header - %s: %s\n", a.Field, a.Key)
    } else {
        q := req.URL.Query()
        q.Set(a.Field, a.Key)
        req.URL.RawQuery = q.Encode()
        fmt.Printf("[DEBUG] Set query parameter - %s: %s\n", a.Field, a.Key)
    }

    // Print all headers for debugging
    fmt.Printf("[DEBUG] All headers after ApplyAuth:\n")
    for k, v := range req.Header {
        fmt.Printf("[DEBUG] %s: %v\n", k, v)
    }

    return nil
}