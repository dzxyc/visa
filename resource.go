// Copyright (c) 2017 The visa developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package visa

import "io"

type Resource interface {
	io.ReadWriteCloser
	WriterString(s string) (n int, err error)
	Query(s string) (value string, err error)
}
