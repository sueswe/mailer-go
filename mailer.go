package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml"
	"gopkg.in/gomail.v2"
)

/*
TODO: more attachments comma-seperated.
*/

var version string = "0.3.4"

var SMTPD string
var SENDER string
var home string = os.Getenv("HOME")

func help() {
	fmt.Println("Usage: ")
	fmt.Println("mailer [-f sender] [-d] [-t recipient,recipient] -s subject -b message [-a \"attachment*\"] ")
	fmt.Println("\n -> use -h for more help.")
}

func main() {

	infoLog := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	//warningLog := log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog := log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	infoLog.Print("mailer, Version ", version)

	config, err := toml.LoadFile(home + "/mailerconfig.toml")
	if err != nil {
		errorLog.Print("Error ", err.Error())
		panic(err)
	}
	SMTPD := config.Get("default.SMTPD").(string)
	SENDER := config.Get("default.SENDER").(string)
	infoLog.Print("Mailserver: ", SMTPD)
	infoLog.Print("Sender: ", SENDER)

	fromPart := flag.String("f", SENDER, "email-sender.")
	toPart := flag.String("t", SENDER, "email-recipient.")
	subjectPart := flag.String("s", "(no subject)", "email-subject.")
	bodyPart := flag.String("b", "(empty)", "email-body.")
	attachPart := flag.String("a", "(none)", "email-attachments.")
	flag.Parse()

	if *subjectPart == "(no subject)" || *bodyPart == "(empty)" {
		errorLog.Print("Sorry, I'm missing a mandatory parameter.")
		help()
		os.Exit(2)
	} else {
		//log.Print("Sender: \t", *fromPart)
		infoLog.Print("Recipient: \t", *toPart)
		infoLog.Print("Subject: \t", *subjectPart)
		infoLog.Print("Body: \t", *bodyPart)
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
		infoLog.Print("Attachment:\t", *attachPart)
	} else {
		//log.Print("globbing attachments: ", *attachPart)
		filenames, err := filepath.Glob(*attachPart)
		if err != nil {
			errorLog.Print("glob-error")
			os.Exit(2)
		}
		if len(filenames) == 0 {
			errorLog.Print("Attachments-Glob ist leer.")
			os.Exit(3)
		}
		for i, fname := range filenames {
			//log.Print(fname)
			_, error := os.Stat(fname)
			// check if error is "file not exists"
			if os.IsNotExist(error) {
				errorLog.Print("file does not exist: ", fname)
				os.Exit(5)
			}
			infoLog.Print("Attaching file ", i, " ,is: ", fname)
			m.Attach(fname)
		}
	}

	infoLog.Print("Trying to send ...")
	d := gomail.Dialer{Host: SMTPD, Port: 25}
	if err := d.DialAndSend(m); err != nil {
		errorLog.Print("Sorry, that didn't work.")
		panic(err)
	} else {
		infoLog.Print("done.")
	}

} //main_end
