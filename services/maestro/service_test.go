package maestro

import (
	"testing"
)

func TestTableDeploy(t *testing.T) {
	var tests = []struct {
		pod       string
		blueprint string
		result    bool
		err       error
	}{
		{"p02", "blueprint", true, nil},
		{"p10", "blueprint", true, nil},
		{"p02verylongpod", "blueprint", true, nil},
	}

	var m Maestro
	m = NewService()

	for _, test := range tests {
		res, err := m.Deploy(test.pod, test.blueprint)

		if err != test.err {
			t.Errorf("Test Failed: %s pod, %s blueprint, %t expected, recived %t",
				test.pod, test.blueprint, test.result, res)
		}
		if res != test.result {
			t.Errorf("Test Failed: %s pod, %s blueprint, %t expected, recived %t",
				test.pod, test.blueprint, test.result, res)
		}
	}
}
