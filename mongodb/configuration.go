package mongodb

import (
	"regexp"
)

type DatabaseConfig struct {
	Name string
	URI  string
}

var configuration DatabaseConfig

func Init(URI *string) error {
	re, err := regexp.Compile(`mongodb\:\/\/(?P<user>([^"]*))\:(?P<password>([^"]*))\@(?P<host>([^"]*))\:(?P<port>([^"]*))\/(?P<database>([^"]*))`)

	if err != nil {
		return err
	}

	matches := re.FindStringSubmatch(*URI)
	configuration = DatabaseConfig{matches[5], *URI}

	return nil
}
