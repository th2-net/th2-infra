all: docs

.PHONY: docs
docs:
	docker run --rm -v "$$(pwd):/helm-docs" -u $$(id -u) jnorwood/helm-docs:v1.8.1 -i /helm-docs/ci/.helmdocsignore

.PHONY: crd-gen
crd-gen:
		mkdir $$(pwd)/chart/crds/tmp/
		docker run -it --user $$(id -u):$$(id -u) \
		-v $$(pwd)/chart/crds/tmp:/opt/crd-docs-generator/output \
		-v $$(pwd)/ci/:/opt/crd-docs-generator/config \
		quay.io/giantswarm/crd-docs-generator:0.10.0 \
		--config /opt/crd-docs-generator/config/crd-gen.yaml
		cat $$(pwd)/chart/crds/tmp/* > ./chart/crds/README.md
		rm -rf $$(pwd)/chart/crds/tmp/

.PHONY: box-chart
box-chart:
	helm package ./box-chart/ --destination ./infra-repo/
	helm repo index ./infra-repo/

.PHONY: infra-repo
infra-repo:
	cp -r ./box-chart/ ./infra-repo
	docker build -t infra-repo:latest --progress=plain -f ./infra-repo/Dockerfile ./infra-repo

.PHONY: lint-infra-chart
lint-infra-chart:
	helm lint ./chart


