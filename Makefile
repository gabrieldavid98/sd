idl:
	docker run --volume "$(pwd):/workspace" --workdir /workspace bufbuild/buf lint
nodo: 
	go run ./node/main.go