package e231

func checkOnesSegment(s string) bool {

	var ok bool
	var zeroAfterOne bool
	for _, r := range s {
		if r == '1' && zeroAfterOne == false {
			if ok == false {
				ok = true
			}
			continue
		} else if r == '1' && zeroAfterOne == true {
			ok = false
			break
		}

		if r == '0' && ok {
			zeroAfterOne = true
		}
	}

	return ok
}
