FROM golang:1.22.2-alpine3.19 AS builder

ENV APP_NAME=lucky

WORKDIR /app
COPY . /app
RUN go mod download
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -trimpath -mod=readonly -a -installsuffix cgo -v -o ${APP_NAME} main.go

FROM alpine:3.19
RUN apk --no-cache add ca-certificates bash curl

ENV APP_NAME=lucky
ENV WORK_DIR=/app

WORKDIR ${WORK_DIR}

COPY --from=builder /app/deployments/docker/*.sh .
COPY --from=builder /app/conf ./conf
COPY --from=builder /app/${APP_NAME} .
RUN ls -la

EXPOSE 8080
ENTRYPOINT ["sh", "-c", "./entrypoint.sh"]
