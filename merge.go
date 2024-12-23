package utilize

// MergeMap merge two map to one, if overwrite is true, the value of m2 will overwrite the value of m1
func MergeMap(m1, m2 map[string]any, overwrite bool) (m3 map[string]any) {
	m3 = make(map[string]any)
	for k, v := range m1 {
		m3[k] = v
	}

	for k, v := range m2 {
		if _, ok := m3[k]; ok && !overwrite {
			continue
		}
		m3[k] = v
	}
	return m3
}
