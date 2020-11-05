FROM xushikuan/alpine-build:2.0 AS builder
ENV PROJECT_REPO=DIV-01/real-time-forum
ENV APP_PATH=/go/src/${PROJECT_REPO}/
ENV GO111MODULE=on
WORKDIR ${APP_PATH}
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . ${APP_PATH}
RUN ls
RUN go build -a -ldflags "-linkmode external -extldflags '-static' -s -w" -o server ./cmd/server


# stage 2
FROM xushikuan/alpine-build:1.0
ENV PROJECT_REPO=DIV-01/real-time-forum
ENV APP_PATH=/go/src/${PROJECT_REPO}
WORKDIR /root/

COPY --from=builder ${APP_PATH}/ ./
RUN ls

ENV PORT 8080
EXPOSE 8080
RUN ls
CMD ["./server"]