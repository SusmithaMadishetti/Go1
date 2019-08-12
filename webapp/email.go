// package main

// import (
// 	"log"
// 	"net/smtp"
// )

// func main() {
// 	send("You can contact me if you have any coding challenges..!")
// }

// func send(body string) {
// 	from := "juttadapavan@gmail.com"
// 	pass := "p@1kumar"
// 	to := "pavanjuttada@gmail.com"

// 	msg := "From: " + from + "\n" +
// 		"To: " + to + "\n" +
// 		"Subject: Go lang master minders\n\n" +
// 		body

// 	err := smtp.SendMail("smtp.gmail.com:587",
// 		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
// 		from, []string{to}, []byte(msg))

// 	if err != nil {
// 		log.Printf("smtp error: %s", err)
// 		return
// 	}

// 	log.Print("sent, It's time to go ..!")
// }
