.PHONY: help
help: # Display this help
	@awk 'BEGIN{FS=":.*#";printf "Usage:\n  make <target>\n\nTargets:\n"}/^[a-zA-Z_-]+:.*?#/{printf"  %-10s %s\n",$$1,$$2}' $(MAKEFILE_LIST)

.PHONY: abigen
abigen: # Generate go files
	rm -rf artifacts/
	npm run compile
	npm run extract-abi
	docker build -f Dockerfile.abigen -t extract-abi .
	rm -rf pkg/*
	CONTAINER=`docker create extract-abi --name extract-abi`; \
	docker cp $$CONTAINER:/bridge pkg/bridge; \
	docker cp $$CONTAINER:/interfaces pkg/interfaces; \
	docker cp $$CONTAINER:/test pkg/test; \
	docker rm -v $$CONTAINER

.PHONY: init
init:
	cp .env.example .env
	cp contracts-<chain-id>.json contracts-5.json
	touch contracts-<chain-id>.json contracts-80001.json

.PHONY: deploy
deploy:
	npx hardhat --network mumbai run scripts/deploy-bridge.ts
	npx hardhat --network goerli run scripts/deploy-bridge.ts

.PHONY: deploy-receiver
deploy-receiver:
	npx hardhat --network mumbai run scripts/deploy-test-receiver.ts
