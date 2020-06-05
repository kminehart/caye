package caye

func (c *Client) Done() {
	id := c.Context.Value(ContextKeyBuildID)
	c.Builds.EndBuild(c.Context, &EndBuildRequest{
		Id: id.(string),
	})
}
