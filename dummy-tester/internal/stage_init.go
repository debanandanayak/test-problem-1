package internal

import (
	"fmt"

	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testInit(stageHarness *test_case_harness.TestCaseHarness) error {
	stageHarness.Logger.Infof("$ ./script.sh echo 1")

	result, err := stageHarness.Executable.Run("echo", "1")
	if err != nil {
		return err
	}

	if result.ExitCode != 0 {
		return fmt.Errorf("expected exit code %v, got %v", 0, result.ExitCode)
	}

	stageHarness.Logger.Successf("✓ Received exit code 0.")

	if string(result.Stdout) != "1\n" {
		return fmt.Errorf("expected stdout %v, got %v", "1\n", result.Stdout)
	}

	stageHarness.Logger.Successf("✓ Received output 1")
	return nil
}
