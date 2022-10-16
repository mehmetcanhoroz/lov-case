#!/bin/bash

# Stop script when ever error is found
set -e

fn=$1
from=$2
where=$3

RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m'

buildapp() {
  echo -e "${GREEN}Building $1...${NC}"
  env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/"$where" ./"$from"
}

buildapp "$1" "$2" "$3"