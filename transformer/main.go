package main

func main() {
	rmqConn := newRabbit("amqp://usr:secret_pass@localhost:15672/", "transformDWG")
	defer rmqConn.conn.Close()
	defer rmqConn.chanL.Close()
	rmqConn.listenMessage()

}
