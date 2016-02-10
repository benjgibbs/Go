package main

import (
	"bytes"
	"compress/gzip"
	"github.com/go-stomp/stomp"
	"log"
	"strconv"
)

const nrePort = 61613

func fromNre(cfg Cfg) {
	addr := server + ":" + strconv.Itoa(nrePort)
	fmt.Println("Stomping trains from:", addr, " with user:", cfg.Nre.User)
	con, err := stomp.Dial("tcp", addr, stomp.ConnOpt.Login(cfg.Nre.User, cfg.Nre.Pass))
	defer con.Disconnect()
	failIf(err)

	sub, err := con.Subscribe("/queue/"+cfg.Nre.Queue, stomp.AckClient)
	failIf(err)

	for {
		msg := <-sub.C
		failIf(msg.Err)
		reader := bytes.NewReader(msg.Body)
		gz, err := gzip.NewReader(reader)
		failIf(err)

		buff := make([]byte, 1024)
		for {
			n, _ := gz.Read(buff)
			if n == 0 {
				break
			}
		}

		log.Println(string(buff))
		err = con.Ack(msg)
		failIf(err)
	}
	err = sub.Unsubscribe()
	failIf(err)
}
