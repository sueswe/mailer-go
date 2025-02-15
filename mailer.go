package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
	"gopkg.in/gomail.v2"
)

var SMTPD = "localhost"
var SENDER = "sueswe@localhost"

func help() {
	fmt.Println("Usage: ")
	fmt.Println("mailer [-f sender] [-d] [-t recipient,recipient] -s subject -b body/message [-a attachments] ")
	fmt.Println("\n -> use -h for more help!")
	//fmt.Print("\nDefault sender and recipient is: ")
	//color.Cyan("rz.om.stp@itsv.at")
}

func details() {
	fmt.Print("Default sender:\t")
	color.Cyan(SENDER)
	fmt.Print("Default SMTPD:\t")
	color.Cyan(SMTPD)
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
		log.Fatal("Sorry, I'm missing something.")
		help()
		os.Exit(1)
	} else {
		fmt.Println("Sender: \t", *fromPart)
		fmt.Println("Recipient: \t", *toPart)
		fmt.Println("Subject: \t", *subjectPart)
		fmt.Println("Body: \t\t", *bodyPart)
		fmt.Println("Attachments: \t", *attachPart)
	}

	m := gomail.NewMessage()
	toSlice := strings.Split(*toPart, ",")
	addresses := make([]string, len(toSlice))
	for i, adress := range toSlice {
		addresses[i] = m.FormatAddress(adress, "")
	}

	m.SetHeader("From", *fromPart)
	m.SetHeader("To", addresses...)
	m.SetHeader("Subject", *subjectPart)
	m.SetBody("text/plain", *bodyPart)

	d := gomail.Dialer{Host: SMTPD, Port: 25}
	if err := d.DialAndSend(m); err != nil {
		log.Fatal("Whoops, that didn't work, pal!")
		panic(err)
	} else {
		log.Print("DONE.")
	}
}
