FROM golang:1.19
WORKDIR /app
COPY go.mod ./
COPY .apiConfig ./
COPY *.go ./
RUN go mod download
RUN go build -o /docker_project
EXPOSE 8090
CMD [ "/docker_project" ]
