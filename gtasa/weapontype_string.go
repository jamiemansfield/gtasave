// Code generated by "stringer -type=WeaponType"; DO NOT EDIT.

package gtasa

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[WEAPON_NONE-0]
	_ = x[WEAPON_BRASS_KNUCKLES-1]
	_ = x[WEAPON_GOLF_CLUB-2]
	_ = x[WEAPON_NIGHTSTICK-3]
	_ = x[WEAPON_KNIFE-4]
	_ = x[WEAPON_BASEBALL_BAT-5]
	_ = x[WEAPON_SHOVEL-6]
	_ = x[WEAPON_POOL_CUE-7]
	_ = x[WEAPON_KATANA-8]
	_ = x[WEAPON_CHAINSAW-9]
	_ = x[WEAPON_DILDO_1-10]
	_ = x[WEAPON_DILDO_2-11]
	_ = x[WEAPON_VIBRATOR_1-12]
	_ = x[WEAPON_VIBRATOR_2-13]
	_ = x[WEAPON_FLOWERS-14]
	_ = x[WEAPON_CANE-15]
	_ = x[WEAPON_GRENADE-16]
	_ = x[WEAPON_TEAR_GAS-17]
	_ = x[WEAPON_MOLOTOV_COCKTAIL-18]
	_ = x[WEAPON_PISTOL_9MM-22]
	_ = x[WEAPON_PISTOL_SILENCED_9MM-23]
	_ = x[WEAPON_PISTOL_DESERT_EAGLE-24]
	_ = x[WEAPON_SHOTGUN-25]
	_ = x[WEAPON_SHOTGUN_SAWNOFF-26]
	_ = x[WEAPON_SHOTGUN_COMBAT-27]
	_ = x[WEAPON_MICRO_SMG-28]
	_ = x[WEAPON_SMG-29]
	_ = x[WEAPON_AK47-30]
	_ = x[WEAPON_M4-31]
	_ = x[WEAPON_TEC9-32]
	_ = x[WEAPON_RIFLE-33]
	_ = x[WEAPON_SNIPER_RIFLE-34]
	_ = x[WEAPON_ROCKER_LAUNCHER-35]
	_ = x[WEAPON_ROCKET_LAUNCHER_HEAT_SINKING-36]
	_ = x[WEAPON_FLAME_THROWER-37]
	_ = x[WEAPON_MINI_GUN-38]
	_ = x[WEAPON_REMOTE_EXPLOSIVE-39]
	_ = x[WEAPON_REMOTE_DETONATOR-40]
	_ = x[WEAPON_SPRAY_CAN-41]
	_ = x[WEAPON_FIRE_EXTINGUISHER-42]
	_ = x[WEAPON_CAMERA-43]
	_ = x[WEAPON_NIGHT_VISION_GOGGLES-44]
	_ = x[WEAPON_THERMAL_GOGGLES-45]
	_ = x[WEAPON_PARACHUTE-46]
}

const (
	_WeaponType_name_0 = "WEAPON_NONEWEAPON_BRASS_KNUCKLESWEAPON_GOLF_CLUBWEAPON_NIGHTSTICKWEAPON_KNIFEWEAPON_BASEBALL_BATWEAPON_SHOVELWEAPON_POOL_CUEWEAPON_KATANAWEAPON_CHAINSAWWEAPON_DILDO_1WEAPON_DILDO_2WEAPON_VIBRATOR_1WEAPON_VIBRATOR_2WEAPON_FLOWERSWEAPON_CANEWEAPON_GRENADEWEAPON_TEAR_GASWEAPON_MOLOTOV_COCKTAIL"
	_WeaponType_name_1 = "WEAPON_PISTOL_9MMWEAPON_PISTOL_SILENCED_9MMWEAPON_PISTOL_DESERT_EAGLEWEAPON_SHOTGUNWEAPON_SHOTGUN_SAWNOFFWEAPON_SHOTGUN_COMBATWEAPON_MICRO_SMGWEAPON_SMGWEAPON_AK47WEAPON_M4WEAPON_TEC9WEAPON_RIFLEWEAPON_SNIPER_RIFLEWEAPON_ROCKER_LAUNCHERWEAPON_ROCKET_LAUNCHER_HEAT_SINKINGWEAPON_FLAME_THROWERWEAPON_MINI_GUNWEAPON_REMOTE_EXPLOSIVEWEAPON_REMOTE_DETONATORWEAPON_SPRAY_CANWEAPON_FIRE_EXTINGUISHERWEAPON_CAMERAWEAPON_NIGHT_VISION_GOGGLESWEAPON_THERMAL_GOGGLESWEAPON_PARACHUTE"
)

var (
	_WeaponType_index_0 = [...]uint16{0, 11, 32, 48, 65, 77, 96, 109, 124, 137, 152, 166, 180, 197, 214, 228, 239, 253, 268, 291}
	_WeaponType_index_1 = [...]uint16{0, 17, 43, 69, 83, 105, 126, 142, 152, 163, 172, 183, 195, 214, 236, 271, 291, 306, 329, 352, 368, 392, 405, 432, 454, 470}
)

func (i WeaponType) String() string {
	switch {
	case 0 <= i && i <= 18:
		return _WeaponType_name_0[_WeaponType_index_0[i]:_WeaponType_index_0[i+1]]
	case 22 <= i && i <= 46:
		i -= 22
		return _WeaponType_name_1[_WeaponType_index_1[i]:_WeaponType_index_1[i+1]]
	default:
		return "WeaponType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}