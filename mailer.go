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

var version string = "0.4.0"

var SMTPD string
var SENDER string
var home string = os.Getenv("HOME")

func help() {
	fmt.Println("Usage: ")
	fmt.Println("mailer [-f sender] -t recipient,recipient -s subject -m message [-a \"attachment*,attachment\"] ")
	fmt.Println("(use -h for extended help.)")
}

func createConfig() {
	infoLog := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	if err := os.WriteFile(home+"/TEMPLATE.mailerconfig.toml", []byte("[default]\nSMTPD = \"localhost\"\nSENDER = \"user@domain.at\"\n"), 0640); err != nil {
		log.Fatal(err)
	}
	infoLog.Print("writing " + home + "/TEMPLATE.mailerconfig.toml")
}

func main() {

	infoLog := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	//warningLog := log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog := log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	infoLog.Print("mailer, Version ", version)

	configPart := flag.Bool("c", false, "Optional: creates a default config file.")
	fromPart := flag.String("f", SENDER, "email-sender.")
	toPart := flag.String("t", SENDER, "email-recipient.")
	subjectPart := flag.String("s", "(no subject)", "email-subject.")
	bodyPart := flag.String("m", "(empty)", "message-body.")
	attachPart := flag.String("a", "(none)", "email-attachments.")
	flag.Parse()

	config, err := toml.LoadFile(home + "/.mailerconfig.toml")
	if err != nil {
		errorLog.Print("Error ", err.Error())
		panic(err)
	}
	SMTPD := config.Get("default.SMTPD").(string)
	SENDER := config.Get("default.SENDER").(string)
	infoLog.Print("Mailserver: ", SMTPD)
	// infoLog.Print("Defaultsender: ", SENDER)

	if *configPart == true {
		infoLog.Print("creating config")
		createConfig()
		os.Exit(0)
	}

	if *subjectPart == "(no subject)" || *bodyPart == "(empty)" {
		errorLog.Print("Sorry, I'm missing a mandatory parameter.")
		help()
		os.Exit(2)
	} else {
		infoLog.Print("Sender: \t", *fromPart)
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

	// prepare the body:
	defBody := *bodyPart + "<br>---<br><pre>(please reply to: " + SENDER + ". Made with GoLang and ❤️. Version " + version + ")</pre>"

	m.SetBody("text/html", defBody)
	// m.AddAlternative("text/html", "<br>-----<br><pre>(please reply to: "+SENDER+")</pre>")

	// Attachment-Parameter verarbeiten:
	if *attachPart == "(none)" {
		infoLog.Print("Attachment:\t", *attachPart)
	} else {
		attSlice := strings.Split(*attachPart, ",")
		for _, a := range attSlice {
			//infoLog.Print("inhalt: ", a)
			//log.Print("globbing attachments: ", *attachPart)
			filenames, err := filepath.Glob(a)
			if err != nil {
				errorLog.Print("glob-error")
				os.Exit(2)
			}
			if len(filenames) == 0 {
				errorLog.Print("attachments not found.")
				os.Exit(3)
			}

			// folgende for wird nie erreicht wenn attachments ohnehin leer sind
			// (deshalb auch ein os:exit bei der Prüfung zuvor).
			// Macht nur Sinn, wenn während der Ausführung ein file verschwindet:
			for _, fname := range filenames {
				_, error := os.Stat(fname)
				// check if error is "file not exists"
				if os.IsNotExist(error) {
					errorLog.Print("file does not exist: ", fname)
					os.Exit(5)
				}

				// befuellen der Email mit Attachments:
				infoLog.Print("Attaching file: ", fname)
				m.Attach(fname)
			}
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
