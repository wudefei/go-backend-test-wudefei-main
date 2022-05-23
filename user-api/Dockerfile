FROM XXXXX AS builder

ARG CGO=0
ENV CGO_ENABLED=${CGO}
ENV GOARCH=amd64
ENV GOOS=linux
ENV GO111MODULE=on

WORKDIR /build/

COPY . /build/
RUN set -ex && \
    cp -f config/conf.yaml conf.yaml && \
    CGO_ENABLED=${CGO_ENABLED} GOARCH=${GOARCH} GOOS=${GOOS} go build -a -o user-api -i cmd/main.go

FROM XXXXX

WORKDIR /work

COPY --from=builder /build/user-api .
COPY --from=builder /build/conf.yaml ./config/

VOLUME ["/work/logs"]
EXPOSE 8080
ENTRYPOINT ["./user-api", "-conf=config/conf.yaml"]
