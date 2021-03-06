package eve

// DL provides a convenient way to write Display List commands.
type DL struct {
	Writer
}

// DL wraps W to return Display List writer. See W for more information.
func (d *Driver) DL(addr int) DL {
	return DL{d.W(addr)}
}

// AlphaFunc sets the alpha test function.
func (dl *DL) AlphaFunc(fun, ref byte) {
	dl.restart(4)
	dl.wr32(ALPHA_FUNC | uint32(fun)<<8 | uint32(ref))
}

// Begin begins drawing a graphics primitive.
func (dl *DL) Begin(prim byte) {
	dl.restart(4)
	dl.wr32(BEGIN | uint32(prim))
}

// BitmapHandle selscts the bitmap handle.
func (dl *DL) BitmapHandle(handle byte) {
	dl.restart(4)
	dl.wr32(BITMAP_HANDLE | uint32(handle))
}

// BitmapLayout sets the bitmap memory format and layout for the current handle.
func (dl *DL) BitmapLayout(format byte, linestride, height int) {
	l := uint32(linestride) & 1023
	h := uint32(height) & 511
	dl.restart(4)
	dl.wr32(BITMAP_LAYOUT | uint32(format)<<19 | l<<9 | h)
	if dl.d.mmap != &eve1 {
		l = uint32(linestride) >> 10 & 3
		h = uint32(height) >> 9 & 3
		dl.addrAdd(4)
		dl.wr32(BITMAP_LAYOUT_H | l<<2 | h)
	}
}

// BitmapSize sets the screen drawing of bitmaps for the current handle.
func (dl *DL) BitmapSize(options byte, width, height int) {
	w := uint32(width) & 511
	h := uint32(height) & 511
	dl.restart(4)
	dl.wr32(BITMAP_SIZE | uint32(options)<<18 | w<<9 | h)
	if dl.d.mmap != &eve1 {
		w = uint32(width) >> 9 & 3
		h = uint32(height) >> 9 & 3
		dl.addrAdd(4)
		dl.wr32(BITMAP_SIZE_H | w<<2 | h)
	}
}

// BitmapSource sets the source address of bitmap data in graphics memory RAM_G.
func (dl *DL) BitmapSource(addr int) {
	dl.restart(4)
	dl.wr32(BITMAP_SOURCE | uint32(addr)&0x3FFFFF)
}

// BitmapTransA sets the A coefficient of the bitmap transform matrix.
func (dl *DL) BitmapTransformA(a int) {
	dl.restart(4)
	dl.wr32(BITMAP_TRANSFORM_A | uint32(a)&0x1FFFF)
}

// BitmapTransformB sets the B coefficient of the bitmap transform matrix.
func (dl *DL) BitmapTransformB(b int) {
	dl.restart(4)
	dl.wr32(BITMAP_TRANSFORM_B | uint32(b)&0x1FFFF)
}

// BitmapTransformC sets the C coefficient of the bitmap transform matrix.
func (dl *DL) BitmapTransformC(c int) {
	dl.restart(4)
	dl.wr32(BITMAP_TRANSFORM_C | uint32(c)&0x1FFFF)
}

// BitmapTransformD sets the D coefficient of the bitmap transform matrix.
func (dl *DL) BitmapTransformD(d int) {
	dl.restart(4)
	dl.wr32(BITMAP_TRANSFORM_D | uint32(d)&0x1FFFF)
}

// BitmapTransformE sets the E coefficient of the bitmap transform matrix.
func (dl *DL) BitmapTransformE(e int) {
	dl.restart(4)
	dl.wr32(BITMAP_TRANSFORM_E | uint32(e)&0x1FFFF)
}

// BitmapTransformF sets the F coefficient of the bitmap transform matrix.
func (dl *DL) BitmapTransformF(f int) {
	dl.restart(4)
	dl.wr32(BITMAP_TRANSFORM_F | uint32(f)&0x1FFFF)
}

// BlendFunc configures pixel arithmetic.
func (dl *DL) BlendFunc(src, dst byte) {
	dl.restart(4)
	dl.wr32(BLEND_FUNC | uint32(src)<<3 | uint32(dst))
}

// Call executes a sequence of commands at another location in the display list.
func (dl *DL) Call(dest int) {
	dl.restart(4)
	dl.wr32(CALL | uint32(dest)&0xFFFF)
}

// Cell sets the bitmap cell number for the Vertex2F command.
func (dl *DL) Cell(cell byte) {
	dl.restart(4)
	dl.wr32(CELL | uint32(cell))
}

// Clear clears buffers to preset values.
func (dl *DL) Clear(cst byte) {
	dl.restart(4)
	dl.wr32(CLEAR | uint32(cst))
}

// ClearColorA sets the clear value for the alpha channel.
func (dl *DL) ClearColorA(alpha byte) {
	dl.restart(4)
	dl.wr32(CLEAR_COLOR_A | uint32(alpha))
}

// ClearColorRGB sets the clear values for red, green and blue channels.
func (dl *DL) ClearColorRGB(r, g, b byte) {
	dl.restart(4)
	dl.wr32(CLEAR_COLOR_RGB | uint32(r)<<16 | uint32(g)<<8 | uint32(b))
}

// ClearStencil sets the clear value for the stencil buffer.
func (dl *DL) ClearStencil(s byte) {
	dl.restart(4)
	dl.wr32(CLEAR_STENCIL | uint32(s))
}

