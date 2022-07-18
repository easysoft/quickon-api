#!/usr/bin/env bash

set -xe

addlicense -f hack/licenses/licenses.tpl -ignore web/** -ignore "**/*.md" -ignore vendor/** -ignore "**/*.yml" -ignore "**/*.yaml" -ignore "**/*.sh" ./**
