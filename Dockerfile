FROM library/golang:1.8.3-alpine

# install curl which is required by healthcheck. without curl,
# the app won't be able to deploy
RUN apk add --no-cache curl

ENV PORT 80

RUN mkdir -p /opt/app 
WORKDIR /opt/app

# Copy files
COPY goserver.go /opt/app
COPY start_app /opt/app
RUN cd /opt/app 
RUN go build goserver.go

ENTRYPOINT ["/opt/app/start_app"]
