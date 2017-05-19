# Tool Lending API
### Simple CRUD API Using Golang

## Overview
Simple JSON API to keep track of your tools and who has borrowed them (deadbeats!).

## Goal
This is a quick experimental project to explore the features and syntax of the Go language.

### Dependencies
**gorilla/mux** - Request router and dispatcher package (https://github.com/gorilla/mux)
**github.com/lib/pq** - Postgres driver implementation for database/sql (https://github.com/lib/pq)
**github.com/DavidHuie/gomigrate** - Migration package for postgres (https://github.com/DavidHuie/gomigrate)

### Routes
**GET /tools** - Index route for Tool resource
**GET /tools/{id}** - Show route for specific Tool
**PUT /tools/{id}** - Update route for modifying and existing Tool
**POST /tools** - Insert route for adding a new Tool
**DELETE /tools/{id}** - Delete route for destroying specific Tool

### Setup
1. install Go (http://sourabhbajaj.com/mac-setup/Go/README.html)
2. install Gorilla
> $ go get github.com/gorilla/mux
3. install pq
> $ go get github.com/lib/pq
4. create database
> $ createdb {database name}
5. run migrations and seeds
> $ go run cmd/migration/migrate.go
> $ go run cmd/seed/seed.go

### Execution
1. Start server
> $ go run *.go
2. Browse to http://localhost:8080
