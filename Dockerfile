FROM golang:1.11-alpine3.7 as builder
RUN set -ex && \
	apk add libc-dev && \
	apk add gcc && \
    apk add git
WORKDIR /go/src/github.com/nullscc/jenkins_requirements
COPY . .
RUN go get -v && \
CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=${VERSION}" -o jenkins_requirements .

FROM alpine
LABEL maintainer="nullscc"
# See https://stackoverflow.com/questions/34729748/installed-go-binary-not-found-in-path-on-alpine-linux-docker
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

WORKDIR /root/
COPY --from=builder /go/src/github.com/nullscc/jenkins_requirements/jenkins_requirements ./

EXPOSE 8000
ENTRYPOINT ["./jenkins_requirements"]
