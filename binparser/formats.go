package binparser

type RawDataInfo struct {
	Exist       bool
	Index       int32
	DataName    [72]int8
	LastUseTick uint32
	Enable      bool
	Data        uint32
	PackOffset  uint32
	DataSize    uint32
	ID          int32
	LoadCnt     uint32
}

type AreaInfo struct {
	Base       RawDataInfo
	Color      uint32
	Music      int32
	EnvColor   uint32
	LightColor uint32
	LightDir   [3]float32
	Type       int8
}

const CharacterNameLength = 32
const CharacterIconNameLength = 17
const CharacterSkinCount = 8
const CharacterItemKindCount = 20
const CharacterBirthEffectCount = 3
const CharacterDeathEffectCount = 3
const CharacterHPEffectCount = 3
const CharacterInitSkillCount = 11
const CharacterInitItemCount = 10
const CharacterGuildNameLength = 33
const CharacterTitleNameLength = 33
const CharacterJobNameLength = 17

type CharacterInfo struct {
	Base           RawDataInfo
	ID             int32
	Name           [CharacterNameLength]int8
	IconName       [CharacterIconNameLength]int8
	ModalType      int8
	CtrlType       int8
	Model          int16
	SuitID         int16
	SuitNum        int16
	SkinInfo       [CharacterSkinCount]int16
	FeffID         [4]int16
	EeffID         int16
	EffectActionID [3]int16
	Shadow         int16
	ActionID       int16
	Diaphaneity    int8
	Footfall       int16
	Whoop          int16
	Dirge          int16
	ControlAble    int8
	Territory      int8
	SeaHeight      int16
	ItemType       [CharacterItemKindCount]int16
	Lengh          float32
	Width          float32
	Height         float32
	Radii          int16
	BirthBehave    [CharacterBirthEffectCount]int8
	DiedBehave     [CharacterDeathEffectCount]int8
	BornEff        int16
	DieEff         int16
	Dormancy       int16
	DieAction      int8
	HPEffect       [CharacterHPEffectCount]int32
	IsFace         bool
	IsCyclone      bool
	Script         int32
	Weapon         int32
	Skill          [CharacterInitSkillCount][2]int32
	Item           [CharacterInitItemCount][2]int32
	TaskItem       [CharacterInitItemCount][2]int32
	MaxShowItem    int32
	AllShow        float32
	Prefix         int32
	AiNo           int32
	CanTurn        int8
	Vision         int32
	Noise          int32
	GetEXP         int32
	Light          bool
	Mobexp         int32
	Lv             int32
	MxHp           int32
	Hp             int32
	MxSp           int32
	Sp             int32
	MnAtk          int32
	MxAtk          int32
	PDef           int32
	Def            int32
	Hit            int32
	Flee           int32
	Crt            int32
	Mf             int32
	HRec           int32
	SRec           int32
	ASpd           int32
	ADis           int32
	CDis           int32
	MSpd           int32
	Col            int32
	Str            int32
	Agi            int32
	Dex            int32
	Con            int32
	Sta            int32
	Luk            int32
	LHandVal       int32
	Guild          [CharacterGuildNameLength]int8
	Title          [CharacterTitleNameLength]int8
	Job            [CharacterJobNameLength]int8
	CExp           int32
	NExp           int32
	Fame           int32
	Ap             int32
	Tp             int32
	Gd             int32
	Spri           int32
	MxSail         int32
	Sail           int32
	Stasa          int32
	Scsm           int32
	TStr           int32
	TAgi           int32
	TDex           int32
	TCon           int32
	TSta           int32
	TLuk           int32
	TMxHp          int32
	TMxSp          int32
	TAtk           int32
	TDef           int32
	THit           int32
	TFlee          int32
	TMf            int32
	TCrt           int32
	THRec          int32
	TSRec          int32
	TASpd          int32
	TADis          int32
	TSpd           int32
	TSpri          int32
	TScsm          int32
	Scaling        [3]float32
	HaveEffectFog  bool
}

