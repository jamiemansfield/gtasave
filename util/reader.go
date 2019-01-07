/*
 * This file is part of GtaSave, licensed under the MIT License (MIT).
 *
 * Copyright (c) Jamie Mansfield <https://www.jamiemansfield.me/>
 * Copyright (c) contributors
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 */

package util

type Reader struct {
	data []byte
	index int
}

// Gets the length of the data
func (r *Reader) Length() int {
	return len(r.data)
}

// Gets the remaining length to be read
func (r *Reader) Remaining() int {
	return len(r.data) - r.index
}

// Gets the current index
func (r *Reader) Index() int {
	return r.index
}

// Establishes whether there is a readable byte at the given offset
func (r *Reader) Readable(offset int) bool {
	return r.index + offset <= len(r.data)
}

// Establishes whether there is data available to be read
func (r *Reader) Available() bool {
	return r.Readable(1)
}

// Skips the given length of values
func (r *Reader) Skip(length int) {
	r.index += length
}

// Peeks at the value of the given offset
func (r *Reader) Peek(offset int) byte {
	return r.data[r.index + offset]
}

// Splices the data
func (r *Reader) Splice(start int, end int) []byte {
	return r.data[start : end]
}

func CreateReader(data []byte) *Reader {
	return &Reader{
		data: data,
		index: 0,
	}
}
