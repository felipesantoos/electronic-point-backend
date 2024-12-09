#!/bin/bash

function primary_mockgen() {
  local actor="$1"

  case "$actor" in
      resources | occurrence | accompaniment | request | student | professor)
          mockgen -source="./src/core/interfaces/primary/${actor}.go" -destination="./src/apps/api/handlers/mocks/${actor}.go" -package="mocks";;
        *)
          echo "Invalid actor. Available actors: <resources|occurrence|accompaniment|request|student|professor>"
          return 1;;
  esac
}

function generate_primary_mocks() {
  local actors=()

  while [[ $# -gt 0 ]]; do
      case $1 in
          -resources | -occurrence | -accompaniment | -request | -student | -professor)
              actors+=("${1#-}")
              shift;;
          *)
              echo "Invalid actor."
              return 1;;
      esac
  done

  for actor in "${actors[@]}"; do
      primary_mockgen "$actor"
  done
}

function secondary_mockgen() {
  local actor="$1"

  case "$actor" in
      occurrence | accompaniment | request | student | resources | attachmentOfRequest | requestMovement | professor | email)
          mockgen -source="./src/core/interfaces/secondary/${actor}.go" -destination="./src/core/services/mocks/${actor}.go" -package="mocks";;
        *)
          echo "Invalid actor. Available actors: <occurrence | accompaniment | request | student | resources | attachmentOfRequest | requestMovement | professor | email>"
          return 1;;
  esac
}

function generate_secondary_mocks() {
  local actors=()

  while [[ $# -gt 0 ]]; do
      case $1 in
          -occurrence | -accompaniment | -request | -student | -resources | -attachmentOfRequest | -requestMovement | -professor | -email)
              actors+=("${1#-}")
              shift;;
          *)
              echo "Invalid actor."
              return 1;;
      esac
  done

  for actor in "${actors[@]}"; do
      secondary_mockgen "$actor"
  done
}

case "$1" in
    -primary)
        shift
        if [[ $# -eq 0 ]]; then
            generate_primary_mocks -resources -occurrence -accompaniment -request -student -professor
        else
            generate_primary_mocks "$@"
        fi;;
    -secondary)
        shift
        if [[ $# -eq 0 ]]; then
            generate_secondary_mocks -resources -ocurrence -request -student -accompaniment -pedagogicalAccompaniment -attachmentOfRequest -requestMovement -professor -email
        else
            generate_secondary_mocks "$@"
        fi;;
    -all)
        generate_primary_mocks -resources -occurrence -accompaniment -request -student -professor
        generate_secondary_mocks -resources -occurrence -accompaniment -request -student -attachmentOfRequest -requestMovement -professor -email;;
    *)
        echo "Invalid option for mock generation. Possible arguments: -mockgen <-primary|-secondary|-all>";;
esac