type CharacterPoseInfo struct {
	Base       RawDataInfo
	RealPoseID [7]int16
}

type ChatIconInfo struct {
	Base      RawDataInfo
	Small     [16]int8
	SmallX    int32
	SmallY    int32
	SmallOff  [16]int8
	SmallOffX int32
	SmallOffY int32
	Big       [16]int8
	BigX      int32
	BigY      int32
	Hint      [32]int8
}

type ElfSkillInfo struct {
	Base   RawDataInfo
	Index  int32
	TypeID int32
}

type EventSoundInfo struct {
	Base    RawDataInfo
	SoundID int32
}

type forgeItem struct {
	Item  uint16
	Num   uint8
	Param uint8
}

const ForgeMaxNumberItems = 6

type ForgeInfo struct {
	Base      RawDataInfo
	Level     uint8
	Failure   uint8
	Rate      uint8
	Param     uint8
	Money     uint32
	ForgeItem [ForgeMaxNumberItems]forgeItem
}

const HairMaxItem = 4
const HairMaxFailItem = 3

type HairInfo struct {
	Base       RawDataInfo
	Color      [10]int8
	NeedItem   [HairMaxItem][2]uint32
	Money      uint32
	FailItemID [HairMaxFailItem]uint32
	IsChaUse   [4]bool
	FailNum    int32
}

const ItemNameLength = 80
const ItemIconNameLength = 17
const ItemModuleCount = 5
const ItemModuleLength = 19
const ItemBodyCount = 4
const MaxJobType = 19
const MaxEquipment = 10
const ItemAttributeEffectNameLength = 33
const ItemBindEffectCount = 8
const ItemDescriptionLength = 257

type ItemInfo struct {
	Base          RawDataInfo
	ID            int32
	Name          [ItemNameLength]int8
	ICON          [ItemIconNameLength]int8
	Module        [ItemModuleCount][ItemModuleLength]int8
	ShipFlag      int16
	ShipType      int16
	Type          int16
	ForgeLv       int8
	ForgeSteady   int8
	ExclusiveID   int8
	IsTrade       int8
	IsPick        int8
	IsThrow       int8
	IsDel         int8
	Price         int32
	Body          [ItemBodyCount]int8
	NeedLv        int16
	Work          [MaxJobType]int8
	PileMax       int32
	Instance      int8
	AbleLink      [MaxEquipment]int8
	NeedLink      [MaxEquipment]int8
	PickTo        int8
	StrCoef       int16
	AgiCoef       int16
	DexCoef       int16
	ConCoef       int16
	StaCoef       int16
	LukCoef       int16
	ASpdCoef      int16
	ADisCoef      int16
	MnAtkCoef     int16
	MxAtkCoef     int16
	DefCoef       int16
	MxHpCoef      int16
	MxSpCoef      int16
	FleeCoef      int16
	HitCoef       int16
	CrtCoef       int16
	MfCoef        int16
	HRecCoef      int16
	SRecCoef      int16
	MSpdCoef      int16
	ColCoef       int16
	StrValu       [2]int16
	AgiValu       [2]int16
	DexValu       [2]int16
	ConValu       [2]int16
	StaValu       [2]int16
	LukValu       [2]int16
	ASpdValu      [2]int16
	ADisValu      [2]int16
	MnAtkValu     [2]int16
	MxAtkValu     [2]int16
	DefValu       [2]int16
	MxHpValu      [2]int16
	MxSpValu      [2]int16
	FleeValu      [2]int16
	HitValu       [2]int16
	CrtValu       [2]int16
	MfValu        [2]int16
	HRecValu      [2]int16
	SRecValu      [2]int16
	MSpdValu      [2]int16
	ColValu       [2]int16
	PDef          [2]int16
	LHandValu     int16
	Endure        [2]int16
	Energy        [2]int16
	Hole          int16
	AttrEffect    [ItemAttributeEffectNameLength]int8
	Drap          int16
	Effect        [ItemBindEffectCount][2]int16
	ItemEffect    [2]int16
	AreaEffect    [2]int16
	UseItemEffect [2]int16
	Descriptor    [ItemDescriptionLength]int8
	EffNum        int16
	IsBody        [5]bool
}

