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
	Ballas GangInfo `gta:"index:0,length:16"`
	GroveStreet GangInfo `gta:"index:16,length:16"`
	LosSantosVagos GangInfo `gta:"index:32,length:16"`
	SanFierroRifa GangInfo `gta:"index:48,length:16"`
	DuNangBoys GangInfo `gta:"index:64,length:16"`
	ItalianMafia GangInfo `gta:"index:80,length:16"`
	Triads GangInfo `gta:"index:96,length:16"`
	VarrioLosAztecas GangInfo `gta:"index:112,length:16"`
}

type GangInfo struct {
	PedTypeOverride int8 `gta:"index:0"`
	Pistol WeaponType `gta:"index:4"`
	MachineGun WeaponType `gta:"index:8"`
	Melee WeaponType `gta:"index:12"`
}
