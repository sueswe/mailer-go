package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"gopkg.in/gomail.v2"
)

var version string = "0.2"

// FIXME: config-file ($HOME/mailer.conf)
var SMTPD = "localhost"
var SENDER = "sueswe@localhost"

func help() {
	fmt.Println("Usage: ")
	fmt.Println("mailer [-f sender] [-d] [-t recipient,recipient] -s subject -b message [-a attachment] ")
	fmt.Println("\n -> use -h for more help!")
}

func details() {
	fmt.Print("Default sender:\t")
	color.Cyan(SENDER)
	fmt.Print("Default SMTPD:\t")
	color.Cyan(SMTPD)
}

func main() {

	log.Print("mailer, Version ", version)

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
		log.Print("Sender: \t", *fromPart)
		log.Print("Recipient: \t", *toPart)
		log.Print("Subject: \t", *subjectPart)
		log.Print("Body: \t", *bodyPart)
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

	if *attachPart == "(none)" {
		log.Print("Attachment:\t", *attachPart)
	} else {

		filenames, err := filepath.Glob(*attachPart)
		if err != nil {
			log.Fatal("Error globbing files.")
			os.Exit(3)
		}
		for i, name := range filenames {
			log.Print("Attaching file ", i, " ,is: ", name)
			m.Attach(name)
		}
		//log.Print(filenames)
	}

	log.Print("Trying to send ...")
	d := gomail.Dialer{Host: SMTPD, Port: 25}
	if err := d.DialAndSend(m); err != nil {
		log.Fatal("Sorry, that didn't work.")
		panic(err)
	} else {
		log.Print("done.")
	}
}
