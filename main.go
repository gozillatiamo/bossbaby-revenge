package main

const (
	GOOD string = "Good boy"
	BAD  string = "Bad boy"
)

// To check the events of Boss baby's revenge
// All operations represent in
// Time complexity O(n)
// Space complexity O(1)
func CheckBossBaby(events string) (string, error) {
	if isInitiateRevenge(events) {
		return BAD, nil
	}

	if remainingShots(events) > 0 {
		return BAD, nil
	}

	return GOOD, nil
}

// To Check the event start with revenged or not
// Time complexity is O(1)
// Space complexity is O(1)
func isInitiateRevenge(events string) bool {
	// e := *event
	return events[0] == 'R'
}

// Time complexity is O(n) depend on events size
// Space complexity is O(1) because of re-using shots variable
func remainingShots(events string) int {
	shots := 0

	for _, event := range events {
		// count every shot received
		if event == 'S' {
			shots += 1
		}

		// decrease when revenged
		if event == 'R' {
			shots -= 1
		}

		// when revenged more then remain shots in a row reset shots count to 0
		if shots < 0 {
			shots = 0
		}
	}

	return shots
}
