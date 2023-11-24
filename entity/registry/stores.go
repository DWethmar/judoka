package registry

import (
	"github.com/dwethmar/judoka/component"
	"github.com/dwethmar/judoka/ids"
)

// Stores is a collection of component stores.
var stores = Stores{
	Transform: NewStore[*component.Transform](
		func(s *Store[*component.Transform]) {
			idGen := ids.New(0)
			s.BeforeAdd = func(c *component.Transform) error {
				c.CID = idGen.Next()
				if err := ValidateComponent(c); err != nil {
					return err
				}

				// unique component
				if len(s.store[c.Entity()]) > 0 {
					return ErrUniqueConstraintFailed
				}

				return nil
			}
		},
	),
	Sprite: NewStore[*component.Sprite](
		func(s *Store[*component.Sprite]) {
			idGen := ids.New(0)
			s.BeforeAdd = func(c *component.Sprite) error {
				c.CID = idGen.Next()
				if err := ValidateComponent(c); err != nil {
					return err
				}

				return nil
			}
		},
	),
	Controller: NewStore[*component.Controller](
		func(s *Store[*component.Controller]) {
			idGen := ids.New(0)
			s.BeforeAdd = func(c *component.Controller) error {
				c.CID = idGen.Next()

				if err := ValidateComponent(c); err != nil {
					return err
				}

				// unique component
				if len(s.store[c.Entity()]) > 0 {
					return ErrUniqueConstraintFailed
				}

				return nil
			}
		},
	),
	Velocity: NewStore[*component.Velocity](
		func(s *Store[*component.Velocity]) {
			idGen := ids.New(0)
			s.BeforeAdd = func(c *component.Velocity) error {
				c.CID = idGen.Next()

				if err := ValidateComponent(c); err != nil {
					return err
				}

				// unique component
				if len(s.store[c.Entity()]) > 0 {
					return ErrUniqueConstraintFailed
				}

				return nil
			}
		},
	),
	Actor: NewStore[*component.Actor](
		func(s *Store[*component.Actor]) {
			idGen := ids.New(0)
			s.BeforeAdd = func(c *component.Actor) error {
				c.CID = idGen.Next()

				if err := ValidateComponent(c); err != nil {
					return err
				}

				// unique component
				if len(s.store[c.Entity()]) > 0 {
					return ErrUniqueConstraintFailed
				}

				return nil
			}
		},
	),
	Chunk: NewStore[*component.Chunk](
		func(s *Store[*component.Chunk]) {
			idGen := ids.New(0)
			s.BeforeAdd = func(c *component.Chunk) error {
				c.CID = idGen.Next()

				if err := ValidateComponent(c); err != nil {
					return err
				}

				// unique component
				if len(s.store[c.Entity()]) > 0 {
					return ErrUniqueConstraintFailed
				}

				return nil
			}
		},
	),
	Layer: NewStore[*component.Layer](
		func(s *Store[*component.Layer]) {
			idGen := ids.New(0)
			s.BeforeAdd = func(c *component.Layer) error {
				c.CID = idGen.Next()

				if err := ValidateComponent(c); err != nil {
					return err
				}

				// unique component
				if len(s.store[c.Entity()]) > 0 {
					return ErrUniqueConstraintFailed
				}

				return nil
			}
		},
	),
}
