# Libpostal-grpc

This project is a grpc api that serves [libpostal](https://github.com/openvenues/libpostal)

Written in golang

## How to use

1. [install libpostal](https://github.com/openvenues/libpostal)
2. run this application with go run . inside the /src folder, it should run main.go that serves the api
3. Use the proto to make grpc calls.

## Tests
`run go test ./... -v`

# Benchmarks
Benchmarks using [ghz](ghz.sh) on a computer with a AMD ryzen 5 3600 with 3200mhz memory, the application was running on a docker container under a wsl2 environment.

### Expand Parser
Summary:<br />
&nbsp;&nbsp;  Count:        10000 <br />
&nbsp;&nbsp;  Total:        2.03 s <br />
&nbsp;&nbsp;  Slowest:      23.26 ms <br />
&nbsp;&nbsp;  Fastest:      0.73 ms <br />
&nbsp;&nbsp;  Average:      9.42 ms <br />
&nbsp;&nbsp;  Requests/sec: 4917.91 <br />

Latency distribution:<br />
&nbsp;&nbsp;  10 % in 7.39 ms <br />
&nbsp;&nbsp;  25 % in 8.35 ms <br />
&nbsp;&nbsp;  50 % in 9.33 ms <br />
&nbsp;&nbsp;  75 % in 10.45 ms <br />
&nbsp;&nbsp;  90 % in 11.67 ms <br />
&nbsp;&nbsp;  95 % in 12.52 ms <br />
&nbsp;&nbsp;  99 % in 14.94 ms <br />

Status code distribution: <br />
&nbsp;&nbsp;  [OK]   10000 responses

### Parse Address
Summary:
&nbsp;&nbsp;  Count:        10000 <br />
&nbsp;&nbsp;  Total:        1.23 s <br />
&nbsp;&nbsp;  Slowest:      17.63 ms <br />
&nbsp;&nbsp;  Fastest:      0.49 ms <br />
&nbsp;&nbsp;  Average:      3.78 ms <br />
&nbsp;&nbsp;  Requests/sec: 8108.28 <br />

Latency distribution:
&nbsp;&nbsp;  10 % in 1.62 ms <br />
&nbsp;&nbsp;  25 % in 2.31 ms <br />
&nbsp;&nbsp;  50 % in 3.42 ms <br />
&nbsp;&nbsp;  75 % in 4.78 ms <br />
&nbsp;&nbsp;  90 % in 6.25 ms <br />
&nbsp;&nbsp;  95 % in 7.53 ms <br />
&nbsp;&nbsp;  99 % in 10.88 ms <br />

Status code distribution: <br />
&nbsp;&nbsp;  [OK]   10000 responses

## Checklist
- [X] Add Parser endpoint
- [X] Add Expander endpoint
- [X] Write tests
- [X] Create benchmarks for routes.
- [X] Dockerfile for easy deployment