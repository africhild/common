// auth/pat.go
package auth

import (
    "fmt"
    "net/http"
)

// PATAuth holds the personal access token
type PATAuth struct {
    Token string
}

// ApplyAuth adds the personal access token to request headers
func (p *PATAuth) ApplyAuth(req *http.Request) error {
    if p.Token == "" {
        return fmt.Errorf("missing personal access token")
    }
    req.Header.Set("Authorization", "Bearer "+p.Token)
    return nil
}