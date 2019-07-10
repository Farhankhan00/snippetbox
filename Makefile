BINARY=snippetbox
VERSION=0.0.1
GITSHA=`git rev-parse HEAD`
BUILD_TIME=`date -u +%Y-%m-%dT%H:%M:%S%Z`
LDFLAGS=""

.DEFAULT_GOAL: build

build:
    gox -osarch="linux/amd64 darwin/amd64 windows/amd64" -output "dist/${BINARY}_{{.OS}}_{{.Arch}}"

install:
    go install 

clean:
    if [ -d dist ] ; then rm -rf dist; fi

.PHONY: clean install

test:
    go test -v ./...

publish:
ifeq ($(BRANCH),master)
    @echo On branch master, publishing release
    ghr \
    -t ${GITHUB_TOKEN} \
    -u ${CIRCLE_PROJECT_USERNAME} \
    -r ${CIRCLE_PROJECT_REPONAME} \
    -n version ${VERSION} \
    --recreate \
    v${VERSION} \
    dist/
else
    @echo On branch $(BRANCH), publishing prerelease
    ghr \
    -t ${GITHUB_TOKEN} \
    -u ${CIRCLE_PROJECT_USERNAME} \
    -r ${CIRCLE_PROJECT_REPONAME} \
    -n version ${VERSION} \
    --recreate \
    --prerelease \
    v${VERSION} \
    dist/
endif