all: docs

.PHONY: docs
docs:
	docker run --rm -v "$$(pwd):/helm-docs" -u $$(id -u) jnorwood/helm-docs:v1.7.0 -i /helm-docs/ci/.helmdocsignore

