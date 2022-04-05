all: docs

.PHONY: docs
docs:
<<<<<<< HEAD
	docker run --rm -v "$$(pwd):/helm-docs" -u $$(id -u) jnorwood/helm-docs:latest -i /helm-docs/ci/.helmdocsignore
=======
	docker run --rm -v "$$(pwd):/helm-docs" -u $$(id -u) jnorwood/helm-docs:v1.7.0 -i /helm-docs/ci/.helmdocsignore
>>>>>>> origin/release-v1.8.0
