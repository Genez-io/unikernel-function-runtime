FROM golang:latest AS go_build
WORKDIR /app
ADD . /app/
RUN go mod tidy
RUN go build -o java_builder

# Place OSv path here
ENV OSV_PATH="/home/cgeorge/genez/unikernel-function-runtime/builders/builder_deps/osv"

FROM eclipse-temurin:latest
WORKDIR /

RUN apt update && apt install -y python3 
COPY --from=go_build /app/java_builder /

CMD [ "./java_builder" ]