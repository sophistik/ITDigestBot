package slices

func Unique(elems []string) []string {
	var uniques []string

	m := make(map[string]bool, 0)

	for _, elem := range elems {
		m[elem] = true
	}

	for k, _ := range m {
		uniques = append(uniques, k)
	}

	return uniques
}

func Delete(elems []string, toDelete []string) []string {
	var deleted []string

	m := make(map[string]bool, 0)

	for _, elem := range toDelete {
		m[elem] = true
	}

	for _, elem := range elems {
		if m[elem] {
			continue
		}
		deleted = append(deleted, elem)
	}

	return deleted
}
