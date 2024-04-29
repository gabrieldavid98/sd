########################## BUILD ##############################
FROM golang:latest AS build-stage

WORKDIR /app/idl

COPY ./idl/go.mod ./idl/go.sum ./
RUN go mod download

COPY ./idl ./

WORKDIR /app/master

COPY ./master/go.mod ./master/go.sum ./
RUN go mod download

COPY ./master ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /master

########################## RELEASE ##############################

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /master /master

EXPOSE 7000
EXPOSE 7070

USER nonroot:nonroot

ENTRYPOINT ["/master"]