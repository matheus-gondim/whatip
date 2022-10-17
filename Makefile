SHELL := /bin/bash

build:
	go build -o whatip
	sudo chmod +x whatip

move_bin:
	sudo mv ./whatip /usr/local/bin/ 
	