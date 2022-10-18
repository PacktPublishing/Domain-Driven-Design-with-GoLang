# builder image
FROM golang:1.18 as builder
RUN mkdir /build
ADD * /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -o partnerships .


# generate clean, final image for end users
FROM alpine
COPY --from=builder /build/partnerships .

# executable
ENTRYPOINT [ "./partnerships" ]
