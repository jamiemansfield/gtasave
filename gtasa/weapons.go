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

//go:generate stringer -type=WeaponType

type WeaponType uint32

// See https://gtasa-savegame-editor.github.io/docs/#/weapons
const (
	WEAPON_NONE WeaponType = 0x00
	WEAPON_BRASS_KNUCKLES WeaponType = 0x01
	WEAPON_GOLF_CLUB WeaponType = 0x02
	WEAPON_NIGHTSTICK WeaponType = 0x03
	WEAPON_KNIFE WeaponType = 0x04
	WEAPON_BASEBALL_BAT WeaponType = 0x05
	WEAPON_SHOVEL WeaponType = 0x06
	WEAPON_POOL_CUE WeaponType = 0x07
	WEAPON_KATANA WeaponType = 0x08
	WEAPON_CHAINSAW WeaponType = 0x09
	WEAPON_DILDO_1 WeaponType = 0x0A
	WEAPON_DILDO_2 WeaponType = 0x0B
	WEAPON_VIBRATOR_1 WeaponType = 0x0C
	WEAPON_VIBRATOR_2 WeaponType = 0x0D
	WEAPON_FLOWERS WeaponType = 0x0E
	WEAPON_CANE WeaponType = 0x0F

	// Throwable
	WEAPON_GRENADE WeaponType = 0x10
	WEAPON_TEAR_GAS WeaponType = 0x11
	WEAPON_MOLOTOV_COCKTAIL WeaponType = 0x12

	// Guns
	WEAPON_PISTOL_9MM WeaponType = 0x16
	WEAPON_PISTOL_SILENCED_9MM WeaponType = 0x17
	WEAPON_PISTOL_DESERT_EAGLE WeaponType = 0x18
	WEAPON_SHOTGUN WeaponType = 0x19
	WEAPON_SHOTGUN_SAWNOFF WeaponType = 0x1A
	WEAPON_SHOTGUN_COMBAT WeaponType = 0x1B
	WEAPON_MICRO_SMG WeaponType = 0x1C
	WEAPON_SMG WeaponType = 0x1D
	WEAPON_AK47 WeaponType = 0x1E
	WEAPON_M4 WeaponType = 0x1F
	WEAPON_TEC9 WeaponType = 0x20
	WEAPON_RIFLE WeaponType = 0x21
	WEAPON_SNIPER_RIFLE WeaponType = 0x22

	WEAPON_ROCKER_LAUNCHER WeaponType = 0x23
	WEAPON_ROCKET_LAUNCHER_HEAT_SINKING WeaponType = 0x24
	WEAPON_FLAME_THROWER WeaponType = 0x25
	WEAPON_MINI_GUN WeaponType = 0x26
	WEAPON_REMOTE_EXPLOSIVE WeaponType = 0x27
	WEAPON_REMOTE_DETONATOR WeaponType = 0x28
	WEAPON_SPRAY_CAN WeaponType = 0x29
	WEAPON_FIRE_EXTINGUISHER WeaponType = 0x2A
	WEAPON_CAMERA WeaponType = 0x2B

	// Wearable
	WEAPON_NIGHT_VISION_GOGGLES WeaponType = 0x2C
	WEAPON_THERMAL_GOGGLES WeaponType = 0x2D
	WEAPON_PARACHUTE WeaponType = 0x2E
)

func (i WeaponType) MarshalJSON() ([]byte, error) {
	return []byte("\"" + i.String() + "\""), nil
}

