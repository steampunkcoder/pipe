package pipe_test

import (
	"fmt"
	"strings"
)

// MockLogger mocks interface pipe.StdLogger
type MockLogger []string

// Printf implements pipe.StdLogger interface for MockLogger
func (mock *MockLogger) Printf(format string, v ...interface{}) {
	mock.appendToLog(fmt.Sprintf(format, v...))
}

// Println implements pipe.StdLogger interface for MockLogger
func (mock *MockLogger) Println(v ...interface{}) {
	mock.appendToLog(fmt.Sprintln(v...))
}

func (mock *MockLogger) appendToLog(loggedStr string) {
	if !strings.HasSuffix(loggedStr, "\n") {
		loggedStr = loggedStr + "\n"
	}
	*mock = append(*mock, loggedStr)
}
