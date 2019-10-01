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

// See https://gtasa-savegame-editor.github.io/docs/#/weapons
const (
	WEAPON_NONE uint32 = 0x00
	WEAPON_BRASS_KNUCKLES uint32 = 0x01
	WEAPON_GOLF_CLUB uint32 = 0x02
	WEAPON_NIGHTSTICK uint32 = 0x03
	WEAPON_KNIFE uint32 = 0x04
	WEAPON_BASEBALL_BAT uint32 = 0x05
	WEAPON_SHOVEL uint32 = 0x06
	WEAPON_POOL_CUE uint32 = 0x07
	WEAPON_KATANA uint32 = 0x08
	WEAPON_CHAINSAW uint32 = 0x09
	WEAPON_DILDO_1 uint32 = 0x0A
	WEAPON_DILDO_2 uint32 = 0x0B
	WEAPON_VIBRATOR_1 uint32 = 0x0C
	WEAPON_VIBRATOR_2 uint32 = 0x0D
	WEAPON_FLOWERS uint32 = 0x0E
	WEAPON_CANE uint32 = 0x0F

	// Throwable
	WEAPON_GRENADE uint32 = 0x10
	WEAPON_TEAR_GAS uint32 = 0x11
	WEAPON_MOLOTOV_COCKTAIL uint32 = 0x12

	// Guns
	WEAPON_PISTOL_9MM uint32 = 0x16
	WEAPON_PISTOL_SILENCED_9MM uint32 = 0x17
	WEAPON_PISTOL_DESERT_EAGLE uint32 = 0x18
	WEAPON_SHOTGUN uint32 = 0x19
	WEAPON_SHOTGUN_SAWNOFF uint32 = 0x1A
	WEAPON_SHOTGUN_COMBAT uint32 = 0x1B
	WEAPON_MICRO_SMG uint32 = 0x1C
	WEAPON_SMG uint32 = 0x1D
	WEAPON_AK47 uint32 = 0x1E
	WEAPON_M4 uint32 = 0x1F
	WEAPON_TEC9 uint32 = 0x20
	WEAPON_RIFLE uint32 = 0x21
	WEAPON_SNIPER_RIFLE uint32 = 0x22

	WEAPON_ROCKER_LAUNCHER uint32 = 0x23
	WEAPON_ROCKET_LAUNCHER_HEAT_SINKING uint32 = 0x24
	WEAPON_FLAME_THROWER uint32 = 0x25
	WEAPON_MINI_GUN uint32 = 0x26
	WEAPON_REMOTE_EXPLOSIVE uint32 = 0x27
	WEAPON_REMOTE_DETONATOR uint32 = 0x28
	WEAPON_SPRAY_CAN uint32 = 0x29
	WEAPON_FIRE_EXTINGUISHER uint32 = 0x2A
	WEAPON_CAMERA uint32 = 0x2B

	// Wearable
	WEAPON_NIGHT_VISION_GOGGLES uint32 = 0x2C
	WEAPON_THERMAL_GOGGLES uint32 = 0x2D
	WEAPON_PARACHUTE uint32 = 0x2E
)
