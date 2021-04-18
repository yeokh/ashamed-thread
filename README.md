Golang CRUD Booster

#### Running Locally

First, install the dependencies
`go get`

The web app requires Postgres DB. If you have docker installed, you can run the
following command to start a new Postgres container

`docker run -e POSTGRESQL_ADMIN_PASSWORD=mysecretpassword -d -p 5432:5432 centos/postgresql-96-centos7`

In this example, the database user is `postgres`, the password is `mysecretpassword` and the database is `postgres`

You can then start the application like this:

`go build && ./golang-http-crud`

Then go to http://localhost:8080
