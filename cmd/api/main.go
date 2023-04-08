package main

import (
	"jagch/backend/cmd/api/boostrap"
	"log"
)

func main() {
	if err := boostrap.Run(); err != nil {
		log.Fatal(err)
	}
}
