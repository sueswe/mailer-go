# mailer-go

> [!IMPORTANT]
> The project is in a process of heavy refactoring.

An mailer tool for the CLI
written in go language.

![build workflow](https://github.com/sueswe/mailer-go/actions/workflows/go.yml/badge.svg?event=push)

## Usage:

~~~sh
 mailer -f sender -t recipient -s subject -b body/message [-a "attachment_a,attachm*,attachment_c"]
~~~

  - -f: sender (optional)
  - -t: (recipients):  foo@server,bar@domain
  - -s: subject
  - -b: message-body
  - -a: (attachments): files, commaseperated, wildcards allowed, like filenam*


## Config-file:

requires: `$HOME/mailerconfig.toml`

* Example:

~~~sh
[default]
SMTPD = "localhost"
SENDER = "sueswe@localhost"
~~~

.
