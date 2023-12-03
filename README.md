# GoRemoteAccess

This is a telegram bot developed as a personal assistant.

At the moment, the public functionality is limited to displaying the schedule of classes for NUST MISIS groups with the possibility to save your group in the database for quick access to your personal schedule.

The project uses the open API of misis.ru to obtain raw data.

![image](https://github.com/LegendLex/GoRemoteAccess/blob/main/example.jpg)

# Getting Started

## Enviroment variables

BOT_TOKEN - The only necessary variable to run

DB_SWITCH - "on" to use DB. The following variables are its parameters.

HOST

PORT

USER

PASSWORD

DBNAME

SSLMODE

CREATE_TABLE - "yes" to create table in database. 

## Docker

To run a project as a Docker image:

1) Build an image from the source directory
```
docker build --tag docker-go-remote .
```

2) Run with suitable environment variables
```
docker run -e TOKEN=your_bot_token -d docker-go-remote
```
