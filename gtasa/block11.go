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

// See https://gtasa-savegame-editor.github.io/docs/#/block11
type Block11 struct {
	// 4 bytes
	Ballas GangWeapons `gta:"index:4,length:12"`
	// 4 bytes
	GroveStreet GangWeapons `gta:"index:20,length:12"`
	// 4 bytes
	LosSantosVagos GangWeapons `gta:"index:36,length:12"`
	// 4 bytes
	SanFierroRifa GangWeapons `gta:"index:52,length:12"`
	// 4 bytes
	DuNangBoys GangWeapons `gta:"index:68,length:12"`
	// 4 bytes
	ItalianMafia GangWeapons `gta:"index:84,length:12"`
	// 4 bytes
	Triads GangWeapons `gta:"index:100,length:12"`
	// 4 bytes
	VarrioLosAztecas GangWeapons `gta:"index:116,length:12"`
}

type GangWeapons struct {
	Pistol WeaponType `gta:"index:0"`
	MachineGun WeaponType `gta:"index:4"`
	Melee WeaponType `gta:"index:8"`
}
