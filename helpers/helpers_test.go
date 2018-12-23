package helpers

import (
	"regexp"
	"testing"

	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
	"github.com/tildetown/tilde-fucksgiven-analytics/fucks"
)

func TestMain(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("GetUpperCaseFucks", func() {
		g.It("should return a slice of uppercase fucks", func() {
			uf, _ := ParseFucks()
			ucf, err := GetUpperCaseFucks(uf.UniqueFucks)
			reg, _ := regexp.Compile("[^a-zA-Z]+")

			s0 := reg.ReplaceAllString(ucf[0], "")
			s1 := reg.ReplaceAllString(ucf[1], "")

			runes0 := []rune(s0)
			runes1 := []rune(s1)

			u0 := NumUppercaseRunes(&runes0)
			u1 := NumUppercaseRunes(&runes1)

			Expect(ucf).NotTo(BeNil())
			Expect(err).To(BeNil())
			Expect(len(runes0)).To(Equal(*u0))
			Expect(len(runes1)).To(Equal(*u1))
		})
	})

	g.Describe("GetLowerCaseFucks", func() {
		g.It("should return a slice of lowercase fucks", func() {
			uf, _ := ParseFucks()
			lcf, err := GetLowerCaseFucks(uf.UniqueFucks)
			reg, _ := regexp.Compile("[^a-zA-Z]+")

			s0 := reg.ReplaceAllString(lcf[0], "")
			s1 := reg.ReplaceAllString(lcf[1], "")

			runes0 := []rune(s0)
			runes1 := []rune(s1)

			u0 := NumLowercaseRunes(&runes0)
			u1 := NumLowercaseRunes(&runes1)

			Expect(lcf).NotTo(BeNil())
			Expect(err).To(BeNil())
			Expect(len(runes0)).To(Equal(*u0))
			Expect(len(runes1)).To(Equal(*u1))
		})
	})

	g.Describe("NumUppercaseRunes", func() {
		g.It("should accurately count the number of uppercase runes in a word", func() {
			word1 := "SHAKWAM"
			word2 := "hellodarknessmyoldfriend"

			runeSlice1 := []rune(word1)
			runeSlice2 := []rune(word2)

			n1 := NumUppercaseRunes(&runeSlice1)
			n2 := NumUppercaseRunes(&runeSlice2)

			Expect(len(word1)).To(Equal(*n1))
			Expect(len(word2)).NotTo(Equal(*n2))
			Expect(*n1).To(BeNumerically("==", 7))
			Expect(*n2).To(BeNumerically("==", 0))
		})
	})

	g.Describe("NumLowercaseRunes", func() {
		g.It("should accurately count the number of lowercase runes in a word", func() {
			word1 := "meow"
			word2 := "meowMEOWmeow"

			runeSlice1 := []rune(word1)
			runeSlice2 := []rune(word2)

			n1 := NumLowercaseRunes(&runeSlice1)
			n2 := NumLowercaseRunes(&runeSlice2)

			Expect(len(word1)).To(Equal(*n1))
			Expect(len(word2)).NotTo(Equal(*n2))
			Expect(*n1).To(BeNumerically("==", 4))
			Expect(*n2).To(BeNumerically("==", 8))
		})
	})

	g.Describe("ParseFucks", func() {
		g.It("should parse all the unique fucks into a struct", func() {
			uf, err := ParseFucks()

			Expect(uf).NotTo(BeNil())
			Expect(len(uf.UniqueFucks)).To(BeNumerically(">=", 277))
			Expect(err).To(BeNil())
		})
	})

	g.Describe("GetKindOfFucks", func() {
		g.It("should return a string representing the kind of fucks requested", func() {
			kind := "uppercase"
			f := fucks.Fucks{}
			uf := f.UniqueFucks
			uf = append(uf, "OOHBABYBABYISSAWILDFUCKINWORLD")

			kof, err := GetKindOfFucks(kind, &uf)

			Expect(kof).NotTo(BeNil())
			Expect(err).To(BeNil())
			Expect(kof).To(ContainSubstring("# of unique fucks: 1"))
			Expect(kof).To(ContainSubstring("# of uppercase fucks: 1"))
			Expect(kof).To(ContainSubstring("percentage: 100"))
		})
	})
}
