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

package parser

import (
	"encoding/binary"
	"github.com/jamiemansfield/gtasave/util"
	"math"
	"reflect"
)

func Parse(data []byte, v interface{}) error {
	// Separate the blocks
	var blocks [34][]byte
	var index = 0
	reader := util.CreateReader(data)
	for reader.Available() {
		if isAtBoundary(reader) {
			// Skip the block boundary
			reader.Skip(5)

			// Read the block
			blocks[index] = readBlock(reader)
			index += 1
		} else {
			reader.Skip(1)
		}
	}

	//  Get type information on the save format
	t := reflect.ValueOf(v)

	for i := 0; i < reflect.TypeOf(v).Elem().NumField(); i++ {
		field := t.Type().Elem().Field(i)

		tag, err := util.GetGtaTag(field.Tag.Get("gta"))
		if err != nil {
			return err
		}

		blockReader := util.CreateReader(blocks[tag.Index])

		// Get type information of the block
		t2 := t.Elem().Field(i)

		for j := 0; j < t2.NumField(); j++ {
			field2 := t2.Type().Field(j)

			tag, err := util.GetGtaTag(field2.Tag.Get("gta"))
			if err != nil {
				return err
			}

			read(blockReader, tag, t2.Field(j))
		}
	}

	return nil
}

func read(blockReader *util.Reader, tag *util.GtaTag, f reflect.Value) {
	if f.Type().Kind() == reflect.Int {
		raw := blockReader.Splice(tag.Index, tag.Index + 4)
		value := binary.LittleEndian.Uint32(raw)
		f.SetInt(int64(value))
	}
	if f.Type().Kind() == reflect.Uint32 {
		raw := blockReader.Splice(tag.Index, tag.Index + 4)
		value := binary.LittleEndian.Uint32(raw)
		f.Set(reflect.ValueOf(value))
	}
	if f.Type().Kind() == reflect.Float32 {
		raw := blockReader.Splice(tag.Index, tag.Index + 4)
		intBits := binary.LittleEndian.Uint32(raw)
		value := math.Float32frombits(intBits)
		f.Set(reflect.ValueOf(value))
	}
	if f.Type().Kind() == reflect.String {
		raw := blockReader.Splice(tag.Index, tag.Index + tag.Length)

		// Encoded in ASCII, padded with /raw/ zeros to make up the length.
		reader := util.CreateReader(raw)

		var length int
		for reader.Available() && reader.Peek(0) != 0 {
			reader.Skip(1)
			length += 1
		}

		f.SetString(string(reader.Splice(0, length)))
	}
	if f.Type().Kind() == reflect.Bool {
		raw := blockReader.Splice(tag.Index, tag.Index + 1)
		f.SetBool(raw[0] != 0)
	}
	if f.Type().Kind() == reflect.Uint8 {
		raw := blockReader.Splice(tag.Index, tag.Index + 4)
		f.Set(reflect.ValueOf(uint8(raw[0])))
	}
}

func isAtBoundary(r *util.Reader) bool {
	return r.Peek(0) == 'B' &&
		r.Peek(1) == 'L' &&
		r.Peek(2) == 'O' &&
		r.Peek(3) == 'C' &&
		r.Peek(4) == 'K'
}

func readBlock(r *util.Reader) []byte {
	var start = r.Index()

	for r.Available() && !isAtBoundary(r) {
		r.Skip(1)
	}

	return r.Splice(start, r.Index())
}
