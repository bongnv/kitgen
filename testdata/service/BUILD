load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "service.go",
        "z_noop_service.go",
    ],
    importpath = "github.com/bongnv/kitgen/testdata/service",
    visibility = ["//visibility:public"],
    deps = ["//testdata/model:go_default_library"],
)
