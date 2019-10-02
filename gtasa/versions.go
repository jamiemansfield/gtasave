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

package gtasa

//go:generate stringer -type=Version

type Version uint32

const (
	VERSION_1_00_UNMODIFIED Version = 0x35da8175
	VERSION_1_00_MODIFIED Version = 0x65f3e583
	VERSION_1_01_UNMODIFIED Version = 0x9a6ebe58
	VERSION_1_01_MODIFIED Version = 0x9345765e
	VERSION_2_00_UNMODIFIED Version = 0xfd148df6
	VERSION_2_00_GERMAN Version = 0x5d31cc22
	VERSION_PS2_V1 Version = 0x641ddc4c
	// TODO(same as VERSION_2_00_UNMODIFIED): VERSION_PS2_V2 Version = 0xfd148df6
)

func (i Version) MarshalJSON() ([]byte, error) {
	return []byte("\"" + i.String() + "\""), nil
}
