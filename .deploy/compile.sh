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

echo ""
echo "compiling: go build mailer.go"
go build mailer.go || {
	echo "Status: $?"
	exit 4
}

echo ""
echo "compiling: GOOS=aix GOARCH=ppc64 go build -o mailer.aix"
GOOS=aix GOARCH=ppc64 go build -o mailer.aix || {
	echo "Status: $?"
	exit 4
}

echo ""

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

cp -v "$HOME"/compile/mailer-go/mailer .  || {
	echo "Status: $?"
	exit 4
}

cp -v "$HOME"/compile/mailer-go/mailer.aix .  || {
	echo "Status: $?"
	exit 4
}

echo "git add ..."
git add . || {
	echo "Status: $?"
	exit 4
}

echo "git commit ..."
git commit -m "recompiled mailer "  || {
	echo "Status: $?"
	exit 4
}

echo "git push ..."
git push origin master || {
	echo "Status: $?"
	exit 4
}

echo ""
echo "DONE!"
