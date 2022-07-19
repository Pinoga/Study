package exercises

func GetStartingGasStation(gas []int, cost []int) int {
	count := 0
	stoppedAtIndex := 0
	for index := 0; count < len(gas); index = getNextIndex(len(gas), stoppedAtIndex) {
		stoppedAtIndex = canGoFromTo(index, index, 0, gas, cost)
		if stoppedAtIndex <= 0 {
			return index
		}
		if stoppedAtIndex > index {
			count += stoppedAtIndex - index + 1
		} else {
			count += len(gas) - 1 - index + stoppedAtIndex
		}
	}
	return -1
}

func getNextIndex(length int, current int) int {
	if current < length-1 {
		return current + 1
	}
	return 0
}

func canGoFromTo(index, target, currentGas int, gas []int, cost []int) int {
	next := getNextIndex(len(gas), index)
	if currentGas+gas[index] < cost[index] {
		return index
	}
	if next == target {
		return -1
	}
	return canGoFromTo(next, target, currentGas+gas[index]-cost[index], gas, cost)
}
