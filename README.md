# mailer-go

An mailer tool for the CLI
written in go language.

## Usage:

~~~sh
 mailer -f sender -t recipient -s subject -b body/message [-a "attachments"]
~~~

  -f: sender (optional)
  -t: (recipients):  foo@server,bar@domain
  -s: subject
  -b: message-body
  -a: (attachments): one file, wildcards allowd, like filenam*


## Config-file:

requires: `$HOME/mailerconfig.toml`

* Example:

~~~sh
[default]
SMTPD = "localhost"
SENDER = "sueswe@localhost"
~~~

.
