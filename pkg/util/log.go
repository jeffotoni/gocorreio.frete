package util

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/jeffotoni/gocorreio.frete/models"
)

func PrintErrorFretev2(msg string, gf *models.PostFretev2) {
	timestamp := time.Now().Format("2006-01-02T15:04:05.000Z")

	reqBody, err := json.Marshal(&gf)
	if err != nil {
		fmt.Println(timestamp, " - ERROR: ", msg)
		return
	}

	fmt.Println(timestamp, " - ERROR: ", msg, " - JSON: ", string(reqBody))
	return
}
