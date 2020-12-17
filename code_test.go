package advent20201207_test

import (
	"fmt"
	"testing"

	advent "github.com/bwarren2/advent20201207"
)

func TestNewHeldByMap(t *testing.T) {
	// spew.Dump(advent.NewHeldByMap("sample.txt"))
}
func TestNewBagHoldingMap(t *testing.T) {
	// spew.Dump(advent.NewBagHoldingMap("sample.txt"))
}
func TestPart1(t *testing.T) {
	fmt.Println(advent.Part1("input.txt", "shiny gold"))
}
func TestPart2(t *testing.T) {
	fmt.Println(advent.Part2("input.txt", "shiny gold"))
}
