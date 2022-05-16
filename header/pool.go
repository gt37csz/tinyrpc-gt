// Copyright 2022 <mzh.scnu@qq.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// 使用syn.Pool构造了一个请求池和响应池，这是上面加锁的原因吗？？？
package header

import "sync"

var (
	RequestPool  sync.Pool
	ResponsePool sync.Pool
)

func init() {
	RequestPool = sync.Pool{New: func() interface{} {
		return &RequestHeader{}
	}}
	ResponsePool = sync.Pool{New: func() interface{} {
		return &ResponseHeader{}
	}}
}
