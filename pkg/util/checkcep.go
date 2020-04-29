package util

import (
	"errors"
	"regexp"
)

func CheckRast(Rast string) error {
	re := regexp.MustCompile(`[^0-9][a-z]{13}`)
	formatedRast := re.ReplaceAllString(Rast, `$1`)

	if len(formatedRast) < 8 {
		return errors.New(`{"msg":"error Rast tem que ser valido"}`)
	}

	return nil
}
