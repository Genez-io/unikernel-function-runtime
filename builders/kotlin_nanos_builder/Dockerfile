FROM golang:latest AS go_build
WORKDIR /app
ADD . /app/
RUN go mod tidy
RUN go build -o java_builder

FROM eclipse-temurin:latest
WORKDIR /
COPY --from=go_build /app/java_builder /

CMD [ "./java_builder" ]