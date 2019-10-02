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

// See https://gtasa-savegame-editor.github.io/docs/#/block02
// See https://gtamods.com/wiki/Saves_(GTA_SA)#Block_2:_Players_.26_Objects
type Block02 struct {
	Players []Player `gta:"index:0,length:548"`
	Objects []Object `gta:"index:552"` // todo: use offset rather than index
}

type Player struct {
	Handle uint32 `gta:"index:0"`
	ModelId uint32 `gta:"index:4"`
	PedType uint32 `gta:"index:8"`
	Position Vector3f `gta:"index:16"`
	Health float32 `gta:"index:28"`
	Armor float32 `gta:"index:32"`
	Weapons [13]Weapon `gta:"index:36,length:28"`
	CurrentWeaponSlot uint8 `gta:"index:401"`
}

type Weapon struct {
	Type uint32 `gta:"index:0"` // todo: WeaponType
	Ammo uint32 `gta:"index:12"`
}

type Object struct {
	Position Vector3f `gta:"index:16"`
}
