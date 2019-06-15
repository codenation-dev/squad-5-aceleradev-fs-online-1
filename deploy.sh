#!/bin/bash
docker-compose down
git pull
docker-compose build --force-rm
docker-compose -f docker-compose.yml -f secrets.yml up -d app swagger-ui static
