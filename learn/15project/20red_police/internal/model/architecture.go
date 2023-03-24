package model

type architecture struct {
	id          string
	name        string
	armList     []string
	bloodVolume int
}

type barracks struct {
	*architecture
}

type MJBarracks struct {
	*architecture
	*barracks
}

type UKBarracks struct {
	*architecture
	*MJBarracks
}

type SJBarracks struct {
	*architecture
	*barracks
}

type IRAQBarracks struct {
	*architecture
	*SJBarracks
}

var b *barracks = &barracks{
	&architecture{
		id:      "1",
		name:    "兵营",
		armList: []string{"工程师"},
	},
}

var mjb *MJBarracks = &MJBarracks{
	architecture: &architecture{
		id:      "2",
		name:    "盟军兵营",
		armList: []string{"盟军大兵,盟军狗,飞行兵,间谍,超时空兵"},
	},
	barracks: b,
}

/**
创建国家时，选择国家可以创建的建筑，需要一个建筑列表
查询某个建筑可以创建的兵种
*/

func NewMJBarracks() *MJBarracks {
	mb := mjb
	return mb
}
