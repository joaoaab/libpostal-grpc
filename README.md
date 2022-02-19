# Libpostal-grpc

This project is a grpc api that serves [libpostal](https://github.com/openvenues/libpostal)

Written in golang

## How to use

1. [install libpostal](https://github.com/openvenues/libpostal)
2. run this application with go run . inside the /src folder, it should run main.go that serves the api
3. Use the proto to make grpc calls.

## Checklist
- [X] Add Parser endpoint
- [ ] Add Expander endpoint
- [ ] Write tests
- [ ] Dockerfile for easy deployment