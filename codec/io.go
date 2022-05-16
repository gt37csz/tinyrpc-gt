// Copyright 2022 <mzh.scnu@qq.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package codec

import (
	"encoding/binary"
	"io"
	"net"
)

// sendFrame 向conn中写入response header
func sendFrame(w io.Writer, data []byte) (err error) {
	var size [binary.MaxVarintLen64]byte

	if len(data) == 0 {
		n := binary.PutUvarint(size[:], uint64(0))
		if err = write(w, size[:n]); err != nil {
			return
		}
		return
	}

	n := binary.PutUvarint(size[:], uint64(len(data)))
	if err = write(w, size[:n]); err != nil {
		return
	}
	if err = write(w, data); err != nil {
		return
	}
	return
}

// recvFrame 将conn传过来的request header读到byte slice中，如果conn中没有数据可读
// 就会返回一个EOF的error，如果有数据但读取失败也会返回error，该error会一直往上传导
func recvFrame(r io.Reader) (data []byte, err error) {
	size, err := binary.ReadUvarint(r.(io.ByteReader))
	if err != nil {
		return nil, err
	}
	if size != 0 {
		data = make([]byte, size)
		if err = read(r, data); err != nil {
			return nil, err
		}
	}
	return data, nil
}

// write 向conn中写入response body，由于data可能比较大，w的缓冲区默认是4096，因此要在for循环中分批读取
func write(w io.Writer, data []byte) error {
	for index := 0; index < len(data); {
		n, err := w.Write(data[index:])
		if _, ok := err.(net.Error); !ok {
			return err
		}
		index += n
	}
	return nil
}

// read 从conn中读取request body
func read(r io.Reader, data []byte) error {
	for index := 0; index < len(data); {
		n, err := r.Read(data[index:])
		if err != nil {
			if _, ok := err.(net.Error); !ok {
				return err
			}
		}
		index += n
	}
	return nil
}
