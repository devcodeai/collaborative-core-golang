FROM golang:1.20-alpine

# set backend-related environment 
ENV HOST=0.0.0.0 \
    PORT=3030 

# set database-related environment 
ENV MYSQL_USER=root \
    MYSQL_PASSWORD=password \
    MYSQL_HOST=mysql \
    MYSQL_PORT=3306 \
    MYSQL_DBNAME=devcode

# set default working dir as '/app'
WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod tidy
RUN go mod download && go mod verify

COPY . ./
RUN go build -o /collaborative-core-golang

# expose port 3030 for the service
EXPOSE 3030

CMD ["/collaborative-core-golang"]
