// auth/basic.go
package auth

import (
    "encoding/base64"
    "fmt"
    "net/http"
)

// BasicAuth holds the credentials for basic authentication
type BasicAuth struct {
    Username string
    Password string
}

// ApplyAuth encodes username and password and adds them to the Authorization header
func (b *BasicAuth) ApplyAuth(req *http.Request) error {
    if b.Username == "" || b.Password == "" {
        return fmt.Errorf("missing username or password")
    }
    credentials := b.Username + ":" + b.Password
    encoded := base64.StdEncoding.EncodeToString([]byte(credentials))
    req.Header.Set("Authorization", "Basic "+encoded)
    return nil
}

// func NewBasicAuthHTTPClient( //nolint:ireturn
// 	ctx context.Context,
// 	user, pass string,
// 	opts ...HeaderAuthClientOption,
// ) (AuthenticatedHTTPClient, error) {
// 	return NewHeaderAuthHTTPClient(ctx, append(opts, WithHeaders(Header{
// 		Key:   "Authorization",
// 		Value: "Basic " + basicAuth(user, pass),
// 	}))...)
// }

// // shamelessly stolen from https://pkg.go.dev/net/http#Request.SetBasicAuth
// func basicAuth(username, password string) string {
// 	auth := username + ":" + password

// 	return base64.StdEncoding.EncodeToString([]byte(auth))
// }