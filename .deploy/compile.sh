#!/bin/bash

function Returncode_check {
	time "$@"
	local status=$?
	if ((status != 0)); then
		echo "error with $1" >&2
		exit 5
	fi
	return $status	
}

Returncode_check cd "$HOME"/compile/mailer-go
echo "Status: $?"
Returncode_check git pull origin master
echo "Status: $?"
echo ""
 env | grep PATH
 env | grep LOADED
echo ""
Returncode_check go build mailer.go
echo "Status: $?"
Returncode_check GOOS=aix GOARCH=ppc64 go build -o mailer.aix
echo "Status: $?"
Returncode_check cd "$HOME"/temp/
echo "Status: $?"
Returncode_check rm -rf globaltools
echo "Status: $?"
Returncode_check git clone git@lvgom01.sozvers.at:repos/globaltools.git
echo "Status: $?"
Returncode_check cd globaltools
echo "Status: $?"
Returncode_check cp "$HOME"/compile/mailer-go/mailer .
echo "Status: $?"
Returncode_check cp "$HOME"/compile/mailer-go/mailer.aix .
echo "Status: $?"
Returncode_check git add .
echo "Status: $?"
Returncode_check git commit -m "recompiled mailer "
echo "Status: $?"
Returncode_check git push origin master
echo "Status: $?"
