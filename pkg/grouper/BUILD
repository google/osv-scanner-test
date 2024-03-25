package(
    default_applicable_licenses = ["//third_party/golang/github_com/google/osv_scanner_test/v/v1:license"],
    default_visibility = ["//third_party/golang/github_com/google/osv_scanner_test/v/v1:__subpackages__"],
    licenses = ["notice"],
)

go_library(
    name = "grouper",
    srcs = [
        "grouper.go",
        "grouper_models.go",
    ],
    deps = [
        "//third_party/golang/github_com/google/osv_scanner_test/v/v1/pkg/models",
        "//third_party/golang/go_exp/maps",
        "//third_party/golang/go_exp/slices",
    ],
)

go_test(
    name = "grouper_test",
    srcs = ["grouper_test.go"],
    library = ":grouper",
    deps = [
        "//third_party/golang/cmp",
        "//third_party/golang/github_com/google/osv_scanner_test/v/v1/pkg/models",
    ],
)
