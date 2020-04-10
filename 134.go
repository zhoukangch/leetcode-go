package main

func canCompleteCircuit(gas []int, cost []int) int {
	if len(gas) == 0 || len(cost) == 0 {
		return -1
	}
	aquire := make([]int, len(gas), len(gas))
	for idx, val := range gas {
		aquire[idx] = val - cost[idx]
	}
	for idx, v := range aquire {
		if v < 0 {
			continue
		}
		sum := 0
		for step := 0; step <= len(gas)-1; step++ {
			cur := idx + step
			if cur >= len(gas) {
				cur = cur - len(gas)
			}
			sum += aquire[cur]
			if sum < 0 {
				break
			}
		}
		if sum >= 0 {
			return idx
		}
	}
	return -1
}
