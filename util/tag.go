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

import (
	"strconv"
	"strings"
)

type SliceLengthType int

const (
	SliceUint16 SliceLengthType = iota
	SliceUint32
)

func (t SliceLengthType) GetLength() int {
	switch t {
	default:
		return 4
	case SliceUint32:
		return 4
	case SliceUint16:
		return 2
	}
}

type GtaTag struct {
	Index int
	Length int

	// Slices
	SliceLengthType SliceLengthType
}

func GetGtaTag(tag string) (*GtaTag, error) {
	args := strings.Split(tag, ",")
	var index int
	var length int
	var sliceType = SliceUint32

	for i := 0; i < len(args); i++ {
		if strings.HasPrefix(args[i], "index:") {
			i, err := strconv.Atoi(args[i][6:])
			if err != nil {
				return nil, err
			}
			index = i
		}
		if strings.HasPrefix(args[i], "length:") {
			i, err := strconv.Atoi(args[i][7:])
			if err != nil {
				return nil, err
			}
			length = i
		}
		if strings.HasPrefix(args[i], "slice_type:") {
			val := args[i][len("slice_type:"):]
			switch val {
			case "uint32":
				sliceType = SliceUint32
			case "uint16":
				sliceType = SliceUint16
			}
		}
	}

	return &GtaTag{
		Index: index,
		Length: length,
		SliceLengthType: sliceType,
	}, nil
}
