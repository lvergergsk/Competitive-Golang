package isomorphic_string_0205

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	//mst: s->t mapping
	//mta: t->assignedOrNot mapping
	mst := make(map[uint8]uint8)
	mta := make(map[uint8]bool)
	for i := 0; i < len(s); i++ {
		if val, ok := mst[s[i]]; ok {
			if val != t[i] {
				return false
			}
		} else {
			if mta[t[i]] {
				return false
			}
			mst[s[i]] = t[i]
			mta[t[i]] = true

		}
	}
	return true
}
