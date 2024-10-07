package transformer

import (
	"strings"
)

func transformLine(line string) string {
	if strings.Contains(line, "OWNER TO") || strings.Contains(line, "GRANT ALL") {
		return "-- " + line
	}

	return line
}

func TransformLines(lines []string) []string {
	transformedLines := make([]string, len(lines))

	for i, line := range lines {
		transformedLines[i] = transformLine(line)
	}

	return transformedLines
}
