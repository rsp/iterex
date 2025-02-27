#!/bin/sh
files=`gofmt -l .`
if [ -z "$files" ]; then
  echo All go files formatted correctly
  exit
else
  echo Files need formatting: $files
  echo
  gofmt -d .
  exit 1
fi
