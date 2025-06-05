package auth

import (
    "net/http"
)

// AuthStrategy is the interface defining how to apply an authentication method
type Strategy interface {
    // ApplyAuth modifies the request to include the authentication credentials
    ApplyAuth(req *http.Request) error
}