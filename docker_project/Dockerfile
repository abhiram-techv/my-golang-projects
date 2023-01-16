#Build Stage

FROM golang:1.19 AS BuildStage
WORKDIR /app
COPY go.mod ./
COPY .apiConfig ./
COPY *.go ./
RUN go mod download
RUN go build -o /docker_project
EXPOSE 8090
CMD [ "/docker_project" ]

#Deploy Stage

FROM alpine:latest
WORKDIR /app
COPY --from=BuildStage /docker_project /docker_multistage
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/docker_multistage"]