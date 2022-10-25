LIB=k8s.io/code-generator
check-generator-version:
	go list -m -versions $(LIB) | tr " " "\n"

VERSION=v0.21.14
install-generator:
	go get k8s.io/code-generator@$(VERSION)

install-apimachinery:
	go get k8s.io/apimachinery@$(VERSION)

install-client-go:
	go get k8s.io/client-go@$(VERSION)

generate:
	hack/update-codegen.sh
