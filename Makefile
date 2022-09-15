all: docs

.PHONY: docs
docs:
	docker run --rm -v "$$(pwd):/helm-docs" -u $$(id -u) jnorwood/helm-docs:v1.8.1 -i /helm-docs/ci/.helmdocsignore

.PHONY: crd-gen
crd-gen:
	rm -rf $$(pwd)/docs/reference
	mkdir $$(pwd)/docs/reference
	docker run --rm -u $$(id -u) \
	-v $$(pwd)/docs/reference:/opt/crd-docs-generator/output \
	-v $$(pwd)/ci/:/opt/crd-docs-generator/config \
	quay.io/giantswarm/crd-docs-generator:0.10.0 \
	--config /opt/crd-docs-generator/config/crd-gen.yaml

.PHONY: box-chart
box-chart:
	helm package ./box-chart/ --destination ./infra-repo/
	helm repo index ./infra-repo/

.PHONY: infra-repo
infra-repo:
	cp -r ./box-chart/ ./infra-repo
	docker build -t infra-repo:latest --progress=plain -f ./infra-repo/Dockerfile ./infra-repo

.PHONY: test-infra-chart
test-infra-chart:
	helm lint ./chart
	helm unittest -3 ./chart --color -o ./test-results/results.xml -t JUnit
