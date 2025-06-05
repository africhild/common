// auth/pat.go
package auth

import (
    "net/http"
)

// NoAuth holds nothing
type NoAuth struct {
 
}

// NoAuth does nothing
func (p *NoAuth) ApplyAuth(req *http.Request) error {
    return nil
}