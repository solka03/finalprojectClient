# Pull base image.
FROM ubuntu:14.04

# tell the port number the container should expose
#EXPOSE 50055

EXPOSE 9090

RUN mkdir -p /go/src/app/uuid-client

ADD /bin /go/src/app/uuid-client

WORKDIR /go/src/app/uuid-client

ENTRYPOINT ./client


