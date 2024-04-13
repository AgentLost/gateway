package env

import (
	"fmt"
	"os"
)

var (
	HttpPort = fmt.Sprintf(":%s", Getter("HTTP_PORT", "8080"))
)

func Getter(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return defaultValue
}
