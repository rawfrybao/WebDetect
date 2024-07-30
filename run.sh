#!/bin/bash

rm -rf db-init
mkdir db-init

echo "Enter the webhook port:"
read webhookport

if [ -z "$webhookport" ]
then
    echo "Exit"
    exit 1
fi

# Ask for the password
echo "Enter your Telegram ID (Usually 10 numbers):"
read tgid

if [ -z "$tgid" ]
then
    echo "Exit"
    exit 1
fi

cp init.sql.example db-init/init.sql
sed -i "s/your_telegram_id/$tgid/g" db-init/init.sql

echo "Enter your Telegram bot token:"
read token

if [ -z "$token" ]
then
    echo "Exit"
    exit 1
fi

echo "Enter the domain name of this server:"
read domain

if [ -z "$domain" ]
then
    echo "Exit"
    exit 1
fi

echo "Enter your PostgreSQL database name:"
read dbname

if [ -z "$dbname" ]
then
    echo "Exit"
    exit 1
fi

echo "Enter your PostgreSQL username:"
read dbuser

if [ -z "$dbuser" ]
then
    echo "Exit"
    exit 1
fi

echo "Enter your PostgreSQL password:"
read dbpass

if [ -z "$dbpass" ]
then
    echo "Exit"
    exit 1
fi

echo "Enter your PostgreSQL port (Empty for default 5432):"
read dbport

if [ -z "$dbport" ]
then
    dbport=5432
fi

rm -rf docker-compose.yml

cp docker-compose.yml.example docker-compose.yml
sed -i "s/your_webhook_port/$webhookport/g" docker-compose.yml
sed -i "s/your_domain_name/$domain/g" docker-compose.yml
sed -i "s/your_bot_token/$token/g" docker-compose.yml
sed -i "s/your_db_name/$dbname/g" docker-compose.yml
sed -i "s/your_db_username/$dbuser/g" docker-compose.yml
sed -i "s/your_db_password/$dbpass/g" docker-compose.yml
sed -i "s/your_db_port/$dbport/g" docker-compose.yml

docker compose build
docker compose up -d
