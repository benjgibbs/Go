package main

import (
	"bytes"
	"compress/gzip"
	"github.com/go-stomp/stomp"
	"log"
	"strconv"
)

const nrePort = 61613

type NREUpdates chan Pport

type NREFeed struct {
	sub *stomp.Subscription
	con *stomp.Conn
}

func (feed *NREFeed) Subscribe(cfg Cfg) NREUpdates {
	addr := server + ":" + strconv.Itoa(nrePort)
	log.Println("Stomping trains from:", addr, " with user:", cfg.Nre.User)
	var err error
	feed.con, err = stomp.Dial("tcp", addr, stomp.ConnOpt.Login(cfg.Nre.User, cfg.Nre.Pass))
	failIf(err)

	feed.sub, err = feed.con.Subscribe("/queue/"+cfg.Nre.Queue, stomp.AckClient)
	failIf(err)
	results := make(NREUpdates)
	go func() {
		for {
			msg := <-feed.sub.C
			failIf(msg.Err)
			reader := bytes.NewReader(msg.Body)
			gz, err := gzip.NewReader(reader)
			failIf(err)

			buff := make([]byte, 1024)
			str := []byte{}
			for {
				n, _ := gz.Read(buff)
				if n == 0 {
					break
				} else {
					str = append(str, buff[:n]...)
				}
			}
			results <- *ToStructs([]byte(str))
			err = feed.con.Ack(msg)
			failIf(err)
		}
	}()
	return results
}

func (feed *NREFeed) Unsubscribe() {
	err := feed.sub.Unsubscribe()
	failIf(err)
	feed.con.Disconnect()
}
