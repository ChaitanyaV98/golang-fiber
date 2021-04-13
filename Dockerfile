FROM golang:alpine



WORKDIR /app
COPY ./ ./

RUN apk add --no-cache git
RUN apk add --no-cache sqlite-libs sqlite-dev
RUN apk add --no-cache build-base
RUN go-wrapper install -ldflags "-linkmode external -extldflags -static" 

CMD ["go", "run", "main.go"]



