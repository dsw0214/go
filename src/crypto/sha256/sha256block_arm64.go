// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sha256

import "internal/cpu"

var k = _K

var hasSHA2 = cpu.ARM64.HasSHA2

func sha256block(h []uint32, p []byte, k []uint32)

func block(dig *digest, p []byte) {
	if !hasSHA2 {
		blockGeneric(dig, p)
	} else {
		h := dig.h[:]
		sha256block(h, p, k)
	}
}
