FROM golang:1.23.4-alpine AS builder

WORKDIR /app

COPY . .

RUN apk add make

RUN make build

FROM scratch AS runner

COPY --from=builder /app/bin/main.out /bin/main.out

ENTRYPOINT ["/bin/main.out"]
