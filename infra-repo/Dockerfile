FROM golang:alpine AS builder
WORKDIR /app/
COPY ./go.mod ./main.go .
RUN go build -v -o infra-repo .

FROM alpine:3.13 as plugin-builder
WORKDIR /app/repo/
COPY grafana-plugins.txt download-plugins.sh .
RUN ./download-plugins.sh

FROM alpine/helm as helm-builder
WORKDIR /app/repo
COPY ./box-chart ./box-chart
RUN helm package ./repo/box-chart/ --destination . && \
    helm repo index . && \
    cp index.yaml .

FROM alpine:3.13
COPY --from=builder /app/infra-repo /app/infra-repo
COPY ./dashboards /app/repo/dashboards
COPY --from=plugin-builder /app/repo/plugins/ /app/repo/plugins/
COPY --from=helm-builder /app/repo /app/repo
ENTRYPOINT ["/app/infra-repo", "-port=8080", "-path=/app/repo"]
