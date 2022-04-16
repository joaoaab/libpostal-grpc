
# Libpostal-grpc

This project is a grpc api that serves [libpostal](https://github.com/openvenues/libpostal)

Written in golang

## How to use

1. [install libpostal](https://github.com/openvenues/libpostal)
2. run the application with `make run`, this will run the main.go inside cmd/libpostal-grpc that starts and serves the api.
3. Use the proto definition to make grpc calls.

## Tests
You can run the test by running the following command

    make test

# Benchmarks
Benchmarks using [ghz](http://ghz.sh) using 50 concurrent requests on a computer with a AMD ryzen 5 3600 with 3200mhz memory, the application was running on a docker container under a wsl2 environment.

### Expand Parser
![image](https://user-images.githubusercontent.com/9062839/163675846-ad0aa00f-f8bf-4579-8ccd-84be0eeefb4a.png)


### Parse Address
![image](https://user-images.githubusercontent.com/9062839/163675818-096632b9-9e10-4605-b26a-351e0f3a6b76.png)


## Checklist
- [X] Add Parser endpoint
- [X] Add Expander endpoint
- [X] Write tests
- [X] Create benchmarks for routes.
- [X] Dockerfile for easy deployment.
- [ ] Add Caching for common used addresses
