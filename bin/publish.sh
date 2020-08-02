#!/bin/bash

rm tags
rsync -e ssh -a ~/git/adriano.fyi/ adriano@mycorp.adriano.fyi:/home/adriano/adriano.fyi/
