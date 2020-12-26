
######################################
#
# Add proto gen stage
#
######################################

######################################
#
# Add wire gen stage
#
######################################

######################################
#
# Compliation build stage
#
######################################

FROM golang:1.13-alpine3.11 as builder

RUN apk update && apk add --no-cache git

RUN mkdir /src
WORKDIR /src

RUN apk add --update --no-cache --repository https://dl-4.alpinelinux.org/alpine/latest-stable/community/ make git

ADD . /src

RUN make deps

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

RUN pwd

RUN ls


FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /src/main .
COPY --from=builder /src/.env .

# Expose port 8080 to the outside world
EXPOSE 8080

#Command to run the executable
CMD ["./main"]
