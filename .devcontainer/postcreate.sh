#!/bin/sh

# start postgres
service postgresql start

# create user on postgres db
psql -U postgres -c "CREATE USER $POSTGRES_USER WITH PASSWORD '$POSTGRES_PASSWORD';"
psql -U postgres -c "ALTER USER $POSTGRES_USER WITH SUPERUSER;"
psql -U postgres -c "CREATE DATABASE $POSTGRES_DB OWNER $POSTGRES_USER;"
psql -U postgres -c "GRANT ALL PRIVILEGES ON DATABASE $POSTGRES_DB TO $POSTGRES_USER;"

# Run schema.sql
psql -U $POSTGRES_USER -d $POSTGRES_DB -a -f ./schema.sql

# Install cosmtrek/air
go install github.com/air-verse/air@latest
