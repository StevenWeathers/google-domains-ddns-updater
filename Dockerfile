############################
# STEP 1 build executable binary
############################
FROM golang:alpine as builderGo
# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates
# Create appuser
RUN adduser -D -g '' appuser
RUN mkdir /static
COPY ./types.go $GOPATH/src/mypackage/myapp/
COPY ./datasrc.go $GOPATH/src/mypackage/myapp/
COPY ./handlers.go $GOPATH/src/mypackage/myapp/
COPY ./gddu.go $GOPATH/src/mypackage/myapp/

WORKDIR $GOPATH/src/mypackage/myapp/
# Fetch dependencies.
# Using go mod with go 1.11
#RUN GO111MODULE=on go mod download
# Using go get.
RUN go get -d -v
# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/gddu
############################
# STEP 2 build the vue.js ui
############################
FROM node:10.14.1-alpine as builderNode

RUN mkdir /vueapp
COPY ./public /vueapp/public
COPY ./src /vueapp/src
COPY ./package.json /vueapp
COPY ./package-lock.json /vueapp
COPY ./babel.config.js /vueapp
WORKDIR /vueapp
# install node packages
RUN npm set progress=false
RUN npm install
# Build the vue.js app
RUN npm run build
############################
# STEP 3 build a small image
############################
FROM scratch
# Import from builder.
COPY --from=builderGo /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builderGo /etc/passwd /etc/passwd
# Copy our static executable
COPY --from=builderGo /go/bin/gddu /go/bin/gddu
# Copy our static assets
COPY --from=builderNode /vueapp/dist /static
# Use an unprivileged user.
USER appuser

ENV CADENCE "@hourly"

# VOLUME ["/data"]

# Run the gddu binary.
ENTRYPOINT ["/go/bin/gddu"]