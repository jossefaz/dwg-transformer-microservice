package main

import "time"

func main() {
	rmqConn := newRabbit("amqp://usr:secret_pass@localhost:15672/", "transformDWG")
	defer rmqConn.conn.Close()
	defer rmqConn.chanL.Close()
	message := pickFile{
		Name: "test",
		Path: "pathtest",
	}
	for i := 0; i < 1000; i++ {
		time.Sleep(time.Second)
		rmqConn.sendMessage(message)
	}

}
