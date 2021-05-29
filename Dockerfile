FROM golang:1.16-alpine as builder
WORKDIR /myapp/

COPY . .
RUN go mod tidy

RUN CGO_ENABLED=0 go build -o /demo/myapp

FROM scratch
COPY --from=builder /demo/myapp /demo/myapp
ENTRYPOINT [ "/demo/myapp" ]
