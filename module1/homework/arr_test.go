package homework

import "testing"

func TestArrRange(t *testing.T) {

	arr := [...]string{"I", "am", "stupid", "and", "weak"}

	for index, _ := range arr {
		if arr[index] == "stupid" {
			arr[index] = "smart"
		}
		if arr[index] == "weak" {
			arr[index] = "strong"
		}
	}

	t.Log(arr)
}
