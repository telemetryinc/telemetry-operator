FROM golang:1.24.9 AS builder

WORKDIR /workspace
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY main.go main.go
COPY api/ api/
COPY controller/ controller/

RUN CGO_ENABLED=0 go build -a -o telemetry-operator main.go

FROM registry.access.redhat.com/ubi9/ubi

ARG VERSION=unknown
LABEL name="telemetry-operator" \
      vendor="Telemetry Inc." \
      maintainer="Telemetry Inc." \
      version=${VERSION} \
      release="1" \
      summary="Telemetry Operator." \
      description="Telemetry Operator container image."

COPY LICENSE /licenses/LICENSE

WORKDIR /
COPY --from=builder /workspace/telemetry-operator /usr/bin/telemetry-operator
USER 65534:65534
ENTRYPOINT ["telemetry-operator"]
