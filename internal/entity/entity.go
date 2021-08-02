package entity

import (
	"strings"

	"github.com/zu1k/nali/pkg/dbif"
)

type EntityType uint

const (
	TypeIPv4   = dbif.TypeIPv4
	TypeIPv6   = dbif.TypeIPv6
	TypeDomain = dbif.TypeDomain

	TypePlain = 100
)

type Entity struct {
	Loc  []int // s[Loc[0]:Loc[1]]
	Type EntityType

	Text string
	Info string
}

func (e Entity) ParseInfo() error {
	return nil
}

type Entities []*Entity

func (es Entities) Len() int {
	return len(es)
}

func (es Entities) Less(i, j int) bool {
	return es[i].Loc[0] < es[j].Loc[0]
}

func (es Entities) Swap(i, j int) {
	es[i], es[j] = es[j], es[i]
}

func (es Entities) String() string {
	var result strings.Builder
	for _, entity := range es {
		result.WriteString(entity.Text)
		if entity.Type != TypePlain && len(entity.Info) > 0 {
			result.WriteString("[" + entity.Info + "] ")
		}
	}
	return result.String()
}
