package main

func checkRecord(s string) bool {
	acount := 0
	lcount := 0
	for _, c := range s {
		if string(c) == "A" {
			acount++
			if lcount <= 2 {
				lcount = 0
			}
		}
		if string(c) != "L" {
			if lcount <= 2 {
				lcount = 0
			}
		} else {
			lcount++
		}
	}
	if acount <= 1 && lcount <= 2 {
		return true
	}
	return false
}
