package interfaces

import "github.com/wliao008/mazing/structs"

type Mazer interface {
	Generate(x, y uint16) ([]structs.Cell, error)
}
