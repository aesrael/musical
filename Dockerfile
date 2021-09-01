# Dockerfile References: https://docs.docker.com/engine/reference/builder/

FROM golang:1.16-alpine

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache bash git make curl ffmpeg py3-pip gcc g++ python3-dev

LABEL maintainer="Israel A."

WORKDIR /app

# install the spotify dl executable
RUN pip3 install spotdl

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8999

CMD ["make", "run"]
