package mongodb

import (
	"regexp"
)

type DatabaseConfig struct {
	Name string
	URI  string
}

var configuration DatabaseConfig

func Init(URI *string) {
	re := regexp.MustCompile(`mongodb\:\/\/(?P<user>([^"]*))\:(?P<password>([^"]*))\@(?P<host>([^"]*))\:(?P<port>([^"]*))\/(?P<database>([^"]*))`)
	matches := re.FindStringSubmatch(*URI)

	configuration = DatabaseConfig{matches[5], *URI}
}
