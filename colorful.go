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

// The ANSI standard color codes can be found in this project: https://github.com/fidian/ansi

import "github.com/logrusorgru/aurora"

var enableColors = true // lowercase, cannot be exported directly
var Colorful aurora.Aurora = aurora.NewAurora(enableColors)

// SetEnableColors allows external control of the color switch
func SetEnableColors(enable bool) {
	enableColors = enable
	Colorful = aurora.NewAurora(enableColors) // reinitialize
}

// ReviseColorfully dynamically clears the color value.
func ReviseColorfully(value aurora.Value, colorfully bool) aurora.Value {
	if colorfully {
		return value
	}
	return Colorful.Reset(value)
}

// Purple is a purple color renderer.
func Purple(text string) aurora.Value {
	// ANSI standard color code; purple is 128
	return Colorful.Index(128, text)
}

// Orange is an orange color renderer.
func Orange(text string) aurora.Value {
	// ANSI standard color code; orange is 203
	return Colorful.Index(208, text)
}

// Gray is a gray color renderer.
func Gray(text string) aurora.Value {
	return Colorful.Index(7, text)
}

type SeverityColorfully func(text string) aurora.Value

// SeverityCritical is the color renderer for critical vulnerabilities.
var SeverityCritical SeverityColorfully = func(text string) aurora.Value {
	return Purple(text)
}

// SeverityHigh is the color renderer for high-severity vulnerabilities.
var SeverityHigh SeverityColorfully = func(text string) aurora.Value {
	return Colorful.Red(text)
}

// SeverityMedium is the color renderer for medium-severity vulnerabilities.
var SeverityMedium SeverityColorfully = func(text string) aurora.Value {
	return Orange(text)
}

// SeverityLow is the color renderer for low-severity vulnerabilities.
var SeverityLow SeverityColorfully = func(text string) aurora.Value {
	return Colorful.Yellow(text)
}

// SeverityInfo is the color renderer for informational vulnerabilities.
var SeverityInfo SeverityColorfully = func(text string) aurora.Value {
	return Colorful.Cyan(text)
}

// SeverityUnknown is the color renderer for unknown-severity vulnerabilities.
var SeverityUnknown SeverityColorfully = func(text string) aurora.Value {
	return Gray(text)
}

var _ SeverityColorfully = SeverityCritical
var _ SeverityColorfully = SeverityHigh
var _ SeverityColorfully = SeverityMedium
var _ SeverityColorfully = SeverityUnknown
var _ SeverityColorfully = SeverityLow
