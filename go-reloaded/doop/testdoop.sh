#!/bin/bash

go build
./doop 861 + 870
./doop 861 - 870
./doop 861 "*" 870
./doop 861 % 870
./doop hello + 1
./doop 1 p 1
./doop 1 / 0
./doop 1 % 0
./doop 1 "*" 1
./doop 9223372036854775807 + 1
./doop 9223372036854775809 - 3
./doop 9223372036854775809 "*" 3
