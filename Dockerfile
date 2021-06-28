# Use official golang alpine base image
FROM golang:1.16-alpine as builder

# Add source code
ADD . /app

# Set the build directory 
WORKDIR /app

# Build app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gologger .

######## Start a new stage from scratch #######
FROM alpine:latest  

RUN apk --no-cache add ca-certificates && apk add --update tzdata

# Set timezone
ENV TZ=America/Monterrey

# Clean APK cache
RUN rm -rf /var/cache/apk/*

# Set working directory
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/gologger .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ./gologger