package main
import "net"
import "fmt"
import "bufio"
import "strings" // only needed below for sample processing

func main() {

	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8081")

	// accept connection on port
	conn, _ := ln.Accept()

	// run loop forever (or until ctrl-c)
	for {
		// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// output message received
		fmt.Print("Message Received:", string(message))


		newmessage := strings.Split(string(message),",")
		// send new string back to client
		len := len(newmessage)
		messageToSend := ""

		for i := len-1; i>=0; i--{
			if i == len-1{
				messageToSend = messageToSend + newmessage[i]
			}else{
				messageToSend = messageToSend + "," + newmessage[i]
			}
		//	fmt.Print(newmessage[i])
		//	fmt.Print(",")
			fmt.Print(messageToSend)
         //messageToSend = messageToSend +","+ newmessage[i]
		}

		//conn.Write([]byte(messageToSend))
		//conn.Write([]byte("\n"))

	}
}
