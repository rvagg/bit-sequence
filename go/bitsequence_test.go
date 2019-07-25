package bitsequence

import (
	"bufio"
	"encoding/hex"
	"os"
	"strconv"
	"strings"
	"testing"
)

const FixturesFile = "../test-fixture.csv"

type TestCase struct {
	bytes                   []byte
	start, length, expected uint32
}

func TestFixtures(t *testing.T) {
	testCases, err := readTestCases(FixturesFile)
	if err != nil {
		panic(err)
	}

	for _, testCase := range testCases {
		actual := BitSequence(testCase.bytes, testCase.start, testCase.length)
		if actual != testCase.expected {
			t.Errorf("Bytes [%s] start=%d length=%d expected %d but got %d",
				hex.EncodeToString(testCase.bytes),
				testCase.start,
				testCase.length,
				testCase.expected,
				actual)
		} /* else {
			fmt.Printf("Bytes [%s] start=%d length=%d expected %d and got %d\n",
				hex.EncodeToString(testCase.bytes),
				testCase.start,
				testCase.length,
				testCase.expected,
				actual)
		} */
	}
}

func readTestCases(path string) ([]TestCase, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var testCases []TestCase
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, ",")

		testCase := TestCase{}
		var err error
		var ii int

		testCase.bytes, err = hex.DecodeString(s[0])
		if err != nil {
			return nil, err
		}
		ii, err = strconv.Atoi(s[1])
		if err != nil {
			return nil, err
		}
		testCase.start = uint32(ii)
		ii, err = strconv.Atoi(s[2])
		if err != nil {
			return nil, err
		}
		testCase.length = uint32(ii)
		ii, err = strconv.Atoi(s[3])
		if err != nil {
			return nil, err
		}
		testCase.expected = uint32(ii)

		testCases = append(testCases, testCase)
	}

	return testCases, nil
}
