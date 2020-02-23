package main

func main() {
	rmqConn := newRabbit("amqp://guest:guest@rabbitmq/", "transformDWG")
	defer rmqConn.conn.Close()
	defer rmqConn.chanL.Close()
	rmqConn.listenMessage()

}
