# Description:
#   Auto-imported from github.com/google/osv-scanner-test

package(
    default_applicable_licenses = ["//third_party/golang/github_com/google/osv_scanner_test/v/v1:license"],
    default_compatible_with = ["//buildenv/target:non_prod"],
    default_visibility = ["//third_party/golang/github_com/google/osv_scanner_test/v/v1:__subpackages__"],
)

licenses(["notice"])

go_library(
    name = "url",
    srcs = ["url.go"],
)

go_test(
    name = "url_test",
    srcs = [
        "url_other_test.go",
        "url_test.go",
        "url_windows_test.go",
    ],
    library = ":url",
)
