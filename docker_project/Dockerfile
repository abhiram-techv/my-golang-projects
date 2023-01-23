#Build Stage

FROM golang:1.19.3-alpine AS BuildStage
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /docker_project


#Deploy Stage

FROM alpine:latest

WORKDIR /
COPY .apiConfig ./
COPY .env ./
COPY --from=BuildStage /docker_project /docker_project
EXPOSE 8090
ENTRYPOINT ["/docker_project"]