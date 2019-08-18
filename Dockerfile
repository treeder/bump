# build stage
FROM golang:1.12-alpine AS build-env
RUN apk --no-cache add build-base git bzr mercurial gcc
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
FROM alpine:3.10
RUN apk add --no-cache ca-certificates curl
WORKDIR /app
COPY --from=build-env /tmp/bump /script/bump
ENTRYPOINT ["/script/bump"]
