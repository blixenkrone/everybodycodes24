package q2

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"

	"blixenkrone/everybody-codes-24/files"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func sanitizeInput(b []byte) ([]string, []string, string) {
	a1, found := bytes.CutPrefix(b, []byte("WORDS:"))
	if !found {
		panic("prefix failed")
	}
	firstHalf := bytes.Index(a1, []byte("\n"))
	words := strings.Split(string(a1[:firstHalf]), ",")
	sentence := a1[firstHalf:]
	// return words, strings.ReplaceAll(strings.TrimSpace(string(sentence)), " ", "")
	return words, strings.Split(string(sentence), " "), strings.TrimSpace(string(sentence))
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
func solve(t *testing.T, r io.Reader) int {
	b, err := io.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}

	var numEngravings int

	words, _, sentence := sanitizeInput(b)
	indexScoreM := make(map[int]int, len(sentence))
	// WORD: THI
	// THIS IS A SENTENCE
	// 10-7 < 3

	for i := 0; i < len(words); i++ {
		word := words[i]
		if len(sentence)-i < len(word) {
			break
		}
		origin := word
		rev := reverse(word)
		warr := []string{origin, string(rev)}

		for j := 0; j < len(sentence)-len(word); j++ {
			for _, w := range warr {
				delta := j + len(word)
				curr := sentence[j:delta]
				if w == curr {
					fmt.Printf("matched %s at idx %d-%d \n", w, j, delta)
					// numEngravings += len(curr)
					for k := j; k < delta; k++ {
						indexScoreM[k] = 1
					}
					break
				}
			}
		}
	}

	for _, v := range indexScoreM {
		numEngravings += v
	}

	return numEngravings
}

func TestSolveP2(t *testing.T) {
	t.Run("Ex", func(t *testing.T) {
		in := `WORDS:THE,OWE,MES,ROD,HER

		AWAKEN THE POWE ADORNED WITH THE FLAMES BRIGHT IRE
		THE FLAME SHIELDED THE HEART OF THE KINGS
		POWE PO WER P OWE R
		THERE IS THE END`
		// in := `WORDS:THE,OWE,MES,ROD,HER

		// THERE IS THE END`
		// wants := []int{15, 9, 6, 7}
		got := solve(t, strings.NewReader(in))
		// assert.Equal(t, wants[3], got)
		assert.Equal(t, 37, got)
	})
	t.Run("Real", func(t *testing.T) {
		f := files.MustOpen(t, "./q2_p2.txt")
		spew.Dump(solve(t, f))
	})
}

func TestSolveP1(t *testing.T) {
	t.Run("Ex", func(t *testing.T) {
		in := `WORDS:THE,OWE,MES,ROD,HER

AWAKEN THE POWER ADORNED WITH THE FLAMES BRIGHT IRE`
		got := solve(t, strings.NewReader(in))
		assert.Equal(t, 4, got)
	})
	t.Run("Real", func(t *testing.T) {
		f := files.MustOpen(t, "./q2_p1.txt")
		spew.Dump(solve(t, f))
	})
}
