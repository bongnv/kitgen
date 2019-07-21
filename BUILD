load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/bongnv/kitgen
# gazelle:build_file_name BUILD,BUILD.bazel
gazelle(
    name = "gazelle",
    prefix = "github.cgom/bongnv/kitgen",
)

go_library(
    name = "go_default_library",
    srcs = ["doc.go"],
    importpath = "github.com/bongnv/kitgen",
    visibility = ["//visibility:public"],
)
