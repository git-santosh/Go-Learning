// Program in GO language with real world example of logging.
//`For example i am testing an SMTP connection is working fine or not. For test case I am going to connect to a SMTP server "smtp.smail.com" which is not exist, hence program will terminate with a log message.`
package main

import (
	"log"
	"net/smtp"
)

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("init started")
}
func main() {
	// Connect to the remote SMTP server.
	client, err := smtp.Dial("smtp.smail.com:25")
	if err != nil {
		log.Fatalln(err)
	}
	client.Data()
}
