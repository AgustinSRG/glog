// Tests

package glog

import "testing"

func TestStandardLogFunction(t *testing.T) {
	StandardLogFunction(SILENT, "test log message")
	StandardLogFunction(ERROR, "test log message")
	StandardLogFunction(WARNING, "test log message")
	StandardLogFunction(INFO, "test log message")
	StandardLogFunction(DEBUG, "test log message")
	StandardLogFunction(TRACE, "test log message")
}
