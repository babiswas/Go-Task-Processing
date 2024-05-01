package HTTP

import (
	"log"
	"os"
)

func handle_error(err error) {
	if err == nil {
		return
	} else {
		log.Fatal(err)
		os.Exit(1)
	}
}
