#!/usr/bin/env bash

mkcert -install; mkcert -cert-file="development.crt" -key-file="development.key" localhost;

if [ $? -eq 127 ];
  then echo "Install mkcert https://github.com/FiloSottile/mkcert";
fi