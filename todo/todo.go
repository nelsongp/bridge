package todo

import "github.com/nelsongp/bridge/list"

type Todo interface {
	SetList(l list.List)
	Add(name string)
	Print()
}
