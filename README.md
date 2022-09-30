# Golang Project Template

This repository is a sample backend project template build with Go. We use this boilerplate for our internal development purposes. Feel free to open issue / PR on this repository

## Project Structure
    .
    ├── .github                 # Github config related
    ├── cmd                     # Runtime entrypoint
    ├── docker                  # Docker config for each env
    ├── pkg                     # Main package. All app & business logic belong here
    ├── tmp                     # Temporary directory
    ├── .env.example            # Where we put environment variables
    ├── CHANGELOG.md
    ├── CONTRIBUTING.md
    └── README.md

> Use short lowercase names at least for the top-level files and folders except
> `LICENSE`, `README.md`

## How to run

#### Prerequisite

- docker >= `20.10.12, build e91ed57` or LTS version
- docker-compose >= `1.29.2, build 5becea4c` or LTS version

If this is your first-time on this project. Please execute this command to prepare all project dependency
```bash
./run.sh setup
```

#### Prepare your .env file
Please copy `.env.example` into `.env` and change the value inside that need to be adjusted.
If you run under docker change `PG_HOST=postgres`. But if you run under you host, change `PG_HOST=localhost`

If everything going well, let's start the server
``` bash
# run under docker
docker-compose up

# run under host with hot-reload
./run.sh hot-serve
```

or Please open `run.sh` to see other supported command 
