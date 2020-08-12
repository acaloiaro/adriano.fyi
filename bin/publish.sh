#!/bin/bash

rsync -e ssh -a --exclude=tags ~/git/adriano.fyi/ adriano@mycorp.adriano.fyi:/home/adriano/adriano.fyi/ --progress
