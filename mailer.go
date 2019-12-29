package main

import (
    "fmt"
    "flag"
)

func main() {
    fromPart    := flag.String("from","rz.om.stp@itsv.at","email-sender")
    toPart      := flag.String("to", "rz.om.stp@itsv.at","email-recipient")
    subjectPart := flag.String("subject", "no subject", "email-subject")
    bodyPart    := flag.String("body", "(empty)" , "email-body")
    attachPart  := flag.String("attachment", "none", "email-attachments")
    flag.Parse()

    fmt.Println("Sender: \t", *fromPart)
    fmt.Println("Recipient: \t", *toPart)
    fmt.Println("Subject: \t", *subjectPart)
    fmt.Println("Body: \t\t", *bodyPart)
    fmt.Println("Attachments: \t", *attachPart)


}
