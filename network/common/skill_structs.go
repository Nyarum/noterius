package common

type Skill struct {
	ID         uint16
	State      uint8
	Lv         uint8
	UseSp      uint16
	UseEndure  uint16
	UseEnergy  uint16
	ResumeTime uint32
	Range      [4][2]byte
}

type SkillState struct {
	ID uint8
	Lv uint8
}
