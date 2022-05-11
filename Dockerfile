# build stage
FROM golang:1.14-alpine AS build-env
RUN apk --no-cache add build-base git mercurial gcc
ENV D=/myapp
WORKDIR $D
# cache dependencies
ADD go.mod $D
ADD go.sum $D
RUN go mod download
# now build
ADD . $D
RUN cd $D && go build -o bump ./cmd && cp bump /tmp/

# final stage
FROM alpine:3.11
RUN apk add --no-cache ca-certificates curl
WORKDIR /app
COPY --from=build-env /tmp/bump /script/bump
ENTRYPOINT ["/script/bump"]
