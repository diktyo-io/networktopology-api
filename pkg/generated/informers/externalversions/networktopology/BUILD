load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["interface.go"],
    importmap = "k8s.io/kubernetes/vendor/github.com/diktyo-io/networktopology-api/pkg/generated/informers/externalversions/networktopology",
    importpath = "github.com/diktyo-io/networktopology-api/pkg/generated/informers/externalversions/networktopology",
    visibility = ["//visibility:public"],
    deps = [
        "//github.com/diktyo-io/networktopology-api/pkg/generated/informers/externalversions/internalinterfaces:go_default_library",
        "//github.com/diktyo-io/networktopology-api/pkg/generated/informers/externalversions/networktopology/v1alpha1:go_default_library",
    ],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [
        ":package-srcs",
        "//github.com/diktyo-io/appgroup-api/pkg/generated/informers/externalversions/appgroup/v1alpha1:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
