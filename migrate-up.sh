#!/bin/bash
eval $(cat ./.env)
cd ./db
go run . postgres "user=$DB_USER dbname=$DB_NAME sslmode=$DB_SSL_MODE port=$DB_PORT" up
cd ../