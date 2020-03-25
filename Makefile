.PHONY: build
build:
	env GOOS=linux go build -o ./bin/demo ./cmd/demo/...

.PHONY: deploy
deploy: build
	cp ./bin/demo ./roles/demo/files/demo
	ansible-playbook -i ./hosts demo.yml
