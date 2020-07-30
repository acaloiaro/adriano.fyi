#!/bin/bash

rm tags
hugo
rsync -e ssh -a public/ adriano@mycorp.adriano.fyi:/var/www/adriano.fyi
