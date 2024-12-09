#!/bin/bash

echo "\
+-------------------------------------------+
| Loading environment variables for test... |
+-------------------------------------------+\
"
source src/utils/tests/envs/.env.test
schema=$(echo $DATABASE_SCHEMA | sed "s/\r//")
user=$(echo $DATABASE_USER | sed "s/\r//")
password=$(echo $DATABASE_PASSWORD | sed "s/\r//")
host=$(echo $DATABASE_HOST | sed "s/\r//")
port=$(echo $DATABASE_PORT | sed "s/\r//")
name=$(echo $DATABASE_NAME | sed "s/\r//")
ssl_mode=$(echo $DATABASE_SSL_MODE | sed "s/\r//")
migrations_path=$(echo $DATABASE_MIGRATIONS_PATH | sed "s/\r//")
uri="$schema://$user:$password@$host:$port/$name?sslmode=$ssl_mode"

echo "\
+--------------------------------+
| Starting database for tests... |
+--------------------------------+\
"
docker compose -f ./src/utils/tests/docker/docker-compose.test.yml up -d database_test

echo "\
+-------------------------------------+
| Downloading project dependencies... |
+-------------------------------------+\
"
go mod tidy

echo "\
+--------------------------------------------------------+
| Waiting 5 seconds so that the database can initiate... |
+--------------------------------------------------------+\
"
echo -n "["
for (( i=0; i<20; i++ )); do
    echo -n "#"
    sleep 0.2
done
echo "]"
sleep 1

echo "\
+---------------------------------+
| Loading migrations for tests... |
+---------------------------------+\
"
migrate -path $migrations_path -database $uri up

echo "\
+------------------+
| Running tests... |
+------------------+\
"
go test ./...
test_result=$?

exit $test_result