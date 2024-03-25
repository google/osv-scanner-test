# Description:
#   Auto-imported from github.com/google/osv-scanner-test

package(
    default_applicable_licenses = ["//third_party/golang/github_com/google/osv_scanner_test/v/v1:license"],
    default_compatible_with = ["//buildenv/target:non_prod"],
    default_visibility = ["//third_party/golang/github_com/google/osv_scanner_test/v/v1:__subpackages__"],
)

licenses(["notice"])

go_library(
    name = "lockfile",
    srcs = [
        "apk-installed.go",
        "csv.go",
        "dpkg-status.go",
        "ecosystems.go",
        "extract.go",
        "extractor.go",
        "osv-vuln-results.go",
        "parse.go",
        "parse-composer-lock.go",
        "parse-conan-lock.go",
        "parse-gemfile-lock.go",
        "parse-go-lock.go",
        "parse-gradle-lock.go",
        "parse-maven-lock.go",
        "parse-mix-lock.go",
        "parse-npm-lock.go",
        "parse-nuget-lock.go",
        "parse-pipenv-lock.go",
        "parse-pnpm-lock.go",
        "parse-pubspec-lock.go",
        "parse-renv-lock.go",
        "parse-requirements-txt.go",
        "parse-yarn-lock.go",
        "types.go",
    ],
    deps = [
        "//third_party/golang/github_com/google/osv_scanner_test/v/v1/internal/cachedregexp",
        "//third_party/golang/github_com/google/osv_scanner_test/v/v1/pkg/models",
        "//third_party/golang/go_mod/modfile",
        "//third_party/golang/yamlv3:yaml",
    ],
)
