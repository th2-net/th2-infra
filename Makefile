all: docs

.PHONY: docs
docs:
	docker run --rm -v "$$(pwd):/helm-docs" -u $$(id -u) jnorwood/helm-docs:v1.8.1 -i /helm-docs/ci/.helmdocsignore
