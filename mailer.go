package main

import (
    "flag"
    "fmt"
    "os"
)

func usage(a int) {
    fmt.Println("Usage:", a)
}

func help() {
    fmt.Println("Help:")
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
}
