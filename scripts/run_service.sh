#!/bin/bash

# read the env variables from the env.all file
source ./configs/env.all
HELLO_SERVICE_PORT=$HELLO_SERVICE_PORT go run "$1"