FROM --platform=$BUILDPLATFORM golang:1.25-alpine AS build
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG TARGETOS TARGETARCH
ENV CGO_ENABLED=0
RUN GOOS=$TARGETOS GOARCH=$TARGETARCH go build -ldflags="-s -w" -o /out/healthcheck ./...

FROM scratch
COPY --from=build /out/healthcheck /healthcheck
EXPOSE 8080
ENTRYPOINT ["/healthcheck"]
