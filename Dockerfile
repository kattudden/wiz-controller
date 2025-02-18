FROM golang:alpine AS builder
WORKDIR /app
COPY /app /app
RUN apk update && apk add --no-cache ca-certificates tzdata && update-ca-certificates

# Create appuser
ENV USER=appuser
ENV UID=10001

# See https://stackoverflow.com/a/55757473/12429735RUN 
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"


RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app/wiz-controller

# Build Runtime Container.
FROM scratch
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

COPY --from=builder /app/wiz-controller /app/wiz-controller
COPY ./app/templates/* /app/templates/
COPY ./app/images/* /app/images/
COPY ./app/static/* /app/static/

# Use an unprivileged user.
USER appuser:appuser

EXPOSE 8080
WORKDIR /app
CMD ["/app/wiz-controller"]