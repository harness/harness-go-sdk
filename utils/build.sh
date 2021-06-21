#!/bin/bash

function preprelease() {
  version=${VERSION:?"VERSION must be set"}

  git checkout -b release/$version
  
  sed -i.bak "s/SDKVersion = .*/SDKVersion = \"$version\"/"  harness/version.go && rm harness/version.go.bak

  git branch
  git status
}

$1