type ItemPrefixInfo struct {
	Base RawDataInfo
}

const RefineEffectCharacterCount = 4
const RefineEffectCount = 4

type ItemRefineEffectInfo struct {
	Base      RawDataInfo
	LightID   int32
	EffectID  [RefineEffectCharacterCount][RefineEffectCount]int16
	Dummy     [RefineEffectCount]int8
	EffectNum [RefineEffectCharacterCount]int32
}

const ItemRefineCount = 14

type ItemRefineInfo struct {
	Base           RawDataInfo
	Value          [ItemRefineCount]int16
	ChaEffectScale [4]float32
}

type ItemTypeInfo struct {
	Base RawDataInfo
}

type MagicGroupInfo struct {
	Base      RawDataInfo
	Name      [32]int8
	TypeNum   int32
	TypeID    [8]int32
	Num       [8]int32
	TotalNum  int32
	RenderIdx int32
}

type MagicSingleInfo struct {
	Base      RawDataInfo
	Name      [32]int8
	ModelNum  int32
	Model     [8][24]int8
	Vel       int32
	ParNum    int32
	Part      [8][24]int8
	Dummy     [8]int32
	RenderIdx int32
	LightID   int32
	Result    [24]int8
}

type MapInfo struct {
	Base         RawDataInfo
	Name         [16]int8
	InitX        int32
	InitY        int32
	LightDir     [3]float32
	LightColor   [3]uint8
	IsShowSwitch bool
}

type MusicInfo struct {
	Base RawDataInfo
	Type int32
}

type NotifyInfo struct {
	Base RawDataInfo
	Type int8
	Info [64]int8
}

const ObjectEventNameLength = 18

type ObjectEventInfo struct {
	Base         RawDataInfo
	ID           int32
	Name         [ObjectEventNameLength]int8
	EventType    int16
	ArouseType   int16
	ArouseRadius int16
	Effect       int16
	Music        int16
	BornEffect   int16
	Cursor       int16
}

type ResourceInfo struct {
	Base RawDataInfo
	Type int32
}

type SceneEffectInfo struct {
	Base       RawDataInfo
	Name       [16]int8
	PhotoName  [16]int8
	PhotoTexID int32
	EffType    int32
	ObjType    int32
	DummyNum   int32
	Dummy      [8]int32
	Dummy2     int32
	HeightOff  int32
	PlayTime   float32
	LightID    int32
	BaseSize   float32
}

type SceneObjectInfo struct {
	Base               RawDataInfo
	Name               [16]int8
	Type               int32
	Point32Color       [3]uint8
	EnvColor           [3]uint8
	FogColor           [3]uint8
	Range              int32
	Attenuation1       float32
	AnimCtrlID         int32
	Style              int32
	AttachEffectID     int32
	EnablePoint32Light bool
	EnableEnvLight     bool
	Flag               int32
	SizeFlag           int32
	EnvSound           [11]int8
	EnvSoundDis        int32
	PhotoTexID         int32
	ShadeFlag          bool
	IsReallyBig        bool
	FadeObjNum         int32
	FadeObjSeq         [16]int32
	FadeCoefficent     float32
}

type SelectCharacterInfo struct {
	Base        RawDataInfo
	Type        uint32
	Bone        uint32
	Hair        [64]uint32
	Face        [64]uint32
	Body        [64]uint32
	Hand        [64]uint32
	Foot        [64]uint32
	HairNum     uint32
	FaceNum     uint32
	BodyNum     uint32
	HandNum     uint32
	FootNum     uint32
	Profession  uint32
	Description [1024]int8
}

const ServerInfoMaxGate = 5

type ServerInfo struct {
	Base         RawDataInfo
	GateIP       [ServerInfoMaxGate][16]int8
	Region       [16]int8
	ValidGateCnt int8
}

