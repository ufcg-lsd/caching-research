#!/bin/sh
FILE=$1
echo "tenant,irr" >> "irr.csv"
go run ./irr.go < $FILE >> "irr.csv"