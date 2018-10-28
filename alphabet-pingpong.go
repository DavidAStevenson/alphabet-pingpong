package main

import (
	"bytes"
	"flag"
	"log"
	"sync"

	"github.com/nats-io/go-nats"
)

func printMsg(m *nats.Msg, i int) {
	log.Printf("[#%d] Received on [%s]: '%s'\n", i, m.Subject, string(m.Data))
}

func alphabetRelay(c byte) byte {
	if c == 'Z' {
		return 'A'
	} else {
		return c+1
	}
}

func main() {

	var url = flag.String("url", nats.DefaultURL, "The NATS server URL to connect to") 
	var letter = flag.String("letter", "A", "The letter that is to processed")
	flag.Parse()
	var let = *letter
	var letr = let[0:1]
	
	log.Printf("alphabet-pingpong connecting to [%s]\n", *url)
	
	nc, err := nats.Connect(*url)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()
	
	// use WaitGroup to keep running indefinitely, won't call Done on it 
	wg := sync.WaitGroup{}
	wg.Add(1)
	
	subj, i := "alphabet." + string(letr), 0
	
	// Subscribe
	if _, err := nc.Subscribe(subj, func(msg *nats.Msg) {
		i += 1
		printMsg(msg, i)
		
		msgChar := msg.Data[0:1]
		if (bytes.Equal(msgChar, []byte(letr))) {
			res := alphabetRelay(byte(msgChar[0]))
			pubsubj, pubmsg := "alphabet." + string(res), []byte(string(res))

			nc.Publish(pubsubj, pubmsg)
			nc.Flush()
		
			if err := nc.LastError(); err != nil {
				log.Fatal(err)
			} else {
				log.Printf("Published [%s] : '%s'\n", pubsubj, pubmsg)
			}		
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