type ShadeInfo struct {
	Base         RawDataInfo
	Name         [16]int8
	PhotoTexID   int32
	Size         float32
	Ani          int32
	Row          int32
	Col          int32
	UseAlphaTest int32
	AlphaType    int32
	ColorR       int32
	ColorG       int32
	ColorB       int32
	ColorA       int32
	Type         int32
}

const ShipNameLength = 64
const ShipDescriptionLength = 128
const ShipMaxParts = 16

type ShipInfo struct {
	Base         RawDataInfo
	Name         [ShipNameLength]int8
	Desp         [ShipDescriptionLength]int8
	ItemID       uint16
	CharID       uint16
	PosID        uint16
	IsUpdate     uint8
	NumHeader    uint16
	NumEngine    uint16
	NumCannon    uint16
	NumEquipment uint16
	Header       [ShipMaxParts]uint16
	Engine       [ShipMaxParts]uint16
	Cannon       [ShipMaxParts]uint16
	Equipment    [ShipMaxParts]uint16
	Body         uint16
	LvLimit      uint16
	NumPfLimit   uint16
	PfLimit      [ShipMaxParts]uint16
	Endure       uint16
	Resume       uint16
	Defence      uint16
	Resist       uint16
	MinAttack    uint16
	MaxAttack    uint16
	Distance     uint16
	Time         uint16
	Scope        uint16
	Capacity     uint16
	Supply       uint16
	Consume      uint16
	CannonSpeed  uint16
	Speed        uint16
	Param        uint16
}

const ShipMaxMotor = 4

type ShipItemInfo struct {
	Base        RawDataInfo
	Name        [ShipNameLength]int8
	Desp        [ShipDescriptionLength]int8
	Model       uint32
	Motor       [ShipMaxMotor]uint16
	Price       uint32
	Endure      uint16
	Resume      uint16
	Defence     uint16
	Resist      uint16
	MinAttack   uint16
	MaxAttack   uint16
	Distance    uint16
	Time        uint16
	Scope       uint16
	Capacity    uint16
	Supply      uint16
	Consume     uint16
	CannonSpeed uint16
	Speed       uint16
	Param       uint16
}

const SkillEffectNameLength = 17
const SkillEffectFunctionNameLength = 32
const SkillEffectActionCount = 3

type SkillEffectInfo struct {
	Base          RawDataInfo
	ID            int8
	Name          [SkillEffectNameLength]int8
	Frequency     int16
	OnTransfer    [SkillEffectFunctionNameLength]int8
	AddState      [SkillEffectFunctionNameLength]int8
	SubState      [SkillEffectFunctionNameLength]int8
	AddType       int8
	CanCancel     bool
	CanMove       bool
	CanMSkill     bool
	CanGSkill     bool
	CanTrade      bool
	CanItem       bool
	CanUnbeatable bool
	CanItemmed    bool
	CanSkilled    bool
	NoHide        bool
	NoShow        bool
	OptItem       bool
	TalkToNPC     bool
	FreeStateID   int8
	Screen        int8
	ActBehave     [SkillEffectActionCount]int8
	ChargeLink    int16
	AreaEffect    int16
	IsShowCenter  bool
	IsDizzy       bool
	Effect        int16
	Dummy1        int16
	BitEffect     int16
	Dummy2        int16
	Icon          int16
	ActNum        int32
}

const SkillNameLength = 17
const SkillJobSelectCount = 9
const SkillItemNeedCount = 8
const SkillPreSkillCount = 3
const SkillRangeFunctionNameLength = 33
const SkillEffectFunctionNameLength = 33
const SkillSelfAttributeCount = 2
const SkillSelfEffectCount = 2
const SkillTargetEffectCount = 2
const SkillExpendItemCount = 2
const SkillOperateCount = 3
const SkillPoseCount = 10
const SkillActionEffectCount = 3
const SkillItemEffectCount = 2
const SkillIconNameLength = 17
const SkillRangeParameterCount = 4

type SelectCha int

