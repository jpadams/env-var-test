#!/bin/bash

# Check if .env file exists
if [ ! -f .env ]; then
  echo ".env file not found!"
  exit 1
fi

# Read .env file and export variables only for the duration of this script
while IFS='=' read -r key value; do
  # Skip comments and empty lines
  if [[ ! "$key" =~ ^# && -n "$key" ]]; then
    export "$key=$value"
  fi
done < .env

dagger call test \
--env1 $ENV1 \
--env2 $ENV2 \
--env3 $ENV3 \
--env4 $ENV4 \
--env5 $ENV5 \
--env6 $ENV6 \
--env7 $ENV7 \
--env8 $ENV8 \
terminal
