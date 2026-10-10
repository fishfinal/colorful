// Copyright (c) 2026 fishfinal
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package colorful

import (
	"fmt"
	"testing"

	"github.com/logrusorgru/aurora"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ColorfulSuite struct {
	suite.Suite
	originalColorful aurora.Aurora
}

func (s *ColorfulSuite) SetupSuite() {
	s.originalColorful = Colorful
}

func (s *ColorfulSuite) TearDownSuite() {
	Colorful = s.originalColorful
}

func (s *ColorfulSuite) SetupTest() {
	SetEnableColors(true) // reset to the default state
}

func TestColorfulSuite(t *testing.T) {
	suite.Run(t, new(ColorfulSuite))
}

func (s *ColorfulSuite) TestSetEnableColors_ChangesBehavior() {
	// Test the color-enabled state
	SetEnableColors(true)
	colored := Colorful.Colorize("text", aurora.RedFg).String()
	assert.NotEqual(s.T(), "text", colored, "text should be colored when colors are enabled")

	// Test the color-disabled state
	SetEnableColors(false)
	plain := Colorful.Colorize("text", aurora.RedFg).String()
	assert.Equal(s.T(), "text", plain, "the original text should be returned when colors are disabled")
}

func (s *ColorfulSuite) TestSetEnableColors_CreatesNewInstance() {
	// Save a characteristic value of the original instance (e.g., the output of a specific method)
	originalOutput := Colorful.Red("test").String()

	// Change the state for the first time
	SetEnableColors(false)
	disabledOutput := Colorful.Red("test").String()
	s.NotEqual(originalOutput, disabledOutput, "the output should differ after disabling colors")

	// Change the state for the second time (restore)
	SetEnableColors(true)
	restoredOutput := Colorful.Red("test").String()
	s.Equal(originalOutput, restoredOutput, "the output after restoring colors should match the original")

	// Verify the number of behavior changes (indirectly verifying instance changes)
	// Here we verify behavior changes rather than directly comparing instances
	s.NotEqual(disabledOutput, restoredOutput, "the outputs in the disabled and enabled states should differ")
}

func (s *ColorfulSuite) TestColorMethodsWhenDisabled() {
	SetEnableColors(false)

	tests := []struct {
		name     string
		colored  string
		expected string
	}{
		{"Red", Colorful.Red("text").String(), "text"},
		{"Green", Colorful.Green("text").String(), "text"},
		{"Bold", Colorful.Bold("text").String(), "text"},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			assert.Equal(s.T(), tt.expected, tt.colored)
		})
	}
}

func (s *ColorfulSuite) TestSeverityColorfully() {
	SetEnableColors(true)

	tests := []struct {
		name     string
		colored  string
		expected string
	}{
		{"Critical", SeverityCritical("critical").String(), "\x1b[38;5;128mcritical\x1b[0m"},
		{"High", SeverityHigh("high").String(), "\x1b[31mhigh\x1b[0m"},
		{"Medium", SeverityMedium("medium").String(), "\x1b[38;5;208mmedium\x1b[0m"},
		{"low", SeverityLow("low").String(), "\x1b[33mlow\x1b[0m"},
		{"Info", SeverityInfo("info").String(), "\x1b[36minfo\x1b[0m"},
		{"Unknown", SeverityUnknown("unknown").String(), "\x1b[37munknown\x1b[0m"},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			assert.Equal(s.T(), tt.expected, tt.colored)
			fmt.Println(tt.expected)
		})
	}
}

func (s *ColorfulSuite) TestColorMethodsWhenEnabled() {
	SetEnableColors(true)

	tests := []struct {
		name     string
		colored  string
		expected string
	}{
		{"Red", Colorful.Red("text").String(), "\x1b[31mtext\x1b[0m"},
		{"Green", Colorful.Green("text").String(), "\x1b[32mtext\x1b[0m"},
		{"Bold", Colorful.Bold("text").String(), "\x1b[1mtext\x1b[0m"},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			assert.Equal(s.T(), tt.expected, tt.colored)
		})
	}
}

func (s *ColorfulSuite) TestReviseColorfully() {
	SetEnableColors(true)

	tests := []struct {
		name     string
		colored  string
		expected string
	}{

		{"Red", ReviseColorfully(Colorful.Red("text"), true).String(), "\x1b[31mtext\x1b[0m"},
		{"None Color", ReviseColorfully(Colorful.Green("text"), false).String(), "text"},
		{"Bold", ReviseColorfully(Colorful.Bold("text"), true).String(), "\x1b[1mtext\x1b[0m"},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			assert.Equal(s.T(), tt.expected, tt.colored)
		})
	}
}
