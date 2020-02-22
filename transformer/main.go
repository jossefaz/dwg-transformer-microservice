package main

func main() {
	rmqConn := newRabbit("amqp://guest:guest@localhost:5672/", "transformDWG")
	defer rmqConn.conn.Close()
	defer rmqConn.chanL.Close()
	rmqConn.listenMessage()

}
