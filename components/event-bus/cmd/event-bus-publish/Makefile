NAME = kyma-project/event-bus-publish
VERSION = 0.1.0
REGISTRY = eu.gcr.io

.PHONY: all clean build tag push

all: clean build docker-build tag

clean:
	rm -rf bin/
	rm -rf docker/

build: vet test compile

compile:
	go build -o bin/event-bus-publish

test:
	go test github.com/kyma-project/kyma/components/event-bus/cmd/event-bus-publish -v

docker-build:
	./dockerBuild.sh  $(NAME) $(VERSION)

tag:
	docker tag $(NAME):$(VERSION) $(REGISTRY)/$(NAME):$(VERSION)

push:
	docker push $(REGISTRY)/$(NAME):$(VERSION)

vet:
	go list ../../... | grep -v generated | xargs go vet
