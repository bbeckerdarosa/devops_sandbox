FROM golang:1.12
RUN go get -u github.com/gorilla/mux
EXPOSE 8080
WORKDIR /app/
COPY ./calculator.go /app
CMD ["go", "run", "calculator.go"] 