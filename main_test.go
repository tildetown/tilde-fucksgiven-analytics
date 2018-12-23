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

	g.Describe("numUppercaseRunes", func() {
		g.It("should accurately count the number of uppercase runes in a word", func() {
			word1 := "SHAKWAM"
			word2 := "hellodarknessmyoldfriend"

			runeSlice1 := []rune(word1)
			runeSlice2 := []rune(word2)

			n1 := numUppercaseRunes(&runeSlice1)
			n2 := numUppercaseRunes(&runeSlice2)

			Expect(len(word1)).To(Equal(*n1))
			Expect(len(word2)).NotTo(Equal(*n2))
			Expect(*n1).To(BeNumerically("==", 7))
			Expect(*n2).To(BeNumerically("==", 0))
		})
	})
}
