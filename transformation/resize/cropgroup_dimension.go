package resize

func (c *CropGroup) Width(width int) *CropGroup {
	c.width = width

	return c
}

func (c *CropGroup) WidthPercent(width float64) *CropGroup {
	c.width = width

	return c
}

func (c *CropGroup) WidthExpr(width string) *CropGroup {
	c.width = width

	return c
}

func (c *CropGroup) Height(height int) *CropGroup {
	c.height = height

	return c
}

func (c *CropGroup) HeightPercent(height float64) *CropGroup {
	c.height = height

	return c
}

func (c *CropGroup) HeightExpr(height string) *CropGroup {
	c.height = height

	return c
}

func (c *CropGroup) AspectRatio(aspectRatio float64) *CropGroup {
	c.aspectRatio = aspectRatio

	return c
}
