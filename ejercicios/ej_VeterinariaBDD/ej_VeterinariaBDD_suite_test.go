package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestEjVeterinariaBDD(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "EjVeterinariaBDD Suite")
}
