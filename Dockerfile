# Builder
FROM golang:1.18.0-stretch as builder

RUN apt --yes --force-yes update && apt --yes --force-yes upgrade && \
    apt --yes --force-yes install git \
    make openssh-client

WORKDIR /app

COPY . .

RUN make eform

# Distribution
FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app /eform

WORKDIR /eform

EXPOSE 9090

COPY --from=builder /app/eform /app

CMD /app/eform http