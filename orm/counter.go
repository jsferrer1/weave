package orm

var _ Model = (*CounterWithID)(nil)

// SetPrimaryKey is a minimal implementation, useful when the ID is a separate protobuf field
func (c *CounterWithID) SetPrimaryKey(pk []byte) error {
	c.PrimaryKey = pk
	return nil
}

// Validate is always succesful
func (c *CounterWithID) Validate() error {
	return nil
}
