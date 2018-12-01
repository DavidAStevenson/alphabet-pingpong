package main

import (
	"bytes"
	"flag"
	"log"
	"sync"
	"time"

	"github.com/nats-io/go-nats"
)

func printMsg(m *nats.Msg, i int) {
	log.Printf("[#%d] Received on [%s]: '%s'\n", i, m.Subject, string(m.Data))
}

func alphabetRelay(c byte) byte {
	if c == 'z' {
		return 'a'
	}
	if c == 'Z' {
		return 'A'
	}
	return c + 1
}

func publishMessage(msg string, nc *nats.Conn) {
	pubsubj, pubmsg := "alphabet."+msg, []byte(msg)

	nc.Publish(pubsubj, pubmsg)
	nc.Flush()

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Published [%s] : '%s'\n", pubsubj, pubmsg)
	}
}

func main() {

	var (
		url    = flag.String("url", nats.DefaultURL, "The NATS server URL to connect to")
		letter = flag.String("letter", "A", "The letter that is to processed")
		seed   = flag.Bool("seed", false, "true/false whether to automatically seed with a letter")
	)
	flag.Parse()

	var let = *letter
	var letr = let[0:1]

	log.Println("Starting alphabet-pingpong...")
	log.Printf("alphabet-pingpong handling letter [%s], auto-seed is %t\n", string(letr), *seed)
	log.Printf("alphabet-pingpong connecting to [%s]\n", *url)

	var nc *nats.Conn
	var err error

	maxAttempts := 10
	for attempts := 0; attempts < maxAttempts; attempts++ {
		nc, err = nats.Connect(*url)
		if err == nil {
			break
		}
		log.Println(err)
		time.Sleep(time.Duration(attempts) * time.Second)
	}
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()
	log.Println("alphabet-pingpong started successfully.")

	seed_needed := *seed

	if *seed {
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()

		go func(msg string, nc *nats.Conn) {
			for {
				select {
				case t := <-ticker.C:
					log.Println("Current time: ", t)
					if seed_needed {
						log.Println("You publish now...")
						publishMessage(msg, nc)
					}
					seed_needed = true // seed next tick, unless action happens
				}
			}
		}(string(letr), nc)
	}

	// use WaitGroup to keep running indefinitely, won't call Done on it
	wg := sync.WaitGroup{}
	wg.Add(1)

	subj, i := "alphabet."+string(letr), 0

	// Subscribe
	if _, err := nc.Subscribe(subj, func(msg *nats.Msg) {
		i++
		printMsg(msg, i)

		msgChar := msg.Data[0:1]
		if bytes.Equal(msgChar, []byte(letr)) {
			// received the letter, no need to seed it 
			if *seed {
				seed_needed = false
			}

			res := alphabetRelay(byte(msgChar[0]))

			publishMessage(string(res), nc)
		} else {
			log.Printf("Nothing to do with '%s'\n", string(msg.Data))
		}
	}); err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on [%s]\n", subj)

	// Wait for the subscriber to receive a message
	wg.Wait()
}
