// Package utils
package utils

import (
	"fmt"
	"strings"
)


// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// StringJoin concatenates the elements of its first argument to create a single string. The separator
// string sep is placed between elements in the resulting string.
func StringJoin(elems []string, sep, lastSep string) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return fmt.Sprintf("%s%s", elems[0], lastSep)
	}
	n := len(sep) * (len(elems) - 1)
	for i := 0; i < len(elems); i++ {
		n += len(elems[i])
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(elems[0])
	for _, s := range elems[1:] {
		b.WriteString(sep)
		b.WriteString(s)
	}

	if lastSep != "" {
		b.WriteString(lastSep)
	}

	return b.String()
}


// StringReplacer replace string contains
func StringReplacer( msg string, replacers map[string]string,) string {
	for k, v := range replacers {
		msg = strings.ReplaceAll(msg, k, v)
	}
	return msg
}