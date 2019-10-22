package promql

import (
	"bufio"
	"os"
	"strings"
)

var name, mapping, mappingReversed = readIdMapping()

func readIdMapping() (string, map[string]string, map[string]string) {
	file, err := os.Open("mapping.txt")
	if err != nil {
		return "", map[string]string{}, map[string]string{}
	}
	scanner := bufio.NewScanner(file)

	result := make(map[string]string)
	resultReversed := make(map[string]string)
	var name string
	if scanner.Scan() {
		name = scanner.Text()
	}
	for scanner.Scan() {
		line := scanner.Text()
		splitted := strings.SplitN(line, " ", 2)
		id, name := splitted[0], splitted[1]
		result[id] = name
		resultReversed[name] = id
	}
	return name, result, resultReversed
}

func TargetName() string {
	return name
}

func MapValue(unmappedValue string) (bool, string) {
	var mapped, exists = mapping[unmappedValue]
	return exists, mapped
}

func UnMapValue(mappedValue string) string {
	unmappedValue, exists := mappingReversed[mappedValue]
	if exists {
		return unmappedValue
	} else {
		return mappedValue
	}
}
