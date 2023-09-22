# Build image
FROM golang:1.19-alpine as build

LABEL org.opencontainers.image.source=https://github.com/postfreely/postfreely
LABEL org.opencontainers.image.description="WriteFreely is a clean, minimalist publishing platform made for writers. Start a blog, share knowledge within your organization, or build a community around the shared act of writing."

RUN apk add --update nodejs npm make g++ git
RUN npm install -g less less-plugin-clean-css

RUN mkdir -p /go/src/github.com/postfreely/postfreely
WORKDIR /go/src/github.com/postfreely/postfreely

COPY . .

RUN cat ossl_legacy.cnf > /etc/ssl/openssl.cnf

ENV GO111MODULE=on
ENV NODE_OPTIONS=--openssl-legacy-provider

RUN make build \
  && make ui
RUN mkdir /stage && \
    cp -R /go/bin \
      /go/src/github.com/postfreely/postfreely/templates \
      /go/src/github.com/postfreely/postfreely/static \
      /go/src/github.com/postfreely/postfreely/pages \
      /go/src/github.com/postfreely/postfreely/keys \
      /go/src/github.com/postfreely/postfreely/cmd \
      /stage

# Final image
FROM alpine:3

RUN apk add --no-cache openssl ca-certificates
COPY --from=build --chown=daemon:daemon /stage /go

WORKDIR /go
VOLUME /go/keys
EXPOSE 8080
EXPOSE 7007
USER daemon

ENTRYPOINT ["cmd/postfreely/postfreely"]
