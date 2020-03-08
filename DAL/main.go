package main

import (
	"dal/config"
	"dal/utils"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/yossefazoulay/go_utils/queue"
	"os"
)


func main() {

	config.GetConfig(os.Args[1])
	queueConf := config.LocalConfig.Queue.Rabbitmq
	rmqConn, err := queue.NewRabbit(queueConf.ConnString, queueConf.QueueNames)
	utils.HandleError(err, "Error Occured when RabbitMQ Init", err != nil)
	defer rmqConn.Conn.Close()
	defer rmqConn.ChanL.Close()
	rmqConn.OpenListening(queueConf.Listennig, utils.MessageReceiver)


}

