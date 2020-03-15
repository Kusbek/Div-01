#!/bin/bash
for i in {1..7}
do
   go run main.go "example0$i.txt"
   printf "\n\n\n\n\n\n"
done
for i in {0..1}
do
   go run main.go "badexample0$i.txt"
   printf "\n\n\n\n\n\n"
done
