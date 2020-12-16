# weebinar

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing
purposes.

### How to Set up The App

1. Install Postgresql: [Postgresql](https://www.postgresql.org/download/)
2. Setup your own configuration (for database and google oauth2) on `weebinar/config/weebinar.yaml`
3. Run database migrations: refer below
4. Run Project by building binary or via Docker: refer below
5. Try the app by going to `localhost:5000`

### Build and Run Project

Running project by building binary:

```$xslt
make run
```
or via Docker:
```$xslt
make docker-run
```

### Migrations

When running in Local, you need to run the db-migrations to setup the app's database for your local machine.

1. Create your own database called `weebinar`
2. Go to directory `weebinar/migrations`
3. Change your database credentials at `main.go` file
4. Run `go run *.go init`
5. Run `go run *.go up`

## Directory Structure

This repository is organized in the following directory structure.

```
weebinar
|-- bin                                    # Contains binary of the built app
|-- cmd                                    # Contains executable codes and serves as entry point of the app
|   |-- gql                                # entry point of the app
|-- config                                 # Configuration files needed for deployment
|-- constant                               # Collections of constants file for each module
|-- internal                               # Go files in this folder represent the Big-Pictures and Contracts of the system
|   |-- app                                # Contains dependency injection of the app and other app's related configs
|   |   |-- config                         # Configuration struct for the app
|   |-- delivery                           # Delivery layer (Controller) of the app
|   |   |-- rest                           # REST API as delivery of the app
|   |   |-- <other_delivery_mechanisms>    # Other delivery mechanisms of the app (eg. GRPC, Console, Web, etc.)
|   |-- entity                             # Enterprise Data structures
|   |   |-- webinar.go                     # Data structure for Webinar
|   |   |-- <other_entities>.go            # Other data structures, preferrably 1 struct 1 file 
|   |-- repo                               # Implementations of Repository-pattern to data-sources
|   |   |-- webinar.go                     # Implementations of the webinar repositoriy with Postgres database
|   |   |-- <other_repos>.go               # Other Repositories implementations based on interfaces on folder internal.
|   |-- usecase                            # Usecases implementations for Application Business Logic (Model)
|   |   |-- teacher.go                     # Business logic for Teacher Usecase
|   |   |-- student.go                     # Business logic for Student Usecase
|   |-- repo.go                            # Interfaces / Contracts of all the repositories (Repository Pattern)
|   |-- usecase.go                         # Interfaces / Contracts of all the use-cases (Application Business Logic)
|-- migrations                             # Contains Database migration files or the system
|-- public                                 # Contains client-side related code (Views)

```

## Tech Stacks

- Go
- Postgresql
- Docker
