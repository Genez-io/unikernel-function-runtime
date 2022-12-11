FROM golang:latest AS go_build
WORKDIR /app
ADD . /app/
RUN go mod tidy
RUN go build -o java_builder

FROM eclipse-temurin:latest
WORKDIR /
# Add unikernel files
ADD nanos /nanos
# copy kotlin compiler
COPY ./kotlin/kotlinc ./kotlinc
# grab builder executable
COPY --from=go_build /app/java_builder /

# temp print jdeps output
CMD [ "./java_builder" ]