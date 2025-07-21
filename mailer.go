package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml"
	"gopkg.in/gomail.v2"
)

var version string = "0.4.9"

var SMTPD string
var SENDER string
var home string = os.Getenv("HOME")

func createConfig(server_address string) {
	infoLog := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	s_a := strings.Split(server_address, ",")
	host := s_a[0]
	mail := s_a[1]

	if err := os.WriteFile(home+"/.mailerconfig.toml", []byte("[default]\nSMTPD = \""+host+"\"\nSENDER = \""+mail+"\"\n"), 0640); err != nil {
		log.Fatal(err)
	}
	infoLog.Print("writing " + home + "/.mailerconfig.toml")
	infoLog.Print("Contains: ")
	fmt.Println("[default]\nSMTPD = \"" + host + "\"\nSENDER = \"" + mail + "\"\n\n")
}

func readConfig(cf string) int {
	file, err := os.Open(cf)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	fmt.Println("$> " + cf + ": ")
	fmt.Println("")
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return 0
}

func main() {

	infoLog := log.New(os.Stdout, " INFO: ", log.Ldate|log.Ltime)
	//warningLog := log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog := log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime)

	infoLog.Print("mailer, Version ", version)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage:\n")
		flag.PrintDefaults()
		fmt.Printf("\nExample:\n\n")
		fmt.Println("$> mailer -t dagobert.duck@entenhausen.de -s \"Ds ist der Betreff\" -m \"<h2>Titel</h2><br>Und hier folgt der Emailbody. Auch als HTML möglich. \" -a attachments* ")
		fmt.Println("")
	}

	//configPart := flag.Bool("c", false, "Optional: creates a default config file.")
	configPart := flag.String("c", "localhost,nobody<at>nowhere.org", "Create configfile with values.")
	fromPart := flag.String("f", SENDER, "email-sender. Default is taken from config.")
	toPart := flag.String("t", SENDER, "email-recipients.")
	subjectPart := flag.String("s", "(no subject)", "email-subject.")
	bodyPart := flag.String("m", "(empty)", "message-body.")
	attachPart := flag.String("a", "(none)", "email-attachments.")
	readConfigPart := flag.Bool("r", false, "view configfile.")
	flag.Parse()

	if strings.Contains(*configPart, "@") {
		infoLog.Print("creating config")
		createConfig(*configPart)
		os.Exit(0)
	}

	if *readConfigPart {
		readConfig(home + "/.mailerconfig.toml")
		os.Exit(0)
	}

	config, err := toml.LoadFile(home + "/.mailerconfig.toml")
	if err != nil {
		errorLog.Print("Error ", err.Error())
		infoLog.Print("please run option -c .")
		//createConfig(*configPart)
		os.Exit(2)
	}
	SMTPD := config.Get("default.SMTPD").(string)
	SENDER := config.Get("default.SENDER").(string)
	infoLog.Print("Mailserver: ", SMTPD)
	infoLog.Print("Defaultsender: ", SENDER)

	if *subjectPart == "(no subject)" || *bodyPart == "(empty)" || *bodyPart == "." {
		errorLog.Print("Sorry, I'm missing a mandatory parameter.")
		errorLog.Println("use -h for help.")
		os.Exit(2)
	} else {
		if len(*fromPart) == 0 {
			infoLog.Print("Sender: \t", SENDER)
		} else {
			infoLog.Print("Sender: \t", *fromPart)
			SENDER = *fromPart
		}
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

	m.SetHeader("From", SENDER)
	m.SetHeader("To", addresses...)
	m.SetHeader("Subject", *subjectPart)

	// prepare the body:
	defBody := *bodyPart + "<br>---<br><pre>(please reply to: " + SENDER + ". mailer written in GoLang, build with ❤. Version " + version + ")</pre>"

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
