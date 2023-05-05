package cmd

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/spf13/cobra"
)

func TestGetImageTag(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
		image    string
		tag      string
	}{
		{"test-image:latest", true, "test-image", "latest"},
		{"test-image", false, "", ""},
		{"test-image:123", true, "test-image", "123"},
		{"", false, "", ""},
	}

	for _, test := range tests {
		i := &Image{}
		result := i.getImageTag(test.input)
		if result != test.expected {
			t.Errorf("Expected %v but got %v for input %s", test.expected, result, test.input)
		}

		if i.imageName != test.image || i.tag != test.tag {
			t.Errorf("Expected imageName: %s and tag: %s but got imageName: %s and tag: %s", test.image, test.tag, i.imageName, i.tag)
		}
	}
}

func executeCommand(cmd *cobra.Command, args ...string) (output string, err error) {
	_, orig := os.Stderr, os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	defer func() {
		os.Stderr = orig
	}()

	cmd.SetArgs(args)
	err = cmd.Execute()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	output = string(out)
	return output, err
}

func TestScanCmd(t *testing.T) {
	// This test only checks if the command can be called with the required flags
	// and returns an error. It does not mock external dependencies and test the
	// actual functionality.
	args := []string{"--image", "test-image:latest", "--region", "us-west-2"}
	_, err := executeCommand(scanCmd, args...)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
