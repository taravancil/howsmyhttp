#!/usr/bin/env bash

glyphhanger http://localhost:4000 --spider --family='Nunito' --subset='src/fonts/*.ttf' --formats=woff,woff2

mv src/fonts/*-subset*.woff* static/fonts

printf "\n→ Copy output unicode range to update the existing unicode range in src/fonts.css"
