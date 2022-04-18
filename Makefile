all: docs

.PHONY: docs
docs:
	docker run --rm -v "$$(pwd):/helm-docs" -u $$(id -u) jnorwood/helm-docs:v1.7.0 -i /helm-docs/ci/.helmdocsignore

.PHONY: box-chart
box-chart:
	helm package ./box-chart/ --destination ./infra-repo/
	helm repo index ./infra-repo/

.PHONY: infra-repo
infra-repo:
	cp -r ./box-chart/ ./infra-repo
	docker build -t infra-repo:latest --progress=plain -f ./infra-repo/Dockerfile ./infra-repo