#!/usr/bin/env bash

docker build --no-cache -t cxbyte/balancer balancer

sed "s/#TIMESTAMP#/$(date +%s)/g" #domain#/app/Dockerfile.tmpl > #domain#/app/Dockerfile
docker build -t cxbyte/#app#_app #domain#/app
rm #domain#/app/Dockerfile

docker-compose up