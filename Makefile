# Copyright 2019 Shanghai JingDuo Information Technology co., Ltd.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

.PHONY: all

IMAGE_REPOSITORY_URL = reg.kpaas.io/kpaas
BUILD_NUMBER = $(shell git rev-parse --short HEAD)
BUILD_TAG := $(shell date +%Y%m%d%H%M%S)

ifneq ($(shell uname), Darwin)
	EXTLDFLAGS = -extldflags "-static" $(null)
else
	EXTLDFLAGS =
endif

all: build_service_cross

.PHONY: service_doc

SED_PATH := $(shell which sed)
service_doc:
	# swaggo need larger than 1.6
	swag init -g cli/service/main.go -o pkg/service/swaggerdocs
	if strings $(SED_PATH) | grep -q 'GNU'; then \
		$(SED_PATH) -i '/^\/\/ This file was generated by swaggo\/swag at/{n;d;}' pkg/service/swaggerdocs/docs.go; \
		$(SED_PATH) -i '/^\/\/ This file was generated by swaggo\/swag at/d' pkg/service/swaggerdocs/docs.go; \
	else \
		$(SED_PATH) -i '' '/^\/\/ This file was generated by swaggo\/swag at/{n;d;}' pkg/service/swaggerdocs/docs.go; \
		$(SED_PATH) -i '' '/^\/\/ This file was generated by swaggo\/swag at/d' pkg/service/swaggerdocs/docs.go; \
	fi

.PHONY: test
test:
	go test -v -cover ./...

.PHONY: build_service_local
build_service_local: service_doc
	mkdir -p builds/debug
	go build -o builds/debug/service -ldflags '${EXTLDFLAGS}-X github.com/kpaas-io/kpaas/pkg/utils/version.VersionDev=build.$(BUILD_NUMBER)' github.com/kpaas-io/kpaas/cli/service

.PHONY: build_service_cross
build_service_cross: service_doc
	mkdir -p builds/release
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o builds/release/service -ldflags '${EXTLDFLAGS}-X github.com/kpaas-io/kpaas/pkg/utils/version.VersionDev=build.$(BUILD_NUMBER)' github.com/kpaas-io/kpaas/cli/service

.PHONY: run
run: build_service_local
	builds/debug/service

build_service_image: build_service_cross
	docker build -t $(IMAGE_REPOSITORY_URL)/service:$(BUILD_TAG) -f builds/docker/service/Dockerfile .

push_service_image: build_service_image
	docker push $(IMAGE_REPOSITORY_URL)/service:$(BUILD_TAG)

.PHONY: build_deploy_protos
build_deploy_protos:
	@sh builds/protos/deploy/protos.sh

.PHONY: build_deploy_local
build_deploy_local: build_deploy_protos
	mkdir -p builds/debug
	go build -o builds/debug/deploy -ldflags '${EXTLDFLAGS}-X github.com/kpaas-io/kpaas/pkg/utils/version.VersionDev=build.$(BUILD_NUMBER)' github.com/kpaas-io/kpaas/cli/deploy

.PHONY: build_deploy_cross
build_deploy_cross: build_deploy_protos
	mkdir -p builds/release
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o builds/release/deploy -ldflags '${EXTLDFLAGS}-X github.com/kpaas-io/kpaas/pkg/utils/version.VersionDev=build.$(BUILD_NUMBER)' github.com/kpaas-io/kpaas/cli/deploy

