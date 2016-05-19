#!/usr/bin/env bash

docker build --no-cache -t cxbyte/balancer balancer

sed "s/#TIMESTAMP#/$(date +%s)/g" test.ru/app/Dockerfile.tmpl > test.ru/app/Dockerfile
docker build -t cxbyte/test_ru_app test.ru/app
rm test.ru/app/Dockerfile

docker-compose up