# kitgen
[![Build Status](https://travis-ci.com/bongnv/kitgen.svg?branch=master)](https://travis-ci.com/bongnv/kitgen)[![codecov](https://codecov.io/gh/bongnv/kitgen/branch/master/graph/badge.svg)](https://codecov.io/gh/bongnv/kitgen)[![Go Report Card](https://goreportcard.com/badge/github.com/bongnv/kitgen)](https://goreportcard.com/report/github.com/bongnv/kitgen)[![GoDoc](https://godoc.org/github.com/bongnv/kitgen?status.svg)](https://godoc.org/github.com/bongnv/kitgen)[![GolangCI](https://golangci.com/badges/github.com/golangci/golangci-lint.svg)](https://golangci.com)

kitgen is a code generation utility that help to speed up building web-application with Golang 

## Working with Bazel
Use following command to update & fix BUILD files:

    $bazel run //:gazelle -- fix

Use the following command to add go_repository for external dependencies

    $bazel run //:gazelle -- update-repos -from_file=go.mod

