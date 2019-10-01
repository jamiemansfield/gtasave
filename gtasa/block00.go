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

const (
	VERSION_1_00_UNMODIFIED uint32 = 0x35da8175
	VERSION_1_00_MODIFIED uint32 = 0x65f3e583
	VERSION_1_01_UNMODIFIED uint32 = 0x9a6ebe58
	VERSION_1_01_MODIFIED uint32 = 0x9345765e
	VERSION_2_00_UNMODIFIED uint32 = 0xfd148df6
	VERSION_2_00_GERMAN uint32 = 0x5d31cc22
	VERSION_PS2_V1 uint32 = 0x641ddc4c
	VERSION_PS2_V2 uint32 = 0xfd148df6
)

// See https://gtasa-savegame-editor.github.io/docs/#/block00
type Block00 struct {
	Version uint32 `gta:"index:0"`
	SaveName string `gta:"index:4,length:100"`
	MissionPack uint32 `gta:"index:104"` // TODO: source?
	CurrentTown uint32 `gta:"index:108"` // TODO: source?
	Camera Vector3f `gta:"index:112,length:12"`
	MinuteLength uint32 `gta:"index:124"`
	WeatherTimer uint32 `gta:"index:128"`

	Cheated bool `gta:"index:144"`

	GlobalTimer uint32 `gta:"index:148"`

	FrenchLanguage bool `gta:"index:236"`
	GermanLanguage bool `gta:"index:237"`
	FreeFromCensoring bool `gta:"index:238"`

	SystemYear uint16 `gta:"index:286"`
	SystemMonth uint16 `gta:"index:288"`
	SystemDayOfWeek uint16 `gta:"index:290"`
	SystemDay uint16 `gta:"index:292"`
	SystemHour uint16 `gta:"index:294"`
	SystemMinute uint16 `gta:"index:296"`
	SystemSecond uint16 `gta:"index:298"`
	SystemMillisecond uint16 `gta:"index:300"`

	TaxiNitro bool `gta:"index:304"`
	PaidByProstitutes bool `gta:"index:305"`
}
