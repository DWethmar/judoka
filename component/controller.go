package component

import "github.com/dwethmar/judoka/entity"

const ControllerType = "controller"

type Controller struct {
	CID    uint32
	entity entity.Entity
	X, Y   int
}

func NewController(id uint32, entity entity.Entity) *Controller {
	return &Controller{
		CID:    id,
		entity: entity,
	}
}

func (c *Controller) ID() uint32            { return c.CID }
func (c *Controller) Type() string          { return ControllerType }
func (c *Controller) Entity() entity.Entity { return c.entity }
