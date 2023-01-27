package main

import (
    "flag"
    "fmt"
    "os"
    //"crypto/tls"
    "gopkg.in/gomail.v2"
    "strings"
    "github.com/fatih/color"
)

var SMTPD = "viruswall.sozvers.at"
var SENDER = "rz.om.stp@itsv.at"

func help() {
    color.Yellow("Usage: ")
    fmt.Println("mailer [-f sender] [-d] [-t recipient,recipient] -s subject -b body/message [-a attachments] ")
    fmt.Println("\n -> use -h for more help!\n")
    //fmt.Print("\nDefault sender and recipient is: ")
    //color.Cyan("rz.om.stp@itsv.at")
}

func details() {
  fmt.Print("Default sender: ")
  color.Cyan(SENDER)
  fmt.Print("Default SMTPD: ")
  color.Cyan(SMTPD)
}

func mailer_single(from , to, subject, body , file string) {
    m := gomail.NewMessage()
    m.SetHeader("From", from)
    m.SetHeader("To", to)
    m.SetHeader("Subject", subject)
    m.SetBody("text/plain", body)

    d := gomail.Dialer{Host: SMTPD, Port: 25}
    if err := d.DialAndSend(m); err != nil {
        color.Red("\nWhoops, that didn't work, pal!")
        panic(err)
    }
}

func main() {

    showDetails := flag.Bool("d" , false, "Show default configuration settings." )

    fromPart := flag.String("f", SENDER, "email-sender.")
    toPart := flag.String("t", SENDER, "email-recipient.")
    subjectPart := flag.String("s", "(no subject)", "email-subject.")
    bodyPart := flag.String("b", "(empty)", "email-body.")
    attachPart := flag.String("a", "(none)", "email-attachments.")
    flag.Parse()
    if  *showDetails == true {
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
    toSlice := strings.Split(*toPart, ",")
    for _, adress := range toSlice {
        fmt.Println("recipient:", adress)
        mailer_single(*fromPart, adress, *subjectPart, *bodyPart, *attachPart)
    }
}
