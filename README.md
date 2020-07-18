# go-starter
A starter template for a golang project based on [golang project layout](https://github.com/golang-standards/project-layout).

## How to build
To build this project, run:
```bash
make tools
```
this will install and build all the tools.
Then run:
```bash
make generate
```
this will generate all the gRPC and protocol buffers files.

## How to run
To run all the executables, use:
```bash
make run
```
it uses [goreman](https://github.com/mattn/goreman) to run all the services defined in the `configs/Procfile`

