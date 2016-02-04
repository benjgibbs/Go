package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/go-stomp/stomp"
	"gopkg.in/gcfg.v1"
	"log"
	"strconv"
)

const server = "datafeeds.nationalrail.co.uk"

type Cfg struct {
	Network struct {
		User, Pass string
	}
	Nre struct {
		User, Pass, Queue string
	}
}

func main() {
	cfg := Cfg{}
	err := gcfg.ReadFileInto(&cfg, "etc/trains.gcfg")
	if err != nil {
		log.Fatal(err)
	}
	fromNre(cfg)
}

func fromNre(cfg Cfg) {
	const nrePort = 61613
	addr := server + ":" + strconv.Itoa(nrePort)
	fmt.Println("Stomping trains from:", addr, " with user:", cfg.Nre.User)
	con, err := stomp.Dial("tcp", addr,
		stomp.ConnOpt.Login(cfg.Nre.User, cfg.Nre.Pass))
	if err != nil {
		log.Fatal(err)
	}
	sub, err := con.Subscribe("/queue/"+cfg.Nre.Queue, stomp.AckClient)
	if err != nil {
		log.Fatal(err)
	}
	for {
		msg := <-sub.C
		if msg.Err != nil {
			log.Fatal(msg.Err)
		}
		reader := bytes.NewReader(msg.Body)
		gz, _ := gzip.NewReader(reader)
		if err != nil {
			log.Fatal(err)
		}
		buff := make([]byte, 1024)
		for {
			n, _ := gz.Read(buff)
			if n == 0 {
				break
			}
		}
		fmt.Println(string(buff))
		err = con.Ack(msg)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = sub.Unsubscribe()
	if err != nil {
		log.Fatal(err)
	}
	con.Disconnect()

}

func fromNetworkRail(cfg Cfg) {
	const nrPort = 61618
	const nrTopic = "/topic/TRAIN_MVT_ALL_TOC"

	addr := server + ":" + strconv.Itoa(nrPort)

	fmt.Println("Stomping trains from:", addr, " with user:", cfg.Network.User)
	con, err := stomp.Dial("tcp", addr,
		stomp.ConnOpt.Login(cfg.Network.User, cfg.Network.Pass))
	if err != nil {
		log.Fatal(err)
	}

	sub, err := con.Subscribe(nrTopic, stomp.AckClient)
	if err != nil {
		log.Fatal(err)
	}

	for {
		msg := <-sub.C
		if msg.Err != nil {
			log.Fatal(msg.Err)
		}
		fmt.Println(msg)
		err = con.Ack(msg)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = sub.Unsubscribe()
	if err != nil {
		log.Fatal(err)
	}
	con.Disconnect()
}
