package exercises

func FindFirstLastPositionOfTarget(arr []int, target int) [2]int {

	firstFoundIndex := -1
	currentFoundIndex := -1

	for currIndex, value := range arr {
		if value != target {
			continue
		}

		firstFoundIndex = currIndex
		for next := currIndex + 1; next < len(arr) && arr[next] == target; next += 1 {
			currentFoundIndex = next
		}
		break

	}

	return [2]int{firstFoundIndex, currentFoundIndex}

}
