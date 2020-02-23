package main

import "time"

func main() {
	rmqConn := newRabbit("amqp://guest:guest@rabbitmq/", "transformDWG")
	defer rmqConn.conn.Close()
	defer rmqConn.chanL.Close()
	root := "./"

	files := listFilesInDir(root)

	for _, file := range files {

		message := pickFile{
			Name: "File Uploaded",
			Path: file,
		}
		time.Sleep(time.Second)
		rmqConn.sendMessage(message)
	}

}
