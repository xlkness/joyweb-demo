package model

type PermBit = int

const (
	DefaultPerm_Read   = 1
	DefaultPerm_Write  = 2
	DefaultPerm_Exec   = 4
	DefaultPerm_Delete = 8
)

type Permission struct {
	SystemID string // 系统唯一id
	Perm     int    // 管理员|写|读
	SubPerms map[string]*Permission
}

func DefaultPermission(systemID string) *Permission {
	p := new(Permission)
	p.SystemID = systemID
	p.SubPerms = make(map[string]*Permission)
	return p
}

func (p *Permission) AddSubPermission(systemID string, subPermission *Permission) *Permission {
	p.SubPerms[systemID] = subPermission
	return p
}

// Permit 校验是否存在systemIDs链的权限perm
// 例如校验是否存在：系统a-模块b-条目c的读权限，Permit(PermBit_Read, "system-a", "module-b", "entry-c")
func (p *Permission) Permit(perm PermBit, systemIDs ...string) bool {
	curPerm := p
	for _, v := range systemIDs {
		subPerm, find := curPerm.SubPerms[v]
		if find {
			curPerm = subPerm
			continue
		}
		return false
	}

	return curPerm.Perm >= perm
}
