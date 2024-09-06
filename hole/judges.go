package hole

import (
	"strings"

	"github.com/agnivade/levenshtein"
	hungarianAlgorithm "github.com/oddg/hungarian-algorithm"
)

// An ultimate task of a Judge is to tell whether a run is passing or not.
// However, it does more. When the output is not recognised as correct,
// it returns a similar output that is.
type Judge func(run Run) string

// A judge for a single output (corresponding to either a single input or a single preset output item)
type SingleOutputJudge func(arg, userOutput, rawExpectedOutput string) string

// Creates a judge by treating an output as a sequence of
// delimited outputs, each corresponding to a respective input.
// The final expected output length's is the bigger of input length & the rawExpectedOutput length.
// This allows two different ways of building the expected output:
//  1. purely from the args
//  2. by transforming the preset raw expected output (from the runs generator)
func perOutputJudge(singleOutputJudge SingleOutputJudge) Judge {
	return func(run Run) string {
		args := run.Args
		userOutputs := strings.Split(run.Stdout, run.OutputDelimiter)
		rawExpectedOutputs := strings.Split(run.Answer, run.OutputDelimiter)
		var expectedOutputs []string
		for i := range max(len(args), len(rawExpectedOutputs)) {
			arg := ""
			if i < len(args) {
				arg = args[i]
			}
			userOutput := ""
			if i < len(userOutputs) {
				userOutput = userOutputs[i]
			}
			rawExpectedOutput := ""
			if i < len(rawExpectedOutputs) {
				rawExpectedOutput = rawExpectedOutputs[i]
			}
			expectedOutput := singleOutputJudge(arg, userOutput, rawExpectedOutput)
			if expectedOutput != "" {
				expectedOutputs = append(expectedOutputs, expectedOutput)
			}
		}
		return strings.Join(expectedOutputs, run.OutputDelimiter)
	}
}

// Creates a judge which checks whether each user output
// corresponds to one of preset outputs corresponding to the respective arg.
func oneOfPerOutputJudge(getAllSolutions func(arg string) []string, caseFold bool) Judge {
	return perOutputJudge(func(arg, userOutput, rawExpectedOutput string) string {
		solutions := getAllSolutions(arg)
		for _, solution := range solutions {
			if caseFold && strings.EqualFold(solution, userOutput) || !caseFold && solution == userOutput {
				return userOutput
			}
		}

		closestSolution := ""
		minDistance := 1 << 24
		if caseFold {
			userOutput = strings.ToLower(userOutput)
		}
		for i, solution := range solutions {
			if caseFold {
				solution = strings.ToLower(solution)
			}
			distance := levenshtein.ComputeDistance(solution, userOutput)
			if distance < minDistance {
				minDistance = distance
				closestSolution = solutions[i]
			}
		}
		return closestSolution
	})
}

func getClosestMultiset(anyAnswer, stdout, itemDelimiter string) string {
	expectedItems := strings.Split(anyAnswer, itemDelimiter)
	expectedItemsReordered := make([]string, len(expectedItems))
	userItems := strings.Split(stdout, itemDelimiter)

	expectedItemsMap := make(map[string]int)
	for _, expected := range expectedItems {
		expectedItemsMap[expected]++
	}

	// Match items that are correct
	matches := 0
	for i, user := range userItems {
		if i < len(expectedItems) && expectedItemsMap[user] > 0 {
			expectedItemsReordered[i] = user
			expectedItemsMap[user]--
			userItems[i] = ""
			matches++
		}
	}

	// Process mismatched items
	if matches < len(expectedItems) {

		// Calculate indices of expected & user items that couldn't be matched be equality
		unmatchedExpectedIndices := []int{}
		unmatchedUserIndices := []int{}

		for i, expected := range expectedItems {
			if expectedItemsMap[expected] > 0 {
				unmatchedExpectedIndices = append(unmatchedExpectedIndices, i)
				expectedItemsMap[expected]--
			}
		}

		for i, user := range userItems {
			if user != "" {
				unmatchedUserIndices = append(unmatchedUserIndices, i)
			}
		}

		n := max(len(unmatchedExpectedIndices), len(unmatchedUserIndices))

		permutation := make([]int, n)
		for i := range permutation {
			permutation[i] = i
		}

		// If there are not many wrong items, try to match them
		// otherwise, use the above identity permutation
		if n <= 32 {
			dist := make([][]int, n)
			for i := range dist {
				dist[i] = make([]int, n)
				for j := range dist {
					if j >= len(unmatchedExpectedIndices) {
						dist[i][j] = len(userItems[unmatchedUserIndices[i]])
					} else if i >= len(unmatchedUserIndices) {
						dist[i][j] = len(expectedItems[unmatchedExpectedIndices[j]])
					} else {
						dist[i][j] = levenshtein.ComputeDistance(expectedItems[unmatchedExpectedIndices[j]], userItems[unmatchedUserIndices[i]])
					}
				}
			}

			permutation, _ = hungarianAlgorithm.Solve(dist)
		}

		k := 0
		for _, i := range permutation {
			if k >= len(expectedItemsReordered) {
				break
			}
			if i < len(unmatchedExpectedIndices) {
				for expectedItemsReordered[k] != "" {
					k++
				}
				expectedItemsReordered[k] = expectedItems[unmatchedExpectedIndices[i]]
			}
		}
	}

	return strings.Join(expectedItemsReordered, itemDelimiter)
}
