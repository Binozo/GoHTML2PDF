FROM golang:alpine3.16 as builder

RUN mkdir build
WORKDIR build
COPY . .

RUN go mod download
WORKDIR cmd/main/
RUN CGO_ENABLED=0 GOOS=linux go build -a -o endpoint .
RUN mv endpoint /go/build/

FROM alpine:3.14
LABEL maintainer=binozoworks
LABEL org.opencontainers.image.source="https://github.com/Binozo/GoHtml2PDF"
LABEL org.opencontainers.image.description="An easy-to-deploy microservice which converts html to a pdf file using chromium"

COPY --from=builder /go/build/endpoint .

RUN apk add --no-cache --update \
    chromium


EXPOSE 7524
CMD ["./endpoint"]

