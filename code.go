package advent20201207

import (
	"bufio"
	"os"
	"regexp"
	"strconv"

	mapset "github.com/deckarep/golang-set"
)

// BagQuantity stores a back quantity
type BagQuantity struct {
	count int64
	kind  string
}

// BagHoldingMap keeps track of what a given bag holds
type BagHoldingMap map[string][]BagQuantity

// HeldByMap keeps track of what can hold a given bag
type HeldByMap map[string][]string

// LinesFromFile produces the lines from a file
func LinesFromFile(filename string) (result []string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return
}

var BagCountRegex = regexp.MustCompile(`(?P<count>\d) (?P<kind>[a-zA-Z ]*) bags?`)
var ContainingBagRegex = regexp.MustCompile(`(?P<kind>[a-zA-Z ]*) bags contain `)
var TerminalBagRegex = regexp.MustCompile(`contain no other bags.`)

// NewHeldByMap creates a mapping dict from a file
func NewHeldByMap(filename string) HeldByMap {
	result := make(HeldByMap)
	lines := LinesFromFile(filename)
	for _, line := range lines {
		container := ContainingBagRegex.FindStringSubmatch(line)[1]
		holds := BagCountRegex.FindAllStringSubmatch(line, -1)

		for _, match := range holds {
			kind := match[2]
			result[kind] = append(result[kind], container)
		}
	}
	return result
}

func IsHeldBy(holdingMap HeldByMap, targetBag string) []string {
	if len(holdingMap[targetBag]) == 0 {
		return []string{targetBag}
	}
	returnList := make([]string, 0)
	returnList = append(returnList, targetBag)
	for _, heldByBag := range holdingMap[targetBag] {
		sublist := IsHeldBy(holdingMap, heldByBag)
		for _, value := range sublist {
			returnList = append(returnList, value)
		}
	}
	// fmt.Println(targetBag, returnList)
	return returnList
}

// Part1 answers part 1
func Part1(filename, targetBag string) (result int) {
	holdingMap := NewHeldByMap(filename)
	seen := mapset.NewSet()
	for _, bag := range IsHeldBy(holdingMap, targetBag) {
		seen.Add(bag)
	}
	seen.Remove(targetBag)
	// spew.Dump(seen)
	return seen.Cardinality()
}

func NewBagHoldingMap(filename string) BagHoldingMap {
	result := make(BagHoldingMap)
	lines := LinesFromFile(filename)
	for _, line := range lines {
		container := ContainingBagRegex.FindStringSubmatch(line)[1]
		holds := BagCountRegex.FindAllStringSubmatch(line, -1)
		for _, match := range holds {
			count, err := strconv.ParseInt(match[1], 10, 64)
			if err != nil {
				panic(err)
			}
			kind := match[2]
			result[container] = append(result[container], BagQuantity{count, kind})
		}
	}
	return result
}

func HoldingCount(holdingMap BagHoldingMap, targetBag string) (count int64) {
	if len(holdingMap[targetBag]) == 0 {
		return 1
	}
	for _, bagQuantity := range holdingMap[targetBag] {
		count += (bagQuantity.count * HoldingCount(holdingMap, bagQuantity.kind))
	}
	return count + 1
}

func Part2(filename, targetBag string) (count int64) {
	return HoldingCount(NewBagHoldingMap(filename), targetBag) - 1
}
