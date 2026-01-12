#!/bin/bash

# Add Go bin to PATH
export PATH=$PATH:$(go env GOPATH)/bin

function primary_mockgen() {
  local actor="$1"
  mockgen -source="./src/core/interfaces/primary/${actor}.go" -destination="./src/apps/api/handlers/mocks/${actor}.go" -package="mocks"
}

function secondary_mockgen() {
  local actor="$1"
  mockgen -source="./src/core/interfaces/secondary/${actor}.go" -destination="./src/core/services/mocks/${base}.go" -package="mocks"
}

# Add all actors here
actors=(
    "account"
    "auth"
    "campus"
    "course"
    "file"
    "institution"
    "internship"
    "internshipLocation"
    "passwordReset"
    "resources"
    "session"
    "student"
    "timeRecord"
    "timeRecordStatus"
)

mkdir -p src/core/services/mocks src/apps/api/handlers/mocks

case "$1" in
    -primary)
        for actor in "${actors[@]}"; do
            primary_mockgen "$actor"
        done;;
    -secondary)
        for actor in "${actors[@]}"; do
            secondary_mockgen "$actor"
        done;;
    -all)
        for actor in "${actors[@]}"; do
            primary_mockgen "$actor"
            secondary_mockgen "$actor"
        done;;
    *)
        echo "Usage: ./mockgen.sh <-primary|-secondary|-all>";;
esac
