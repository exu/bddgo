package bddgo_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestBddgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bddgo Suite")
}
