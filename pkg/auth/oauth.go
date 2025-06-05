// auth/oauth.go
package auth

import (
    "fmt"
    "net/http"
)

// OAuth2Auth holds the OAuth2 bearer token
type OAuth2Auth struct {
    Token string
}

// ApplyAuth adds the OAuth2 bearer token to request headers
func (o *OAuth2Auth) ApplyAuth(req *http.Request) error {
    if o.Token == "" {
        return fmt.Errorf("missing OAuth2 token")
    }
    req.Header.Set("Authorization", "Bearer "+o.Token)
    return nil
}