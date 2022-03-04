package util

import (
	"encoding/json"
	"fmt"
	"log"
)

func PrintJSON(d interface{}) {
	data, err := json.MarshalIndent(d, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(data))
}
