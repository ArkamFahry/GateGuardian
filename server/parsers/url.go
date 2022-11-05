package parsers

import (
	"net/url"
	"strings"

	"github.com/ArkamFahry/GateGuardian/server/constants"
	"github.com/ArkamFahry/GateGuardian/server/env"
	"github.com/gin-gonic/gin"
)

func GetHost(c *gin.Context) string {
	authorizerURL, err := env.GetEnvByKey(constants.GateGuardianURL)
	if err != nil {
		authorizerURL = ""
	}
	if authorizerURL != "" {
		return authorizerURL
	}

	authorizerURL = c.Request.Header.Get("X-Authorizer-URL")
	if authorizerURL != "" {
		return authorizerURL
	}

	scheme := c.Request.Header.Get("X-Forwarded-Proto")
	if scheme != "https" {
		scheme = "http"
	}
	host := c.Request.Host
	return scheme + "://" + host
}

func GetHostParts(uri string) (string, string) {
	tempURI := uri
	if !strings.HasPrefix(tempURI, "http://") && !strings.HasPrefix(tempURI, "https://") {
		tempURI = "https://" + tempURI
	}

	u, err := url.Parse(tempURI)
	if err != nil {
		return "localhost", "8080"
	}

	host := u.Hostname()
	port := u.Port()

	return host, port
}

func GetDomainName(uri string) string {
	tempURI := uri
	if !strings.HasPrefix(tempURI, "http://") && !strings.HasPrefix(tempURI, "https://") {
		tempURI = "https://" + tempURI
	}

	u, err := url.Parse(tempURI)
	if err != nil {
		return `localhost`
	}

	host := u.Hostname()

	// code to get root domain in case of sub-domains
	hostParts := strings.Split(host, ".")
	hostPartsLen := len(hostParts)

	if hostPartsLen == 1 {
		return host
	}

	if hostPartsLen == 2 {
		if hostParts[0] == "www" {
			return hostParts[1]
		} else {
			return host
		}
	}

	if hostPartsLen > 2 {
		return strings.Join(hostParts[hostPartsLen-2:], ".")
	}

	return host
}

func GetAppURL(gc *gin.Context) string {
	envAppURL, err := env.GetEnvByKey(constants.AppURL)
	if envAppURL == "" || err != nil {
		envAppURL = GetHost(gc) + "/app"
	}
	return envAppURL
}
