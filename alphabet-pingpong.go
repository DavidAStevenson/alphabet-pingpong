package main

import (
	"log"

	"github.com/nats-io/go-nats"
)

func main() {
	nc, err := nats.Connect("nats://192.168.99.100:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()
	
	subj, msg := "alphabet.A", []byte("A")

	nc.Publish(subj, msg)
	nc.Flush()
	
	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Published [%s] : '%s'\n", subj, msg)
	}
}

