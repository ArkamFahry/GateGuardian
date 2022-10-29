// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Env struct {
	DatabaseURL       *string `json:"DATABASE_URL"`
	DatabaseName      *string `json:"DATABASE_NAME"`
	DatabaseNamespace *string `json:"DATABASE_NAMESPACE"`
	DatabaseUsername  *string `json:"DATABASE_USERNAME"`
	Port              *string `json:"PORT"`
	JwtType           *string `json:"JWT_TYPE"`
	JwtSecret         *string `json:"JWT_SECRET"`
	ClientID          *string `json:"CLIENT_ID"`
}
