#!/usr/bin/env bash

go mod tidy
go mod vendor
retVal=$?
if [ $retVal -ne 0 ]; then
    exit $retVal
fi

set -ex
TMP_DIR=$(mktemp -d)
mkdir -p "${TMP_DIR}"/src/github.com/easysoft/quickon-api
cp -r ./{qucheng,hack,vendor,go.mod} "${TMP_DIR}"/src/github.com/easysoft/quickon-api

(cd "${TMP_DIR}"/src/github.com/easysoft/quickon-api; \
    GOPATH=${TMP_DIR} GO111MODULE=off /bin/bash vendor/k8s.io/code-generator/generate-groups.sh all \
    github.com/easysoft/quickon-api/client github.com/easysoft/quickon-api "qucheng:v1beta1" -h ./hack/boilerplate.go.txt ;
    )

rm -rf ./client/{clientset,informers,listers}
tree "${TMP_DIR}"/src/github.com/easysoft/quickon-api/client/
mv "${TMP_DIR}"/src/github.com/easysoft/quickon-api/client .
rm -rf ${TMP_DIR}
rm -rf vendor
