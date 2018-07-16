FROM golang:1.10.3-alpine3.7 as builder

LABEL maintainer="antonputra@icloud.com"

RUN apk add --update git make

WORKDIR /go/src/conductor

COPY . .

RUN make

FROM alpine:3.7

RUN addgroup -S conductor && \
    adduser -S -G conductor conductor

COPY --from=builder /go/src/conductor/conductor /home/conductor/

RUN chown -R conductor:conductor /home/conductor

USER conductor

ENTRYPOINT ["/home/conductor/conductor"]