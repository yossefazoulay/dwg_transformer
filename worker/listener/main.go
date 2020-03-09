package main

import (
	"github.com/yossefazoulay/go_utils/queue"
	"listener/config"
	"listener/utils"
	"os"
)

func main() {
	config.GetConfig(os.Args[1])
	queueConf := config.LocalConfig.Queue.Rabbitmq
	rmqConn, err := queue.NewRabbit(queueConf.ConnString, queueConf.QueueNames)
	utils.HandleError(err, "Error Occured when RabbitMQ Init", err != nil)
	defer rmqConn.Conn.Close()
	defer rmqConn.ChanL.Close()
	go rmqConn.OpenListening(queueConf.Listennig, utils.MessageReceiver)
}
