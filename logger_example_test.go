package logger

import "testing"

func Test_ExampleLog(t *testing.T) {
	Log().Info("This is a test log")
}

func Test_GetObservedLogs(t *testing.T) {
	Log().Info("This is a test log")
	Log().Error("This is a test error log")

	observedLogs := GetObservedLogs()

	// range over the observed logs and print them
	for _, l := range observedLogs.All() {
		t.Log(l)

		// check if the log level is error
		if l.Entry.Level == ErrorLevel() {
			t.Log("error found in logs", l.Message)
		}
	}
}
