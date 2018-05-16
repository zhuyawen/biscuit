// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bytealg

// Empirical data shows that using IndexShortStr can get better
// performance when len(s) <= 16.
const MaxBruteForce = 16

func init() {
	// 8 bytes can be completely loaded into 1 register.
	MaxLen = 8
}

// Cutover reports the number of failures of IndexByte we should tolerate
// before switching over to IndexShortStr.
// n is the number of bytes processed so far.
// See the bytes.Index implementation for details.
func Cutover(n int) int {
	// 1 error per 16 characters, plus a few slop to start.
	return 4 + n>>4
}
