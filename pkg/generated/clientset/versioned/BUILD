load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "clientset.go",
        "doc.go",
    ],
    importmap = "k8s.io/kubernetes/vendor/github.com/diktyo-io/networktopology-api/pkg/generated/clientset/versioned",
    importpath = "github.com/diktyo-io/networktopology-api/pkg/generated/clientset/versioned",
    visibility = ["//visibility:public"],
    deps = [
        "//staging/src/k8s.io/client-go/discovery:go_default_library",
        "//staging/src/k8s.io/client-go/rest:go_default_library",
        "//staging/src/k8s.io/client-go/util/flowcontrol:go_default_library",
        "//github.com/diktyo-io/networktopology-api/pkg/generated/clientset/versioned/typed/networktopology/v1alpha1:go_default_library",
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
        "//github.com/diktyo-io/networktopology-api/pkg/generated/clientset/versioned/fake:all-srcs",
        "//github.com/diktyo-io/networktopology-api/pkg/generated/clientset/versioned/scheme:all-srcs",
        "//github.com/diktyo-io/networktopology-api/pkg/generated/clientset/versioned/typed/networktopology/v1alpha1:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
