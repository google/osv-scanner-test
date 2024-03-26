package lockfile

import (
	"path/filepath"
)

const TestEcosystem Ecosystem = "Test"

type TestLockExtractor struct{}

func (e TestLockExtractor) ShouldExtract(path string) bool {
	return filepath.Base(path) == "test.lock"
}

func (e TestLockExtractor) Extract(f DepFile) ([]PackageDetails, error) {
        return []PackageDetails{}, nil
}

var _ Extractor = TestLockExtractor{}

//nolint:gochecknoinits
func init() {
	registerExtractor("test.lock", TestLockExtractor{})
}

func ParseTestLock(pathToLockfile string) ([]PackageDetails, error) {
	return extractFromFile(pathToLockfile, TestLockExtractor{})
}
