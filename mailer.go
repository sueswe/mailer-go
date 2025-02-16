package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/pelletier/go-toml"
	"gopkg.in/gomail.v2"
)

var version string = "0.3.1"

var SMTPD string
var SENDER string

func help() {
	fmt.Println("Usage: ")
	fmt.Println("mailer [-f sender] [-d] [-t recipient,recipient] -s subject -b message [-a \"attachment*\"] ")
	fmt.Println("\n -> use -h for more help.")
}

func details() {
	fmt.Print("Default sender:\t")
	color.Cyan(SENDER)
	fmt.Print("Default SMTPD:\t")
	color.Cyan(SMTPD)
}

func main() {

	log.Print("mailer, Version ", version)

	config, err := toml.LoadFile("/home/sueswe/mailerconfig.toml")
	if err != nil {
		fmt.Println("Error ", err.Error())
		panic(err)
	}
	SMTPD := config.Get("default.SMTPD").(string)
	SENDER := config.Get("default.SENDER").(string)
	log.Print("Mailserver: ", SMTPD)
	log.Print("Sender: ", SENDER)

	showConfig := flag.Bool("c", false, "Show default configuration settings from configfile.")
	fromPart := flag.String("f", SENDER, "email-sender.")
	toPart := flag.String("t", SENDER, "email-recipient.")
	subjectPart := flag.String("s", "(no subject)", "email-subject.")
	bodyPart := flag.String("b", "(empty)", "email-body.")
	attachPart := flag.String("a", "(none)", "email-attachments.")
	flag.Parse()

	if *showConfig == true {
		details()
		os.Exit(3)
	}

	if *subjectPart == "(no subject)" || *bodyPart == "(empty)" {
		log.Print("Sorry, I'm missing a mandatory parameter.")
		help()
		os.Exit(2)
	} else {
		log.Print("Sender: \t", *fromPart)
		log.Print("Recipient: \t", *toPart)
		log.Print("Subject: \t", *subjectPart)
		log.Print("Body: \t", *bodyPart)
	}

	m := gomail.NewMessage()

	// Recipients zerlegen und adden:
	toSlice := strings.Split(*toPart, ",")
	addresses := make([]string, len(toSlice))
	for i, adress := range toSlice {
		addresses[i] = m.FormatAddress(adress, "")
	}

	m.SetHeader("From", *fromPart)
	m.SetHeader("To", addresses...)
	m.SetHeader("Subject", *subjectPart)
	m.SetBody("text/plain", *bodyPart)

	// Attachment-Parameter verarbeiten:
	if *attachPart == "(none)" {
		log.Print("Attachment:\t", *attachPart)
	} else {
		log.Print("globbing attachments: ", *attachPart)
		filenames, err := filepath.Glob(*attachPart)
		if err != nil {
			log.Print("glob-error")
			os.Exit(2)
		}
		if len(filenames) == 0 {
			log.Print("Attachments-Glob ist leer.")
			os.Exit(3)
		}
		for i, fname := range filenames {
			log.Print(fname)
			_, error := os.Stat(fname)
			// check if error is "file not exists"
			if os.IsNotExist(error) {
				log.Print("file does not exist: ", fname)
				os.Exit(5)
			}
			log.Print("Attaching file ", i, " ,is: ", fname)
			m.Attach(fname)
		}
	}

	log.Print("Trying to send ...")
	d := gomail.Dialer{Host: SMTPD, Port: 25}
	if err := d.DialAndSend(m); err != nil {
		log.Print("Sorry, that didn't work.")
		panic(err)
	} else {
		log.Print("done.")
	}

} //main_end
