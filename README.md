# Go Shop

Simple online shop simulation with Golang and html templates. Nothing too fancy!

## Built with
- Go
- HTML
- Bootstrap
- PostgreSQL

## Getting started

### Prerequisites

- [Go](https://go.dev/doc/install) (check installation guide in docs)
- [PostgreSQL](https://www.postgresql.org/)

### Installation
1. Clone repo

    `git clone https://github.com/CylonSam/go-shop.git`

2. Configure database (PostgresSQL)
   
   You need to create a database and a user with access to it, then you have to fill the database information inside the variable `connection` in the `db/db.go` file.
  
## Usage
To run the project, just type `go run main.go` in your terminal inside the project root. Then you can access the shop at http://localhost:4000 .

The shop allows users to add, edit, delete and view products stored inside the database.
