package pascaltriangle

func pascalTriangle(lineNumber int) []int {
	currentLine := []int{1}
	currentLineSize := lineNumber + 1

	for index := 1; index < currentLineSize; index++ {
		currentLine = append(currentLine, currentLine[index-1]*(lineNumber-index+1)/index)
	}

	return currentLine
}
