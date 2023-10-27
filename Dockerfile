FROM golang:1.21.1-alpine
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN cd cmd/ && CGO_ENABLED=0 GOOS=linux go build -o /monitora-site
CMD [ "/monitora-site" ]