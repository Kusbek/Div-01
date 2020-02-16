FROM golang:1.13 as builder
ENV PROJECT_REPO=ascii_web
ENV APP_PATH=/go/src/${PROJECT_REPO}/
WORKDIR ${APP_PATH}
COPY . ${APP_PATH}
RUN go build -o ascii_web
EXPOSE 8080
# CMD ["tail", "-f","/dev/null"]
ENTRYPOINT ["./ascii_web"]