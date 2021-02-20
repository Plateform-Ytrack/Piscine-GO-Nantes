#!/bin/bash

curl -s https://ytrack.learn.ynov.com/assets/superhero/all.json | jq --argjson ID $HERO_ID '.[] | select(.id==$ID) .connections.relatives '| tr -d "\""
