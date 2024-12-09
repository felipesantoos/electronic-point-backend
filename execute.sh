#!/bin/bash

function select_environment() {
  local environment="$1"

  case "$environment" in
      -development)
          run_development_environment;;
      -test)
          run_test_environment;;
      -all)
          if run_test_environment; then
              run_development_environment
          else
              echo "Tests failed. Skipping development environment startup."
          fi;;
      *)
          echo "Invalid option for environment generation. Possible arguments: -environment <-development|-test|-all>"
          return 1;;
  esac
}

function run_development_environment() {
    chmod +x ./tools/executables/environment/development.sh
    ./tools/executables/environment/development.sh
}

function run_test_environment() {
    chmod +x ./tools/executables/environment/test.sh
    ./tools/executables/environment/test.sh
}

function run_mockgen() {
    chmod +x ./tools/executables/mock/mockgen.sh
    ./tools/executables/mock/mockgen.sh "$@"
}

case "$1" in
    -environment)
        shift
        select_environment "$@";;
    -mockgen)
        shift
        run_mockgen "$@";;
    *)
       echo "Invalid option. Possible arguments: ./execute.sh <-environment|-mockgen>";;
esac