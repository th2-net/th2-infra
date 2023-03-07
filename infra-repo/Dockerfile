FROM golang:alpine AS builder
WORKDIR /app/
COPY ./go.mod ./main.go .
RUN go build -v -o infra-repo .

FROM alpine:3.13 as plugin-builder
WORKDIR /app/repo/
COPY grafana-plugins.txt download-plugins.sh .
RUN ./download-plugins.sh

FROM alpine/helm as helm-builder
WORKDIR /app/repo/
COPY ./chart ./th2-infra
COPY ./box-chart ./box-chart
RUN helm package ./th2-infra --destination . && helm package ./box-chart/ --destination . && \
    helm repo index .

FROM python:3
WORKDIR /app/repo/
COPY requirements.txt ./
RUN mkdir -p python-dependencies && \
    pip install --no-cache-dir --upgrade pip && \
    pip download -r requirements.txt -d ./python-dependencies

FROM alpine:3.13
COPY --from=builder /app/infra-repo /app/infra-repo
COPY ./entry-page/ /app/repo/
COPY ./dashboards /app/repo/dashboards
COPY --from=plugin-builder /app/repo/plugins/ /app/repo/plugins/
COPY --from=helm-builder /app/repo /app/repo
ENTRYPOINT ["/app/infra-repo", "-port=8080", "-path=/app/repo"]
