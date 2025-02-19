# mailer-go

![build workflow](https://github.com/sueswe/mailer-go/actions/workflows/go.yml/badge.svg?event=push)

> An mailer tool for the CLI written in go language.

> [!IMPORTANT]
> The project is in a process of heavy refactoring.

## Usage:

~~~sh
 mailer [-f sender] -t recipient,recipient -s subject -m messagebody [-a "attachment_a,attachm*,attachment_c"]
~~~

  - -f: sender (optional)
  - -t: (recipients):  foo@server,bar@domain
  - -s: subject
  - -m: message-body
  - -a: (attachments): files, commaseparated, wildcards allowed, like filenam*


## Config-file:

* Position: `$HOME/.mailerconfig.toml`

* Examplecontent:

~~~sh
[default]
SMTPD = "localhost"
SENDER = "sueswe@localhost"
~~~

.
