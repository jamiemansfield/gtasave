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

// See https://gtasa-savegame-editor.github.io/docs/#/block16
type Block16 struct {
	ProgressMade float32 `gta:"index:0"`
	MaxProgress float32 `gta:"index:4"`

	DistanceTravelledByFoot float32 `gta:"index:12"`
	DistanceTravelledByCar float32 `gta:"index:16"`
	DistanceTravelledByMotorbike float32 `gta:"index:20"`
	DistanceTravelledByBoat float32 `gta:"index:24"`
	DistanceTravelledByGolfCart float32 `gta:"index:28"`
	DistanceTravelledByHelicopter float32 `gta:"index:32"`
	DistanceTravelledByPlane float32 `gta:"index:36"`
	LongestWheelieDistance float32 `gta:"index:40"`
	LongestStoppieDistance float32 `gta:"index:44"`
	LongestTwoWheelsDistance float32 `gta:"index:48"`
	BudgetWeapons float32 `gta:"index:52"`
	BudgetFashion float32 `gta:"index:56"`
	BudgetProperty float32 `gta:"index:60"`
	BudgetAutoRepair float32 `gta:"index:64"`
	LongestWheelieTime float32 `gta:"index:68"`
	LongestStoppieTime float32 `gta:"index:72"`
	LongestTwoWheelsTime float32 `gta:"index:76"`
	BudgetFood float32 `gta:"index:80"`
	Fat float32 `gta:"index:84"`
	Stamina float32 `gta:"index:88"`
	Muscle float32 `gta:"index:92"`
	MaxHealth float32 `gta:"index:100"`
	// 104: unknown
	DistanceTravelledBySwimming float32 `gta:"index:108"`
	DistanceTravelledByBicycle float32 `gta:"index:112"`
	DistanceTravelledByTreadmill float32 `gta:"index:116"`
	DistanceTravelledByExerciseBike float32 `gta:"index:120"`
	BudgetTattoo float32 `gta:"index:124"`
	BudgetHairdressing float32 `gta:"index:128"`
	// 132: unknown
	BudgetProstitute float32 `gta:"index:136"`
	// 140: unknown
	GamblingExpenditure float32 `gta:"index:144"`
	PimpRevenue float32 `gta:"index:148"`
	GamblingRevenue float32 `gta:"index:152"`
	GamblingLargestWin float32 `gta:"index:156"`
	GamblingLargestLoss float32 `gta:"index:160"`
	BurglaryLargestSwag float32 `gta:"index:164"`
	BurglaryRevenue float32 `gta:"index:168"`

	WastedByOthers int `gta:"index:328"`
	Wasted int `gta:"index:332"`
	LandVehiclesDestroyed int `gta:"index:336"`
	SeaVehiclesDestroyed int `gta:"index:340"`
	AirVehiclesDestroyed int `gta:"index:344"`
	PropertyDamage int `gta:"index:348"`
	BulletsFired int `gta:"index:352"`
	// TODO: explosives used?
	BulletsHit int `gta:"index:360"`
	TiresPoppedByGunfire int `gta:"index:364"`
	Headshots int `gta:"index:368"`
	TotalWantedStars int `gta:"index:372"`
	TotalWantedStarsEvaded int `gta:"index:376"`
	Busted int `gta:"index:380"`
	DaysPassed int `gta:"index:384"`

	MissionsAttempted int `gta:"index:432"`
	MissionsPassed int `gta:"index:436"`

	TaxiRevenue int `gta:"index:444"`
	TaxiPassengers int `gta:"index:448"`
	ParamedicPeopleSaved int `gta:"index:452"`
	VigilanteCriminalsKilled int `gta:"index:456"`
	FirefighterFiresExtinguished int `gta:"index:460"`

	VigilanteHighestLevel int `gta:"index:476"`
	ParamedicHighestLevel int `gta:"index:480"`
	FirefighterHighestLevel int `gta:"index:484"`
	SkillDriving int `gta:"index:488"`

	PimpGirls int `gta:"index:532"`

	PimpHighestLevel int `gta:"index:688"`

	SkillFlying int `gta:"index:740"`

	SkillMotorbike int `gta:"index:764"`
	SkillCycling int `gta:"index:768"`
}

// Gets the progress made, as a percentage value
func (b *Block16) GetProgress() float32 {
	return b.ProgressMade / b.MaxProgress * 100
}
