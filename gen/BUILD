load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "data_parser.go",
        "gen.go",
        "generator.go",
        "helpers.go",
        "writer.go",
    ],
    importpath = "github.com/bongnv/kitgen/gen",
    visibility = ["//visibility:public"],
    deps = [
        "//templates:go_default_library",
        "@org_golang_x_tools//go/packages:go_default_library",
        "@org_golang_x_tools//imports:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "gen_test.go",
        "helpers_test.go",
        "writer_test.go",
    ],
    embed = [":go_default_library"],
    deps = [
        "@com_github_pallinder_go_randomdata//:go_default_library",
        "@com_github_stretchr_testify//require:go_default_library",
    ],
)
