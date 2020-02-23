package main

import "fmt"

func main() {
	// rmqConn := newRabbit("amqp://guest:guest@rabbitmq/", "transformDWG")
	// defer rmqConn.conn.Close()
	// defer rmqConn.chanL.Close()

	// for i := 0; i < 1000; i++ {
	// 	time.Sleep(time.Second)
	// 	rmqConn.sendMessage(message)
	// }

	root := "./"

	files := listFilesInDir(root)

	for _, file := range files {

		message := pickFile{
			Name: "File Uploaded",
			Path: file,
		}
		// rmqConn.sendMessage(message)
		fmt.Println(message)
	}

}
