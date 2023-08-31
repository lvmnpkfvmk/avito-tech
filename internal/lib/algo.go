package algo

import "github.com/lvmnpkfvmk/avito-tech/internal/model"

func FilterSegments(first *model.Segments, second *model.Segments) *model.Segments {
	nameSet := make(map[string]bool)

	for _, seg := range *second {
		nameSet[seg.Name] = true
	}

	var uniqueFirst model.Segments
	for _, seg := range *first {
		if !nameSet[seg.Name] {
			uniqueFirst = append(uniqueFirst, seg)
			nameSet[seg.Name] = true
		}
	}

	var filteredNames model.Segments
	for _, seg := range uniqueFirst {
		filteredNames = append(filteredNames, seg)
	}

	return &filteredNames
}
