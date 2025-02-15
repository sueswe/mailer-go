package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"gopkg.in/gomail.v2"
)

var SMTPD = "localhost"
var SENDER = "sueswe@carbon"

func help() {
	color.Yellow("Usage: ")
	fmt.Println("mailer [-f sender] [-d] [-t recipient,recipient] -s subject -b body/message [-a attachments] ")
	fmt.Println("\n -> use -h for more help!")
	//fmt.Print("\nDefault sender and recipient is: ")
	//color.Cyan("rz.om.stp@itsv.at")
}

func details() {
	fmt.Print("Default sender: ")
	color.Cyan(SENDER)
	fmt.Print("Default SMTPD: ")
	color.Cyan(SMTPD)
}

func mailer_single(from, to, subject, body, file string) {
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.Dialer{Host: SMTPD, Port: 25}
	if err := d.DialAndSend(m); err != nil {
		log.Fatal("\nWhoops, that didn't work, pal!")
		panic(err)
	}
}

func main() {

	showDetails := flag.Bool("d", false, "Show default configuration settings.")

	fromPart := flag.String("f", SENDER, "email-sender.")
	toPart := flag.String("t", SENDER, "email-recipient.")
	subjectPart := flag.String("s", "(no subject)", "email-subject.")
	bodyPart := flag.String("b", "(empty)", "email-body.")
	attachPart := flag.String("a", "(none)", "email-attachments.")
	flag.Parse()
	if *showDetails == true {
		details()
		os.Exit(0)
	}
	if *subjectPart == "no subject" || *bodyPart == "(empty)" {
		//usage(5)
		color.Red("Sorry, I'm missing something.")
		help()
		os.Exit(1)
	} else {
		fmt.Println("Sender: \t", *fromPart)
		fmt.Println("Recipient: \t", *toPart)
		fmt.Println("Subject: \t", *subjectPart)
		fmt.Println("Body: \t\t", *bodyPart)
		fmt.Println("Attachments: \t", *attachPart)
	}

	/*#toSlice := strings.Split(*toPart, ",")
	#for _, adress := range toSlice {
	#	fmt.Println("recipient:", adress)
	#	mailer_single(*fromPart, adress, *subjectPart, *bodyPart, *attachPart)
	#}*/

	m := gomail.NewMessage()
	m.SetHeader("From", *fromPart)
	m.SetHeader("To", *toPart)
	m.SetHeader("Subject", *subjectPart)
	m.SetBody("text/plain", *bodyPart)

	d := gomail.Dialer{Host: SMTPD, Port: 25}
	if err := d.DialAndSend(m); err != nil {
		log.Fatal("Whoops, that didn't work, pal!")
		panic(err)
	}
}
