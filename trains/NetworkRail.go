package main

import (
	"fmt"
	"github.com/go-stomp/stomp"
	"strconv"
)

const nrPort = 61618
const nrTopic = "/topic/TRAIN_MVT_ALL_TOC"

func fromNetworkRail(cfg Cfg) {

	addr := server + ":" + strconv.Itoa(nrPort)

	fmt.Println("Stomping trains from:", addr, " with user:", cfg.Network.User)
	loginOpt := stomp.ConnOpt.Login(cfg.Network.User, cfg.Network.Pass)
	con, err := stomp.Dial("tcp", addr, loginOpt)
	defer con.Disconnect()
	failIf(err)

	sub, err := con.Subscribe(nrTopic, stomp.AckClient)
	failIf(err)

	for {
		msg := <-sub.C
		failIf(msg.Err)
		fmt.Println(msg)
		err = con.Ack(msg)
		failIf(err)
	}
	err = sub.Unsubscribe()
	failIf(err)
}
