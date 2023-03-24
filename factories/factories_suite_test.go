package factories_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

func TestFactories(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Factories Suite")
}