const (
	SCNone SelectCha = iota
	SCAll
	SCPlayer
	SCEnemy
	SCPlayerAshes
	SCMonster
	SCMonsterRepairable
	SCMonsterTree
	SCMonsterMine
	SCMonsterFish
	SCMonsterBoat
	SCSelf
	SCTeam
)

type skillGrid struct {
	State      int8
	Level      int8
	ID         int16
	UseSP      int16
	UseEndure  int16
	UseEnergy  int16
	ResumeTime int32
	Range      [SkillRangeParameterCount]int16
}

type SkillInfo struct {
	Base                     RawDataInfo
	ID                       int16
	Name                     [SkillNameLength]int8
	FightType                int8
	JobSelect                [SkillJobSelectCount][2]int8
	ConchNeed                [SkillItemNeedCount][3]int16
	Phase                    int8
	Type                     int8
	LevelDemand              int16
	PremissSkill             [SkillPreSkillCount][2]int16
	Point32Expend            int8
	SrcType                  int8
	TarType                  int8
	ApplyDistance            int16
	ApplyTarget              int8
	ApplyType                int8
	Helpful                  int8
	Angle                    int16
	Radii                    int16
	Range                    int8
	Prepare                  [SkillRangeFunctionNameLength]int8
	UseSP                    [SkillEffectFunctionNameLength]int8
	UseEndure                [SkillEffectFunctionNameLength]int8
	UseEnergy                [SkillEffectFunctionNameLength]int8
	SetRange                 [SkillEffectFunctionNameLength]int8
	RangeState               [SkillEffectFunctionNameLength]int8
	Use                      [SkillEffectFunctionNameLength]int8
	Effect                   [SkillEffectFunctionNameLength]int8
	Active                   [SkillEffectFunctionNameLength]int8
	Inactive                 [SkillEffectFunctionNameLength]int8
	StateID                  int32
	SelfAttr                 [SkillSelfAttributeCount]int16
	SelfEffect               [SkillSelfEffectCount]int16
	ItemExpend               [SkillExpendItemCount][2]int16
	BeingTime                int16
	TargetAttr               [SkillTargetEffectCount]int16
	SplashPara               int16
	TargetEffect             int16
	SplashEffect             int16
	Variation                int16
	Summon                   int16
	PreTime                  int16
	FireSpeed                [SkillEffectFunctionNameLength]int8
	Operate                  [SkillOperateCount]int8
	ActionHarm               int16
	ActionPlayType           int8
	ActionPose               [SkillPoseCount]int16
	ActionKeyFrme            int16
	Whop                     int16
	ActionDummyLink          [SkillActionEffectCount]int16
	ActionEffect             [SkillActionEffectCount]int16
	ActionEffectType         [SkillActionEffectCount]int16
	ItemDummyLink            int16
	ItemEffect1              [SkillItemEffectCount]int16
	ItemEffect2              [SkillItemEffectCount]int16
	SkyEffectActionKeyFrame  int16
	SkyEffectActionDummyLink int16
	SkyEffectItemDummyLink   int16
	SkyEffect                int16
	SkySpd                   int16
	Whoped                   int16
	TargetDummyLink          int16
	TargetEffectID           int16
	AgroundEffectID          int16
	WaterEffectID            int16
	ICON                     [SkillIconNameLength]int8
	PlayTime                 int8
	DescribeHint32           [128]int8
	EffectHint32             [128]int8
	ExpendHint32             [128]int8
	Skill                    skillGrid
	Upgrade                  int32
	IsActive                 bool
	dwAttackTime             uint32
	SelectCha                SelectCha
	PoseNum                  int32
}

const StoneEquipMax = 3

type StoneInfo struct {
	Base         RawDataInfo
	ItemID       int32
	EquipPos     [StoneEquipMax]int32
	Type         int32
	HintFunction [64]int8
}

type StringInfo map[int]string

type TerrainInfo struct {
	Base      RawDataInfo
	Type      int8
	TextureID int32
	Attr      int8
}
