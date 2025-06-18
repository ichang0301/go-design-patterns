package pipeline

import "testing"

func TestLaunchPipeline(t *testing.T) {
	tableTests := [][]int{
		{3, 14},
		{5, 55},
		{10, 385},
		{100, 338350},
	}

	var res int
	for _, test := range tableTests {
		res = LaunchPipeline(test[0])
		if res != test[1] {
			t.Errorf("LaunchPipeline(%d) = %d, but want = %d", test[0], res, test[1])
		}

		t.Logf("LaunchPipeline(%d) = %d", test[0], res)
	}
}
