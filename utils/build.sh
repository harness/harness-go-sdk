#!/bin/bash

function preprelease() {
  version=${VERSION:?"VERSION must be set"}
  # release_branch="release/$version"

  # git checkout -b $release_branch
  # git config user.email "harness-ci-automation@harness.io"
  # git config user.name "Harness CI Automation"
  sed -i.bak "s/SDKVersion = .*/SDKVersion = \"$version\"/"  harness/version.go && rm harness/version.go.bak

  # git branch
  # git status
  # cat .git/config

  # git add -A
  # git commit -am"Release v${version} $(date +%m-%d-%Y)"
  # git push origin $release_branch
}

$1
