package main

import (
	"testing"
	"regexp"

	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestMain(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("parseFucks", func() {
		g.It("should parse all the unique fucks into a struct", func() {
			uf, err := parseFucks()

			Expect(uf).NotTo(BeNil())
			Expect(len(uf.UniqueFucks)).To(BeNumerically(">=", 277))
			Expect(err).To(BeNil())
		})
	})

	g.Describe("getUpperCaseFucks", func() {
		g.It("should return a slice of uppercase fucks", func() {
			uf, _ := parseFucks()
			ucf, err := getUpperCaseFucks(uf.UniqueFucks)
			reg, _ := regexp.Compile("[^a-zA-Z]+")

			s0 := reg.ReplaceAllString(ucf[0], "")
			s1 := reg.ReplaceAllString(ucf[1], "")

			runes0 := []rune(s0)
			runes1 := []rune(s1)

			u0 := numUppercaseRunes(&runes0)
			u1 := numUppercaseRunes(&runes1)

			Expect(ucf).NotTo(BeNil())
			Expect(err).To(BeNil())
			Expect(len(runes0)).To(Equal(*u0))
			Expect(len(runes1)).To(Equal(*u1))
		})
	})
}
