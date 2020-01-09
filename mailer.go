package main

import (
    "flag"
    "fmt"
    "os"
    //"crypto/tls"
    "gopkg.in/gomail.v2"
)

func usage(a int) {
    fmt.Println("Usage:", a)
}

func help() {
    fmt.Println("Help: ")
    fmt.Println("    mailer -f sender -t recipient -s subject -b body/message -a attachments \n")
}

func mailing() {
    m := gomail.NewMessage()
    m.SetHeader("From", "werner.suess@itsv.at")
    m.SetHeader("To", "werner.suess@itsv.at")
    m.SetHeader("Subject", "Hello, just a test")
    m.SetBody("text/plain", "Ahoihoi!")

    d := gomail.Dialer{Host: "viruswall.sozvers.at", Port: 25}
    if err := d.DialAndSend(m); err != nil {
        panic(err)
    }
}

func main() {
    fromPart := flag.String("f", "rz.om.stp@itsv.at", "email-sender")
    toPart := flag.String("t", "rz.om.stp@itsv.at", "email-recipient")
    subjectPart := flag.String("s", "no subject", "email-subject")
    bodyPart := flag.String("b", "(empty)", "email-body")
    attachPart := flag.String("a", "none", "email-attachments")
    flag.Parse()
    if *subjectPart == "no subject" || *bodyPart == "(empty)" {
        //usage(5)
        help()
        os.Exit(1)
    } else {
        fmt.Println("Sender: \t", *fromPart)
        fmt.Println("Recipient: \t", *toPart)
        fmt.Println("Subject: \t", *subjectPart)
        fmt.Println("Body: \t\t", *bodyPart)
        fmt.Println("Attachments: \t", *attachPart)
    }
    mailing()
}
