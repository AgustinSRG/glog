// Tests

package glog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateLoggerConfigurationFromLevel(t *testing.T) {
	var config LoggerConfiguration

	config = CreateLoggerConfigurationFromLevel(SILENT)

	assert.Equal(t, config.ErrorEnabled, false)
	assert.Equal(t, config.WarningEnabled, false)
	assert.Equal(t, config.InfoEnabled, false)
	assert.Equal(t, config.DebugEnabled, false)
	assert.Equal(t, config.TraceEnabled, false)

	config = CreateLoggerConfigurationFromLevel(ERROR)

	assert.Equal(t, config.ErrorEnabled, true)
	assert.Equal(t, config.WarningEnabled, false)
	assert.Equal(t, config.InfoEnabled, false)
	assert.Equal(t, config.DebugEnabled, false)
	assert.Equal(t, config.TraceEnabled, false)

	config = CreateLoggerConfigurationFromLevel(WARNING)

	assert.Equal(t, config.ErrorEnabled, true)
	assert.Equal(t, config.WarningEnabled, true)
	assert.Equal(t, config.InfoEnabled, false)
	assert.Equal(t, config.DebugEnabled, false)
	assert.Equal(t, config.TraceEnabled, false)

	config = CreateLoggerConfigurationFromLevel(INFO)

	assert.Equal(t, config.ErrorEnabled, true)
	assert.Equal(t, config.WarningEnabled, true)
	assert.Equal(t, config.InfoEnabled, true)
	assert.Equal(t, config.DebugEnabled, false)
	assert.Equal(t, config.TraceEnabled, false)

	config = CreateLoggerConfigurationFromLevel(DEBUG)

	assert.Equal(t, config.ErrorEnabled, true)
	assert.Equal(t, config.WarningEnabled, true)
	assert.Equal(t, config.InfoEnabled, true)
	assert.Equal(t, config.DebugEnabled, true)
	assert.Equal(t, config.TraceEnabled, false)

	config = CreateLoggerConfigurationFromLevel(TRACE)

	assert.Equal(t, config.ErrorEnabled, true)
	assert.Equal(t, config.WarningEnabled, true)
	assert.Equal(t, config.InfoEnabled, true)
	assert.Equal(t, config.DebugEnabled, true)
	assert.Equal(t, config.TraceEnabled, true)
}
