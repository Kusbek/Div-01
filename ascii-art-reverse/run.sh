#!/bin/bash
for i in {0..7}
do
   go run main.go --reverse="example0$i.txt"
   printf "\n"
done


