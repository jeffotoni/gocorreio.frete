package token

import (
	"time"

	"github.com/jeffotoni/gocorreio.frete/pkg/util"
)

func CheckTokenExpired(expiraEm string) (hasExpired bool, err error) {
	expiraEm = util.Concat(expiraEm, ".000Z")

	layout := "2006-01-02T15:04:05.000Z"
	timeToken, err := time.Parse(layout, expiraEm)
	if err != nil {
		return
	}

	timeToken = timeToken.UTC()
	timeNow := time.Now().UTC()

	// UTC
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		return
	}

	expiresAt := timeToken.In(loc)

	diff := timeNow.Sub(expiresAt)

	// 30 minutos antes das 24h um novo token pode ser gerado
	// Tempo de 30 minutos definido pelo Correios
	if diff.Minutes() >= -30 {
		hasExpired = true
	} else {
		hasExpired = false
	}
	return
}
