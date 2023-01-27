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

func help() {
    color.Yellow("Usage: ")
    fmt.Println("mailer [-f sender] [-t recipient,recipient] -s subject -b body/message [-a attachments] ")
    fmt.Print("\nDefault sender and recipient is: ")
    color.Cyan("rz.om.stp@itsv.at")
}

func mailer_single(from , to, subject, body , file string) {
    m := gomail.NewMessage()
    m.SetHeader("From", from)
    m.SetHeader("To", to)
    m.SetHeader("Subject", subject)
    m.SetBody("text/plain", body)

    d := gomail.Dialer{Host: "viruswall.sozvers.at", Port: 25}
    if err := d.DialAndSend(m); err != nil {
        color.Red("\nWhoops, that didn't work, pal!")
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
        color.Red("I'm missing something, take a look at:")
        help()
        print("\n");
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
