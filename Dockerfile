FROM golang:1.17.7-bullseye

RUN apt-get update && apt-get install -y \
    autoconf automake build-essential curl git libsnappy-dev libtool pkg-config

# clone libpostal
RUN git clone https://github.com/openvenues/libpostal /code/libpostal

COPY ./*.sh /code/libpostal/

WORKDIR /code/libpostal
# install libpostal
RUN ./build_libpostal.sh

WORKDIR $GOPATH/src/github.com/joaoaab/libpostal-grpc

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 50051
ENTRYPOINT ["libpostal-grpc"]

