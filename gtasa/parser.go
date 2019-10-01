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

import (
	"github.com/jamiemansfield/gtasave/io"
	"github.com/jamiemansfield/gtasave/util"
	"math"
	"reflect"
)

func Parse(data []byte, v interface{}) error {
	// Separate the blocks
	var blocks [34][]byte
	var index = 0
	reader := io.CreateReader(data)
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

		blockReader := io.CreateReader(blocks[tag.Index])

		// Get type information of the block
		t2 := t.Elem().Field(i)

		// Read the struct
		readStruct(blockReader, t2.Type(), t2)
	}

	return nil
}

func read(reader *io.Reader, tag *util.GtaTag, f reflect.Value) {
	if f.Type().Kind() == reflect.Int {
		value := reader.ReadUInt32(tag.Index)
		f.SetInt(int64(value))
	}
	if f.Type().Kind() == reflect.Uint8 {
		value := reader.ReadUInt8(tag.Index)
		f.Set(reflect.ValueOf(value))
	}
	if f.Type().Kind() == reflect.Uint16 {
		value := reader.ReadUInt16(tag.Index)
		f.Set(reflect.ValueOf(value))
	}
	if f.Type().Kind() == reflect.Uint32 {
		value := reader.ReadUInt32(tag.Index)
		f.Set(reflect.ValueOf(value))
	}
	if f.Type().Kind() == reflect.Float32 {
		intBits := reader.ReadUInt32(tag.Index)
		value := math.Float32frombits(intBits)
		f.Set(reflect.ValueOf(value))
	}
	if f.Type().Kind() == reflect.String {
		raw := reader.Splice(tag.Index, tag.Length)

		// Encoded in ASCII, padded with /raw/ zeros to make up the length.
		reader := io.CreateReader(raw)

		var length int
		for reader.Available() && reader.Peek(0) != 0 {
			reader.Skip(1)
			length += 1
		}

		f.SetString(string(reader.Splice(-length, length)))
	}
	if f.Type().Kind() == reflect.Bool {
		f.SetBool(reader.ReadBool(tag.Index))
	}
	if f.Type().Kind() == reflect.Struct {
		raw := reader.Splice(tag.Index, tag.Length)
		reader := io.CreateReader(raw)

		readStruct(reader, f.Type(), f)
	}
	if f.Type().Kind() == reflect.Slice {
		// 4 bytes for length of array
		length := int(reader.ReadUInt32(tag.Index))

		raw := reader.Splice(tag.Index + 4, tag.Length * length)
		reader := io.CreateReader(raw)

		f.Set(reflect.MakeSlice(f.Type(), length, length))

		for i := 0; i < length; i++ {
			subReader := io.CreateReader(reader.Splice(i * length, length))

			read(subReader, &util.GtaTag{
				// We're using a new reader, so need to start fresh
				Index: 0,
				Length: tag.Length,
			}, f.Index(i))
		}
	}
	if f.Type().Kind() == reflect.Array {
		raw := reader.Splice(tag.Index, f.Len() * tag.Length)
		reader := io.CreateReader(raw)

		for i := 0; i < f.Len(); i++ {
			subReader := io.CreateReader(reader.Splice(i * f.Len(), f.Len()))

			read(subReader, &util.GtaTag{
				// We're using a new reader, so need to start fresh
				Index: 0,
				Length: tag.Length,
			}, f.Index(i))
		}
	}
}

// This method populates a struct's fields with values from the reader.
func readStruct(reader *io.Reader, structType reflect.Type, structValue reflect.Value) {
	// Iterate over the structs fields
	for i := 0; i < structType.NumField(); i++ {
		fieldType := structType.Field(i)
		fieldValue := structValue.Field(i)

		tag, err := util.GetGtaTag(fieldType.Tag.Get("gta"))
		if err != nil {
			return
		}

		read(reader, tag, fieldValue)
	}
}

func isAtBoundary(r *io.Reader) bool {
	return r.Peek(0) == 'B' &&
		r.Peek(1) == 'L' &&
		r.Peek(2) == 'O' &&
		r.Peek(3) == 'C' &&
		r.Peek(4) == 'K'
}

func readBlock(r *io.Reader) []byte {
	var start = r.Index()

	for r.Available() && !isAtBoundary(r) {
		r.Skip(1)
	}

	return r.Splice(start - r.Index(), r.Index() - start)
}
