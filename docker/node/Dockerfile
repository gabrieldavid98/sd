########################## BUILD ##############################
FROM golang:latest AS build-stage

WORKDIR /app/idl

COPY ./idl/go.mod ./idl/go.sum ./
RUN go mod download

COPY ./idl ./

WORKDIR /app/node

COPY ./node/go.mod ./node/go.sum ./
RUN go mod download

COPY ./node ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /node

########################## RELEASE ##############################

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /node /node

EXPOSE 7001

USER nonroot:nonroot

ENTRYPOINT ["/node"]