// ClearTag sets the clear value for the stencil buffer.
func (dl *DL) ClearTag(t int) {
	dl.restart(4)
	dl.wr32(CLEAR_TAG | uint32(uint16(t)))
}

// ColorA sets the current color alpha.
func (dl *DL) ColorA(alpha byte) {
	dl.restart(4)
	dl.wr32(COLOR_A | uint32(alpha))
}

// ColorMask enables or disables writing of color components.
func (dl *DL) ColorMask(rgba byte) {
	dl.restart(4)
	dl.wr32(COLOR_MASK | uint32(rgba))
}

// ColorRGB sets the current color red, green and blue.
func (dl *DL) ColorRGB(r, g, b byte) {
	dl.restart(4)
	dl.wr32(COLOR_RGB | uint32(r)<<16 | uint32(g)<<8 | uint32(b))
}

// Display ends the display list (following command will be ignored).
func (dl *DL) Display() {
	dl.restart(4)
	dl.wr32(DISPLAY)
}

// End ends drawing a graphics primitive.
func (dl *DL) End() {
	dl.restart(4)
	dl.wr32(END)
}

// Jump executes commands at another location in the display list. Dest is the
// command number in display list (address = RAM_DL + dest*4).
func (dl *DL) Jump(dest int) {
	dl.restart(4)
	dl.wr32(JUMP | uint32(dest)&0xFFFF)
}

// LineWidth sets the width of lines to be drawn with primitive LINES in 1/16
// pixel precision.
func (dl *DL) LineWidth(width int) {
	dl.restart(4)
	dl.wr32(LINE_WIDTH | uint32(width)&0xFFF)
}

// Macro executes a single command from a macro register.
func (dl *DL) Macro(m int) {
	dl.restart(4)
	dl.wr32(MACRO | uint32(m&1))
}

// Nop does nothing.
func (dl *DL) Nop() {
	dl.restart(4)
	dl.wr32(NOP)
}

// PaletteSource sets the base address of the palette (EVE2).
func (dl *DL) PaletteSource(addr int) {
	dl.restart(4)
	dl.wr32(PALETTE_SOURCE | uint32(addr)&0x3FFFFF)
}

// PointSize sets the radius of points.
func (dl *DL) PointSize(size int) {
	dl.restart(4)
	dl.wr32(POINT_SIZE | uint32(size)&0x1FFF)
}

// RestoreContext restores the current graphics context from the context stack.
func (dl *DL) RestoreContext() {
	dl.restart(4)
	dl.wr32(RESTORE_CONTEXT)
}

// Return returns from a previous CALL command.
func (dl *DL) Return() {
	dl.restart(4)
	dl.wr32(RETURN)
}

// SaveContext pushes the current graphics context on the context stack.
func (dl *DL) SaveContext() {
	dl.restart(4)
	dl.wr32(SAVE_CONTEXT)
}

// ScissorSize sets the size of the scissor clip rectangle.
func (dl *DL) ScissorSize(width, height int) {
	dl.restart(4)
	dl.wr32(SCISSOR_SIZE | uint32(width)&0xFFF<<12 | uint32(height)&0xFFF)
}

// ScissorXY sets the size of the scissor clip rectangle.
func (dl *DL) ScissorXY(x, y int) {
	dl.restart(4)
	dl.wr32(SCISSOR_XY | uint32(x)&0x7FF<<11 | uint32(y)&0x7FF)
}

// StencilFunc sets function and reference value for stencil testing.
func (dl *DL) StencilFunc(fun, ref, mask byte) {
	dl.restart(4)
	dl.wr32(STENCIL_FUNC | uint32(fun)<<16 | uint32(ref)<<8 | uint32(mask))
}

// StencilMask controls the writing of individual bits in the stencil planes.
func (dl *DL) StencilMask(mask byte) {
	dl.restart(4)
	dl.wr32(STENCIL_MASK | uint32(mask))
}

// StencilOp sets stencil test actions.
func (dl *DL) StencilOp(sfail, spass byte) {
	dl.restart(4)
	dl.wr32(STENCIL_OP | uint32(sfail)<<3 | uint32(spass))
}

// Tag attaches the tag value for the following graphics objects drawn on the
// screen. The initial tag buffer value is 255.
func (dl *DL) Tag(t int) {
	dl.restart(4)
	dl.wr32(TAG | uint32(uint16(t)))
}

// TagMask controls the writing of the tag buffer.
func (dl *DL) TagMask(mask byte) {
	dl.restart(4)
	dl.wr32(TAG_MASK | uint32(mask))
}

// Vertex2f starts the operation of graphics primitives at the specified screen
// coordinate, in the pixel precision set by VertexFormat (default: 1/16 pixel).
func (dl *DL) Vertex2f(x, y int) {
	dl.restart(4)
	dl.wr32(VERTEX2F | uint32(x)&0x7FFF<<15 | uint32(y)&0x7FFF)
}

// Vertex2II starts the operation of graphics primitive at the specified
// coordinates in pixel precision.
func (dl *DL) Vertex2ii(x, y int, handle, cell byte) {
	dl.restart(4)
	dl.wr32(VERTEX2II | uint32(x)&511<<21 | uint32(y)&511<<12 |
		uint32(handle)<<7 | uint32(cell))
}

// VertexFormat sets the precision of Vertex2F coordinates (EVE2).
func (dl *DL) VertexFormat(frac uint) {
	dl.restart(4)
	dl.wr32(VERTEX_FORMAT | uint32(frac)&7)
}
