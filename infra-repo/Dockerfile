FROM golang:alpine AS builder
WORKDIR /app/
COPY ./go.mod ./main.go .
RUN go build -v -o infra-repo .

FROM alpine/helm as helm-builder
WORKDIR /app/
COPY ./box-chart ./box-chart
RUN helm package ./box-chart/ --destination ./repo && \
	helm repo index . && \
    cp index.yaml ./repo

FROM alpine:3.13
COPY --from=builder /app/infra-repo /app/infra-repo
COPY ./dashboards /app/repo/dashboards
COPY ./plugins /app/repo/plugins
COPY --from=helm-builder /app/repo /app/repo
ENTRYPOINT ["/app/infra-repo", "-port=8080", "-path=/app/repo"]
