FROM ubuntu
COPY main /app/main
WORKDIR /app
EXPOSE 8080
CMD ["./main"]
