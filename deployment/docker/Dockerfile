FROM alpine:1.25-alpine AS build
WORKDIR /app 
COPY . . 
RUN go build -o test_project cmd/test_project/main.go 

FROM alpine:3
WORKDIR /app
COPY --from=build test_project .
CMD [./test_project]