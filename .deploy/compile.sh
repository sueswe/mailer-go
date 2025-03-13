#!/bin/bash

source ~/.profile


cd "$HOME"/compile/mailer-go || {
	echo "Status: $?"
	exit 4
}

git pull origin master|| {
	echo "Status: $?"
	exit 4
}

echo "------------------------------------"
env | grep PATH
env | grep LOADED
echo "------------------------------------"

go build mailer.go|| {
	echo "Status: $?"
	exit 4
}

GOOS=aix GOARCH=ppc64 go build -o mailer.aix|| {
	echo "Status: $?"
	exit 4
}

cd "$HOME"/temp/|| {
	echo "Status: $?"
	exit 4
}

rm -rf globaltools|| {
	echo "Status: $?"
	exit 4
}

git clone git@lvgom01.sozvers.at:repos/globaltools.git|| {
	echo "Status: $?"
	exit 4
}

cd globaltools|| {
	echo "Status: $?"
	exit 4
}

cp "$HOME"/compile/mailer-go/mailer .  || {
	echo "Status: $?"
	exit 4
}

cp "$HOME"/compile/mailer-go/mailer.aix .  || {
	echo "Status: $?"
	exit 4
}

git add . || {
	echo "Status: $?"
	exit 4
}

git commit -m "recompiled mailer "  || {
	echo "Status: $?"
	exit 4
}

git push origin master || {
	echo "Status: $?"
	exit 4
}

