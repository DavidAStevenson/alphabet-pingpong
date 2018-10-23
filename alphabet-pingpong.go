package main

import (
	"log"
//	"sync"
	"runtime"

	"github.com/nats-io/go-nats"
)

func printMsg(m *nats.Msg, i int) {
	log.Printf("[#%d] Received on [%s]: '%s'\n", i, m.Subject, string(m.Data))
}

func testDefer() {
	log.Printf("testDefer called\n")
}

func main() {
	nc, err := nats.Connect("nats://192.168.99.100:4222")
	if err != nil {
		log.Fatal(err)
	}
//	defer testDefer()
//	defer nc.Close()
//	defer testDefer()

	
	subj, i := "alphabet.*", 0
	
	// Subscribe
	if _, err := nc.Subscribe(subj, func(msg *nats.Msg) {

		i += 1
		printMsg(msg, i)
		
		if string(msg.Data) == "A" {
			pubsubj, pubmsg := "alphabet.B", []byte("B")

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
	
	runtime.Goexit()
}

/*	
	// publisher
	subj, msg := "alphabet.A", []byte("A")

	nc.Publish(subj, msg)
	nc.Flush()
	
	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Published [%s] : '%s'\n", subj, msg)
	}
*/

//		wg.Done()

/*
	// as per https://github.com/nats-io/go-nats-examples/blob/master/api-examples/subscribe_async/main.go
	wg := sync.WaitGroup{}
	wg.Add(1)
*/

/*	
	// Wait for a message to come in
	wg.Wait()
*/
	
	/*
	
	nc.Subscribe(subj, func(msg *nats.Msg) {
		i += 1
		printMsg(msg, i)
	})
	nc.Flush()
	
	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	}
	*/
