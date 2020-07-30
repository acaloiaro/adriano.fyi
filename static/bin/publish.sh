#!/bin/bash

rm tags
jekyll build
rsync -e ssh -a _site/ root@mycorp.adriano.fyi:/var/www/adriano.fyi
