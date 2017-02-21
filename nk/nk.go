
package nk

/*
#include "nuklear.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import "unsafe"

// NkInitDefault function as declared in nk/nuklear.h:696
func NkInitDefault(arg0 *Context, arg1 *UserFont) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := arg1.PassRef()
	__ret := C.nk_init_default(carg0, carg1)
	__v := (int32)(__ret)
	return __v
}

// NkInitFixed function as declared in nk/nuklear.h:698
func NkInitFixed(arg0 *Context, memory unsafe.Pointer, size Size, arg3 *UserFont) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cmemory, _ := (unsafe.Pointer)(unsafe.Pointer(memory)), cgoAllocsUnknown
	csize, _ := (C.nk_size)(size), cgoAllocsUnknown
	carg3, _ := arg3.PassRef()
	__ret := C.nk_init_fixed(carg0, cmemory, csize, carg3)
	__v := (int32)(__ret)
	return __v
}

// NkInit function as declared in nk/nuklear.h:699
func NkInit(arg0 *Context, arg1 *Allocator, arg2 *UserFont) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (*C.struct_nk_allocator)(unsafe.Pointer(arg1)), cgoAllocsUnknown
	carg2, _ := arg2.PassRef()
	__ret := C.nk_init(carg0, carg1, carg2)
	__v := (int32)(__ret)
	return __v
}

// NkInitCustom function as declared in nk/nuklear.h:700
func NkInitCustom(arg0 *Context, cmds *Buffer, pool *Buffer, arg3 *UserFont) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	ccmds, _ := (*C.struct_nk_buffer)(unsafe.Pointer(cmds)), cgoAllocsUnknown
	cpool, _ := (*C.struct_nk_buffer)(unsafe.Pointer(pool)), cgoAllocsUnknown
	carg3, _ := arg3.PassRef()
	__ret := C.nk_init_custom(carg0, ccmds, cpool, carg3)
	__v := (int32)(__ret)
	return __v
}

// NkClear function as declared in nk/nuklear.h:701
func NkClear(arg0 *Context) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	C.nk_clear(carg0)
}

// NkFree function as declared in nk/nuklear.h:702
func NkFree(arg0 *Context) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	C.nk_free(carg0)
}

// NkBegin function as declared in nk/nuklear.h:708
func NkBegin(arg0 *Context, title string, bounds Rect, flags Flags) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	ctitle, _ := unpackPCharString(title)
	cbounds, _ := *(*C.struct_nk_rect)(unsafe.Pointer(&bounds)), cgoAllocsUnknown
	cflags, _ := (C.nk_flags)(flags), cgoAllocsUnknown
	__ret := C.nk_begin(carg0, ctitle, cbounds, cflags)
	__v := (int32)(__ret)
	return __v
}

// NkBeginTitled function as declared in nk/nuklear.h:709
func NkBeginTitled(arg0 *Context, name string, title string, bounds Rect, flags Flags) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cname, _ := unpackPCharString(name)
	ctitle, _ := unpackPCharString(title)
	cbounds, _ := *(*C.struct_nk_rect)(unsafe.Pointer(&bounds)), cgoAllocsUnknown
	cflags, _ := (C.nk_flags)(flags), cgoAllocsUnknown
	__ret := C.nk_begin_titled(carg0, cname, ctitle, cbounds, cflags)
	__v := (int32)(__ret)
	return __v
}

// NkEnd function as declared in nk/nuklear.h:710
func NkEnd(arg0 *Context) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	C.nk_end(carg0)
}

// NkWindowFind function as declared in nk/nuklear.h:712
func NkWindowFind(ctx *Context, name string) *Window {
	cctx, _ := (*C.struct_nk_context)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	cname, _ := unpackPCharString(name)
	__ret := C.nk_window_find(cctx, cname)
	__v := *(**Window)(unsafe.Pointer(&__ret))
	return __v
}

// NkWindowGetBounds function as declared in nk/nuklear.h:713
func NkWindowGetBounds(arg0 *Context) Rect {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	__ret := C.nk_window_get_bounds(carg0)
	__v := *(*Rect)(unsafe.Pointer(&__ret))
	return __v
}

// NkWindowGetPosition function as declared in nk/nuklear.h:714
func NkWindowGetPosition(arg0 *Context) Vec2 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	__ret := C.nk_window_get_position(carg0)
	__v := *(*Vec2)(unsafe.Pointer(&__ret))
	return __v
}

// NkWindowGetSize function as declared in nk/nuklear.h:715
func NkWindowGetSize(arg0 *Context) Vec2 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	__ret := C.nk_window_get_size(carg0)
	__v := *(*Vec2)(unsafe.Pointer(&__ret))
	return __v
}

// NkWindowGetWidth function as declared in nk/nuklear.h:716
func NkWindowGetWidth(arg0 *Context) float32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	__ret := C.nk_window_get_width(carg0)
	__v := (float32)(__ret)
	return __v
}

// NkWindowGetHeight function as declared in nk/nuklear.h:717
func NkWindowGetHeight(arg0 *Context) float32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	__ret := C.nk_window_get_height(carg0)
	__v := (float32)(__ret)
	return __v
}

// NkWindowGetPanel function as declared in nk/nuklear.h:718
func NkWindowGetPanel(arg0 *Context) *Panel {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	__ret := C.nk_window_get_panel(carg0)
	__v := *(**Panel)(unsafe.Pointer(&__ret))
	return __v
}

// NkWindowGetContentRegion function as declared in nk/nuklear.h:719
func NkWindowGetContentRegion(arg0 *Context) Rect {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	__ret := C.nk_window_get_content_region(carg0)
	__v := *(*Rect)(unsafe.Pointer(&__ret))
	return __v
}

// NkWindowGetContentRegionMin function as declared in nk/nuklear.h:720
func NkWindowGetContentRegionMin(arg0 *Context) Vec2 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	__ret := C.nk_window_get_content_region_min(carg0)
	__v := *(*Vec2)(unsafe.Pointer(&__ret))
	return __v
}

// NkWindowGetContentRegionMax function as declared in nk/nuklear.h:721
func NkWindowGetContentRegionMax(arg0 *Context) Vec2 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	__ret := C.nk_window_get_content_region_max(carg0)
	__v := *(*Vec2)(unsafe.Pointer(&__ret))
	return __v
}

// NkWindowGetContentRegionSize function as declared in nk/nuklear.h:722
func NkWindowGetContentRegionSize(arg0 *Context) Vec2 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	__ret := C.nk_window_get_content_region_size(carg0)
	__v := *(*Vec2)(unsafe.Pointer(&__ret))
	return __v
}

// NkWindowGetCanvas function as declared in nk/nuklear.h:723
func NkWindowGetCanvas(arg0 *Context) *CommandBuffer {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	__ret := C.nk_window_get_canvas(carg0)
	__v := *(**CommandBuffer)(unsafe.Pointer(&__ret))
	return __v
}

// NkWindowHasFocus function as declared in nk/nuklear.h:725
func NkWindowHasFocus(arg0 *Context) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	__ret := C.nk_window_has_focus(carg0)
	__v := (int32)(__ret)
	return __v
}

// NkWindowIsCollapsed function as declared in nk/nuklear.h:726
func NkWindowIsCollapsed(arg0 *Context, arg1 string) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	__ret := C.nk_window_is_collapsed(carg0, carg1)
	__v := (int32)(__ret)
	return __v
}

// NkWindowIsClosed function as declared in nk/nuklear.h:727
func NkWindowIsClosed(arg0 *Context, arg1 string) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	__ret := C.nk_window_is_closed(carg0, carg1)
	__v := (int32)(__ret)
	return __v
}

// NkWindowIsHidden function as declared in nk/nuklear.h:728
func NkWindowIsHidden(arg0 *Context, arg1 string) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	__ret := C.nk_window_is_hidden(carg0, carg1)
	__v := (int32)(__ret)
	return __v
}

// NkWindowIsActive function as declared in nk/nuklear.h:729
func NkWindowIsActive(arg0 *Context, arg1 string) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	__ret := C.nk_window_is_active(carg0, carg1)
	__v := (int32)(__ret)
	return __v
}

// NkWindowIsHovered function as declared in nk/nuklear.h:730
func NkWindowIsHovered(arg0 *Context) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	__ret := C.nk_window_is_hovered(carg0)
	__v := (int32)(__ret)
	return __v
}

// NkWindowIsAnyHovered function as declared in nk/nuklear.h:731
func NkWindowIsAnyHovered(arg0 *Context) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	__ret := C.nk_window_is_any_hovered(carg0)
	__v := (int32)(__ret)
	return __v
}

// NkItemIsAnyActive function as declared in nk/nuklear.h:732
func NkItemIsAnyActive(arg0 *Context) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	__ret := C.nk_item_is_any_active(carg0)
	__v := (int32)(__ret)
	return __v
}

// NkWindowSetBounds function as declared in nk/nuklear.h:734
func NkWindowSetBounds(arg0 *Context, arg1 Rect) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := *(*C.struct_nk_rect)(unsafe.Pointer(&arg1)), cgoAllocsUnknown
	C.nk_window_set_bounds(carg0, carg1)
}

// NkWindowSetPosition function as declared in nk/nuklear.h:735
func NkWindowSetPosition(arg0 *Context, arg1 Vec2) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&arg1)), cgoAllocsUnknown
	C.nk_window_set_position(carg0, carg1)
}

// NkWindowSetSize function as declared in nk/nuklear.h:736
func NkWindowSetSize(arg0 *Context, arg1 Vec2) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&arg1)), cgoAllocsUnknown
	C.nk_window_set_size(carg0, carg1)
}

// NkWindowSetFocus function as declared in nk/nuklear.h:737
func NkWindowSetFocus(arg0 *Context, name string) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cname, _ := unpackPCharString(name)
	C.nk_window_set_focus(carg0, cname)
}

// NkWindowClose function as declared in nk/nuklear.h:739
func NkWindowClose(ctx *Context, name string) {
	cctx, _ := (*C.struct_nk_context)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	cname, _ := unpackPCharString(name)
	C.nk_window_close(cctx, cname)
}

// NkWindowCollapse function as declared in nk/nuklear.h:740
func NkWindowCollapse(arg0 *Context, name string, arg2 CollapseStates) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cname, _ := unpackPCharString(name)
	carg2, _ := (C.enum_nk_collapse_states)(arg2), cgoAllocsUnknown
	C.nk_window_collapse(carg0, cname, carg2)
}

// NkWindowCollapseIf function as declared in nk/nuklear.h:741
func NkWindowCollapseIf(arg0 *Context, name string, arg2 CollapseStates, cond int32) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cname, _ := unpackPCharString(name)
	carg2, _ := (C.enum_nk_collapse_states)(arg2), cgoAllocsUnknown
	ccond, _ := (C.int)(cond), cgoAllocsUnknown
	C.nk_window_collapse_if(carg0, cname, carg2, ccond)
}

// NkWindowShow function as declared in nk/nuklear.h:742
func NkWindowShow(arg0 *Context, name string, arg2 ShowStates) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cname, _ := unpackPCharString(name)
	carg2, _ := (C.enum_nk_show_states)(arg2), cgoAllocsUnknown
	C.nk_window_show(carg0, cname, carg2)
}

// NkWindowShowIf function as declared in nk/nuklear.h:743
func NkWindowShowIf(arg0 *Context, name string, arg2 ShowStates, cond int32) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cname, _ := unpackPCharString(name)
	carg2, _ := (C.enum_nk_show_states)(arg2), cgoAllocsUnknown
	ccond, _ := (C.int)(cond), cgoAllocsUnknown
	C.nk_window_show_if(carg0, cname, carg2, ccond)
}

// NkLayoutRowDynamic function as declared in nk/nuklear.h:746
func NkLayoutRowDynamic(arg0 *Context, height float32, cols int32) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cheight, _ := (C.float)(height), cgoAllocsUnknown
	ccols, _ := (C.int)(cols), cgoAllocsUnknown
	C.nk_layout_row_dynamic(carg0, cheight, ccols)
}

// NkLayoutRowStatic function as declared in nk/nuklear.h:747
func NkLayoutRowStatic(arg0 *Context, height float32, itemWidth int32, cols int32) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cheight, _ := (C.float)(height), cgoAllocsUnknown
	citemWidth, _ := (C.int)(itemWidth), cgoAllocsUnknown
	ccols, _ := (C.int)(cols), cgoAllocsUnknown
	C.nk_layout_row_static(carg0, cheight, citemWidth, ccols)
}

// NkLayoutRowBegin function as declared in nk/nuklear.h:749
func NkLayoutRowBegin(arg0 *Context, arg1 LayoutFormat, rowHeight float32, cols int32) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.enum_nk_layout_format)(arg1), cgoAllocsUnknown
	crowHeight, _ := (C.float)(rowHeight), cgoAllocsUnknown
	ccols, _ := (C.int)(cols), cgoAllocsUnknown
	C.nk_layout_row_begin(carg0, carg1, crowHeight, ccols)
}

// NkLayoutRowPush function as declared in nk/nuklear.h:750
func NkLayoutRowPush(arg0 *Context, value float32) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cvalue, _ := (C.float)(value), cgoAllocsUnknown
	C.nk_layout_row_push(carg0, cvalue)
}

// NkLayoutRowEnd function as declared in nk/nuklear.h:751
func NkLayoutRowEnd(arg0 *Context) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	C.nk_layout_row_end(carg0)
}

// NkLayoutRow function as declared in nk/nuklear.h:752
func NkLayoutRow(arg0 *Context, arg1 LayoutFormat, height float32, cols int32, ratio []float32) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.enum_nk_layout_format)(arg1), cgoAllocsUnknown
	cheight, _ := (C.float)(height), cgoAllocsUnknown
	ccols, _ := (C.int)(cols), cgoAllocsUnknown
	cratio, _ := (*C.float)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&ratio)).Data)), cgoAllocsUnknown
	C.nk_layout_row(carg0, carg1, cheight, ccols, cratio)
}

// NkLayoutSpaceBegin function as declared in nk/nuklear.h:754
func NkLayoutSpaceBegin(arg0 *Context, arg1 LayoutFormat, height float32, widgetCount int32) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.enum_nk_layout_format)(arg1), cgoAllocsUnknown
	cheight, _ := (C.float)(height), cgoAllocsUnknown
	cwidgetCount, _ := (C.int)(widgetCount), cgoAllocsUnknown
	C.nk_layout_space_begin(carg0, carg1, cheight, cwidgetCount)
}

// NkLayoutSpacePush function as declared in nk/nuklear.h:755
func NkLayoutSpacePush(arg0 *Context, arg1 Rect) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := *(*C.struct_nk_rect)(unsafe.Pointer(&arg1)), cgoAllocsUnknown
	C.nk_layout_space_push(carg0, carg1)
}

// NkLayoutSpaceEnd function as declared in nk/nuklear.h:756
func NkLayoutSpaceEnd(arg0 *Context) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	C.nk_layout_space_end(carg0)
}

// NkLayoutSpaceBounds function as declared in nk/nuklear.h:758
func NkLayoutSpaceBounds(arg0 *Context) Rect {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	__ret := C.nk_layout_space_bounds(carg0)
	__v := *(*Rect)(unsafe.Pointer(&__ret))
	return __v
}

// NkLayoutSpaceToScreen function as declared in nk/nuklear.h:759
func NkLayoutSpaceToScreen(arg0 *Context, arg1 Vec2) Vec2 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&arg1)), cgoAllocsUnknown
	__ret := C.nk_layout_space_to_screen(carg0, carg1)
	__v := *(*Vec2)(unsafe.Pointer(&__ret))
	return __v
}

// NkLayoutSpaceToLocal function as declared in nk/nuklear.h:760
func NkLayoutSpaceToLocal(arg0 *Context, arg1 Vec2) Vec2 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&arg1)), cgoAllocsUnknown
	__ret := C.nk_layout_space_to_local(carg0, carg1)
	__v := *(*Vec2)(unsafe.Pointer(&__ret))
	return __v
}

// NkLayoutSpaceRectToScreen function as declared in nk/nuklear.h:761
func NkLayoutSpaceRectToScreen(arg0 *Context, arg1 Rect) Rect {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := *(*C.struct_nk_rect)(unsafe.Pointer(&arg1)), cgoAllocsUnknown
	__ret := C.nk_layout_space_rect_to_screen(carg0, carg1)
	__v := *(*Rect)(unsafe.Pointer(&__ret))
	return __v
}

// NkLayoutSpaceRectToLocal function as declared in nk/nuklear.h:762
func NkLayoutSpaceRectToLocal(arg0 *Context, arg1 Rect) Rect {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := *(*C.struct_nk_rect)(unsafe.Pointer(&arg1)), cgoAllocsUnknown
	__ret := C.nk_layout_space_rect_to_local(carg0, carg1)
	__v := *(*Rect)(unsafe.Pointer(&__ret))
	return __v
}

// NkLayoutRatioFromPixel function as declared in nk/nuklear.h:763
func NkLayoutRatioFromPixel(arg0 *Context, pixelWidth float32) float32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cpixelWidth, _ := (C.float)(pixelWidth), cgoAllocsUnknown
	__ret := C.nk_layout_ratio_from_pixel(carg0, cpixelWidth)
	__v := (float32)(__ret)
	return __v
}

// NkGroupBegin function as declared in nk/nuklear.h:766
func NkGroupBegin(arg0 *Context, title string, arg2 Flags) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	ctitle, _ := unpackPCharString(title)
	carg2, _ := (C.nk_flags)(arg2), cgoAllocsUnknown
	__ret := C.nk_group_begin(carg0, ctitle, carg2)
	__v := (int32)(__ret)
	return __v
}

// NkGroupScrolledOffsetBegin function as declared in nk/nuklear.h:767
func NkGroupScrolledOffsetBegin(arg0 *Context, xOffset *Uint, yOffset *Uint, arg3 string, arg4 Flags) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cxOffset, _ := (*C.nk_uint)(unsafe.Pointer(xOffset)), cgoAllocsUnknown
	cyOffset, _ := (*C.nk_uint)(unsafe.Pointer(yOffset)), cgoAllocsUnknown
	carg3, _ := unpackPCharString(arg3)
	carg4, _ := (C.nk_flags)(arg4), cgoAllocsUnknown
	__ret := C.nk_group_scrolled_offset_begin(carg0, cxOffset, cyOffset, carg3, carg4)
	__v := (int32)(__ret)
	return __v
}

// NkGroupScrolledBegin function as declared in nk/nuklear.h:768
func NkGroupScrolledBegin(arg0 *Context, arg1 *Scroll, title string, arg3 Flags) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (*C.struct_nk_scroll)(unsafe.Pointer(arg1)), cgoAllocsUnknown
	ctitle, _ := unpackPCharString(title)
	carg3, _ := (C.nk_flags)(arg3), cgoAllocsUnknown
	__ret := C.nk_group_scrolled_begin(carg0, carg1, ctitle, carg3)
	__v := (int32)(__ret)
	return __v
}

// NkGroupScrolledEnd function as declared in nk/nuklear.h:769
func NkGroupScrolledEnd(arg0 *Context) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	C.nk_group_scrolled_end(carg0)
}

// NkGroupEnd function as declared in nk/nuklear.h:770
func NkGroupEnd(arg0 *Context) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	C.nk_group_end(carg0)
}

// NkListViewBegin function as declared in nk/nuklear.h:772
func NkListViewBegin(arg0 *Context, out *ListView, id string, arg3 Flags, rowHeight int32, rowCount int32) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cout, _ := (*C.struct_nk_list_view)(unsafe.Pointer(out)), cgoAllocsUnknown
	cid, _ := unpackPCharString(id)
	carg3, _ := (C.nk_flags)(arg3), cgoAllocsUnknown
	crowHeight, _ := (C.int)(rowHeight), cgoAllocsUnknown
	crowCount, _ := (C.int)(rowCount), cgoAllocsUnknown
	__ret := C.nk_list_view_begin(carg0, cout, cid, carg3, crowHeight, crowCount)
	__v := (int32)(__ret)
	return __v
}

// NkListViewEnd function as declared in nk/nuklear.h:773
func NkListViewEnd(arg0 *ListView) {
	carg0, _ := (*C.struct_nk_list_view)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	C.nk_list_view_end(carg0)
}

// NkTreePushHashed function as declared in nk/nuklear.h:778
func NkTreePushHashed(arg0 *Context, arg1 TreeType, title string, initialState CollapseStates, hash string, len int32, seed int32) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.enum_nk_tree_type)(arg1), cgoAllocsUnknown
	ctitle, _ := unpackPCharString(title)
	cinitialState, _ := (C.enum_nk_collapse_states)(initialState), cgoAllocsUnknown
	chash, _ := unpackPCharString(hash)
	clen, _ := (C.int)(len), cgoAllocsUnknown
	cseed, _ := (C.int)(seed), cgoAllocsUnknown
	__ret := C.nk_tree_push_hashed(carg0, carg1, ctitle, cinitialState, chash, clen, cseed)
	__v := (int32)(__ret)
	return __v
}

// NkTreeImagePushHashed function as declared in nk/nuklear.h:781
func NkTreeImagePushHashed(arg0 *Context, arg1 TreeType, arg2 Image, title string, initialState CollapseStates, hash string, len int32, seed int32) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.enum_nk_tree_type)(arg1), cgoAllocsUnknown
	carg2, _ := *(*C.struct_nk_image)(unsafe.Pointer(&arg2)), cgoAllocsUnknown
	ctitle, _ := unpackPCharString(title)
	cinitialState, _ := (C.enum_nk_collapse_states)(initialState), cgoAllocsUnknown
	chash, _ := unpackPCharString(hash)
	clen, _ := (C.int)(len), cgoAllocsUnknown
	cseed, _ := (C.int)(seed), cgoAllocsUnknown
	__ret := C.nk_tree_image_push_hashed(carg0, carg1, carg2, ctitle, cinitialState, chash, clen, cseed)
	__v := (int32)(__ret)
	return __v
}

// NkTreePop function as declared in nk/nuklear.h:782
func NkTreePop(arg0 *Context) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	C.nk_tree_pop(carg0)
}

// NkTreeStatePush function as declared in nk/nuklear.h:784
func NkTreeStatePush(arg0 *Context, arg1 TreeType, title string, state *CollapseStates) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.enum_nk_tree_type)(arg1), cgoAllocsUnknown
	ctitle, _ := unpackPCharString(title)
	cstate, _ := (*C.enum_nk_collapse_states)(unsafe.Pointer(state)), cgoAllocsUnknown
	__ret := C.nk_tree_state_push(carg0, carg1, ctitle, cstate)
	__v := (int32)(__ret)
	return __v
}

// NkTreeStateImagePush function as declared in nk/nuklear.h:785
func NkTreeStateImagePush(arg0 *Context, arg1 TreeType, arg2 Image, title string, state []CollapseStates) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.enum_nk_tree_type)(arg1), cgoAllocsUnknown
	carg2, _ := *(*C.struct_nk_image)(unsafe.Pointer(&arg2)), cgoAllocsUnknown
	ctitle, _ := unpackPCharString(title)
	cstate, _ := (*C.enum_nk_collapse_states)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&state)).Data)), cgoAllocsUnknown
	__ret := C.nk_tree_state_image_push(carg0, carg1, carg2, ctitle, cstate)
	__v := (int32)(__ret)
	return __v
}

// NkTreeStatePop function as declared in nk/nuklear.h:786
func NkTreeStatePop(arg0 *Context) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	C.nk_tree_state_pop(carg0)
}

// NkText function as declared in nk/nuklear.h:789
func NkText(arg0 *Context, arg1 string, arg2 int32, arg3 Flags) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	carg2, _ := (C.int)(arg2), cgoAllocsUnknown
	carg3, _ := (C.nk_flags)(arg3), cgoAllocsUnknown
	C.nk_text(carg0, carg1, carg2, carg3)
}

// NkTextColored function as declared in nk/nuklear.h:790
func NkTextColored(arg0 *Context, arg1 string, arg2 int32, arg3 Flags, arg4 Color) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	carg2, _ := (C.int)(arg2), cgoAllocsUnknown
	carg3, _ := (C.nk_flags)(arg3), cgoAllocsUnknown
	carg4, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg4)), cgoAllocsUnknown
	C.nk_text_colored(carg0, carg1, carg2, carg3, carg4)
}

// NkTextWrap function as declared in nk/nuklear.h:791
func NkTextWrap(arg0 *Context, arg1 string, arg2 int32) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	carg2, _ := (C.int)(arg2), cgoAllocsUnknown
	C.nk_text_wrap(carg0, carg1, carg2)
}

// NkTextWrapColored function as declared in nk/nuklear.h:792
func NkTextWrapColored(arg0 *Context, arg1 string, arg2 int32, arg3 Color) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	carg2, _ := (C.int)(arg2), cgoAllocsUnknown
	carg3, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg3)), cgoAllocsUnknown
	C.nk_text_wrap_colored(carg0, carg1, carg2, carg3)
}

// NkLabel function as declared in nk/nuklear.h:794
func NkLabel(arg0 *Context, arg1 string, align Flags) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	calign, _ := (C.nk_flags)(align), cgoAllocsUnknown
	C.nk_label(carg0, carg1, calign)
}

// NkLabelColored function as declared in nk/nuklear.h:795
func NkLabelColored(arg0 *Context, arg1 string, align Flags, arg3 Color) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	calign, _ := (C.nk_flags)(align), cgoAllocsUnknown
	carg3, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg3)), cgoAllocsUnknown
	C.nk_label_colored(carg0, carg1, calign, carg3)
}

// NkLabelWrap function as declared in nk/nuklear.h:796
func NkLabelWrap(arg0 *Context, arg1 string) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	C.nk_label_wrap(carg0, carg1)
}

// NkLabelColoredWrap function as declared in nk/nuklear.h:797
func NkLabelColoredWrap(arg0 *Context, arg1 string, arg2 Color) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	carg2, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg2)), cgoAllocsUnknown
	C.nk_label_colored_wrap(carg0, carg1, carg2)
}

// NkImage function as declared in nk/nuklear.h:798
func NkImage(arg0 *Context, arg1 Image) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := *(*C.struct_nk_image)(unsafe.Pointer(&arg1)), cgoAllocsUnknown
	C.nk_image(carg0, carg1)
}

// NkButtonSetBehavior function as declared in nk/nuklear.h:815
func NkButtonSetBehavior(arg0 *Context, arg1 ButtonBehavior) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.enum_nk_button_behavior)(arg1), cgoAllocsUnknown
	C.nk_button_set_behavior(carg0, carg1)
}

// NkButtonPushBehavior function as declared in nk/nuklear.h:816
func NkButtonPushBehavior(arg0 *Context, arg1 ButtonBehavior) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.enum_nk_button_behavior)(arg1), cgoAllocsUnknown
	__ret := C.nk_button_push_behavior(carg0, carg1)
	__v := (int32)(__ret)
	return __v
}

// NkButtonPopBehavior function as declared in nk/nuklear.h:817
func NkButtonPopBehavior(arg0 *Context) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	__ret := C.nk_button_pop_behavior(carg0)
	__v := (int32)(__ret)
	return __v
}

// NkButtonText function as declared in nk/nuklear.h:819
func NkButtonText(arg0 *Context, title string, len int32) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	ctitle, _ := unpackPCharString(title)
	clen, _ := (C.int)(len), cgoAllocsUnknown
	__ret := C.nk_button_text(carg0, ctitle, clen)
	__v := (int32)(__ret)
	return __v
}

// NkButtonLabel function as declared in nk/nuklear.h:820
func NkButtonLabel(arg0 *Context, title string) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	ctitle, _ := unpackPCharString(title)
	__ret := C.nk_button_label(carg0, ctitle)
	__v := (int32)(__ret)
	return __v
}

// NkButtonColor function as declared in nk/nuklear.h:821
func NkButtonColor(arg0 *Context, arg1 Color) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg1)), cgoAllocsUnknown
	__ret := C.nk_button_color(carg0, carg1)
	__v := (int32)(__ret)
	return __v
}

// NkButtonSymbol function as declared in nk/nuklear.h:822
func NkButtonSymbol(arg0 *Context, arg1 SymbolType) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.enum_nk_symbol_type)(arg1), cgoAllocsUnknown
	__ret := C.nk_button_symbol(carg0, carg1)
	__v := (int32)(__ret)
	return __v
}

// NkButtonImage function as declared in nk/nuklear.h:823
func NkButtonImage(arg0 *Context, img Image) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cimg, _ := *(*C.struct_nk_image)(unsafe.Pointer(&img)), cgoAllocsUnknown
	__ret := C.nk_button_image(carg0, cimg)
	__v := (int32)(__ret)
	return __v
}

// NkButtonSymbolLabel function as declared in nk/nuklear.h:824
func NkButtonSymbolLabel(arg0 *Context, arg1 SymbolType, arg2 string, textAlignment Flags) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.enum_nk_symbol_type)(arg1), cgoAllocsUnknown
	carg2, _ := unpackPCharString(arg2)
	ctextAlignment, _ := (C.nk_flags)(textAlignment), cgoAllocsUnknown
	__ret := C.nk_button_symbol_label(carg0, carg1, carg2, ctextAlignment)
	__v := (int32)(__ret)
	return __v
}

// NkButtonSymbolText function as declared in nk/nuklear.h:825
func NkButtonSymbolText(arg0 *Context, arg1 SymbolType, arg2 string, arg3 int32, alignment Flags) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.enum_nk_symbol_type)(arg1), cgoAllocsUnknown
	carg2, _ := unpackPCharString(arg2)
	carg3, _ := (C.int)(arg3), cgoAllocsUnknown
	calignment, _ := (C.nk_flags)(alignment), cgoAllocsUnknown
	__ret := C.nk_button_symbol_text(carg0, carg1, carg2, carg3, calignment)
	__v := (int32)(__ret)
	return __v
}

// NkButtonImageLabel function as declared in nk/nuklear.h:826
func NkButtonImageLabel(arg0 *Context, img Image, arg2 string, textAlignment Flags) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cimg, _ := *(*C.struct_nk_image)(unsafe.Pointer(&img)), cgoAllocsUnknown
	carg2, _ := unpackPCharString(arg2)
	ctextAlignment, _ := (C.nk_flags)(textAlignment), cgoAllocsUnknown
	__ret := C.nk_button_image_label(carg0, cimg, carg2, ctextAlignment)
	__v := (int32)(__ret)
	return __v
}

// NkButtonImageText function as declared in nk/nuklear.h:827
func NkButtonImageText(arg0 *Context, img Image, arg2 string, arg3 int32, alignment Flags) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cimg, _ := *(*C.struct_nk_image)(unsafe.Pointer(&img)), cgoAllocsUnknown
	carg2, _ := unpackPCharString(arg2)
	carg3, _ := (C.int)(arg3), cgoAllocsUnknown
	calignment, _ := (C.nk_flags)(alignment), cgoAllocsUnknown
	__ret := C.nk_button_image_text(carg0, cimg, carg2, carg3, calignment)
	__v := (int32)(__ret)
	return __v
}

// NkButtonTextStyled function as declared in nk/nuklear.h:829
func NkButtonTextStyled(arg0 *Context, arg1 *StyleButton, title string, len int32) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (*C.struct_nk_style_button)(unsafe.Pointer(arg1)), cgoAllocsUnknown
	ctitle, _ := unpackPCharString(title)
	clen, _ := (C.int)(len), cgoAllocsUnknown
	__ret := C.nk_button_text_styled(carg0, carg1, ctitle, clen)
	__v := (int32)(__ret)
	return __v
}

// NkButtonLabelStyled function as declared in nk/nuklear.h:830
func NkButtonLabelStyled(arg0 *Context, arg1 *StyleButton, title string) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (*C.struct_nk_style_button)(unsafe.Pointer(arg1)), cgoAllocsUnknown
	ctitle, _ := unpackPCharString(title)
	__ret := C.nk_button_label_styled(carg0, carg1, ctitle)
	__v := (int32)(__ret)
	return __v
}

// NkButtonSymbolStyled function as declared in nk/nuklear.h:831
func NkButtonSymbolStyled(arg0 *Context, arg1 *StyleButton, arg2 SymbolType) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (*C.struct_nk_style_button)(unsafe.Pointer(arg1)), cgoAllocsUnknown
	carg2, _ := (C.enum_nk_symbol_type)(arg2), cgoAllocsUnknown
	__ret := C.nk_button_symbol_styled(carg0, carg1, carg2)
	__v := (int32)(__ret)
	return __v
}

// NkButtonImageStyled function as declared in nk/nuklear.h:832
func NkButtonImageStyled(arg0 *Context, arg1 *StyleButton, img Image) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (*C.struct_nk_style_button)(unsafe.Pointer(arg1)), cgoAllocsUnknown
	cimg, _ := *(*C.struct_nk_image)(unsafe.Pointer(&img)), cgoAllocsUnknown
	__ret := C.nk_button_image_styled(carg0, carg1, cimg)
	__v := (int32)(__ret)
	return __v
}

// NkButtonSymbolLabelStyled function as declared in nk/nuklear.h:833
func NkButtonSymbolLabelStyled(arg0 *Context, arg1 *StyleButton, arg2 SymbolType, arg3 string, textAlignment Flags) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (*C.struct_nk_style_button)(unsafe.Pointer(arg1)), cgoAllocsUnknown
	carg2, _ := (C.enum_nk_symbol_type)(arg2), cgoAllocsUnknown
	carg3, _ := unpackPCharString(arg3)
	ctextAlignment, _ := (C.nk_flags)(textAlignment), cgoAllocsUnknown
	__ret := C.nk_button_symbol_label_styled(carg0, carg1, carg2, carg3, ctextAlignment)
	__v := (int32)(__ret)
	return __v
}

// NkButtonSymbolTextStyled function as declared in nk/nuklear.h:834
func NkButtonSymbolTextStyled(arg0 *Context, arg1 *StyleButton, arg2 SymbolType, arg3 string, arg4 int32, alignment Flags) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (*C.struct_nk_style_button)(unsafe.Pointer(arg1)), cgoAllocsUnknown
	carg2, _ := (C.enum_nk_symbol_type)(arg2), cgoAllocsUnknown
	carg3, _ := unpackPCharString(arg3)
	carg4, _ := (C.int)(arg4), cgoAllocsUnknown
	calignment, _ := (C.nk_flags)(alignment), cgoAllocsUnknown
	__ret := C.nk_button_symbol_text_styled(carg0, carg1, carg2, carg3, carg4, calignment)
	__v := (int32)(__ret)
	return __v
}

// NkButtonImageLabelStyled function as declared in nk/nuklear.h:835
func NkButtonImageLabelStyled(arg0 *Context, arg1 *StyleButton, img Image, arg3 string, textAlignment Flags) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (*C.struct_nk_style_button)(unsafe.Pointer(arg1)), cgoAllocsUnknown
	cimg, _ := *(*C.struct_nk_image)(unsafe.Pointer(&img)), cgoAllocsUnknown
	carg3, _ := unpackPCharString(arg3)
	ctextAlignment, _ := (C.nk_flags)(textAlignment), cgoAllocsUnknown
	__ret := C.nk_button_image_label_styled(carg0, carg1, cimg, carg3, ctextAlignment)
	__v := (int32)(__ret)
	return __v
}

// NkButtonImageTextStyled function as declared in nk/nuklear.h:836
func NkButtonImageTextStyled(arg0 *Context, arg1 *StyleButton, img Image, arg3 string, arg4 int32, alignment Flags) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (*C.struct_nk_style_button)(unsafe.Pointer(arg1)), cgoAllocsUnknown
	cimg, _ := *(*C.struct_nk_image)(unsafe.Pointer(&img)), cgoAllocsUnknown
	carg3, _ := unpackPCharString(arg3)
	carg4, _ := (C.int)(arg4), cgoAllocsUnknown
	calignment, _ := (C.nk_flags)(alignment), cgoAllocsUnknown
	__ret := C.nk_button_image_text_styled(carg0, carg1, cimg, carg3, carg4, calignment)
	__v := (int32)(__ret)
	return __v
}

// NkCheckLabel function as declared in nk/nuklear.h:839
func NkCheckLabel(arg0 *Context, arg1 string, active int32) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	cactive, _ := (C.int)(active), cgoAllocsUnknown
	__ret := C.nk_check_label(carg0, carg1, cactive)
	__v := (int32)(__ret)
	return __v
}

// NkCheckText function as declared in nk/nuklear.h:840
func NkCheckText(arg0 *Context, arg1 string, arg2 int32, active int32) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	carg2, _ := (C.int)(arg2), cgoAllocsUnknown
	cactive, _ := (C.int)(active), cgoAllocsUnknown
	__ret := C.nk_check_text(carg0, carg1, carg2, cactive)
	__v := (int32)(__ret)
	return __v
}

// NkCheckFlagsLabel function as declared in nk/nuklear.h:841
func NkCheckFlagsLabel(arg0 *Context, arg1 string, flags uint32, value uint32) uint32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	cflags, _ := (C.uint)(flags), cgoAllocsUnknown
	cvalue, _ := (C.uint)(value), cgoAllocsUnknown
	__ret := C.nk_check_flags_label(carg0, carg1, cflags, cvalue)
	__v := (uint32)(__ret)
	return __v
}

// NkCheckFlagsText function as declared in nk/nuklear.h:842
func NkCheckFlagsText(arg0 *Context, arg1 string, arg2 int32, flags uint32, value uint32) uint32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	carg2, _ := (C.int)(arg2), cgoAllocsUnknown
	cflags, _ := (C.uint)(flags), cgoAllocsUnknown
	cvalue, _ := (C.uint)(value), cgoAllocsUnknown
	__ret := C.nk_check_flags_text(carg0, carg1, carg2, cflags, cvalue)
	__v := (uint32)(__ret)
	return __v
}

// NkCheckboxLabel function as declared in nk/nuklear.h:843
func NkCheckboxLabel(arg0 *Context, arg1 string, active *int32) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	cactive, _ := (*C.int)(unsafe.Pointer(active)), cgoAllocsUnknown
	__ret := C.nk_checkbox_label(carg0, carg1, cactive)
	__v := (int32)(__ret)
	return __v
}

// NkCheckboxText function as declared in nk/nuklear.h:844
func NkCheckboxText(arg0 *Context, arg1 string, arg2 int32, active *int32) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	carg2, _ := (C.int)(arg2), cgoAllocsUnknown
	cactive, _ := (*C.int)(unsafe.Pointer(active)), cgoAllocsUnknown
	__ret := C.nk_checkbox_text(carg0, carg1, carg2, cactive)
	__v := (int32)(__ret)
	return __v
}

// NkCheckboxFlagsLabel function as declared in nk/nuklear.h:845
func NkCheckboxFlagsLabel(arg0 *Context, arg1 string, flags *uint32, value uint32) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	cflags, _ := (*C.uint)(unsafe.Pointer(flags)), cgoAllocsUnknown
	cvalue, _ := (C.uint)(value), cgoAllocsUnknown
	__ret := C.nk_checkbox_flags_label(carg0, carg1, cflags, cvalue)
	__v := (int32)(__ret)
	return __v
}

// NkCheckboxFlagsText function as declared in nk/nuklear.h:846
func NkCheckboxFlagsText(arg0 *Context, arg1 string, arg2 int32, flags *uint32, value uint32) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	carg2, _ := (C.int)(arg2), cgoAllocsUnknown
	cflags, _ := (*C.uint)(unsafe.Pointer(flags)), cgoAllocsUnknown
	cvalue, _ := (C.uint)(value), cgoAllocsUnknown
	__ret := C.nk_checkbox_flags_text(carg0, carg1, carg2, cflags, cvalue)
	__v := (int32)(__ret)
	return __v
}

// NkRadioLabel function as declared in nk/nuklear.h:849
func NkRadioLabel(arg0 *Context, arg1 string, active *int32) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	cactive, _ := (*C.int)(unsafe.Pointer(active)), cgoAllocsUnknown
	__ret := C.nk_radio_label(carg0, carg1, cactive)
	__v := (int32)(__ret)
	return __v
}

// NkRadioText function as declared in nk/nuklear.h:850
func NkRadioText(arg0 *Context, arg1 string, arg2 int32, active *int32) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	carg2, _ := (C.int)(arg2), cgoAllocsUnknown
	cactive, _ := (*C.int)(unsafe.Pointer(active)), cgoAllocsUnknown
	__ret := C.nk_radio_text(carg0, carg1, carg2, cactive)
	__v := (int32)(__ret)
	return __v
}

// NkOptionLabel function as declared in nk/nuklear.h:851
func NkOptionLabel(arg0 *Context, arg1 string, active int32) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	cactive, _ := (C.int)(active), cgoAllocsUnknown
	__ret := C.nk_option_label(carg0, carg1, cactive)
	__v := (int32)(__ret)
	return __v
}

// NkOptionText function as declared in nk/nuklear.h:852
func NkOptionText(arg0 *Context, arg1 string, arg2 int32, active int32) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	carg2, _ := (C.int)(arg2), cgoAllocsUnknown
	cactive, _ := (C.int)(active), cgoAllocsUnknown
	__ret := C.nk_option_text(carg0, carg1, carg2, cactive)
	__v := (int32)(__ret)
	return __v
}

// NkSelectableLabel function as declared in nk/nuklear.h:855
func NkSelectableLabel(arg0 *Context, arg1 string, align Flags, value *int32) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	calign, _ := (C.nk_flags)(align), cgoAllocsUnknown
	cvalue, _ := (*C.int)(unsafe.Pointer(value)), cgoAllocsUnknown
	__ret := C.nk_selectable_label(carg0, carg1, calign, cvalue)
	__v := (int32)(__ret)
	return __v
}

// NkSelectableText function as declared in nk/nuklear.h:856
func NkSelectableText(arg0 *Context, arg1 string, arg2 int32, align Flags, value []int32) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	carg2, _ := (C.int)(arg2), cgoAllocsUnknown
	calign, _ := (C.nk_flags)(align), cgoAllocsUnknown
	cvalue, _ := (*C.int)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&value)).Data)), cgoAllocsUnknown
	__ret := C.nk_selectable_text(carg0, carg1, carg2, calign, cvalue)
	__v := (int32)(__ret)
	return __v
}

// NkSelectableImageLabel function as declared in nk/nuklear.h:857
func NkSelectableImageLabel(arg0 *Context, arg1 Image, arg2 string, align Flags, value []int32) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := *(*C.struct_nk_image)(unsafe.Pointer(&arg1)), cgoAllocsUnknown
	carg2, _ := unpackPCharString(arg2)
	calign, _ := (C.nk_flags)(align), cgoAllocsUnknown
	cvalue, _ := (*C.int)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&value)).Data)), cgoAllocsUnknown
	__ret := C.nk_selectable_image_label(carg0, carg1, carg2, calign, cvalue)
	__v := (int32)(__ret)
	return __v
}

// NkSelectableImageText function as declared in nk/nuklear.h:858
func NkSelectableImageText(arg0 *Context, arg1 Image, arg2 string, arg3 int32, align Flags, value []int32) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := *(*C.struct_nk_image)(unsafe.Pointer(&arg1)), cgoAllocsUnknown
	carg2, _ := unpackPCharString(arg2)
	carg3, _ := (C.int)(arg3), cgoAllocsUnknown
	calign, _ := (C.nk_flags)(align), cgoAllocsUnknown
	cvalue, _ := (*C.int)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&value)).Data)), cgoAllocsUnknown
	__ret := C.nk_selectable_image_text(carg0, carg1, carg2, carg3, calign, cvalue)
	__v := (int32)(__ret)
	return __v
}

// NkSelectLabel function as declared in nk/nuklear.h:860
func NkSelectLabel(arg0 *Context, arg1 string, align Flags, value int32) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	calign, _ := (C.nk_flags)(align), cgoAllocsUnknown
	cvalue, _ := (C.int)(value), cgoAllocsUnknown
	__ret := C.nk_select_label(carg0, carg1, calign, cvalue)
	__v := (int32)(__ret)
	return __v
}

// NkSelectText function as declared in nk/nuklear.h:861
func NkSelectText(arg0 *Context, arg1 string, arg2 int32, align Flags, value int32) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	carg2, _ := (C.int)(arg2), cgoAllocsUnknown
	calign, _ := (C.nk_flags)(align), cgoAllocsUnknown
	cvalue, _ := (C.int)(value), cgoAllocsUnknown
	__ret := C.nk_select_text(carg0, carg1, carg2, calign, cvalue)
	__v := (int32)(__ret)
	return __v
}

// NkSelectImageLabel function as declared in nk/nuklear.h:862
func NkSelectImageLabel(arg0 *Context, arg1 Image, arg2 string, align Flags, value int32) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := *(*C.struct_nk_image)(unsafe.Pointer(&arg1)), cgoAllocsUnknown
	carg2, _ := unpackPCharString(arg2)
	calign, _ := (C.nk_flags)(align), cgoAllocsUnknown
	cvalue, _ := (C.int)(value), cgoAllocsUnknown
	__ret := C.nk_select_image_label(carg0, carg1, carg2, calign, cvalue)
	__v := (int32)(__ret)
	return __v
}

// NkSelectImageText function as declared in nk/nuklear.h:863
func NkSelectImageText(arg0 *Context, arg1 Image, arg2 string, arg3 int32, align Flags, value int32) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := *(*C.struct_nk_image)(unsafe.Pointer(&arg1)), cgoAllocsUnknown
	carg2, _ := unpackPCharString(arg2)
	carg3, _ := (C.int)(arg3), cgoAllocsUnknown
	calign, _ := (C.nk_flags)(align), cgoAllocsUnknown
	cvalue, _ := (C.int)(value), cgoAllocsUnknown
	__ret := C.nk_select_image_text(carg0, carg1, carg2, carg3, calign, cvalue)
	__v := (int32)(__ret)
	return __v
}

// NkSlideFloat function as declared in nk/nuklear.h:866
func NkSlideFloat(arg0 *Context, min float32, val float32, max float32, step float32) float32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cmin, _ := (C.float)(min), cgoAllocsUnknown
	cval, _ := (C.float)(val), cgoAllocsUnknown
	cmax, _ := (C.float)(max), cgoAllocsUnknown
	cstep, _ := (C.float)(step), cgoAllocsUnknown
	__ret := C.nk_slide_float(carg0, cmin, cval, cmax, cstep)
	__v := (float32)(__ret)
	return __v
}

// NkSlideInt function as declared in nk/nuklear.h:867
func NkSlideInt(arg0 *Context, min int32, val int32, max int32, step int32) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cmin, _ := (C.int)(min), cgoAllocsUnknown
	cval, _ := (C.int)(val), cgoAllocsUnknown
	cmax, _ := (C.int)(max), cgoAllocsUnknown
	cstep, _ := (C.int)(step), cgoAllocsUnknown
	__ret := C.nk_slide_int(carg0, cmin, cval, cmax, cstep)
	__v := (int32)(__ret)
	return __v
}

// NkSliderFloat function as declared in nk/nuklear.h:868
func NkSliderFloat(arg0 *Context, min float32, val *float32, max float32, step float32) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cmin, _ := (C.float)(min), cgoAllocsUnknown
	cval, _ := (*C.float)(unsafe.Pointer(val)), cgoAllocsUnknown
	cmax, _ := (C.float)(max), cgoAllocsUnknown
	cstep, _ := (C.float)(step), cgoAllocsUnknown
	__ret := C.nk_slider_float(carg0, cmin, cval, cmax, cstep)
	__v := (int32)(__ret)
	return __v
}

// NkSliderInt function as declared in nk/nuklear.h:869
func NkSliderInt(arg0 *Context, min int32, val *int32, max int32, step int32) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cmin, _ := (C.int)(min), cgoAllocsUnknown
	cval, _ := (*C.int)(unsafe.Pointer(val)), cgoAllocsUnknown
	cmax, _ := (C.int)(max), cgoAllocsUnknown
	cstep, _ := (C.int)(step), cgoAllocsUnknown
	__ret := C.nk_slider_int(carg0, cmin, cval, cmax, cstep)
	__v := (int32)(__ret)
	return __v
}

// NkProgress function as declared in nk/nuklear.h:872
func NkProgress(arg0 *Context, cur *Size, max Size, modifyable int32) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	ccur, _ := (*C.nk_size)(unsafe.Pointer(cur)), cgoAllocsUnknown
	cmax, _ := (C.nk_size)(max), cgoAllocsUnknown
	cmodifyable, _ := (C.int)(modifyable), cgoAllocsUnknown
	__ret := C.nk_progress(carg0, ccur, cmax, cmodifyable)
	__v := (int32)(__ret)
	return __v
}

// NkProg function as declared in nk/nuklear.h:873
func NkProg(arg0 *Context, cur Size, max Size, modifyable int32) Size {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	ccur, _ := (C.nk_size)(cur), cgoAllocsUnknown
	cmax, _ := (C.nk_size)(max), cgoAllocsUnknown
	cmodifyable, _ := (C.int)(modifyable), cgoAllocsUnknown
	__ret := C.nk_prog(carg0, ccur, cmax, cmodifyable)
	__v := (Size)(__ret)
	return __v
}

// NkColorPicker function as declared in nk/nuklear.h:876
func NkColorPicker(arg0 *Context, arg1 Color, arg2 ColorFormat) Color {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg1)), cgoAllocsUnknown
	carg2, _ := (C.enum_nk_color_format)(arg2), cgoAllocsUnknown
	__ret := C.nk_color_picker(carg0, carg1, carg2)
	__v := *(*Color)(unsafe.Pointer(&__ret))
	return __v
}

// NkColorPick function as declared in nk/nuklear.h:877
func NkColorPick(arg0 *Context, arg1 *Color, arg2 ColorFormat) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (*C.struct_nk_color)(unsafe.Pointer(arg1)), cgoAllocsUnknown
	carg2, _ := (C.enum_nk_color_format)(arg2), cgoAllocsUnknown
	__ret := C.nk_color_pick(carg0, carg1, carg2)
	__v := (int32)(__ret)
	return __v
}

// NkPropertyInt function as declared in nk/nuklear.h:880
func NkPropertyInt(arg0 *Context, name string, min int32, val *int32, max int32, step int32, incPerPixel float32) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cname, _ := unpackPCharString(name)
	cmin, _ := (C.int)(min), cgoAllocsUnknown
	cval, _ := (*C.int)(unsafe.Pointer(val)), cgoAllocsUnknown
	cmax, _ := (C.int)(max), cgoAllocsUnknown
	cstep, _ := (C.int)(step), cgoAllocsUnknown
	cincPerPixel, _ := (C.float)(incPerPixel), cgoAllocsUnknown
	C.nk_property_int(carg0, cname, cmin, cval, cmax, cstep, cincPerPixel)
}

// NkPropertyFloat function as declared in nk/nuklear.h:881
func NkPropertyFloat(arg0 *Context, name string, min float32, val *float32, max float32, step float32, incPerPixel float32) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cname, _ := unpackPCharString(name)
	cmin, _ := (C.float)(min), cgoAllocsUnknown
	cval, _ := (*C.float)(unsafe.Pointer(val)), cgoAllocsUnknown
	cmax, _ := (C.float)(max), cgoAllocsUnknown
	cstep, _ := (C.float)(step), cgoAllocsUnknown
	cincPerPixel, _ := (C.float)(incPerPixel), cgoAllocsUnknown
	C.nk_property_float(carg0, cname, cmin, cval, cmax, cstep, cincPerPixel)
}

// NkPropertyDouble function as declared in nk/nuklear.h:882
func NkPropertyDouble(arg0 *Context, name string, min float64, val *float64, max float64, step float64, incPerPixel float32) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cname, _ := unpackPCharString(name)
	cmin, _ := (C.double)(min), cgoAllocsUnknown
	cval, _ := (*C.double)(unsafe.Pointer(val)), cgoAllocsUnknown
	cmax, _ := (C.double)(max), cgoAllocsUnknown
	cstep, _ := (C.double)(step), cgoAllocsUnknown
	cincPerPixel, _ := (C.float)(incPerPixel), cgoAllocsUnknown
	C.nk_property_double(carg0, cname, cmin, cval, cmax, cstep, cincPerPixel)
}

// NkPropertyi function as declared in nk/nuklear.h:883
func NkPropertyi(arg0 *Context, name string, min int32, val int32, max int32, step int32, incPerPixel float32) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cname, _ := unpackPCharString(name)
	cmin, _ := (C.int)(min), cgoAllocsUnknown
	cval, _ := (C.int)(val), cgoAllocsUnknown
	cmax, _ := (C.int)(max), cgoAllocsUnknown
	cstep, _ := (C.int)(step), cgoAllocsUnknown
	cincPerPixel, _ := (C.float)(incPerPixel), cgoAllocsUnknown
	__ret := C.nk_propertyi(carg0, cname, cmin, cval, cmax, cstep, cincPerPixel)
	__v := (int32)(__ret)
	return __v
}

// NkPropertyf function as declared in nk/nuklear.h:884
func NkPropertyf(arg0 *Context, name string, min float32, val float32, max float32, step float32, incPerPixel float32) float32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cname, _ := unpackPCharString(name)
	cmin, _ := (C.float)(min), cgoAllocsUnknown
	cval, _ := (C.float)(val), cgoAllocsUnknown
	cmax, _ := (C.float)(max), cgoAllocsUnknown
	cstep, _ := (C.float)(step), cgoAllocsUnknown
	cincPerPixel, _ := (C.float)(incPerPixel), cgoAllocsUnknown
	__ret := C.nk_propertyf(carg0, cname, cmin, cval, cmax, cstep, cincPerPixel)
	__v := (float32)(__ret)
	return __v
}

// NkPropertyd function as declared in nk/nuklear.h:885
func NkPropertyd(arg0 *Context, name string, min float64, val float64, max float64, step float64, incPerPixel float32) float64 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cname, _ := unpackPCharString(name)
	cmin, _ := (C.double)(min), cgoAllocsUnknown
	cval, _ := (C.double)(val), cgoAllocsUnknown
	cmax, _ := (C.double)(max), cgoAllocsUnknown
	cstep, _ := (C.double)(step), cgoAllocsUnknown
	cincPerPixel, _ := (C.float)(incPerPixel), cgoAllocsUnknown
	__ret := C.nk_propertyd(carg0, cname, cmin, cval, cmax, cstep, cincPerPixel)
	__v := (float64)(__ret)
	return __v
}

// NkEditFocus function as declared in nk/nuklear.h:888
func NkEditFocus(ctx *Context, flags Flags) {
	cctx, _ := (*C.struct_nk_context)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	cflags, _ := (C.nk_flags)(flags), cgoAllocsUnknown
	C.nk_edit_focus(cctx, cflags)
}

// NkEditString function as declared in nk/nuklear.h:889
func NkEditString(arg0 *Context, arg1 Flags, buffer []byte, len *int32, max int32, arg5 PluginFilter) Flags {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.nk_flags)(arg1), cgoAllocsUnknown
	cbuffer, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&buffer)).Data)), cgoAllocsUnknown
	clen, _ := (*C.int)(unsafe.Pointer(len)), cgoAllocsUnknown
	cmax, _ := (C.int)(max), cgoAllocsUnknown
	carg5, _ := arg5.PassValue()
	__ret := C.nk_edit_string(carg0, carg1, cbuffer, clen, cmax, carg5)
	__v := (Flags)(__ret)
	return __v
}

// NkEditBuffer function as declared in nk/nuklear.h:890
func NkEditBuffer(arg0 *Context, arg1 Flags, arg2 *TextEdit, arg3 PluginFilter) Flags {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.nk_flags)(arg1), cgoAllocsUnknown
	carg2, _ := (*C.struct_nk_text_edit)(unsafe.Pointer(arg2)), cgoAllocsUnknown
	carg3, _ := arg3.PassValue()
	__ret := C.nk_edit_buffer(carg0, carg1, carg2, carg3)
	__v := (Flags)(__ret)
	return __v
}

// NkEditStringZeroTerminated function as declared in nk/nuklear.h:891
func NkEditStringZeroTerminated(arg0 *Context, arg1 Flags, buffer []byte, max int32, arg4 PluginFilter) Flags {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.nk_flags)(arg1), cgoAllocsUnknown
	cbuffer, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&buffer)).Data)), cgoAllocsUnknown
	cmax, _ := (C.int)(max), cgoAllocsUnknown
	carg4, _ := arg4.PassValue()
	__ret := C.nk_edit_string_zero_terminated(carg0, carg1, cbuffer, cmax, carg4)
	__v := (Flags)(__ret)
	return __v
}

// NkChartBegin function as declared in nk/nuklear.h:894
func NkChartBegin(arg0 *Context, arg1 ChartType, num int32, min float32, max float32) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.enum_nk_chart_type)(arg1), cgoAllocsUnknown
	cnum, _ := (C.int)(num), cgoAllocsUnknown
	cmin, _ := (C.float)(min), cgoAllocsUnknown
	cmax, _ := (C.float)(max), cgoAllocsUnknown
	__ret := C.nk_chart_begin(carg0, carg1, cnum, cmin, cmax)
	__v := (int32)(__ret)
	return __v
}

// NkChartBeginColored function as declared in nk/nuklear.h:895
func NkChartBeginColored(arg0 *Context, arg1 ChartType, arg2 Color, active Color, num int32, min float32, max float32) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.enum_nk_chart_type)(arg1), cgoAllocsUnknown
	carg2, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg2)), cgoAllocsUnknown
	cactive, _ := *(*C.struct_nk_color)(unsafe.Pointer(&active)), cgoAllocsUnknown
	cnum, _ := (C.int)(num), cgoAllocsUnknown
	cmin, _ := (C.float)(min), cgoAllocsUnknown
	cmax, _ := (C.float)(max), cgoAllocsUnknown
	__ret := C.nk_chart_begin_colored(carg0, carg1, carg2, cactive, cnum, cmin, cmax)
	__v := (int32)(__ret)
	return __v
}

// NkChartAddSlot function as declared in nk/nuklear.h:896
func NkChartAddSlot(ctx *Context, arg1 ChartType, count int32, minValue float32, maxValue float32) {
	cctx, _ := (*C.struct_nk_context)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	carg1, _ := (C.enum_nk_chart_type)(arg1), cgoAllocsUnknown
	ccount, _ := (C.int)(count), cgoAllocsUnknown
	cminValue, _ := (C.float)(minValue), cgoAllocsUnknown
	cmaxValue, _ := (C.float)(maxValue), cgoAllocsUnknown
	C.nk_chart_add_slot(cctx, carg1, ccount, cminValue, cmaxValue)
}

// NkChartAddSlotColored function as declared in nk/nuklear.h:897
func NkChartAddSlotColored(ctx *Context, arg1 ChartType, arg2 Color, active Color, count int32, minValue float32, maxValue float32) {
	cctx, _ := (*C.struct_nk_context)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	carg1, _ := (C.enum_nk_chart_type)(arg1), cgoAllocsUnknown
	carg2, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg2)), cgoAllocsUnknown
	cactive, _ := *(*C.struct_nk_color)(unsafe.Pointer(&active)), cgoAllocsUnknown
	ccount, _ := (C.int)(count), cgoAllocsUnknown
	cminValue, _ := (C.float)(minValue), cgoAllocsUnknown
	cmaxValue, _ := (C.float)(maxValue), cgoAllocsUnknown
	C.nk_chart_add_slot_colored(cctx, carg1, carg2, cactive, ccount, cminValue, cmaxValue)
}

// NkChartPush function as declared in nk/nuklear.h:898
func NkChartPush(arg0 *Context, arg1 float32) Flags {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.float)(arg1), cgoAllocsUnknown
	__ret := C.nk_chart_push(carg0, carg1)
	__v := (Flags)(__ret)
	return __v
}

// NkChartPushSlot function as declared in nk/nuklear.h:899
func NkChartPushSlot(arg0 *Context, arg1 float32, arg2 int32) Flags {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.float)(arg1), cgoAllocsUnknown
	carg2, _ := (C.int)(arg2), cgoAllocsUnknown
	__ret := C.nk_chart_push_slot(carg0, carg1, carg2)
	__v := (Flags)(__ret)
	return __v
}

// NkChartEnd function as declared in nk/nuklear.h:900
func NkChartEnd(arg0 *Context) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	C.nk_chart_end(carg0)
}

// NkPlot function as declared in nk/nuklear.h:901
func NkPlot(arg0 *Context, arg1 ChartType, values *float32, count int32, offset int32) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.enum_nk_chart_type)(arg1), cgoAllocsUnknown
	cvalues, _ := (*C.float)(unsafe.Pointer(values)), cgoAllocsUnknown
	ccount, _ := (C.int)(count), cgoAllocsUnknown
	coffset, _ := (C.int)(offset), cgoAllocsUnknown
	C.nk_plot(carg0, carg1, cvalues, ccount, coffset)
}

// NkPopupBegin function as declared in nk/nuklear.h:905
func NkPopupBegin(arg0 *Context, arg1 PopupType, arg2 string, arg3 Flags, bounds Rect) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.enum_nk_popup_type)(arg1), cgoAllocsUnknown
	carg2, _ := unpackPCharString(arg2)
	carg3, _ := (C.nk_flags)(arg3), cgoAllocsUnknown
	cbounds, _ := *(*C.struct_nk_rect)(unsafe.Pointer(&bounds)), cgoAllocsUnknown
	__ret := C.nk_popup_begin(carg0, carg1, carg2, carg3, cbounds)
	__v := (int32)(__ret)
	return __v
}

// NkPopupClose function as declared in nk/nuklear.h:906
func NkPopupClose(arg0 *Context) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	C.nk_popup_close(carg0)
}

// NkPopupEnd function as declared in nk/nuklear.h:907
func NkPopupEnd(arg0 *Context) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	C.nk_popup_end(carg0)
}

// NkCombo function as declared in nk/nuklear.h:910
func NkCombo(arg0 *Context, items []string, count int32, selected int32, itemHeight int32, size Vec2) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	citems, _ := unpackArgSString(items)
	ccount, _ := (C.int)(count), cgoAllocsUnknown
	cselected, _ := (C.int)(selected), cgoAllocsUnknown
	citemHeight, _ := (C.int)(itemHeight), cgoAllocsUnknown
	csize, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&size)), cgoAllocsUnknown
	__ret := C.nk_combo(carg0, citems, ccount, cselected, citemHeight, csize)
	packSString(items, citems)
	__v := (int32)(__ret)
	return __v
}

// NkComboSeparator function as declared in nk/nuklear.h:911
func NkComboSeparator(arg0 *Context, itemsSeparatedBySeparator string, separator int32, selected int32, count int32, itemHeight int32, size Vec2) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	citemsSeparatedBySeparator, _ := unpackPCharString(itemsSeparatedBySeparator)
	cseparator, _ := (C.int)(separator), cgoAllocsUnknown
	cselected, _ := (C.int)(selected), cgoAllocsUnknown
	ccount, _ := (C.int)(count), cgoAllocsUnknown
	citemHeight, _ := (C.int)(itemHeight), cgoAllocsUnknown
	csize, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&size)), cgoAllocsUnknown
	__ret := C.nk_combo_separator(carg0, citemsSeparatedBySeparator, cseparator, cselected, ccount, citemHeight, csize)
	__v := (int32)(__ret)
	return __v
}

// NkComboString function as declared in nk/nuklear.h:912
func NkComboString(arg0 *Context, itemsSeparatedByZeros string, selected int32, count int32, itemHeight int32, size Vec2) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	citemsSeparatedByZeros, _ := unpackPCharString(itemsSeparatedByZeros)
	cselected, _ := (C.int)(selected), cgoAllocsUnknown
	ccount, _ := (C.int)(count), cgoAllocsUnknown
	citemHeight, _ := (C.int)(itemHeight), cgoAllocsUnknown
	csize, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&size)), cgoAllocsUnknown
	__ret := C.nk_combo_string(carg0, citemsSeparatedByZeros, cselected, ccount, citemHeight, csize)
	__v := (int32)(__ret)
	return __v
}

// NkCombobox function as declared in nk/nuklear.h:914
func NkCombobox(arg0 *Context, items []string, count int32, selected *int32, itemHeight int32, size Vec2) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	citems, _ := unpackArgSString(items)
	ccount, _ := (C.int)(count), cgoAllocsUnknown
	cselected, _ := (*C.int)(unsafe.Pointer(selected)), cgoAllocsUnknown
	citemHeight, _ := (C.int)(itemHeight), cgoAllocsUnknown
	csize, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&size)), cgoAllocsUnknown
	C.nk_combobox(carg0, citems, ccount, cselected, citemHeight, csize)
	packSString(items, citems)
}

// NkComboboxString function as declared in nk/nuklear.h:915
func NkComboboxString(arg0 *Context, itemsSeparatedByZeros string, selected *int32, count int32, itemHeight int32, size Vec2) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	citemsSeparatedByZeros, _ := unpackPCharString(itemsSeparatedByZeros)
	cselected, _ := (*C.int)(unsafe.Pointer(selected)), cgoAllocsUnknown
	ccount, _ := (C.int)(count), cgoAllocsUnknown
	citemHeight, _ := (C.int)(itemHeight), cgoAllocsUnknown
	csize, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&size)), cgoAllocsUnknown
	C.nk_combobox_string(carg0, citemsSeparatedByZeros, cselected, ccount, citemHeight, csize)
}

// NkComboboxSeparator function as declared in nk/nuklear.h:916
func NkComboboxSeparator(arg0 *Context, itemsSeparatedBySeparator string, separator int32, selected *int32, count int32, itemHeight int32, size Vec2) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	citemsSeparatedBySeparator, _ := unpackPCharString(itemsSeparatedBySeparator)
	cseparator, _ := (C.int)(separator), cgoAllocsUnknown
	cselected, _ := (*C.int)(unsafe.Pointer(selected)), cgoAllocsUnknown
	ccount, _ := (C.int)(count), cgoAllocsUnknown
	citemHeight, _ := (C.int)(itemHeight), cgoAllocsUnknown
	csize, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&size)), cgoAllocsUnknown
	C.nk_combobox_separator(carg0, citemsSeparatedBySeparator, cseparator, cselected, ccount, citemHeight, csize)
}

// NkComboBeginText function as declared in nk/nuklear.h:920
func NkComboBeginText(arg0 *Context, selected string, arg2 int32, size Vec2) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cselected, _ := unpackPCharString(selected)
	carg2, _ := (C.int)(arg2), cgoAllocsUnknown
	csize, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&size)), cgoAllocsUnknown
	__ret := C.nk_combo_begin_text(carg0, cselected, carg2, csize)
	__v := (int32)(__ret)
	return __v
}

// NkComboBeginLabel function as declared in nk/nuklear.h:921
func NkComboBeginLabel(arg0 *Context, selected string, size Vec2) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cselected, _ := unpackPCharString(selected)
	csize, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&size)), cgoAllocsUnknown
	__ret := C.nk_combo_begin_label(carg0, cselected, csize)
	__v := (int32)(__ret)
	return __v
}

// NkComboBeginColor function as declared in nk/nuklear.h:922
func NkComboBeginColor(arg0 *Context, color Color, size Vec2) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	ccolor, _ := *(*C.struct_nk_color)(unsafe.Pointer(&color)), cgoAllocsUnknown
	csize, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&size)), cgoAllocsUnknown
	__ret := C.nk_combo_begin_color(carg0, ccolor, csize)
	__v := (int32)(__ret)
	return __v
}

// NkComboBeginSymbol function as declared in nk/nuklear.h:923
func NkComboBeginSymbol(arg0 *Context, arg1 SymbolType, size Vec2) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.enum_nk_symbol_type)(arg1), cgoAllocsUnknown
	csize, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&size)), cgoAllocsUnknown
	__ret := C.nk_combo_begin_symbol(carg0, carg1, csize)
	__v := (int32)(__ret)
	return __v
}

// NkComboBeginSymbolLabel function as declared in nk/nuklear.h:924
func NkComboBeginSymbolLabel(arg0 *Context, selected string, arg2 SymbolType, size Vec2) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cselected, _ := unpackPCharString(selected)
	carg2, _ := (C.enum_nk_symbol_type)(arg2), cgoAllocsUnknown
	csize, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&size)), cgoAllocsUnknown
	__ret := C.nk_combo_begin_symbol_label(carg0, cselected, carg2, csize)
	__v := (int32)(__ret)
	return __v
}

// NkComboBeginSymbolText function as declared in nk/nuklear.h:925
func NkComboBeginSymbolText(arg0 *Context, selected string, arg2 int32, arg3 SymbolType, size Vec2) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cselected, _ := unpackPCharString(selected)
	carg2, _ := (C.int)(arg2), cgoAllocsUnknown
	carg3, _ := (C.enum_nk_symbol_type)(arg3), cgoAllocsUnknown
	csize, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&size)), cgoAllocsUnknown
	__ret := C.nk_combo_begin_symbol_text(carg0, cselected, carg2, carg3, csize)
	__v := (int32)(__ret)
	return __v
}

// NkComboBeginImage function as declared in nk/nuklear.h:926
func NkComboBeginImage(arg0 *Context, img Image, size Vec2) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cimg, _ := *(*C.struct_nk_image)(unsafe.Pointer(&img)), cgoAllocsUnknown
	csize, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&size)), cgoAllocsUnknown
	__ret := C.nk_combo_begin_image(carg0, cimg, csize)
	__v := (int32)(__ret)
	return __v
}

// NkComboBeginImageLabel function as declared in nk/nuklear.h:927
func NkComboBeginImageLabel(arg0 *Context, selected string, arg2 Image, size Vec2) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cselected, _ := unpackPCharString(selected)
	carg2, _ := *(*C.struct_nk_image)(unsafe.Pointer(&arg2)), cgoAllocsUnknown
	csize, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&size)), cgoAllocsUnknown
	__ret := C.nk_combo_begin_image_label(carg0, cselected, carg2, csize)
	__v := (int32)(__ret)
	return __v
}

// NkComboBeginImageText function as declared in nk/nuklear.h:928
func NkComboBeginImageText(arg0 *Context, selected string, arg2 int32, arg3 Image, size Vec2) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cselected, _ := unpackPCharString(selected)
	carg2, _ := (C.int)(arg2), cgoAllocsUnknown
	carg3, _ := *(*C.struct_nk_image)(unsafe.Pointer(&arg3)), cgoAllocsUnknown
	csize, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&size)), cgoAllocsUnknown
	__ret := C.nk_combo_begin_image_text(carg0, cselected, carg2, carg3, csize)
	__v := (int32)(__ret)
	return __v
}

// NkComboItemLabel function as declared in nk/nuklear.h:929
func NkComboItemLabel(arg0 *Context, arg1 string, alignment Flags) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	calignment, _ := (C.nk_flags)(alignment), cgoAllocsUnknown
	__ret := C.nk_combo_item_label(carg0, carg1, calignment)
	__v := (int32)(__ret)
	return __v
}

// NkComboItemText function as declared in nk/nuklear.h:930
func NkComboItemText(arg0 *Context, arg1 string, arg2 int32, alignment Flags) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	carg2, _ := (C.int)(arg2), cgoAllocsUnknown
	calignment, _ := (C.nk_flags)(alignment), cgoAllocsUnknown
	__ret := C.nk_combo_item_text(carg0, carg1, carg2, calignment)
	__v := (int32)(__ret)
	return __v
}

// NkComboItemImageLabel function as declared in nk/nuklear.h:931
func NkComboItemImageLabel(arg0 *Context, arg1 Image, arg2 string, alignment Flags) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := *(*C.struct_nk_image)(unsafe.Pointer(&arg1)), cgoAllocsUnknown
	carg2, _ := unpackPCharString(arg2)
	calignment, _ := (C.nk_flags)(alignment), cgoAllocsUnknown
	__ret := C.nk_combo_item_image_label(carg0, carg1, carg2, calignment)
	__v := (int32)(__ret)
	return __v
}

// NkComboItemImageText function as declared in nk/nuklear.h:932
func NkComboItemImageText(arg0 *Context, arg1 Image, arg2 string, arg3 int32, alignment Flags) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := *(*C.struct_nk_image)(unsafe.Pointer(&arg1)), cgoAllocsUnknown
	carg2, _ := unpackPCharString(arg2)
	carg3, _ := (C.int)(arg3), cgoAllocsUnknown
	calignment, _ := (C.nk_flags)(alignment), cgoAllocsUnknown
	__ret := C.nk_combo_item_image_text(carg0, carg1, carg2, carg3, calignment)
	__v := (int32)(__ret)
	return __v
}

// NkComboItemSymbolLabel function as declared in nk/nuklear.h:933
func NkComboItemSymbolLabel(arg0 *Context, arg1 SymbolType, arg2 string, alignment Flags) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.enum_nk_symbol_type)(arg1), cgoAllocsUnknown
	carg2, _ := unpackPCharString(arg2)
	calignment, _ := (C.nk_flags)(alignment), cgoAllocsUnknown
	__ret := C.nk_combo_item_symbol_label(carg0, carg1, carg2, calignment)
	__v := (int32)(__ret)
	return __v
}

// NkComboItemSymbolText function as declared in nk/nuklear.h:934
func NkComboItemSymbolText(arg0 *Context, arg1 SymbolType, arg2 string, arg3 int32, alignment Flags) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.enum_nk_symbol_type)(arg1), cgoAllocsUnknown
	carg2, _ := unpackPCharString(arg2)
	carg3, _ := (C.int)(arg3), cgoAllocsUnknown
	calignment, _ := (C.nk_flags)(alignment), cgoAllocsUnknown
	__ret := C.nk_combo_item_symbol_text(carg0, carg1, carg2, carg3, calignment)
	__v := (int32)(__ret)
	return __v
}

// NkComboClose function as declared in nk/nuklear.h:935
func NkComboClose(arg0 *Context) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	C.nk_combo_close(carg0)
}

// NkComboEnd function as declared in nk/nuklear.h:936
func NkComboEnd(arg0 *Context) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	C.nk_combo_end(carg0)
}

// NkContextualBegin function as declared in nk/nuklear.h:939
func NkContextualBegin(arg0 *Context, arg1 Flags, arg2 Vec2, triggerBounds Rect) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.nk_flags)(arg1), cgoAllocsUnknown
	carg2, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&arg2)), cgoAllocsUnknown
	ctriggerBounds, _ := *(*C.struct_nk_rect)(unsafe.Pointer(&triggerBounds)), cgoAllocsUnknown
	__ret := C.nk_contextual_begin(carg0, carg1, carg2, ctriggerBounds)
	__v := (int32)(__ret)
	return __v
}

// NkContextualItemText function as declared in nk/nuklear.h:940
func NkContextualItemText(arg0 *Context, arg1 string, arg2 int32, align Flags) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	carg2, _ := (C.int)(arg2), cgoAllocsUnknown
	calign, _ := (C.nk_flags)(align), cgoAllocsUnknown
	__ret := C.nk_contextual_item_text(carg0, carg1, carg2, calign)
	__v := (int32)(__ret)
	return __v
}

// NkContextualItemLabel function as declared in nk/nuklear.h:941
func NkContextualItemLabel(arg0 *Context, arg1 string, align Flags) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	calign, _ := (C.nk_flags)(align), cgoAllocsUnknown
	__ret := C.nk_contextual_item_label(carg0, carg1, calign)
	__v := (int32)(__ret)
	return __v
}

// NkContextualItemImageLabel function as declared in nk/nuklear.h:942
func NkContextualItemImageLabel(arg0 *Context, arg1 Image, arg2 string, alignment Flags) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := *(*C.struct_nk_image)(unsafe.Pointer(&arg1)), cgoAllocsUnknown
	carg2, _ := unpackPCharString(arg2)
	calignment, _ := (C.nk_flags)(alignment), cgoAllocsUnknown
	__ret := C.nk_contextual_item_image_label(carg0, carg1, carg2, calignment)
	__v := (int32)(__ret)
	return __v
}

// NkContextualItemImageText function as declared in nk/nuklear.h:943
func NkContextualItemImageText(arg0 *Context, arg1 Image, arg2 string, len int32, alignment Flags) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := *(*C.struct_nk_image)(unsafe.Pointer(&arg1)), cgoAllocsUnknown
	carg2, _ := unpackPCharString(arg2)
	clen, _ := (C.int)(len), cgoAllocsUnknown
	calignment, _ := (C.nk_flags)(alignment), cgoAllocsUnknown
	__ret := C.nk_contextual_item_image_text(carg0, carg1, carg2, clen, calignment)
	__v := (int32)(__ret)
	return __v
}

// NkContextualItemSymbolLabel function as declared in nk/nuklear.h:944
func NkContextualItemSymbolLabel(arg0 *Context, arg1 SymbolType, arg2 string, alignment Flags) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.enum_nk_symbol_type)(arg1), cgoAllocsUnknown
	carg2, _ := unpackPCharString(arg2)
	calignment, _ := (C.nk_flags)(alignment), cgoAllocsUnknown
	__ret := C.nk_contextual_item_symbol_label(carg0, carg1, carg2, calignment)
	__v := (int32)(__ret)
	return __v
}

// NkContextualItemSymbolText function as declared in nk/nuklear.h:945
func NkContextualItemSymbolText(arg0 *Context, arg1 SymbolType, arg2 string, arg3 int32, alignment Flags) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.enum_nk_symbol_type)(arg1), cgoAllocsUnknown
	carg2, _ := unpackPCharString(arg2)
	carg3, _ := (C.int)(arg3), cgoAllocsUnknown
	calignment, _ := (C.nk_flags)(alignment), cgoAllocsUnknown
	__ret := C.nk_contextual_item_symbol_text(carg0, carg1, carg2, carg3, calignment)
	__v := (int32)(__ret)
	return __v
}

// NkContextualClose function as declared in nk/nuklear.h:946
func NkContextualClose(arg0 *Context) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	C.nk_contextual_close(carg0)
}

// NkContextualEnd function as declared in nk/nuklear.h:947
func NkContextualEnd(arg0 *Context) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	C.nk_contextual_end(carg0)
}

// NkTooltip function as declared in nk/nuklear.h:950
func NkTooltip(arg0 *Context, arg1 string) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	C.nk_tooltip(carg0, carg1)
}

// NkTooltipBegin function as declared in nk/nuklear.h:951
func NkTooltipBegin(arg0 *Context, width float32) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cwidth, _ := (C.float)(width), cgoAllocsUnknown
	__ret := C.nk_tooltip_begin(carg0, cwidth)
	__v := (int32)(__ret)
	return __v
}

// NkTooltipEnd function as declared in nk/nuklear.h:952
func NkTooltipEnd(arg0 *Context) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	C.nk_tooltip_end(carg0)
}

// NkMenubarBegin function as declared in nk/nuklear.h:955
func NkMenubarBegin(arg0 *Context) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	C.nk_menubar_begin(carg0)
}

// NkMenubarEnd function as declared in nk/nuklear.h:956
func NkMenubarEnd(arg0 *Context) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	C.nk_menubar_end(carg0)
}

// NkMenuBeginText function as declared in nk/nuklear.h:958
func NkMenuBeginText(arg0 *Context, title string, titleLen int32, align Flags, size Vec2) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	ctitle, _ := unpackPCharString(title)
	ctitleLen, _ := (C.int)(titleLen), cgoAllocsUnknown
	calign, _ := (C.nk_flags)(align), cgoAllocsUnknown
	csize, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&size)), cgoAllocsUnknown
	__ret := C.nk_menu_begin_text(carg0, ctitle, ctitleLen, calign, csize)
	__v := (int32)(__ret)
	return __v
}

// NkMenuBeginLabel function as declared in nk/nuklear.h:959
func NkMenuBeginLabel(arg0 *Context, arg1 string, align Flags, size Vec2) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	calign, _ := (C.nk_flags)(align), cgoAllocsUnknown
	csize, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&size)), cgoAllocsUnknown
	__ret := C.nk_menu_begin_label(carg0, carg1, calign, csize)
	__v := (int32)(__ret)
	return __v
}

// NkMenuBeginImage function as declared in nk/nuklear.h:960
func NkMenuBeginImage(arg0 *Context, arg1 string, arg2 Image, size Vec2) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	carg2, _ := *(*C.struct_nk_image)(unsafe.Pointer(&arg2)), cgoAllocsUnknown
	csize, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&size)), cgoAllocsUnknown
	__ret := C.nk_menu_begin_image(carg0, carg1, carg2, csize)
	__v := (int32)(__ret)
	return __v
}

// NkMenuBeginImageText function as declared in nk/nuklear.h:961
func NkMenuBeginImageText(arg0 *Context, arg1 string, arg2 int32, align Flags, arg4 Image, size Vec2) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	carg2, _ := (C.int)(arg2), cgoAllocsUnknown
	calign, _ := (C.nk_flags)(align), cgoAllocsUnknown
	carg4, _ := *(*C.struct_nk_image)(unsafe.Pointer(&arg4)), cgoAllocsUnknown
	csize, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&size)), cgoAllocsUnknown
	__ret := C.nk_menu_begin_image_text(carg0, carg1, carg2, calign, carg4, csize)
	__v := (int32)(__ret)
	return __v
}

// NkMenuBeginImageLabel function as declared in nk/nuklear.h:962
func NkMenuBeginImageLabel(arg0 *Context, arg1 string, align Flags, arg3 Image, size Vec2) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	calign, _ := (C.nk_flags)(align), cgoAllocsUnknown
	carg3, _ := *(*C.struct_nk_image)(unsafe.Pointer(&arg3)), cgoAllocsUnknown
	csize, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&size)), cgoAllocsUnknown
	__ret := C.nk_menu_begin_image_label(carg0, carg1, calign, carg3, csize)
	__v := (int32)(__ret)
	return __v
}

// NkMenuBeginSymbol function as declared in nk/nuklear.h:963
func NkMenuBeginSymbol(arg0 *Context, arg1 string, arg2 SymbolType, size Vec2) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	carg2, _ := (C.enum_nk_symbol_type)(arg2), cgoAllocsUnknown
	csize, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&size)), cgoAllocsUnknown
	__ret := C.nk_menu_begin_symbol(carg0, carg1, carg2, csize)
	__v := (int32)(__ret)
	return __v
}

// NkMenuBeginSymbolText function as declared in nk/nuklear.h:964
func NkMenuBeginSymbolText(arg0 *Context, arg1 string, arg2 int32, align Flags, arg4 SymbolType, size Vec2) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	carg2, _ := (C.int)(arg2), cgoAllocsUnknown
	calign, _ := (C.nk_flags)(align), cgoAllocsUnknown
	carg4, _ := (C.enum_nk_symbol_type)(arg4), cgoAllocsUnknown
	csize, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&size)), cgoAllocsUnknown
	__ret := C.nk_menu_begin_symbol_text(carg0, carg1, carg2, calign, carg4, csize)
	__v := (int32)(__ret)
	return __v
}

// NkMenuBeginSymbolLabel function as declared in nk/nuklear.h:965
func NkMenuBeginSymbolLabel(arg0 *Context, arg1 string, align Flags, arg3 SymbolType, size Vec2) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	calign, _ := (C.nk_flags)(align), cgoAllocsUnknown
	carg3, _ := (C.enum_nk_symbol_type)(arg3), cgoAllocsUnknown
	csize, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&size)), cgoAllocsUnknown
	__ret := C.nk_menu_begin_symbol_label(carg0, carg1, calign, carg3, csize)
	__v := (int32)(__ret)
	return __v
}

// NkMenuItemText function as declared in nk/nuklear.h:966
func NkMenuItemText(arg0 *Context, arg1 string, arg2 int32, align Flags) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	carg2, _ := (C.int)(arg2), cgoAllocsUnknown
	calign, _ := (C.nk_flags)(align), cgoAllocsUnknown
	__ret := C.nk_menu_item_text(carg0, carg1, carg2, calign)
	__v := (int32)(__ret)
	return __v
}

// NkMenuItemLabel function as declared in nk/nuklear.h:967
func NkMenuItemLabel(arg0 *Context, arg1 string, alignment Flags) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	calignment, _ := (C.nk_flags)(alignment), cgoAllocsUnknown
	__ret := C.nk_menu_item_label(carg0, carg1, calignment)
	__v := (int32)(__ret)
	return __v
}

// NkMenuItemImageLabel function as declared in nk/nuklear.h:968
func NkMenuItemImageLabel(arg0 *Context, arg1 Image, arg2 string, alignment Flags) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := *(*C.struct_nk_image)(unsafe.Pointer(&arg1)), cgoAllocsUnknown
	carg2, _ := unpackPCharString(arg2)
	calignment, _ := (C.nk_flags)(alignment), cgoAllocsUnknown
	__ret := C.nk_menu_item_image_label(carg0, carg1, carg2, calignment)
	__v := (int32)(__ret)
	return __v
}

// NkMenuItemImageText function as declared in nk/nuklear.h:969
func NkMenuItemImageText(arg0 *Context, arg1 Image, arg2 string, len int32, alignment Flags) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := *(*C.struct_nk_image)(unsafe.Pointer(&arg1)), cgoAllocsUnknown
	carg2, _ := unpackPCharString(arg2)
	clen, _ := (C.int)(len), cgoAllocsUnknown
	calignment, _ := (C.nk_flags)(alignment), cgoAllocsUnknown
	__ret := C.nk_menu_item_image_text(carg0, carg1, carg2, clen, calignment)
	__v := (int32)(__ret)
	return __v
}

// NkMenuItemSymbolText function as declared in nk/nuklear.h:970
func NkMenuItemSymbolText(arg0 *Context, arg1 SymbolType, arg2 string, arg3 int32, alignment Flags) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.enum_nk_symbol_type)(arg1), cgoAllocsUnknown
	carg2, _ := unpackPCharString(arg2)
	carg3, _ := (C.int)(arg3), cgoAllocsUnknown
	calignment, _ := (C.nk_flags)(alignment), cgoAllocsUnknown
	__ret := C.nk_menu_item_symbol_text(carg0, carg1, carg2, carg3, calignment)
	__v := (int32)(__ret)
	return __v
}

// NkMenuItemSymbolLabel function as declared in nk/nuklear.h:971
func NkMenuItemSymbolLabel(arg0 *Context, arg1 SymbolType, arg2 string, alignment Flags) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.enum_nk_symbol_type)(arg1), cgoAllocsUnknown
	carg2, _ := unpackPCharString(arg2)
	calignment, _ := (C.nk_flags)(alignment), cgoAllocsUnknown
	__ret := C.nk_menu_item_symbol_label(carg0, carg1, carg2, calignment)
	__v := (int32)(__ret)
	return __v
}

// NkMenuClose function as declared in nk/nuklear.h:972
func NkMenuClose(arg0 *Context) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	C.nk_menu_close(carg0)
}

// NkMenuEnd function as declared in nk/nuklear.h:973
func NkMenuEnd(arg0 *Context) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	C.nk_menu_end(carg0)
}

// NkConvert function as declared in nk/nuklear.h:978
func NkConvert(arg0 *Context, cmds *Buffer, vertices *Buffer, elements *Buffer, arg4 *ConvertConfig) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	ccmds, _ := (*C.struct_nk_buffer)(unsafe.Pointer(cmds)), cgoAllocsUnknown
	cvertices, _ := (*C.struct_nk_buffer)(unsafe.Pointer(vertices)), cgoAllocsUnknown
	celements, _ := (*C.struct_nk_buffer)(unsafe.Pointer(elements)), cgoAllocsUnknown
	carg4, _ := arg4.PassRef()
	C.nk_convert(carg0, ccmds, cvertices, celements, carg4)
}

// Nk_DrawBegin function as declared in nk/nuklear.h:981
func Nk_DrawBegin(arg0 *Context, arg1 *Buffer) *DrawCommand {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (*C.struct_nk_buffer)(unsafe.Pointer(arg1)), cgoAllocsUnknown
	__ret := C.nk__draw_begin(carg0, carg1)
	__v := *(**DrawCommand)(unsafe.Pointer(&__ret))
	return __v
}

// Nk_DrawEnd function as declared in nk/nuklear.h:982
func Nk_DrawEnd(arg0 *Context, arg1 *Buffer) *DrawCommand {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (*C.struct_nk_buffer)(unsafe.Pointer(arg1)), cgoAllocsUnknown
	__ret := C.nk__draw_end(carg0, carg1)
	__v := *(**DrawCommand)(unsafe.Pointer(&__ret))
	return __v
}

// Nk_DrawNext function as declared in nk/nuklear.h:983
func Nk_DrawNext(arg0 *DrawCommand, arg1 *Buffer, arg2 *Context) *DrawCommand {
	carg0, _ := (*C.struct_nk_draw_command)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (*C.struct_nk_buffer)(unsafe.Pointer(arg1)), cgoAllocsUnknown
	carg2, _ := (*C.struct_nk_context)(unsafe.Pointer(arg2)), cgoAllocsUnknown
	__ret := C.nk__draw_next(carg0, carg1, carg2)
	__v := *(**DrawCommand)(unsafe.Pointer(&__ret))
	return __v
}

// NkInputBegin function as declared in nk/nuklear.h:987
func NkInputBegin(arg0 *Context) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	C.nk_input_begin(carg0)
}

// NkInputMotion function as declared in nk/nuklear.h:988
func NkInputMotion(arg0 *Context, x int32, y int32) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cx, _ := (C.int)(x), cgoAllocsUnknown
	cy, _ := (C.int)(y), cgoAllocsUnknown
	C.nk_input_motion(carg0, cx, cy)
}

// NkInputKey function as declared in nk/nuklear.h:989
func NkInputKey(arg0 *Context, arg1 Keys, down int32) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.enum_nk_keys)(arg1), cgoAllocsUnknown
	cdown, _ := (C.int)(down), cgoAllocsUnknown
	C.nk_input_key(carg0, carg1, cdown)
}

// NkInputButton function as declared in nk/nuklear.h:990
func NkInputButton(arg0 *Context, arg1 Buttons, x int32, y int32, down int32) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.enum_nk_buttons)(arg1), cgoAllocsUnknown
	cx, _ := (C.int)(x), cgoAllocsUnknown
	cy, _ := (C.int)(y), cgoAllocsUnknown
	cdown, _ := (C.int)(down), cgoAllocsUnknown
	C.nk_input_button(carg0, carg1, cx, cy, cdown)
}

// NkInputScroll function as declared in nk/nuklear.h:991
func NkInputScroll(arg0 *Context, y float32) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cy, _ := (C.float)(y), cgoAllocsUnknown
	C.nk_input_scroll(carg0, cy)
}

// NkInputChar function as declared in nk/nuklear.h:992
func NkInputChar(arg0 *Context, arg1 byte) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.char)(arg1), cgoAllocsUnknown
	C.nk_input_char(carg0, carg1)
}

// NkInputUnicode function as declared in nk/nuklear.h:994
func NkInputUnicode(arg0 *Context, arg1 Rune) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.nk_rune)(arg1), cgoAllocsUnknown
	C.nk_input_unicode(carg0, carg1)
}

// NkInputEnd function as declared in nk/nuklear.h:995
func NkInputEnd(arg0 *Context) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	C.nk_input_end(carg0)
}

// NkStyleDefault function as declared in nk/nuklear.h:998
func NkStyleDefault(arg0 *Context) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	C.nk_style_default(carg0)
}

// NkStyleFromTable function as declared in nk/nuklear.h:999
func NkStyleFromTable(arg0 *Context, arg1 *Color) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (*C.struct_nk_color)(unsafe.Pointer(arg1)), cgoAllocsUnknown
	C.nk_style_from_table(carg0, carg1)
}

// NkStyleLoadCursor function as declared in nk/nuklear.h:1000
func NkStyleLoadCursor(arg0 *Context, arg1 StyleCursor, arg2 *Cursor) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.enum_nk_style_cursor)(arg1), cgoAllocsUnknown
	carg2, _ := (*C.struct_nk_cursor)(unsafe.Pointer(arg2)), cgoAllocsUnknown
	C.nk_style_load_cursor(carg0, carg1, carg2)
}

// NkStyleLoadAllCursors function as declared in nk/nuklear.h:1001
func NkStyleLoadAllCursors(arg0 *Context, arg1 *Cursor) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (*C.struct_nk_cursor)(unsafe.Pointer(arg1)), cgoAllocsUnknown
	C.nk_style_load_all_cursors(carg0, carg1)
}

// NkStyleGetColorByName function as declared in nk/nuklear.h:1002
func NkStyleGetColorByName(arg0 StyleColors) string {
	carg0, _ := (C.enum_nk_style_colors)(arg0), cgoAllocsUnknown
	__ret := C.nk_style_get_color_by_name(carg0)
	__v := packPCharString(__ret)
	return __v
}

// NkStyleSetFont function as declared in nk/nuklear.h:1003
func NkStyleSetFont(arg0 *Context, arg1 *UserFont) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := arg1.PassRef()
	C.nk_style_set_font(carg0, carg1)
}

// NkStyleSetCursor function as declared in nk/nuklear.h:1004
func NkStyleSetCursor(arg0 *Context, arg1 StyleCursor) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.enum_nk_style_cursor)(arg1), cgoAllocsUnknown
	__ret := C.nk_style_set_cursor(carg0, carg1)
	__v := (int32)(__ret)
	return __v
}

// NkStyleShowCursor function as declared in nk/nuklear.h:1005
func NkStyleShowCursor(arg0 *Context) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	C.nk_style_show_cursor(carg0)
}

// NkStyleHideCursor function as declared in nk/nuklear.h:1006
func NkStyleHideCursor(arg0 *Context) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	C.nk_style_hide_cursor(carg0)
}

// NkStylePushFont function as declared in nk/nuklear.h:1009
func NkStylePushFont(arg0 *Context, arg1 *UserFont) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := arg1.PassRef()
	__ret := C.nk_style_push_font(carg0, carg1)
	__v := (int32)(__ret)
	return __v
}

// NkStylePushFloat function as declared in nk/nuklear.h:1010
func NkStylePushFloat(arg0 *Context, arg1 *float32, arg2 float32) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (*C.float)(unsafe.Pointer(arg1)), cgoAllocsUnknown
	carg2, _ := (C.float)(arg2), cgoAllocsUnknown
	__ret := C.nk_style_push_float(carg0, carg1, carg2)
	__v := (int32)(__ret)
	return __v
}

// NkStylePushVec2 function as declared in nk/nuklear.h:1011
func NkStylePushVec2(arg0 *Context, arg1 *Vec2, arg2 Vec2) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (*C.struct_nk_vec2)(unsafe.Pointer(arg1)), cgoAllocsUnknown
	carg2, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&arg2)), cgoAllocsUnknown
	__ret := C.nk_style_push_vec2(carg0, carg1, carg2)
	__v := (int32)(__ret)
	return __v
}

// NkStylePushStyleItem function as declared in nk/nuklear.h:1012
func NkStylePushStyleItem(arg0 *Context, arg1 *StyleItem, arg2 StyleItem) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (*C.struct_nk_style_item)(unsafe.Pointer(arg1)), cgoAllocsUnknown
	carg2, _ := *(*C.struct_nk_style_item)(unsafe.Pointer(&arg2)), cgoAllocsUnknown
	__ret := C.nk_style_push_style_item(carg0, carg1, carg2)
	__v := (int32)(__ret)
	return __v
}

// NkStylePushFlags function as declared in nk/nuklear.h:1013
func NkStylePushFlags(arg0 *Context, arg1 *Flags, arg2 Flags) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (*C.nk_flags)(unsafe.Pointer(arg1)), cgoAllocsUnknown
	carg2, _ := (C.nk_flags)(arg2), cgoAllocsUnknown
	__ret := C.nk_style_push_flags(carg0, carg1, carg2)
	__v := (int32)(__ret)
	return __v
}

// NkStylePushColor function as declared in nk/nuklear.h:1014
func NkStylePushColor(arg0 *Context, arg1 *Color, arg2 Color) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (*C.struct_nk_color)(unsafe.Pointer(arg1)), cgoAllocsUnknown
	carg2, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg2)), cgoAllocsUnknown
	__ret := C.nk_style_push_color(carg0, carg1, carg2)
	__v := (int32)(__ret)
	return __v
}

// NkStylePopFont function as declared in nk/nuklear.h:1016
func NkStylePopFont(arg0 *Context) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	__ret := C.nk_style_pop_font(carg0)
	__v := (int32)(__ret)
	return __v
}

// NkStylePopFloat function as declared in nk/nuklear.h:1017
func NkStylePopFloat(arg0 *Context) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	__ret := C.nk_style_pop_float(carg0)
	__v := (int32)(__ret)
	return __v
}

// NkStylePopVec2 function as declared in nk/nuklear.h:1018
func NkStylePopVec2(arg0 *Context) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	__ret := C.nk_style_pop_vec2(carg0)
	__v := (int32)(__ret)
	return __v
}

// NkStylePopStyleItem function as declared in nk/nuklear.h:1019
func NkStylePopStyleItem(arg0 *Context) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	__ret := C.nk_style_pop_style_item(carg0)
	__v := (int32)(__ret)
	return __v
}

// NkStylePopFlags function as declared in nk/nuklear.h:1020
func NkStylePopFlags(arg0 *Context) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	__ret := C.nk_style_pop_flags(carg0)
	__v := (int32)(__ret)
	return __v
}

// NkStylePopColor function as declared in nk/nuklear.h:1021
func NkStylePopColor(arg0 *Context) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	__ret := C.nk_style_pop_color(carg0)
	__v := (int32)(__ret)
	return __v
}

// NkWidgetBounds function as declared in nk/nuklear.h:1024
func NkWidgetBounds(arg0 *Context) Rect {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	__ret := C.nk_widget_bounds(carg0)
	__v := *(*Rect)(unsafe.Pointer(&__ret))
	return __v
}

// NkWidgetPosition function as declared in nk/nuklear.h:1025
func NkWidgetPosition(arg0 *Context) Vec2 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	__ret := C.nk_widget_position(carg0)
	__v := *(*Vec2)(unsafe.Pointer(&__ret))
	return __v
}

// NkWidgetSize function as declared in nk/nuklear.h:1026
func NkWidgetSize(arg0 *Context) Vec2 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	__ret := C.nk_widget_size(carg0)
	__v := *(*Vec2)(unsafe.Pointer(&__ret))
	return __v
}

// NkWidgetWidth function as declared in nk/nuklear.h:1027
func NkWidgetWidth(arg0 *Context) float32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	__ret := C.nk_widget_width(carg0)
	__v := (float32)(__ret)
	return __v
}

// NkWidgetHeight function as declared in nk/nuklear.h:1028
func NkWidgetHeight(arg0 *Context) float32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	__ret := C.nk_widget_height(carg0)
	__v := (float32)(__ret)
	return __v
}

// NkWidgetIsHovered function as declared in nk/nuklear.h:1029
func NkWidgetIsHovered(arg0 *Context) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	__ret := C.nk_widget_is_hovered(carg0)
	__v := (int32)(__ret)
	return __v
}

// NkWidgetIsMouseClicked function as declared in nk/nuklear.h:1030
func NkWidgetIsMouseClicked(arg0 *Context, arg1 Buttons) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.enum_nk_buttons)(arg1), cgoAllocsUnknown
	__ret := C.nk_widget_is_mouse_clicked(carg0, carg1)
	__v := (int32)(__ret)
	return __v
}

// NkWidgetHasMouseClickDown function as declared in nk/nuklear.h:1031
func NkWidgetHasMouseClickDown(arg0 *Context, arg1 Buttons, down int32) int32 {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.enum_nk_buttons)(arg1), cgoAllocsUnknown
	cdown, _ := (C.int)(down), cgoAllocsUnknown
	__ret := C.nk_widget_has_mouse_click_down(carg0, carg1, cdown)
	__v := (int32)(__ret)
	return __v
}

// NkSpacing function as declared in nk/nuklear.h:1032
func NkSpacing(arg0 *Context, cols int32) {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	ccols, _ := (C.int)(cols), cgoAllocsUnknown
	C.nk_spacing(carg0, ccols)
}

// NkWidget function as declared in nk/nuklear.h:1035
func NkWidget(arg0 *Rect, arg1 *Context) WidgetLayoutStates {
	carg0, _ := (*C.struct_nk_rect)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (*C.struct_nk_context)(unsafe.Pointer(arg1)), cgoAllocsUnknown
	__ret := C.nk_widget(carg0, carg1)
	__v := (WidgetLayoutStates)(__ret)
	return __v
}

// NkWidgetFitting function as declared in nk/nuklear.h:1036
func NkWidgetFitting(arg0 *Rect, arg1 *Context, arg2 Vec2) WidgetLayoutStates {
	carg0, _ := (*C.struct_nk_rect)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (*C.struct_nk_context)(unsafe.Pointer(arg1)), cgoAllocsUnknown
	carg2, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&arg2)), cgoAllocsUnknown
	__ret := C.nk_widget_fitting(carg0, carg1, carg2)
	__v := (WidgetLayoutStates)(__ret)
	return __v
}

// NkRgb function as declared in nk/nuklear.h:1039
func NkRgb(r int32, g int32, b int32) Color {
	cr, _ := (C.int)(r), cgoAllocsUnknown
	cg, _ := (C.int)(g), cgoAllocsUnknown
	cb, _ := (C.int)(b), cgoAllocsUnknown
	__ret := C.nk_rgb(cr, cg, cb)
	__v := *(*Color)(unsafe.Pointer(&__ret))
	return __v
}

// NkRgbIv function as declared in nk/nuklear.h:1040
func NkRgbIv(rgb *int32) Color {
	crgb, _ := (*C.int)(unsafe.Pointer(rgb)), cgoAllocsUnknown
	__ret := C.nk_rgb_iv(crgb)
	__v := *(*Color)(unsafe.Pointer(&__ret))
	return __v
}

// NkRgbBv function as declared in nk/nuklear.h:1041
func NkRgbBv(rgb string) Color {
	crgb, _ := unpackPByteString(rgb)
	__ret := C.nk_rgb_bv(crgb)
	__v := *(*Color)(unsafe.Pointer(&__ret))
	return __v
}

// NkRgbF function as declared in nk/nuklear.h:1042
func NkRgbF(r float32, g float32, b float32) Color {
	cr, _ := (C.float)(r), cgoAllocsUnknown
	cg, _ := (C.float)(g), cgoAllocsUnknown
	cb, _ := (C.float)(b), cgoAllocsUnknown
	__ret := C.nk_rgb_f(cr, cg, cb)
	__v := *(*Color)(unsafe.Pointer(&__ret))
	return __v
}

// NkRgbFv function as declared in nk/nuklear.h:1043
func NkRgbFv(rgb *float32) Color {
	crgb, _ := (*C.float)(unsafe.Pointer(rgb)), cgoAllocsUnknown
	__ret := C.nk_rgb_fv(crgb)
	__v := *(*Color)(unsafe.Pointer(&__ret))
	return __v
}

// NkRgbHex function as declared in nk/nuklear.h:1044
func NkRgbHex(rgb string) Color {
	crgb, _ := unpackPCharString(rgb)
	__ret := C.nk_rgb_hex(crgb)
	__v := *(*Color)(unsafe.Pointer(&__ret))
	return __v
}

// NkRgba function as declared in nk/nuklear.h:1046
func NkRgba(r int32, g int32, b int32, a int32) Color {
	cr, _ := (C.int)(r), cgoAllocsUnknown
	cg, _ := (C.int)(g), cgoAllocsUnknown
	cb, _ := (C.int)(b), cgoAllocsUnknown
	ca, _ := (C.int)(a), cgoAllocsUnknown
	__ret := C.nk_rgba(cr, cg, cb, ca)
	__v := *(*Color)(unsafe.Pointer(&__ret))
	return __v
}

// NkRgbaU32 function as declared in nk/nuklear.h:1047
func NkRgbaU32(arg0 Uint) Color {
	carg0, _ := (C.nk_uint)(arg0), cgoAllocsUnknown
	__ret := C.nk_rgba_u32(carg0)
	__v := *(*Color)(unsafe.Pointer(&__ret))
	return __v
}

// NkRgbaIv function as declared in nk/nuklear.h:1048
func NkRgbaIv(rgba *int32) Color {
	crgba, _ := (*C.int)(unsafe.Pointer(rgba)), cgoAllocsUnknown
	__ret := C.nk_rgba_iv(crgba)
	__v := *(*Color)(unsafe.Pointer(&__ret))
	return __v
}

// NkRgbaBv function as declared in nk/nuklear.h:1049
func NkRgbaBv(rgba string) Color {
	crgba, _ := unpackPByteString(rgba)
	__ret := C.nk_rgba_bv(crgba)
	__v := *(*Color)(unsafe.Pointer(&__ret))
	return __v
}

// NkRgbaF function as declared in nk/nuklear.h:1050
func NkRgbaF(r float32, g float32, b float32, a float32) Color {
	cr, _ := (C.float)(r), cgoAllocsUnknown
	cg, _ := (C.float)(g), cgoAllocsUnknown
	cb, _ := (C.float)(b), cgoAllocsUnknown
	ca, _ := (C.float)(a), cgoAllocsUnknown
	__ret := C.nk_rgba_f(cr, cg, cb, ca)
	__v := *(*Color)(unsafe.Pointer(&__ret))
	return __v
}

// NkRgbaFv function as declared in nk/nuklear.h:1051
func NkRgbaFv(rgba *float32) Color {
	crgba, _ := (*C.float)(unsafe.Pointer(rgba)), cgoAllocsUnknown
	__ret := C.nk_rgba_fv(crgba)
	__v := *(*Color)(unsafe.Pointer(&__ret))
	return __v
}

// NkRgbaHex function as declared in nk/nuklear.h:1052
func NkRgbaHex(rgb string) Color {
	crgb, _ := unpackPCharString(rgb)
	__ret := C.nk_rgba_hex(crgb)
	__v := *(*Color)(unsafe.Pointer(&__ret))
	return __v
}

// NkHsv function as declared in nk/nuklear.h:1054
func NkHsv(h int32, s int32, v int32) Color {
	ch, _ := (C.int)(h), cgoAllocsUnknown
	cs, _ := (C.int)(s), cgoAllocsUnknown
	cv, _ := (C.int)(v), cgoAllocsUnknown
	__ret := C.nk_hsv(ch, cs, cv)
	__v := *(*Color)(unsafe.Pointer(&__ret))
	return __v
}

// NkHsvIv function as declared in nk/nuklear.h:1055
func NkHsvIv(hsv *int32) Color {
	chsv, _ := (*C.int)(unsafe.Pointer(hsv)), cgoAllocsUnknown
	__ret := C.nk_hsv_iv(chsv)
	__v := *(*Color)(unsafe.Pointer(&__ret))
	return __v
}

// NkHsvBv function as declared in nk/nuklear.h:1056
func NkHsvBv(hsv string) Color {
	chsv, _ := unpackPByteString(hsv)
	__ret := C.nk_hsv_bv(chsv)
	__v := *(*Color)(unsafe.Pointer(&__ret))
	return __v
}

// NkHsvF function as declared in nk/nuklear.h:1057
func NkHsvF(h float32, s float32, v float32) Color {
	ch, _ := (C.float)(h), cgoAllocsUnknown
	cs, _ := (C.float)(s), cgoAllocsUnknown
	cv, _ := (C.float)(v), cgoAllocsUnknown
	__ret := C.nk_hsv_f(ch, cs, cv)
	__v := *(*Color)(unsafe.Pointer(&__ret))
	return __v
}

// NkHsvFv function as declared in nk/nuklear.h:1058
func NkHsvFv(hsv *float32) Color {
	chsv, _ := (*C.float)(unsafe.Pointer(hsv)), cgoAllocsUnknown
	__ret := C.nk_hsv_fv(chsv)
	__v := *(*Color)(unsafe.Pointer(&__ret))
	return __v
}

// NkHsva function as declared in nk/nuklear.h:1060
func NkHsva(h int32, s int32, v int32, a int32) Color {
	ch, _ := (C.int)(h), cgoAllocsUnknown
	cs, _ := (C.int)(s), cgoAllocsUnknown
	cv, _ := (C.int)(v), cgoAllocsUnknown
	ca, _ := (C.int)(a), cgoAllocsUnknown
	__ret := C.nk_hsva(ch, cs, cv, ca)
	__v := *(*Color)(unsafe.Pointer(&__ret))
	return __v
}

// NkHsvaIv function as declared in nk/nuklear.h:1061
func NkHsvaIv(hsva *int32) Color {
	chsva, _ := (*C.int)(unsafe.Pointer(hsva)), cgoAllocsUnknown
	__ret := C.nk_hsva_iv(chsva)
	__v := *(*Color)(unsafe.Pointer(&__ret))
	return __v
}

// NkHsvaBv function as declared in nk/nuklear.h:1062
func NkHsvaBv(hsva string) Color {
	chsva, _ := unpackPByteString(hsva)
	__ret := C.nk_hsva_bv(chsva)
	__v := *(*Color)(unsafe.Pointer(&__ret))
	return __v
}

// NkHsvaF function as declared in nk/nuklear.h:1063
func NkHsvaF(h float32, s float32, v float32, a float32) Color {
	ch, _ := (C.float)(h), cgoAllocsUnknown
	cs, _ := (C.float)(s), cgoAllocsUnknown
	cv, _ := (C.float)(v), cgoAllocsUnknown
	ca, _ := (C.float)(a), cgoAllocsUnknown
	__ret := C.nk_hsva_f(ch, cs, cv, ca)
	__v := *(*Color)(unsafe.Pointer(&__ret))
	return __v
}

// NkHsvaFv function as declared in nk/nuklear.h:1064
func NkHsvaFv(hsva *float32) Color {
	chsva, _ := (*C.float)(unsafe.Pointer(hsva)), cgoAllocsUnknown
	__ret := C.nk_hsva_fv(chsva)
	__v := *(*Color)(unsafe.Pointer(&__ret))
	return __v
}

// NkColorF function as declared in nk/nuklear.h:1067
func NkColorF(r *float32, g *float32, b *float32, a *float32, arg4 Color) {
	cr, _ := (*C.float)(unsafe.Pointer(r)), cgoAllocsUnknown
	cg, _ := (*C.float)(unsafe.Pointer(g)), cgoAllocsUnknown
	cb, _ := (*C.float)(unsafe.Pointer(b)), cgoAllocsUnknown
	ca, _ := (*C.float)(unsafe.Pointer(a)), cgoAllocsUnknown
	carg4, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg4)), cgoAllocsUnknown
	C.nk_color_f(cr, cg, cb, ca, carg4)
}

// NkColorFv function as declared in nk/nuklear.h:1068
func NkColorFv(rgbaOut []float32, arg1 Color) {
	crgbaOut, _ := (*C.float)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&rgbaOut)).Data)), cgoAllocsUnknown
	carg1, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg1)), cgoAllocsUnknown
	C.nk_color_fv(crgbaOut, carg1)
}

// NkColorD function as declared in nk/nuklear.h:1069
func NkColorD(r *float64, g *float64, b *float64, a *float64, arg4 Color) {
	cr, _ := (*C.double)(unsafe.Pointer(r)), cgoAllocsUnknown
	cg, _ := (*C.double)(unsafe.Pointer(g)), cgoAllocsUnknown
	cb, _ := (*C.double)(unsafe.Pointer(b)), cgoAllocsUnknown
	ca, _ := (*C.double)(unsafe.Pointer(a)), cgoAllocsUnknown
	carg4, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg4)), cgoAllocsUnknown
	C.nk_color_d(cr, cg, cb, ca, carg4)
}

// NkColorDv function as declared in nk/nuklear.h:1070
func NkColorDv(rgbaOut []float64, arg1 Color) {
	crgbaOut, _ := (*C.double)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&rgbaOut)).Data)), cgoAllocsUnknown
	carg1, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg1)), cgoAllocsUnknown
	C.nk_color_dv(crgbaOut, carg1)
}

// NkColorU32 function as declared in nk/nuklear.h:1072
func NkColorU32(arg0 Color) Uint {
	carg0, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg0)), cgoAllocsUnknown
	__ret := C.nk_color_u32(carg0)
	__v := (Uint)(__ret)
	return __v
}

// NkColorHexRgba function as declared in nk/nuklear.h:1073
func NkColorHexRgba(output []byte, arg1 Color) {
	coutput, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&output)).Data)), cgoAllocsUnknown
	carg1, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg1)), cgoAllocsUnknown
	C.nk_color_hex_rgba(coutput, carg1)
}

// NkColorHexRgb function as declared in nk/nuklear.h:1074
func NkColorHexRgb(output []byte, arg1 Color) {
	coutput, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&output)).Data)), cgoAllocsUnknown
	carg1, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg1)), cgoAllocsUnknown
	C.nk_color_hex_rgb(coutput, carg1)
}

// NkColorHsvI function as declared in nk/nuklear.h:1076
func NkColorHsvI(outH *int32, outS *int32, outV *int32, arg3 Color) {
	coutH, _ := (*C.int)(unsafe.Pointer(outH)), cgoAllocsUnknown
	coutS, _ := (*C.int)(unsafe.Pointer(outS)), cgoAllocsUnknown
	coutV, _ := (*C.int)(unsafe.Pointer(outV)), cgoAllocsUnknown
	carg3, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg3)), cgoAllocsUnknown
	C.nk_color_hsv_i(coutH, coutS, coutV, carg3)
}

// NkColorHsvB function as declared in nk/nuklear.h:1077
func NkColorHsvB(outH *byte, outS *byte, outV *byte, arg3 Color) {
	coutH, _ := (*C.nk_byte)(unsafe.Pointer(outH)), cgoAllocsUnknown
	coutS, _ := (*C.nk_byte)(unsafe.Pointer(outS)), cgoAllocsUnknown
	coutV, _ := (*C.nk_byte)(unsafe.Pointer(outV)), cgoAllocsUnknown
	carg3, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg3)), cgoAllocsUnknown
	C.nk_color_hsv_b(coutH, coutS, coutV, carg3)
}

// NkColorHsvIv function as declared in nk/nuklear.h:1078
func NkColorHsvIv(hsvOut []int32, arg1 Color) {
	chsvOut, _ := (*C.int)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&hsvOut)).Data)), cgoAllocsUnknown
	carg1, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg1)), cgoAllocsUnknown
	C.nk_color_hsv_iv(chsvOut, carg1)
}

// NkColorHsvBv function as declared in nk/nuklear.h:1079
func NkColorHsvBv(hsvOut []byte, arg1 Color) {
	chsvOut, _ := (*C.nk_byte)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&hsvOut)).Data)), cgoAllocsUnknown
	carg1, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg1)), cgoAllocsUnknown
	C.nk_color_hsv_bv(chsvOut, carg1)
}

// NkColorHsvF function as declared in nk/nuklear.h:1080
func NkColorHsvF(outH *float32, outS *float32, outV *float32, arg3 Color) {
	coutH, _ := (*C.float)(unsafe.Pointer(outH)), cgoAllocsUnknown
	coutS, _ := (*C.float)(unsafe.Pointer(outS)), cgoAllocsUnknown
	coutV, _ := (*C.float)(unsafe.Pointer(outV)), cgoAllocsUnknown
	carg3, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg3)), cgoAllocsUnknown
	C.nk_color_hsv_f(coutH, coutS, coutV, carg3)
}

// NkColorHsvFv function as declared in nk/nuklear.h:1081
func NkColorHsvFv(hsvOut []float32, arg1 Color) {
	chsvOut, _ := (*C.float)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&hsvOut)).Data)), cgoAllocsUnknown
	carg1, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg1)), cgoAllocsUnknown
	C.nk_color_hsv_fv(chsvOut, carg1)
}

// NkColorHsvaI function as declared in nk/nuklear.h:1083
func NkColorHsvaI(h *int32, s *int32, v *int32, a *int32, arg4 Color) {
	ch, _ := (*C.int)(unsafe.Pointer(h)), cgoAllocsUnknown
	cs, _ := (*C.int)(unsafe.Pointer(s)), cgoAllocsUnknown
	cv, _ := (*C.int)(unsafe.Pointer(v)), cgoAllocsUnknown
	ca, _ := (*C.int)(unsafe.Pointer(a)), cgoAllocsUnknown
	carg4, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg4)), cgoAllocsUnknown
	C.nk_color_hsva_i(ch, cs, cv, ca, carg4)
}

// NkColorHsvaB function as declared in nk/nuklear.h:1084
func NkColorHsvaB(h *byte, s *byte, v *byte, a *Byte, arg4 Color) {
	ch, _ := (*C.nk_byte)(unsafe.Pointer(h)), cgoAllocsUnknown
	cs, _ := (*C.nk_byte)(unsafe.Pointer(s)), cgoAllocsUnknown
	cv, _ := (*C.nk_byte)(unsafe.Pointer(v)), cgoAllocsUnknown
	ca, _ := (*C.nk_byte)(unsafe.Pointer(a)), cgoAllocsUnknown
	carg4, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg4)), cgoAllocsUnknown
	C.nk_color_hsva_b(ch, cs, cv, ca, carg4)
}

// NkColorHsvaIv function as declared in nk/nuklear.h:1085
func NkColorHsvaIv(hsvaOut []int32, arg1 Color) {
	chsvaOut, _ := (*C.int)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&hsvaOut)).Data)), cgoAllocsUnknown
	carg1, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg1)), cgoAllocsUnknown
	C.nk_color_hsva_iv(chsvaOut, carg1)
}

// NkColorHsvaBv function as declared in nk/nuklear.h:1086
func NkColorHsvaBv(hsvaOut []byte, arg1 Color) {
	chsvaOut, _ := (*C.nk_byte)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&hsvaOut)).Data)), cgoAllocsUnknown
	carg1, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg1)), cgoAllocsUnknown
	C.nk_color_hsva_bv(chsvaOut, carg1)
}

// NkColorHsvaF function as declared in nk/nuklear.h:1087
func NkColorHsvaF(outH *float32, outS *float32, outV *float32, outA *float32, arg4 Color) {
	coutH, _ := (*C.float)(unsafe.Pointer(outH)), cgoAllocsUnknown
	coutS, _ := (*C.float)(unsafe.Pointer(outS)), cgoAllocsUnknown
	coutV, _ := (*C.float)(unsafe.Pointer(outV)), cgoAllocsUnknown
	coutA, _ := (*C.float)(unsafe.Pointer(outA)), cgoAllocsUnknown
	carg4, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg4)), cgoAllocsUnknown
	C.nk_color_hsva_f(coutH, coutS, coutV, coutA, carg4)
}

// NkColorHsvaFv function as declared in nk/nuklear.h:1088
func NkColorHsvaFv(hsvaOut []float32, arg1 Color) {
	chsvaOut, _ := (*C.float)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&hsvaOut)).Data)), cgoAllocsUnknown
	carg1, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg1)), cgoAllocsUnknown
	C.nk_color_hsva_fv(chsvaOut, carg1)
}

// NkHandlePtr function as declared in nk/nuklear.h:1091
func NkHandlePtr(arg0 unsafe.Pointer) Handle {
	carg0, _ := (unsafe.Pointer)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	__ret := C.nk_handle_ptr(carg0)
	__v := *(*Handle)(unsafe.Pointer(&__ret))
	return __v
}

// NkHandleId function as declared in nk/nuklear.h:1092
func NkHandleId(arg0 int32) Handle {
	carg0, _ := (C.int)(arg0), cgoAllocsUnknown
	__ret := C.nk_handle_id(carg0)
	__v := *(*Handle)(unsafe.Pointer(&__ret))
	return __v
}

// NkImageHandle function as declared in nk/nuklear.h:1093
func NkImageHandle(arg0 Handle) Image {
	carg0, _ := *(*C.nk_handle)(unsafe.Pointer(&arg0)), cgoAllocsUnknown
	__ret := C.nk_image_handle(carg0)
	__v := *(*Image)(unsafe.Pointer(&__ret))
	return __v
}

// NkImagePtr function as declared in nk/nuklear.h:1094
func NkImagePtr(arg0 unsafe.Pointer) Image {
	carg0, _ := (unsafe.Pointer)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	__ret := C.nk_image_ptr(carg0)
	__v := *(*Image)(unsafe.Pointer(&__ret))
	return __v
}

// NkImageId function as declared in nk/nuklear.h:1095
func NkImageId(arg0 int32) Image {
	carg0, _ := (C.int)(arg0), cgoAllocsUnknown
	__ret := C.nk_image_id(carg0)
	__v := *(*Image)(unsafe.Pointer(&__ret))
	return __v
}

// NkImageIsSubimage function as declared in nk/nuklear.h:1096
func NkImageIsSubimage(img *Image) int32 {
	cimg, _ := (*C.struct_nk_image)(unsafe.Pointer(img)), cgoAllocsUnknown
	__ret := C.nk_image_is_subimage(cimg)
	__v := (int32)(__ret)
	return __v
}

// NkSubimagePtr function as declared in nk/nuklear.h:1097
func NkSubimagePtr(arg0 unsafe.Pointer, w uint16, h uint16, subRegion Rect) Image {
	carg0, _ := (unsafe.Pointer)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cw, _ := (C.ushort)(w), cgoAllocsUnknown
	ch, _ := (C.ushort)(h), cgoAllocsUnknown
	csubRegion, _ := *(*C.struct_nk_rect)(unsafe.Pointer(&subRegion)), cgoAllocsUnknown
	__ret := C.nk_subimage_ptr(carg0, cw, ch, csubRegion)
	__v := *(*Image)(unsafe.Pointer(&__ret))
	return __v
}

// NkSubimageId function as declared in nk/nuklear.h:1098
func NkSubimageId(arg0 int32, w uint16, h uint16, subRegion Rect) Image {
	carg0, _ := (C.int)(arg0), cgoAllocsUnknown
	cw, _ := (C.ushort)(w), cgoAllocsUnknown
	ch, _ := (C.ushort)(h), cgoAllocsUnknown
	csubRegion, _ := *(*C.struct_nk_rect)(unsafe.Pointer(&subRegion)), cgoAllocsUnknown
	__ret := C.nk_subimage_id(carg0, cw, ch, csubRegion)
	__v := *(*Image)(unsafe.Pointer(&__ret))
	return __v
}

// NkSubimageHandle function as declared in nk/nuklear.h:1099
func NkSubimageHandle(arg0 Handle, w uint16, h uint16, subRegion Rect) Image {
	carg0, _ := *(*C.nk_handle)(unsafe.Pointer(&arg0)), cgoAllocsUnknown
	cw, _ := (C.ushort)(w), cgoAllocsUnknown
	ch, _ := (C.ushort)(h), cgoAllocsUnknown
	csubRegion, _ := *(*C.struct_nk_rect)(unsafe.Pointer(&subRegion)), cgoAllocsUnknown
	__ret := C.nk_subimage_handle(carg0, cw, ch, csubRegion)
	__v := *(*Image)(unsafe.Pointer(&__ret))
	return __v
}

// NkMurmurHash function as declared in nk/nuklear.h:1102
func NkMurmurHash(key unsafe.Pointer, len int32, seed Hash) Hash {
	ckey, _ := (unsafe.Pointer)(unsafe.Pointer(key)), cgoAllocsUnknown
	clen, _ := (C.int)(len), cgoAllocsUnknown
	cseed, _ := (C.nk_hash)(seed), cgoAllocsUnknown
	__ret := C.nk_murmur_hash(ckey, clen, cseed)
	__v := (Hash)(__ret)
	return __v
}

// NkTriangleFromDirection function as declared in nk/nuklear.h:1103
func NkTriangleFromDirection(result *Vec2, r Rect, padX float32, padY float32, arg4 Heading) {
	cresult, _ := (*C.struct_nk_vec2)(unsafe.Pointer(result)), cgoAllocsUnknown
	cr, _ := *(*C.struct_nk_rect)(unsafe.Pointer(&r)), cgoAllocsUnknown
	cpadX, _ := (C.float)(padX), cgoAllocsUnknown
	cpadY, _ := (C.float)(padY), cgoAllocsUnknown
	carg4, _ := (C.enum_nk_heading)(arg4), cgoAllocsUnknown
	C.nk_triangle_from_direction(cresult, cr, cpadX, cpadY, carg4)
}

// NkVec2 function as declared in nk/nuklear.h:1105
func NkVec2(x float32, y float32) Vec2 {
	cx, _ := (C.float)(x), cgoAllocsUnknown
	cy, _ := (C.float)(y), cgoAllocsUnknown
	__ret := C.nk_vec2(cx, cy)
	__v := *(*Vec2)(unsafe.Pointer(&__ret))
	return __v
}

// NkVec2i function as declared in nk/nuklear.h:1106
func NkVec2i(x int32, y int32) Vec2 {
	cx, _ := (C.int)(x), cgoAllocsUnknown
	cy, _ := (C.int)(y), cgoAllocsUnknown
	__ret := C.nk_vec2i(cx, cy)
	__v := *(*Vec2)(unsafe.Pointer(&__ret))
	return __v
}

// NkVec2v function as declared in nk/nuklear.h:1107
func NkVec2v(xy *float32) Vec2 {
	cxy, _ := (*C.float)(unsafe.Pointer(xy)), cgoAllocsUnknown
	__ret := C.nk_vec2v(cxy)
	__v := *(*Vec2)(unsafe.Pointer(&__ret))
	return __v
}

// NkVec2iv function as declared in nk/nuklear.h:1108
func NkVec2iv(xy *int32) Vec2 {
	cxy, _ := (*C.int)(unsafe.Pointer(xy)), cgoAllocsUnknown
	__ret := C.nk_vec2iv(cxy)
	__v := *(*Vec2)(unsafe.Pointer(&__ret))
	return __v
}

// NkGetNullRect function as declared in nk/nuklear.h:1110
func NkGetNullRect() Rect {
	__ret := C.nk_get_null_rect()
	__v := *(*Rect)(unsafe.Pointer(&__ret))
	return __v
}

// NkRect function as declared in nk/nuklear.h:1111
func NkRect(x float32, y float32, w float32, h float32) Rect {
	cx, _ := (C.float)(x), cgoAllocsUnknown
	cy, _ := (C.float)(y), cgoAllocsUnknown
	cw, _ := (C.float)(w), cgoAllocsUnknown
	ch, _ := (C.float)(h), cgoAllocsUnknown
	__ret := C.nk_rect(cx, cy, cw, ch)
	__v := *(*Rect)(unsafe.Pointer(&__ret))
	return __v
}

// NkRecti function as declared in nk/nuklear.h:1112
func NkRecti(x int32, y int32, w int32, h int32) Rect {
	cx, _ := (C.int)(x), cgoAllocsUnknown
	cy, _ := (C.int)(y), cgoAllocsUnknown
	cw, _ := (C.int)(w), cgoAllocsUnknown
	ch, _ := (C.int)(h), cgoAllocsUnknown
	__ret := C.nk_recti(cx, cy, cw, ch)
	__v := *(*Rect)(unsafe.Pointer(&__ret))
	return __v
}

// NkRecta function as declared in nk/nuklear.h:1113
func NkRecta(pos Vec2, size Vec2) Rect {
	cpos, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&pos)), cgoAllocsUnknown
	csize, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&size)), cgoAllocsUnknown
	__ret := C.nk_recta(cpos, csize)
	__v := *(*Rect)(unsafe.Pointer(&__ret))
	return __v
}

// NkRectv function as declared in nk/nuklear.h:1114
func NkRectv(xywh *float32) Rect {
	cxywh, _ := (*C.float)(unsafe.Pointer(xywh)), cgoAllocsUnknown
	__ret := C.nk_rectv(cxywh)
	__v := *(*Rect)(unsafe.Pointer(&__ret))
	return __v
}

// NkRectiv function as declared in nk/nuklear.h:1115
func NkRectiv(xywh *int32) Rect {
	cxywh, _ := (*C.int)(unsafe.Pointer(xywh)), cgoAllocsUnknown
	__ret := C.nk_rectiv(cxywh)
	__v := *(*Rect)(unsafe.Pointer(&__ret))
	return __v
}

// NkRectPos function as declared in nk/nuklear.h:1116
func NkRectPos(arg0 Rect) Vec2 {
	carg0, _ := *(*C.struct_nk_rect)(unsafe.Pointer(&arg0)), cgoAllocsUnknown
	__ret := C.nk_rect_pos(carg0)
	__v := *(*Vec2)(unsafe.Pointer(&__ret))
	return __v
}

// NkRectSize function as declared in nk/nuklear.h:1117
func NkRectSize(arg0 Rect) Vec2 {
	carg0, _ := *(*C.struct_nk_rect)(unsafe.Pointer(&arg0)), cgoAllocsUnknown
	__ret := C.nk_rect_size(carg0)
	__v := *(*Vec2)(unsafe.Pointer(&__ret))
	return __v
}

// NkStrlen function as declared in nk/nuklear.h:1120
func NkStrlen(str string) int32 {
	cstr, _ := unpackPCharString(str)
	__ret := C.nk_strlen(cstr)
	__v := (int32)(__ret)
	return __v
}

// NkStricmp function as declared in nk/nuklear.h:1121
func NkStricmp(s1 string, s2 string) int32 {
	cs1, _ := unpackPCharString(s1)
	cs2, _ := unpackPCharString(s2)
	__ret := C.nk_stricmp(cs1, cs2)
	__v := (int32)(__ret)
	return __v
}

// NkStricmpn function as declared in nk/nuklear.h:1122
func NkStricmpn(s1 string, s2 string, n int32) int32 {
	cs1, _ := unpackPCharString(s1)
	cs2, _ := unpackPCharString(s2)
	cn, _ := (C.int)(n), cgoAllocsUnknown
	__ret := C.nk_stricmpn(cs1, cs2, cn)
	__v := (int32)(__ret)
	return __v
}

// NkStrtoi function as declared in nk/nuklear.h:1123
func NkStrtoi(str string, endptr []*byte) int32 {
	cstr, _ := unpackPCharString(str)
	cendptr, _ := (**C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&endptr)).Data)), cgoAllocsUnknown
	__ret := C.nk_strtoi(cstr, cendptr)
	__v := (int32)(__ret)
	return __v
}

// NkStrtof function as declared in nk/nuklear.h:1124
func NkStrtof(str string, endptr []*byte) float32 {
	cstr, _ := unpackPCharString(str)
	cendptr, _ := (**C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&endptr)).Data)), cgoAllocsUnknown
	__ret := C.nk_strtof(cstr, cendptr)
	__v := (float32)(__ret)
	return __v
}

// NkStrtod function as declared in nk/nuklear.h:1125
func NkStrtod(str string, endptr []*byte) float64 {
	cstr, _ := unpackPCharString(str)
	cendptr, _ := (**C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&endptr)).Data)), cgoAllocsUnknown
	__ret := C.nk_strtod(cstr, cendptr)
	__v := (float64)(__ret)
	return __v
}

// NkStrfilter function as declared in nk/nuklear.h:1126
func NkStrfilter(text string, regexp string) int32 {
	ctext, _ := unpackPCharString(text)
	cregexp, _ := unpackPCharString(regexp)
	__ret := C.nk_strfilter(ctext, cregexp)
	__v := (int32)(__ret)
	return __v
}

// NkStrmatchFuzzyString function as declared in nk/nuklear.h:1127
func NkStrmatchFuzzyString(str string, pattern string, outScore *int32) int32 {
	cstr, _ := unpackPCharString(str)
	cpattern, _ := unpackPCharString(pattern)
	coutScore, _ := (*C.int)(unsafe.Pointer(outScore)), cgoAllocsUnknown
	__ret := C.nk_strmatch_fuzzy_string(cstr, cpattern, coutScore)
	__v := (int32)(__ret)
	return __v
}

// NkStrmatchFuzzyText function as declared in nk/nuklear.h:1128
func NkStrmatchFuzzyText(txt string, txtLen int32, pattern string, outScore *int32) int32 {
	ctxt, _ := unpackPCharString(txt)
	ctxtLen, _ := (C.int)(txtLen), cgoAllocsUnknown
	cpattern, _ := unpackPCharString(pattern)
	coutScore, _ := (*C.int)(unsafe.Pointer(outScore)), cgoAllocsUnknown
	__ret := C.nk_strmatch_fuzzy_text(ctxt, ctxtLen, cpattern, coutScore)
	__v := (int32)(__ret)
	return __v
}

// NkUtfDecode function as declared in nk/nuklear.h:1131
func NkUtfDecode(arg0 string, arg1 *Rune, arg2 int32) int32 {
	carg0, _ := unpackPCharString(arg0)
	carg1, _ := (*C.nk_rune)(unsafe.Pointer(arg1)), cgoAllocsUnknown
	carg2, _ := (C.int)(arg2), cgoAllocsUnknown
	__ret := C.nk_utf_decode(carg0, carg1, carg2)
	__v := (int32)(__ret)
	return __v
}

// NkUtfEncode function as declared in nk/nuklear.h:1132
func NkUtfEncode(arg0 Rune, arg1 *byte, arg2 int32) int32 {
	carg0, _ := (C.nk_rune)(arg0), cgoAllocsUnknown
	carg1, _ := (*C.char)(unsafe.Pointer(arg1)), cgoAllocsUnknown
	carg2, _ := (C.int)(arg2), cgoAllocsUnknown
	__ret := C.nk_utf_encode(carg0, carg1, carg2)
	__v := (int32)(__ret)
	return __v
}

// NkUtfLen function as declared in nk/nuklear.h:1133
func NkUtfLen(arg0 string, byteLen int32) int32 {
	carg0, _ := unpackPCharString(arg0)
	cbyteLen, _ := (C.int)(byteLen), cgoAllocsUnknown
	__ret := C.nk_utf_len(carg0, cbyteLen)
	__v := (int32)(__ret)
	return __v
}

// NkUtfAt function as declared in nk/nuklear.h:1134
func NkUtfAt(buffer string, length int32, index int32, unicode *Rune, len []int32) string {
	cbuffer, _ := unpackPCharString(buffer)
	clength, _ := (C.int)(length), cgoAllocsUnknown
	cindex, _ := (C.int)(index), cgoAllocsUnknown
	cunicode, _ := (*C.nk_rune)(unsafe.Pointer(unicode)), cgoAllocsUnknown
	clen, _ := (*C.int)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&len)).Data)), cgoAllocsUnknown
	__ret := C.nk_utf_at(cbuffer, clength, cindex, cunicode, clen)
	__v := packPCharString(__ret)
	return __v
}

// NkFontDefaultGlyphRanges function as declared in nk/nuklear.h:1416
func NkFontDefaultGlyphRanges() *Rune {
	__ret := C.nk_font_default_glyph_ranges()
	__v := *(**Rune)(unsafe.Pointer(&__ret))
	return __v
}

// NkFontChineseGlyphRanges function as declared in nk/nuklear.h:1417
func NkFontChineseGlyphRanges() *Rune {
	__ret := C.nk_font_chinese_glyph_ranges()
	__v := *(**Rune)(unsafe.Pointer(&__ret))
	return __v
}

// NkFontCyrillicGlyphRanges function as declared in nk/nuklear.h:1418
func NkFontCyrillicGlyphRanges() *Rune {
	__ret := C.nk_font_cyrillic_glyph_ranges()
	__v := *(**Rune)(unsafe.Pointer(&__ret))
	return __v
}

// NkFontKoreanGlyphRanges function as declared in nk/nuklear.h:1419
func NkFontKoreanGlyphRanges() *Rune {
	__ret := C.nk_font_korean_glyph_ranges()
	__v := *(**Rune)(unsafe.Pointer(&__ret))
	return __v
}

// NkFontAtlasInitDefault function as declared in nk/nuklear.h:1422
func NkFontAtlasInitDefault(arg0 *FontAtlas) {
	carg0, _ := (*C.struct_nk_font_atlas)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	C.nk_font_atlas_init_default(carg0)
}

// NkFontAtlasInit function as declared in nk/nuklear.h:1424
func NkFontAtlasInit(arg0 *FontAtlas, arg1 *Allocator) {
	carg0, _ := (*C.struct_nk_font_atlas)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (*C.struct_nk_allocator)(unsafe.Pointer(arg1)), cgoAllocsUnknown
	C.nk_font_atlas_init(carg0, carg1)
}

// NkFontAtlasInitCustom function as declared in nk/nuklear.h:1425
func NkFontAtlasInitCustom(arg0 *FontAtlas, persistent *Allocator, transient *Allocator) {
	carg0, _ := (*C.struct_nk_font_atlas)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cpersistent, _ := (*C.struct_nk_allocator)(unsafe.Pointer(persistent)), cgoAllocsUnknown
	ctransient, _ := (*C.struct_nk_allocator)(unsafe.Pointer(transient)), cgoAllocsUnknown
	C.nk_font_atlas_init_custom(carg0, cpersistent, ctransient)
}

// NkFontAtlasBegin function as declared in nk/nuklear.h:1426
func NkFontAtlasBegin(arg0 *FontAtlas) {
	carg0, _ := (*C.struct_nk_font_atlas)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	C.nk_font_atlas_begin(carg0)
}

// NkFontConfig function as declared in nk/nuklear.h:1427
func NkFontConfig(pixelHeight float32) FontConfig {
	cpixelHeight, _ := (C.float)(pixelHeight), cgoAllocsUnknown
	__ret := C.nk_font_config(cpixelHeight)
	__v := *(*FontConfig)(unsafe.Pointer(&__ret))
	return __v
}

// NkFontAtlasAdd function as declared in nk/nuklear.h:1428
func NkFontAtlasAdd(arg0 *FontAtlas, arg1 *FontConfig) *Font {
	carg0, _ := (*C.struct_nk_font_atlas)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (*C.struct_nk_font_config)(unsafe.Pointer(arg1)), cgoAllocsUnknown
	__ret := C.nk_font_atlas_add(carg0, carg1)
	__v := *(**Font)(unsafe.Pointer(&__ret))
	return __v
}

// NkFontAtlasAddDefault function as declared in nk/nuklear.h:1430
func NkFontAtlasAddDefault(arg0 *FontAtlas, height float32, arg2 *FontConfig) *Font {
	carg0, _ := (*C.struct_nk_font_atlas)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cheight, _ := (C.float)(height), cgoAllocsUnknown
	carg2, _ := (*C.struct_nk_font_config)(unsafe.Pointer(arg2)), cgoAllocsUnknown
	__ret := C.nk_font_atlas_add_default(carg0, cheight, carg2)
	__v := *(**Font)(unsafe.Pointer(&__ret))
	return __v
}

// NkFontAtlasAddFromMemory function as declared in nk/nuklear.h:1432
func NkFontAtlasAddFromMemory(atlas *FontAtlas, memory unsafe.Pointer, size Size, height float32, config *FontConfig) *Font {
	catlas, _ := (*C.struct_nk_font_atlas)(unsafe.Pointer(atlas)), cgoAllocsUnknown
	cmemory, _ := (unsafe.Pointer)(unsafe.Pointer(memory)), cgoAllocsUnknown
	csize, _ := (C.nk_size)(size), cgoAllocsUnknown
	cheight, _ := (C.float)(height), cgoAllocsUnknown
	cconfig, _ := (*C.struct_nk_font_config)(unsafe.Pointer(config)), cgoAllocsUnknown
	__ret := C.nk_font_atlas_add_from_memory(catlas, cmemory, csize, cheight, cconfig)
	__v := *(**Font)(unsafe.Pointer(&__ret))
	return __v
}

// NkFontAtlasAddFromFile function as declared in nk/nuklear.h:1434
func NkFontAtlasAddFromFile(atlas *FontAtlas, filePath string, height float32, arg3 *FontConfig) *Font {
	catlas, _ := (*C.struct_nk_font_atlas)(unsafe.Pointer(atlas)), cgoAllocsUnknown
	cfilePath, _ := unpackPCharString(filePath)
	cheight, _ := (C.float)(height), cgoAllocsUnknown
	carg3, _ := (*C.struct_nk_font_config)(unsafe.Pointer(arg3)), cgoAllocsUnknown
	__ret := C.nk_font_atlas_add_from_file(catlas, cfilePath, cheight, carg3)
	__v := *(**Font)(unsafe.Pointer(&__ret))
	return __v
}

// NkFontAtlasAddCompressed function as declared in nk/nuklear.h:1436
func NkFontAtlasAddCompressed(arg0 *FontAtlas, memory unsafe.Pointer, size Size, height float32, arg4 []FontConfig) *Font {
	carg0, _ := (*C.struct_nk_font_atlas)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cmemory, _ := (unsafe.Pointer)(unsafe.Pointer(memory)), cgoAllocsUnknown
	csize, _ := (C.nk_size)(size), cgoAllocsUnknown
	cheight, _ := (C.float)(height), cgoAllocsUnknown
	carg4, _ := (*C.struct_nk_font_config)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&arg4)).Data)), cgoAllocsUnknown
	__ret := C.nk_font_atlas_add_compressed(carg0, cmemory, csize, cheight, carg4)
	__v := *(**Font)(unsafe.Pointer(&__ret))
	return __v
}

// NkFontAtlasAddCompressedBase85 function as declared in nk/nuklear.h:1437
func NkFontAtlasAddCompressedBase85(arg0 *FontAtlas, data string, height float32, config *FontConfig) *Font {
	carg0, _ := (*C.struct_nk_font_atlas)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cdata, _ := unpackPCharString(data)
	cheight, _ := (C.float)(height), cgoAllocsUnknown
	cconfig, _ := (*C.struct_nk_font_config)(unsafe.Pointer(config)), cgoAllocsUnknown
	__ret := C.nk_font_atlas_add_compressed_base85(carg0, cdata, cheight, cconfig)
	__v := *(**Font)(unsafe.Pointer(&__ret))
	return __v
}

// NkFontAtlasBake function as declared in nk/nuklear.h:1438
func NkFontAtlasBake(arg0 *FontAtlas, width *int32, height *int32, arg3 FontAtlasFormat) unsafe.Pointer {
	carg0, _ := (*C.struct_nk_font_atlas)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cwidth, _ := (*C.int)(unsafe.Pointer(width)), cgoAllocsUnknown
	cheight, _ := (*C.int)(unsafe.Pointer(height)), cgoAllocsUnknown
	carg3, _ := (C.enum_nk_font_atlas_format)(arg3), cgoAllocsUnknown
	__ret := C.nk_font_atlas_bake(carg0, cwidth, cheight, carg3)
	__v := *(*unsafe.Pointer)(unsafe.Pointer(&__ret))
	return __v
}

// NkFontAtlasEnd function as declared in nk/nuklear.h:1439
func NkFontAtlasEnd(arg0 *FontAtlas, tex Handle, arg2 *DrawNullTexture) {
	carg0, _ := (*C.struct_nk_font_atlas)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	ctex, _ := *(*C.nk_handle)(unsafe.Pointer(&tex)), cgoAllocsUnknown
	carg2, _ := (*C.struct_nk_draw_null_texture)(unsafe.Pointer(arg2)), cgoAllocsUnknown
	C.nk_font_atlas_end(carg0, ctex, carg2)
}

// NkFontFindGlyph function as declared in nk/nuklear.h:1440
func NkFontFindGlyph(arg0 *Font, unicode Rune) *FontGlyph {
	carg0, _ := (*C.struct_nk_font)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cunicode, _ := (C.nk_rune)(unicode), cgoAllocsUnknown
	__ret := C.nk_font_find_glyph(carg0, cunicode)
	__v := *(**FontGlyph)(unsafe.Pointer(&__ret))
	return __v
}

// NkFontAtlasCleanup function as declared in nk/nuklear.h:1441
func NkFontAtlasCleanup(atlas *FontAtlas) {
	catlas, _ := (*C.struct_nk_font_atlas)(unsafe.Pointer(atlas)), cgoAllocsUnknown
	C.nk_font_atlas_cleanup(catlas)
}

// NkFontAtlasClear function as declared in nk/nuklear.h:1442
func NkFontAtlasClear(arg0 *FontAtlas) {
	carg0, _ := (*C.struct_nk_font_atlas)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	C.nk_font_atlas_clear(carg0)
}

// NkBufferInitDefault function as declared in nk/nuklear.h:1527
func NkBufferInitDefault(arg0 *Buffer) {
	carg0, _ := (*C.struct_nk_buffer)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	C.nk_buffer_init_default(carg0)
}

// NkBufferInit function as declared in nk/nuklear.h:1529
func NkBufferInit(arg0 *Buffer, arg1 *Allocator, size Size) {
	carg0, _ := (*C.struct_nk_buffer)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (*C.struct_nk_allocator)(unsafe.Pointer(arg1)), cgoAllocsUnknown
	csize, _ := (C.nk_size)(size), cgoAllocsUnknown
	C.nk_buffer_init(carg0, carg1, csize)
}

// NkBufferInitFixed function as declared in nk/nuklear.h:1530
func NkBufferInitFixed(arg0 *Buffer, memory unsafe.Pointer, size Size) {
	carg0, _ := (*C.struct_nk_buffer)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cmemory, _ := (unsafe.Pointer)(unsafe.Pointer(memory)), cgoAllocsUnknown
	csize, _ := (C.nk_size)(size), cgoAllocsUnknown
	C.nk_buffer_init_fixed(carg0, cmemory, csize)
}

// NkBufferInfo function as declared in nk/nuklear.h:1531
func NkBufferInfo(arg0 *MemoryStatus, arg1 *Buffer) {
	carg0, _ := (*C.struct_nk_memory_status)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (*C.struct_nk_buffer)(unsafe.Pointer(arg1)), cgoAllocsUnknown
	C.nk_buffer_info(carg0, carg1)
}

// NkBufferPush function as declared in nk/nuklear.h:1532
func NkBufferPush(arg0 *Buffer, kind BufferAllocationType, memory unsafe.Pointer, size Size, align Size) {
	carg0, _ := (*C.struct_nk_buffer)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	ckind, _ := (C.enum_nk_buffer_allocation_type)(kind), cgoAllocsUnknown
	cmemory, _ := (unsafe.Pointer)(unsafe.Pointer(memory)), cgoAllocsUnknown
	csize, _ := (C.nk_size)(size), cgoAllocsUnknown
	calign, _ := (C.nk_size)(align), cgoAllocsUnknown
	C.nk_buffer_push(carg0, ckind, cmemory, csize, calign)
}

// NkBufferMark function as declared in nk/nuklear.h:1533
func NkBufferMark(arg0 *Buffer, kind BufferAllocationType) {
	carg0, _ := (*C.struct_nk_buffer)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	ckind, _ := (C.enum_nk_buffer_allocation_type)(kind), cgoAllocsUnknown
	C.nk_buffer_mark(carg0, ckind)
}

// NkBufferReset function as declared in nk/nuklear.h:1534
func NkBufferReset(arg0 *Buffer, kind BufferAllocationType) {
	carg0, _ := (*C.struct_nk_buffer)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	ckind, _ := (C.enum_nk_buffer_allocation_type)(kind), cgoAllocsUnknown
	C.nk_buffer_reset(carg0, ckind)
}

// NkBufferClear function as declared in nk/nuklear.h:1535
func NkBufferClear(arg0 *Buffer) {
	carg0, _ := (*C.struct_nk_buffer)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	C.nk_buffer_clear(carg0)
}

// NkBufferFree function as declared in nk/nuklear.h:1536
func NkBufferFree(arg0 *Buffer) {
	carg0, _ := (*C.struct_nk_buffer)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	C.nk_buffer_free(carg0)
}

// NkBufferMemory function as declared in nk/nuklear.h:1537
func NkBufferMemory(arg0 *Buffer) unsafe.Pointer {
	carg0, _ := (*C.struct_nk_buffer)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	__ret := C.nk_buffer_memory(carg0)
	__v := *(*unsafe.Pointer)(unsafe.Pointer(&__ret))
	return __v
}

// NkBufferMemoryConst function as declared in nk/nuklear.h:1538
func NkBufferMemoryConst(arg0 *Buffer) unsafe.Pointer {
	carg0, _ := (*C.struct_nk_buffer)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	__ret := C.nk_buffer_memory_const(carg0)
	__v := *(*unsafe.Pointer)(unsafe.Pointer(&__ret))
	return __v
}

// NkBufferTotal function as declared in nk/nuklear.h:1539
func NkBufferTotal(arg0 *Buffer) Size {
	carg0, _ := (*C.struct_nk_buffer)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	__ret := C.nk_buffer_total(carg0)
	__v := (Size)(__ret)
	return __v
}

// NkStrInitDefault function as declared in nk/nuklear.h:1557
func NkStrInitDefault(arg0 *Str) {
	carg0, _ := (*C.struct_nk_str)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	C.nk_str_init_default(carg0)
}

// NkStrInit function as declared in nk/nuklear.h:1559
func NkStrInit(arg0 *Str, arg1 *Allocator, size Size) {
	carg0, _ := (*C.struct_nk_str)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (*C.struct_nk_allocator)(unsafe.Pointer(arg1)), cgoAllocsUnknown
	csize, _ := (C.nk_size)(size), cgoAllocsUnknown
	C.nk_str_init(carg0, carg1, csize)
}

// NkStrInitFixed function as declared in nk/nuklear.h:1560
func NkStrInitFixed(arg0 *Str, memory unsafe.Pointer, size Size) {
	carg0, _ := (*C.struct_nk_str)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cmemory, _ := (unsafe.Pointer)(unsafe.Pointer(memory)), cgoAllocsUnknown
	csize, _ := (C.nk_size)(size), cgoAllocsUnknown
	C.nk_str_init_fixed(carg0, cmemory, csize)
}

// NkStrClear function as declared in nk/nuklear.h:1561
func NkStrClear(arg0 *Str) {
	carg0, _ := (*C.struct_nk_str)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	C.nk_str_clear(carg0)
}

// NkStrFree function as declared in nk/nuklear.h:1562
func NkStrFree(arg0 *Str) {
	carg0, _ := (*C.struct_nk_str)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	C.nk_str_free(carg0)
}

// NkStrAppendTextChar function as declared in nk/nuklear.h:1564
func NkStrAppendTextChar(arg0 *Str, arg1 string, arg2 int32) int32 {
	carg0, _ := (*C.struct_nk_str)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	carg2, _ := (C.int)(arg2), cgoAllocsUnknown
	__ret := C.nk_str_append_text_char(carg0, carg1, carg2)
	__v := (int32)(__ret)
	return __v
}

// NkStrAppendStrChar function as declared in nk/nuklear.h:1565
func NkStrAppendStrChar(arg0 *Str, arg1 string) int32 {
	carg0, _ := (*C.struct_nk_str)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	__ret := C.nk_str_append_str_char(carg0, carg1)
	__v := (int32)(__ret)
	return __v
}

// NkStrAppendTextUtf8 function as declared in nk/nuklear.h:1566
func NkStrAppendTextUtf8(arg0 *Str, arg1 string, arg2 int32) int32 {
	carg0, _ := (*C.struct_nk_str)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	carg2, _ := (C.int)(arg2), cgoAllocsUnknown
	__ret := C.nk_str_append_text_utf8(carg0, carg1, carg2)
	__v := (int32)(__ret)
	return __v
}

// NkStrAppendStrUtf8 function as declared in nk/nuklear.h:1567
func NkStrAppendStrUtf8(arg0 *Str, arg1 string) int32 {
	carg0, _ := (*C.struct_nk_str)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	__ret := C.nk_str_append_str_utf8(carg0, carg1)
	__v := (int32)(__ret)
	return __v
}

// NkStrAppendTextRunes function as declared in nk/nuklear.h:1568
func NkStrAppendTextRunes(arg0 *Str, arg1 *Rune, arg2 int32) int32 {
	carg0, _ := (*C.struct_nk_str)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (*C.nk_rune)(unsafe.Pointer(arg1)), cgoAllocsUnknown
	carg2, _ := (C.int)(arg2), cgoAllocsUnknown
	__ret := C.nk_str_append_text_runes(carg0, carg1, carg2)
	__v := (int32)(__ret)
	return __v
}

// NkStrAppendStrRunes function as declared in nk/nuklear.h:1569
func NkStrAppendStrRunes(arg0 *Str, arg1 *Rune) int32 {
	carg0, _ := (*C.struct_nk_str)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (*C.nk_rune)(unsafe.Pointer(arg1)), cgoAllocsUnknown
	__ret := C.nk_str_append_str_runes(carg0, carg1)
	__v := (int32)(__ret)
	return __v
}

// NkStrInsertAtChar function as declared in nk/nuklear.h:1571
func NkStrInsertAtChar(arg0 *Str, pos int32, arg2 string, arg3 int32) int32 {
	carg0, _ := (*C.struct_nk_str)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cpos, _ := (C.int)(pos), cgoAllocsUnknown
	carg2, _ := unpackPCharString(arg2)
	carg3, _ := (C.int)(arg3), cgoAllocsUnknown
	__ret := C.nk_str_insert_at_char(carg0, cpos, carg2, carg3)
	__v := (int32)(__ret)
	return __v
}

// NkStrInsertAtRune function as declared in nk/nuklear.h:1572
func NkStrInsertAtRune(arg0 *Str, pos int32, arg2 string, arg3 int32) int32 {
	carg0, _ := (*C.struct_nk_str)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cpos, _ := (C.int)(pos), cgoAllocsUnknown
	carg2, _ := unpackPCharString(arg2)
	carg3, _ := (C.int)(arg3), cgoAllocsUnknown
	__ret := C.nk_str_insert_at_rune(carg0, cpos, carg2, carg3)
	__v := (int32)(__ret)
	return __v
}

// NkStrInsertTextChar function as declared in nk/nuklear.h:1574
func NkStrInsertTextChar(arg0 *Str, pos int32, arg2 string, arg3 int32) int32 {
	carg0, _ := (*C.struct_nk_str)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cpos, _ := (C.int)(pos), cgoAllocsUnknown
	carg2, _ := unpackPCharString(arg2)
	carg3, _ := (C.int)(arg3), cgoAllocsUnknown
	__ret := C.nk_str_insert_text_char(carg0, cpos, carg2, carg3)
	__v := (int32)(__ret)
	return __v
}

// NkStrInsertStrChar function as declared in nk/nuklear.h:1575
func NkStrInsertStrChar(arg0 *Str, pos int32, arg2 string) int32 {
	carg0, _ := (*C.struct_nk_str)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cpos, _ := (C.int)(pos), cgoAllocsUnknown
	carg2, _ := unpackPCharString(arg2)
	__ret := C.nk_str_insert_str_char(carg0, cpos, carg2)
	__v := (int32)(__ret)
	return __v
}

// NkStrInsertTextUtf8 function as declared in nk/nuklear.h:1576
func NkStrInsertTextUtf8(arg0 *Str, pos int32, arg2 string, arg3 int32) int32 {
	carg0, _ := (*C.struct_nk_str)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cpos, _ := (C.int)(pos), cgoAllocsUnknown
	carg2, _ := unpackPCharString(arg2)
	carg3, _ := (C.int)(arg3), cgoAllocsUnknown
	__ret := C.nk_str_insert_text_utf8(carg0, cpos, carg2, carg3)
	__v := (int32)(__ret)
	return __v
}

// NkStrInsertStrUtf8 function as declared in nk/nuklear.h:1577
func NkStrInsertStrUtf8(arg0 *Str, pos int32, arg2 string) int32 {
	carg0, _ := (*C.struct_nk_str)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cpos, _ := (C.int)(pos), cgoAllocsUnknown
	carg2, _ := unpackPCharString(arg2)
	__ret := C.nk_str_insert_str_utf8(carg0, cpos, carg2)
	__v := (int32)(__ret)
	return __v
}

// NkStrInsertTextRunes function as declared in nk/nuklear.h:1578
func NkStrInsertTextRunes(arg0 *Str, pos int32, arg2 *Rune, arg3 int32) int32 {
	carg0, _ := (*C.struct_nk_str)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cpos, _ := (C.int)(pos), cgoAllocsUnknown
	carg2, _ := (*C.nk_rune)(unsafe.Pointer(arg2)), cgoAllocsUnknown
	carg3, _ := (C.int)(arg3), cgoAllocsUnknown
	__ret := C.nk_str_insert_text_runes(carg0, cpos, carg2, carg3)
	__v := (int32)(__ret)
	return __v
}

// NkStrInsertStrRunes function as declared in nk/nuklear.h:1579
func NkStrInsertStrRunes(arg0 *Str, pos int32, arg2 *Rune) int32 {
	carg0, _ := (*C.struct_nk_str)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cpos, _ := (C.int)(pos), cgoAllocsUnknown
	carg2, _ := (*C.nk_rune)(unsafe.Pointer(arg2)), cgoAllocsUnknown
	__ret := C.nk_str_insert_str_runes(carg0, cpos, carg2)
	__v := (int32)(__ret)
	return __v
}

// NkStrRemoveChars function as declared in nk/nuklear.h:1581
func NkStrRemoveChars(arg0 *Str, len int32) {
	carg0, _ := (*C.struct_nk_str)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	clen, _ := (C.int)(len), cgoAllocsUnknown
	C.nk_str_remove_chars(carg0, clen)
}

// NkStrRemoveRunes function as declared in nk/nuklear.h:1582
func NkStrRemoveRunes(str *Str, len int32) {
	cstr, _ := (*C.struct_nk_str)(unsafe.Pointer(str)), cgoAllocsUnknown
	clen, _ := (C.int)(len), cgoAllocsUnknown
	C.nk_str_remove_runes(cstr, clen)
}

// NkStrDeleteChars function as declared in nk/nuklear.h:1583
func NkStrDeleteChars(arg0 *Str, pos int32, len int32) {
	carg0, _ := (*C.struct_nk_str)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cpos, _ := (C.int)(pos), cgoAllocsUnknown
	clen, _ := (C.int)(len), cgoAllocsUnknown
	C.nk_str_delete_chars(carg0, cpos, clen)
}

// NkStrDeleteRunes function as declared in nk/nuklear.h:1584
func NkStrDeleteRunes(arg0 *Str, pos int32, len int32) {
	carg0, _ := (*C.struct_nk_str)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cpos, _ := (C.int)(pos), cgoAllocsUnknown
	clen, _ := (C.int)(len), cgoAllocsUnknown
	C.nk_str_delete_runes(carg0, cpos, clen)
}

// NkStrAtChar function as declared in nk/nuklear.h:1586
func NkStrAtChar(arg0 *Str, pos int32) *byte {
	carg0, _ := (*C.struct_nk_str)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cpos, _ := (C.int)(pos), cgoAllocsUnknown
	__ret := C.nk_str_at_char(carg0, cpos)
	__v := *(**byte)(unsafe.Pointer(&__ret))
	return __v
}

// NkStrAtRune function as declared in nk/nuklear.h:1587
func NkStrAtRune(arg0 *Str, pos int32, unicode *Rune, len *int32) *byte {
	carg0, _ := (*C.struct_nk_str)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cpos, _ := (C.int)(pos), cgoAllocsUnknown
	cunicode, _ := (*C.nk_rune)(unsafe.Pointer(unicode)), cgoAllocsUnknown
	clen, _ := (*C.int)(unsafe.Pointer(len)), cgoAllocsUnknown
	__ret := C.nk_str_at_rune(carg0, cpos, cunicode, clen)
	__v := *(**byte)(unsafe.Pointer(&__ret))
	return __v
}

// NkStrRuneAt function as declared in nk/nuklear.h:1588
func NkStrRuneAt(arg0 *Str, pos int32) Rune {
	carg0, _ := (*C.struct_nk_str)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cpos, _ := (C.int)(pos), cgoAllocsUnknown
	__ret := C.nk_str_rune_at(carg0, cpos)
	__v := (Rune)(__ret)
	return __v
}

// NkStrAtCharConst function as declared in nk/nuklear.h:1589
func NkStrAtCharConst(arg0 *Str, pos int32) string {
	carg0, _ := (*C.struct_nk_str)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cpos, _ := (C.int)(pos), cgoAllocsUnknown
	__ret := C.nk_str_at_char_const(carg0, cpos)
	__v := packPCharString(__ret)
	return __v
}

// NkStrAtConst function as declared in nk/nuklear.h:1590
func NkStrAtConst(arg0 *Str, pos int32, unicode *Rune, len *int32) string {
	carg0, _ := (*C.struct_nk_str)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cpos, _ := (C.int)(pos), cgoAllocsUnknown
	cunicode, _ := (*C.nk_rune)(unsafe.Pointer(unicode)), cgoAllocsUnknown
	clen, _ := (*C.int)(unsafe.Pointer(len)), cgoAllocsUnknown
	__ret := C.nk_str_at_const(carg0, cpos, cunicode, clen)
	__v := packPCharString(__ret)
	return __v
}

// NkStrGet function as declared in nk/nuklear.h:1592
func NkStrGet(arg0 *Str) *byte {
	carg0, _ := (*C.struct_nk_str)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	__ret := C.nk_str_get(carg0)
	__v := *(**byte)(unsafe.Pointer(&__ret))
	return __v
}

// NkStrGetConst function as declared in nk/nuklear.h:1593
func NkStrGetConst(arg0 *Str) string {
	carg0, _ := (*C.struct_nk_str)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	__ret := C.nk_str_get_const(carg0)
	__v := packPCharString(__ret)
	return __v
}

// NkStrLen function as declared in nk/nuklear.h:1594
func NkStrLen(arg0 *Str) int32 {
	carg0, _ := (*C.struct_nk_str)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	__ret := C.nk_str_len(carg0)
	__v := (int32)(__ret)
	return __v
}

// NkStrLenChar function as declared in nk/nuklear.h:1595
func NkStrLenChar(arg0 *Str) int32 {
	carg0, _ := (*C.struct_nk_str)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	__ret := C.nk_str_len_char(carg0)
	__v := (int32)(__ret)
	return __v
}

// NkFilterDefault function as declared in nk/nuklear.h:1688
func NkFilterDefault(arg0 *TextEdit, unicode Rune) int32 {
	carg0, _ := (*C.struct_nk_text_edit)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cunicode, _ := (C.nk_rune)(unicode), cgoAllocsUnknown
	__ret := C.nk_filter_default(carg0, cunicode)
	__v := (int32)(__ret)
	return __v
}

// NkFilterAscii function as declared in nk/nuklear.h:1689
func NkFilterAscii(arg0 *TextEdit, unicode Rune) int32 {
	carg0, _ := (*C.struct_nk_text_edit)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cunicode, _ := (C.nk_rune)(unicode), cgoAllocsUnknown
	__ret := C.nk_filter_ascii(carg0, cunicode)
	__v := (int32)(__ret)
	return __v
}

// NkFilterFloat function as declared in nk/nuklear.h:1690
func NkFilterFloat(arg0 *TextEdit, unicode Rune) int32 {
	carg0, _ := (*C.struct_nk_text_edit)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cunicode, _ := (C.nk_rune)(unicode), cgoAllocsUnknown
	__ret := C.nk_filter_float(carg0, cunicode)
	__v := (int32)(__ret)
	return __v
}

// NkFilterDecimal function as declared in nk/nuklear.h:1691
func NkFilterDecimal(arg0 *TextEdit, unicode Rune) int32 {
	carg0, _ := (*C.struct_nk_text_edit)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cunicode, _ := (C.nk_rune)(unicode), cgoAllocsUnknown
	__ret := C.nk_filter_decimal(carg0, cunicode)
	__v := (int32)(__ret)
	return __v
}

// NkFilterHex function as declared in nk/nuklear.h:1692
func NkFilterHex(arg0 *TextEdit, unicode Rune) int32 {
	carg0, _ := (*C.struct_nk_text_edit)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cunicode, _ := (C.nk_rune)(unicode), cgoAllocsUnknown
	__ret := C.nk_filter_hex(carg0, cunicode)
	__v := (int32)(__ret)
	return __v
}

// NkFilterOct function as declared in nk/nuklear.h:1693
func NkFilterOct(arg0 *TextEdit, unicode Rune) int32 {
	carg0, _ := (*C.struct_nk_text_edit)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cunicode, _ := (C.nk_rune)(unicode), cgoAllocsUnknown
	__ret := C.nk_filter_oct(carg0, cunicode)
	__v := (int32)(__ret)
	return __v
}

// NkFilterBinary function as declared in nk/nuklear.h:1694
func NkFilterBinary(arg0 *TextEdit, unicode Rune) int32 {
	carg0, _ := (*C.struct_nk_text_edit)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cunicode, _ := (C.nk_rune)(unicode), cgoAllocsUnknown
	__ret := C.nk_filter_binary(carg0, cunicode)
	__v := (int32)(__ret)
	return __v
}

// NkTexteditInitDefault function as declared in nk/nuklear.h:1698
func NkTexteditInitDefault(arg0 *TextEdit) {
	carg0, _ := (*C.struct_nk_text_edit)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	C.nk_textedit_init_default(carg0)
}

// NkTexteditInit function as declared in nk/nuklear.h:1700
func NkTexteditInit(arg0 *TextEdit, arg1 *Allocator, size Size) {
	carg0, _ := (*C.struct_nk_text_edit)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (*C.struct_nk_allocator)(unsafe.Pointer(arg1)), cgoAllocsUnknown
	csize, _ := (C.nk_size)(size), cgoAllocsUnknown
	C.nk_textedit_init(carg0, carg1, csize)
}

// NkTexteditInitFixed function as declared in nk/nuklear.h:1701
func NkTexteditInitFixed(arg0 *TextEdit, memory unsafe.Pointer, size Size) {
	carg0, _ := (*C.struct_nk_text_edit)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cmemory, _ := (unsafe.Pointer)(unsafe.Pointer(memory)), cgoAllocsUnknown
	csize, _ := (C.nk_size)(size), cgoAllocsUnknown
	C.nk_textedit_init_fixed(carg0, cmemory, csize)
}

// NkTexteditFree function as declared in nk/nuklear.h:1702
func NkTexteditFree(arg0 *TextEdit) {
	carg0, _ := (*C.struct_nk_text_edit)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	C.nk_textedit_free(carg0)
}

// NkTexteditText function as declared in nk/nuklear.h:1703
func NkTexteditText(arg0 *TextEdit, arg1 string, totalLen int32) {
	carg0, _ := (*C.struct_nk_text_edit)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	ctotalLen, _ := (C.int)(totalLen), cgoAllocsUnknown
	C.nk_textedit_text(carg0, carg1, ctotalLen)
}

// NkTexteditDelete function as declared in nk/nuklear.h:1704
func NkTexteditDelete(arg0 *TextEdit, where int32, len int32) {
	carg0, _ := (*C.struct_nk_text_edit)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cwhere, _ := (C.int)(where), cgoAllocsUnknown
	clen, _ := (C.int)(len), cgoAllocsUnknown
	C.nk_textedit_delete(carg0, cwhere, clen)
}

// NkTexteditDeleteSelection function as declared in nk/nuklear.h:1705
func NkTexteditDeleteSelection(arg0 *TextEdit) {
	carg0, _ := (*C.struct_nk_text_edit)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	C.nk_textedit_delete_selection(carg0)
}

// NkTexteditSelectAll function as declared in nk/nuklear.h:1706
func NkTexteditSelectAll(arg0 *TextEdit) {
	carg0, _ := (*C.struct_nk_text_edit)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	C.nk_textedit_select_all(carg0)
}

// NkTexteditCut function as declared in nk/nuklear.h:1707
func NkTexteditCut(arg0 *TextEdit) int32 {
	carg0, _ := (*C.struct_nk_text_edit)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	__ret := C.nk_textedit_cut(carg0)
	__v := (int32)(__ret)
	return __v
}

// NkTexteditPaste function as declared in nk/nuklear.h:1708
func NkTexteditPaste(arg0 *TextEdit, arg1 string, len int32) int32 {
	carg0, _ := (*C.struct_nk_text_edit)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := unpackPCharString(arg1)
	clen, _ := (C.int)(len), cgoAllocsUnknown
	__ret := C.nk_textedit_paste(carg0, carg1, clen)
	__v := (int32)(__ret)
	return __v
}

// NkTexteditUndo function as declared in nk/nuklear.h:1709
func NkTexteditUndo(arg0 *TextEdit) {
	carg0, _ := (*C.struct_nk_text_edit)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	C.nk_textedit_undo(carg0)
}

// NkTexteditRedo function as declared in nk/nuklear.h:1710
func NkTexteditRedo(arg0 *TextEdit) {
	carg0, _ := (*C.struct_nk_text_edit)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	C.nk_textedit_redo(carg0)
}

// NkStrokeLine function as declared in nk/nuklear.h:1947
func NkStrokeLine(b *CommandBuffer, x0 float32, y0 float32, x1 float32, y1 float32, lineThickness float32, arg6 Color) {
	cb, _ := (*C.struct_nk_command_buffer)(unsafe.Pointer(b)), cgoAllocsUnknown
	cx0, _ := (C.float)(x0), cgoAllocsUnknown
	cy0, _ := (C.float)(y0), cgoAllocsUnknown
	cx1, _ := (C.float)(x1), cgoAllocsUnknown
	cy1, _ := (C.float)(y1), cgoAllocsUnknown
	clineThickness, _ := (C.float)(lineThickness), cgoAllocsUnknown
	carg6, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg6)), cgoAllocsUnknown
	C.nk_stroke_line(cb, cx0, cy0, cx1, cy1, clineThickness, carg6)
}

// NkStrokeCurve function as declared in nk/nuklear.h:1948
func NkStrokeCurve(arg0 *CommandBuffer, arg1 float32, arg2 float32, arg3 float32, arg4 float32, arg5 float32, arg6 float32, arg7 float32, arg8 float32, lineThickness float32, arg10 Color) {
	carg0, _ := (*C.struct_nk_command_buffer)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.float)(arg1), cgoAllocsUnknown
	carg2, _ := (C.float)(arg2), cgoAllocsUnknown
	carg3, _ := (C.float)(arg3), cgoAllocsUnknown
	carg4, _ := (C.float)(arg4), cgoAllocsUnknown
	carg5, _ := (C.float)(arg5), cgoAllocsUnknown
	carg6, _ := (C.float)(arg6), cgoAllocsUnknown
	carg7, _ := (C.float)(arg7), cgoAllocsUnknown
	carg8, _ := (C.float)(arg8), cgoAllocsUnknown
	clineThickness, _ := (C.float)(lineThickness), cgoAllocsUnknown
	carg10, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg10)), cgoAllocsUnknown
	C.nk_stroke_curve(carg0, carg1, carg2, carg3, carg4, carg5, carg6, carg7, carg8, clineThickness, carg10)
}

// NkStrokeRect function as declared in nk/nuklear.h:1949
func NkStrokeRect(arg0 *CommandBuffer, arg1 Rect, rounding float32, lineThickness float32, arg4 Color) {
	carg0, _ := (*C.struct_nk_command_buffer)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := *(*C.struct_nk_rect)(unsafe.Pointer(&arg1)), cgoAllocsUnknown
	crounding, _ := (C.float)(rounding), cgoAllocsUnknown
	clineThickness, _ := (C.float)(lineThickness), cgoAllocsUnknown
	carg4, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg4)), cgoAllocsUnknown
	C.nk_stroke_rect(carg0, carg1, crounding, clineThickness, carg4)
}

// NkStrokeCircle function as declared in nk/nuklear.h:1950
func NkStrokeCircle(arg0 *CommandBuffer, arg1 Rect, lineThickness float32, arg3 Color) {
	carg0, _ := (*C.struct_nk_command_buffer)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := *(*C.struct_nk_rect)(unsafe.Pointer(&arg1)), cgoAllocsUnknown
	clineThickness, _ := (C.float)(lineThickness), cgoAllocsUnknown
	carg3, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg3)), cgoAllocsUnknown
	C.nk_stroke_circle(carg0, carg1, clineThickness, carg3)
}

// NkStrokeArc function as declared in nk/nuklear.h:1951
func NkStrokeArc(arg0 *CommandBuffer, cx float32, cy float32, radius float32, aMin float32, aMax float32, lineThickness float32, arg7 Color) {
	carg0, _ := (*C.struct_nk_command_buffer)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	ccx, _ := (C.float)(cx), cgoAllocsUnknown
	ccy, _ := (C.float)(cy), cgoAllocsUnknown
	cradius, _ := (C.float)(radius), cgoAllocsUnknown
	caMin, _ := (C.float)(aMin), cgoAllocsUnknown
	caMax, _ := (C.float)(aMax), cgoAllocsUnknown
	clineThickness, _ := (C.float)(lineThickness), cgoAllocsUnknown
	carg7, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg7)), cgoAllocsUnknown
	C.nk_stroke_arc(carg0, ccx, ccy, cradius, caMin, caMax, clineThickness, carg7)
}

// NkStrokeTriangle function as declared in nk/nuklear.h:1952
func NkStrokeTriangle(arg0 *CommandBuffer, arg1 float32, arg2 float32, arg3 float32, arg4 float32, arg5 float32, arg6 float32, lineThichness float32, arg8 Color) {
	carg0, _ := (*C.struct_nk_command_buffer)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.float)(arg1), cgoAllocsUnknown
	carg2, _ := (C.float)(arg2), cgoAllocsUnknown
	carg3, _ := (C.float)(arg3), cgoAllocsUnknown
	carg4, _ := (C.float)(arg4), cgoAllocsUnknown
	carg5, _ := (C.float)(arg5), cgoAllocsUnknown
	carg6, _ := (C.float)(arg6), cgoAllocsUnknown
	clineThichness, _ := (C.float)(lineThichness), cgoAllocsUnknown
	carg8, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg8)), cgoAllocsUnknown
	C.nk_stroke_triangle(carg0, carg1, carg2, carg3, carg4, carg5, carg6, clineThichness, carg8)
}

// NkStrokePolyline function as declared in nk/nuklear.h:1953
func NkStrokePolyline(arg0 *CommandBuffer, points *float32, pointCount int32, lineThickness float32, col Color) {
	carg0, _ := (*C.struct_nk_command_buffer)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cpoints, _ := (*C.float)(unsafe.Pointer(points)), cgoAllocsUnknown
	cpointCount, _ := (C.int)(pointCount), cgoAllocsUnknown
	clineThickness, _ := (C.float)(lineThickness), cgoAllocsUnknown
	ccol, _ := *(*C.struct_nk_color)(unsafe.Pointer(&col)), cgoAllocsUnknown
	C.nk_stroke_polyline(carg0, cpoints, cpointCount, clineThickness, ccol)
}

// NkStrokePolygon function as declared in nk/nuklear.h:1954
func NkStrokePolygon(arg0 *CommandBuffer, arg1 *float32, pointCount int32, lineThickness float32, arg4 Color) {
	carg0, _ := (*C.struct_nk_command_buffer)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (*C.float)(unsafe.Pointer(arg1)), cgoAllocsUnknown
	cpointCount, _ := (C.int)(pointCount), cgoAllocsUnknown
	clineThickness, _ := (C.float)(lineThickness), cgoAllocsUnknown
	carg4, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg4)), cgoAllocsUnknown
	C.nk_stroke_polygon(carg0, carg1, cpointCount, clineThickness, carg4)
}

// NkFillRect function as declared in nk/nuklear.h:1957
func NkFillRect(arg0 *CommandBuffer, arg1 Rect, rounding float32, arg3 Color) {
	carg0, _ := (*C.struct_nk_command_buffer)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := *(*C.struct_nk_rect)(unsafe.Pointer(&arg1)), cgoAllocsUnknown
	crounding, _ := (C.float)(rounding), cgoAllocsUnknown
	carg3, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg3)), cgoAllocsUnknown
	C.nk_fill_rect(carg0, carg1, crounding, carg3)
}

// NkFillRectMultiColor function as declared in nk/nuklear.h:1958
func NkFillRectMultiColor(arg0 *CommandBuffer, arg1 Rect, left Color, top Color, right Color, bottom Color) {
	carg0, _ := (*C.struct_nk_command_buffer)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := *(*C.struct_nk_rect)(unsafe.Pointer(&arg1)), cgoAllocsUnknown
	cleft, _ := *(*C.struct_nk_color)(unsafe.Pointer(&left)), cgoAllocsUnknown
	ctop, _ := *(*C.struct_nk_color)(unsafe.Pointer(&top)), cgoAllocsUnknown
	cright, _ := *(*C.struct_nk_color)(unsafe.Pointer(&right)), cgoAllocsUnknown
	cbottom, _ := *(*C.struct_nk_color)(unsafe.Pointer(&bottom)), cgoAllocsUnknown
	C.nk_fill_rect_multi_color(carg0, carg1, cleft, ctop, cright, cbottom)
}

// NkFillCircle function as declared in nk/nuklear.h:1959
func NkFillCircle(arg0 *CommandBuffer, arg1 Rect, arg2 Color) {
	carg0, _ := (*C.struct_nk_command_buffer)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := *(*C.struct_nk_rect)(unsafe.Pointer(&arg1)), cgoAllocsUnknown
	carg2, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg2)), cgoAllocsUnknown
	C.nk_fill_circle(carg0, carg1, carg2)
}

// NkFillArc function as declared in nk/nuklear.h:1960
func NkFillArc(arg0 *CommandBuffer, cx float32, cy float32, radius float32, aMin float32, aMax float32, arg6 Color) {
	carg0, _ := (*C.struct_nk_command_buffer)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	ccx, _ := (C.float)(cx), cgoAllocsUnknown
	ccy, _ := (C.float)(cy), cgoAllocsUnknown
	cradius, _ := (C.float)(radius), cgoAllocsUnknown
	caMin, _ := (C.float)(aMin), cgoAllocsUnknown
	caMax, _ := (C.float)(aMax), cgoAllocsUnknown
	carg6, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg6)), cgoAllocsUnknown
	C.nk_fill_arc(carg0, ccx, ccy, cradius, caMin, caMax, carg6)
}

// NkFillTriangle function as declared in nk/nuklear.h:1961
func NkFillTriangle(arg0 *CommandBuffer, x0 float32, y0 float32, x1 float32, y1 float32, x2 float32, y2 float32, arg7 Color) {
	carg0, _ := (*C.struct_nk_command_buffer)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cx0, _ := (C.float)(x0), cgoAllocsUnknown
	cy0, _ := (C.float)(y0), cgoAllocsUnknown
	cx1, _ := (C.float)(x1), cgoAllocsUnknown
	cy1, _ := (C.float)(y1), cgoAllocsUnknown
	cx2, _ := (C.float)(x2), cgoAllocsUnknown
	cy2, _ := (C.float)(y2), cgoAllocsUnknown
	carg7, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg7)), cgoAllocsUnknown
	C.nk_fill_triangle(carg0, cx0, cy0, cx1, cy1, cx2, cy2, carg7)
}

// NkFillPolygon function as declared in nk/nuklear.h:1962
func NkFillPolygon(arg0 *CommandBuffer, arg1 *float32, pointCount int32, arg3 Color) {
	carg0, _ := (*C.struct_nk_command_buffer)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (*C.float)(unsafe.Pointer(arg1)), cgoAllocsUnknown
	cpointCount, _ := (C.int)(pointCount), cgoAllocsUnknown
	carg3, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg3)), cgoAllocsUnknown
	C.nk_fill_polygon(carg0, carg1, cpointCount, carg3)
}

// NkPushScissor function as declared in nk/nuklear.h:1965
func NkPushScissor(arg0 *CommandBuffer, arg1 Rect) {
	carg0, _ := (*C.struct_nk_command_buffer)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := *(*C.struct_nk_rect)(unsafe.Pointer(&arg1)), cgoAllocsUnknown
	C.nk_push_scissor(carg0, carg1)
}

// NkDrawImage function as declared in nk/nuklear.h:1966
func NkDrawImage(arg0 *CommandBuffer, arg1 Rect, arg2 *Image, arg3 Color) {
	carg0, _ := (*C.struct_nk_command_buffer)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := *(*C.struct_nk_rect)(unsafe.Pointer(&arg1)), cgoAllocsUnknown
	carg2, _ := (*C.struct_nk_image)(unsafe.Pointer(arg2)), cgoAllocsUnknown
	carg3, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg3)), cgoAllocsUnknown
	C.nk_draw_image(carg0, carg1, carg2, carg3)
}

// NkDrawText function as declared in nk/nuklear.h:1967
func NkDrawText(arg0 *CommandBuffer, arg1 Rect, text string, len int32, arg4 []UserFont, arg5 Color, arg6 Color) {
	carg0, _ := (*C.struct_nk_command_buffer)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := *(*C.struct_nk_rect)(unsafe.Pointer(&arg1)), cgoAllocsUnknown
	ctext, _ := unpackPCharString(text)
	clen, _ := (C.int)(len), cgoAllocsUnknown
	carg4, _ := unpackArgSUserFont(arg4)
	carg5, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg5)), cgoAllocsUnknown
	carg6, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg6)), cgoAllocsUnknown
	C.nk_draw_text(carg0, carg1, ctext, clen, carg4, carg5, carg6)
	packSUserFont(arg4, carg4)
}

// Nk_Next function as declared in nk/nuklear.h:1968
func Nk_Next(arg0 *Context, arg1 *Command) *Command {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (*C.struct_nk_command)(unsafe.Pointer(arg1)), cgoAllocsUnknown
	__ret := C.nk__next(carg0, carg1)
	__v := *(**Command)(unsafe.Pointer(&__ret))
	return __v
}

// Nk_Begin function as declared in nk/nuklear.h:1969
func Nk_Begin(arg0 *Context) *Command {
	carg0, _ := (*C.struct_nk_context)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	__ret := C.nk__begin(carg0)
	__v := *(**Command)(unsafe.Pointer(&__ret))
	return __v
}

// NkInputHasMouseClick function as declared in nk/nuklear.h:2007
func NkInputHasMouseClick(arg0 *Input, arg1 Buttons) int32 {
	carg0, _ := (*C.struct_nk_input)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.enum_nk_buttons)(arg1), cgoAllocsUnknown
	__ret := C.nk_input_has_mouse_click(carg0, carg1)
	__v := (int32)(__ret)
	return __v
}

// NkInputHasMouseClickInRect function as declared in nk/nuklear.h:2008
func NkInputHasMouseClickInRect(arg0 *Input, arg1 Buttons, arg2 Rect) int32 {
	carg0, _ := (*C.struct_nk_input)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.enum_nk_buttons)(arg1), cgoAllocsUnknown
	carg2, _ := *(*C.struct_nk_rect)(unsafe.Pointer(&arg2)), cgoAllocsUnknown
	__ret := C.nk_input_has_mouse_click_in_rect(carg0, carg1, carg2)
	__v := (int32)(__ret)
	return __v
}

// NkInputHasMouseClickDownInRect function as declared in nk/nuklear.h:2009
func NkInputHasMouseClickDownInRect(arg0 *Input, arg1 Buttons, arg2 Rect, down int32) int32 {
	carg0, _ := (*C.struct_nk_input)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.enum_nk_buttons)(arg1), cgoAllocsUnknown
	carg2, _ := *(*C.struct_nk_rect)(unsafe.Pointer(&arg2)), cgoAllocsUnknown
	cdown, _ := (C.int)(down), cgoAllocsUnknown
	__ret := C.nk_input_has_mouse_click_down_in_rect(carg0, carg1, carg2, cdown)
	__v := (int32)(__ret)
	return __v
}

// NkInputIsMouseClickInRect function as declared in nk/nuklear.h:2010
func NkInputIsMouseClickInRect(arg0 *Input, arg1 Buttons, arg2 Rect) int32 {
	carg0, _ := (*C.struct_nk_input)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.enum_nk_buttons)(arg1), cgoAllocsUnknown
	carg2, _ := *(*C.struct_nk_rect)(unsafe.Pointer(&arg2)), cgoAllocsUnknown
	__ret := C.nk_input_is_mouse_click_in_rect(carg0, carg1, carg2)
	__v := (int32)(__ret)
	return __v
}

// NkInputIsMouseClickDownInRect function as declared in nk/nuklear.h:2011
func NkInputIsMouseClickDownInRect(i *Input, id Buttons, b Rect, down int32) int32 {
	ci, _ := (*C.struct_nk_input)(unsafe.Pointer(i)), cgoAllocsUnknown
	cid, _ := (C.enum_nk_buttons)(id), cgoAllocsUnknown
	cb, _ := *(*C.struct_nk_rect)(unsafe.Pointer(&b)), cgoAllocsUnknown
	cdown, _ := (C.int)(down), cgoAllocsUnknown
	__ret := C.nk_input_is_mouse_click_down_in_rect(ci, cid, cb, cdown)
	__v := (int32)(__ret)
	return __v
}

// NkInputAnyMouseClickInRect function as declared in nk/nuklear.h:2012
func NkInputAnyMouseClickInRect(arg0 *Input, arg1 Rect) int32 {
	carg0, _ := (*C.struct_nk_input)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := *(*C.struct_nk_rect)(unsafe.Pointer(&arg1)), cgoAllocsUnknown
	__ret := C.nk_input_any_mouse_click_in_rect(carg0, carg1)
	__v := (int32)(__ret)
	return __v
}

// NkInputIsMousePrevHoveringRect function as declared in nk/nuklear.h:2013
func NkInputIsMousePrevHoveringRect(arg0 *Input, arg1 Rect) int32 {
	carg0, _ := (*C.struct_nk_input)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := *(*C.struct_nk_rect)(unsafe.Pointer(&arg1)), cgoAllocsUnknown
	__ret := C.nk_input_is_mouse_prev_hovering_rect(carg0, carg1)
	__v := (int32)(__ret)
	return __v
}

// NkInputIsMouseHoveringRect function as declared in nk/nuklear.h:2014
func NkInputIsMouseHoveringRect(arg0 *Input, arg1 Rect) int32 {
	carg0, _ := (*C.struct_nk_input)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := *(*C.struct_nk_rect)(unsafe.Pointer(&arg1)), cgoAllocsUnknown
	__ret := C.nk_input_is_mouse_hovering_rect(carg0, carg1)
	__v := (int32)(__ret)
	return __v
}

// NkInputMouseClicked function as declared in nk/nuklear.h:2015
func NkInputMouseClicked(arg0 *Input, arg1 Buttons, arg2 Rect) int32 {
	carg0, _ := (*C.struct_nk_input)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.enum_nk_buttons)(arg1), cgoAllocsUnknown
	carg2, _ := *(*C.struct_nk_rect)(unsafe.Pointer(&arg2)), cgoAllocsUnknown
	__ret := C.nk_input_mouse_clicked(carg0, carg1, carg2)
	__v := (int32)(__ret)
	return __v
}

// NkInputIsMouseDown function as declared in nk/nuklear.h:2016
func NkInputIsMouseDown(arg0 *Input, arg1 Buttons) int32 {
	carg0, _ := (*C.struct_nk_input)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.enum_nk_buttons)(arg1), cgoAllocsUnknown
	__ret := C.nk_input_is_mouse_down(carg0, carg1)
	__v := (int32)(__ret)
	return __v
}

// NkInputIsMousePressed function as declared in nk/nuklear.h:2017
func NkInputIsMousePressed(arg0 *Input, arg1 Buttons) int32 {
	carg0, _ := (*C.struct_nk_input)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.enum_nk_buttons)(arg1), cgoAllocsUnknown
	__ret := C.nk_input_is_mouse_pressed(carg0, carg1)
	__v := (int32)(__ret)
	return __v
}

// NkInputIsMouseReleased function as declared in nk/nuklear.h:2018
func NkInputIsMouseReleased(arg0 *Input, arg1 Buttons) int32 {
	carg0, _ := (*C.struct_nk_input)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.enum_nk_buttons)(arg1), cgoAllocsUnknown
	__ret := C.nk_input_is_mouse_released(carg0, carg1)
	__v := (int32)(__ret)
	return __v
}

// NkInputIsKeyPressed function as declared in nk/nuklear.h:2019
func NkInputIsKeyPressed(arg0 *Input, arg1 Keys) int32 {
	carg0, _ := (*C.struct_nk_input)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.enum_nk_keys)(arg1), cgoAllocsUnknown
	__ret := C.nk_input_is_key_pressed(carg0, carg1)
	__v := (int32)(__ret)
	return __v
}

// NkInputIsKeyReleased function as declared in nk/nuklear.h:2020
func NkInputIsKeyReleased(arg0 *Input, arg1 Keys) int32 {
	carg0, _ := (*C.struct_nk_input)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.enum_nk_keys)(arg1), cgoAllocsUnknown
	__ret := C.nk_input_is_key_released(carg0, carg1)
	__v := (int32)(__ret)
	return __v
}

// NkInputIsKeyDown function as declared in nk/nuklear.h:2021
func NkInputIsKeyDown(arg0 *Input, arg1 Keys) int32 {
	carg0, _ := (*C.struct_nk_input)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (C.enum_nk_keys)(arg1), cgoAllocsUnknown
	__ret := C.nk_input_is_key_down(carg0, carg1)
	__v := (int32)(__ret)
	return __v
}

// NkDrawListInit function as declared in nk/nuklear.h:2126
func NkDrawListInit(arg0 *DrawList) {
	carg0, _ := (*C.struct_nk_draw_list)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	C.nk_draw_list_init(carg0)
}

// NkDrawListSetup function as declared in nk/nuklear.h:2127
func NkDrawListSetup(arg0 *DrawList, arg1 *ConvertConfig, cmds *Buffer, vertices *Buffer, elements []Buffer) {
	carg0, _ := (*C.struct_nk_draw_list)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := arg1.PassRef()
	ccmds, _ := (*C.struct_nk_buffer)(unsafe.Pointer(cmds)), cgoAllocsUnknown
	cvertices, _ := (*C.struct_nk_buffer)(unsafe.Pointer(vertices)), cgoAllocsUnknown
	celements, _ := (*C.struct_nk_buffer)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&elements)).Data)), cgoAllocsUnknown
	C.nk_draw_list_setup(carg0, carg1, ccmds, cvertices, celements)
}

// Nk_DrawListBegin function as declared in nk/nuklear.h:2132
func Nk_DrawListBegin(arg0 *DrawList, arg1 *Buffer) *DrawCommand {
	carg0, _ := (*C.struct_nk_draw_list)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (*C.struct_nk_buffer)(unsafe.Pointer(arg1)), cgoAllocsUnknown
	__ret := C.nk__draw_list_begin(carg0, carg1)
	__v := *(**DrawCommand)(unsafe.Pointer(&__ret))
	return __v
}

// Nk_DrawListNext function as declared in nk/nuklear.h:2133
func Nk_DrawListNext(arg0 *DrawCommand, arg1 *Buffer, arg2 *DrawList) *DrawCommand {
	carg0, _ := (*C.struct_nk_draw_command)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (*C.struct_nk_buffer)(unsafe.Pointer(arg1)), cgoAllocsUnknown
	carg2, _ := (*C.struct_nk_draw_list)(unsafe.Pointer(arg2)), cgoAllocsUnknown
	__ret := C.nk__draw_list_next(carg0, carg1, carg2)
	__v := *(**DrawCommand)(unsafe.Pointer(&__ret))
	return __v
}

// Nk_DrawListEnd function as declared in nk/nuklear.h:2134
func Nk_DrawListEnd(arg0 *DrawList, arg1 *Buffer) *DrawCommand {
	carg0, _ := (*C.struct_nk_draw_list)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := (*C.struct_nk_buffer)(unsafe.Pointer(arg1)), cgoAllocsUnknown
	__ret := C.nk__draw_list_end(carg0, carg1)
	__v := *(**DrawCommand)(unsafe.Pointer(&__ret))
	return __v
}

// NkDrawListPathClear function as declared in nk/nuklear.h:2138
func NkDrawListPathClear(arg0 *DrawList) {
	carg0, _ := (*C.struct_nk_draw_list)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	C.nk_draw_list_path_clear(carg0)
}

// NkDrawListPathLineTo function as declared in nk/nuklear.h:2139
func NkDrawListPathLineTo(arg0 *DrawList, pos Vec2) {
	carg0, _ := (*C.struct_nk_draw_list)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cpos, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&pos)), cgoAllocsUnknown
	C.nk_draw_list_path_line_to(carg0, cpos)
}

// NkDrawListPathArcToFast function as declared in nk/nuklear.h:2140
func NkDrawListPathArcToFast(arg0 *DrawList, center Vec2, radius float32, aMin int32, aMax int32) {
	carg0, _ := (*C.struct_nk_draw_list)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	ccenter, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&center)), cgoAllocsUnknown
	cradius, _ := (C.float)(radius), cgoAllocsUnknown
	caMin, _ := (C.int)(aMin), cgoAllocsUnknown
	caMax, _ := (C.int)(aMax), cgoAllocsUnknown
	C.nk_draw_list_path_arc_to_fast(carg0, ccenter, cradius, caMin, caMax)
}

// NkDrawListPathArcTo function as declared in nk/nuklear.h:2141
func NkDrawListPathArcTo(arg0 *DrawList, center Vec2, radius float32, aMin float32, aMax float32, segments uint32) {
	carg0, _ := (*C.struct_nk_draw_list)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	ccenter, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&center)), cgoAllocsUnknown
	cradius, _ := (C.float)(radius), cgoAllocsUnknown
	caMin, _ := (C.float)(aMin), cgoAllocsUnknown
	caMax, _ := (C.float)(aMax), cgoAllocsUnknown
	csegments, _ := (C.uint)(segments), cgoAllocsUnknown
	C.nk_draw_list_path_arc_to(carg0, ccenter, cradius, caMin, caMax, csegments)
}

// NkDrawListPathRectTo function as declared in nk/nuklear.h:2142
func NkDrawListPathRectTo(arg0 *DrawList, a Vec2, b Vec2, rounding float32) {
	carg0, _ := (*C.struct_nk_draw_list)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	ca, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&a)), cgoAllocsUnknown
	cb, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&b)), cgoAllocsUnknown
	crounding, _ := (C.float)(rounding), cgoAllocsUnknown
	C.nk_draw_list_path_rect_to(carg0, ca, cb, crounding)
}

// NkDrawListPathCurveTo function as declared in nk/nuklear.h:2143
func NkDrawListPathCurveTo(arg0 *DrawList, p2 Vec2, p3 Vec2, p4 Vec2, numSegments uint32) {
	carg0, _ := (*C.struct_nk_draw_list)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cp2, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&p2)), cgoAllocsUnknown
	cp3, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&p3)), cgoAllocsUnknown
	cp4, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&p4)), cgoAllocsUnknown
	cnumSegments, _ := (C.uint)(numSegments), cgoAllocsUnknown
	C.nk_draw_list_path_curve_to(carg0, cp2, cp3, cp4, cnumSegments)
}

// NkDrawListPathFill function as declared in nk/nuklear.h:2144
func NkDrawListPathFill(arg0 *DrawList, arg1 Color) {
	carg0, _ := (*C.struct_nk_draw_list)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg1)), cgoAllocsUnknown
	C.nk_draw_list_path_fill(carg0, carg1)
}

// NkDrawListPathStroke function as declared in nk/nuklear.h:2145
func NkDrawListPathStroke(arg0 *DrawList, arg1 Color, closed DrawListStroke, thickness float32) {
	carg0, _ := (*C.struct_nk_draw_list)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg1)), cgoAllocsUnknown
	cclosed, _ := (C.enum_nk_draw_list_stroke)(closed), cgoAllocsUnknown
	cthickness, _ := (C.float)(thickness), cgoAllocsUnknown
	C.nk_draw_list_path_stroke(carg0, carg1, cclosed, cthickness)
}

// NkDrawListStrokeLine function as declared in nk/nuklear.h:2148
func NkDrawListStrokeLine(arg0 *DrawList, a Vec2, b Vec2, arg3 Color, thickness float32) {
	carg0, _ := (*C.struct_nk_draw_list)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	ca, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&a)), cgoAllocsUnknown
	cb, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&b)), cgoAllocsUnknown
	carg3, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg3)), cgoAllocsUnknown
	cthickness, _ := (C.float)(thickness), cgoAllocsUnknown
	C.nk_draw_list_stroke_line(carg0, ca, cb, carg3, cthickness)
}

// NkDrawListStrokeRect function as declared in nk/nuklear.h:2149
func NkDrawListStrokeRect(arg0 *DrawList, rect Rect, arg2 Color, rounding float32, thickness float32) {
	carg0, _ := (*C.struct_nk_draw_list)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	crect, _ := *(*C.struct_nk_rect)(unsafe.Pointer(&rect)), cgoAllocsUnknown
	carg2, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg2)), cgoAllocsUnknown
	crounding, _ := (C.float)(rounding), cgoAllocsUnknown
	cthickness, _ := (C.float)(thickness), cgoAllocsUnknown
	C.nk_draw_list_stroke_rect(carg0, crect, carg2, crounding, cthickness)
}

// NkDrawListStrokeTriangle function as declared in nk/nuklear.h:2150
func NkDrawListStrokeTriangle(arg0 *DrawList, a Vec2, b Vec2, c Vec2, arg4 Color, thickness float32) {
	carg0, _ := (*C.struct_nk_draw_list)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	ca, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&a)), cgoAllocsUnknown
	cb, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&b)), cgoAllocsUnknown
	cc, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&c)), cgoAllocsUnknown
	carg4, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg4)), cgoAllocsUnknown
	cthickness, _ := (C.float)(thickness), cgoAllocsUnknown
	C.nk_draw_list_stroke_triangle(carg0, ca, cb, cc, carg4, cthickness)
}

// NkDrawListStrokeCircle function as declared in nk/nuklear.h:2151
func NkDrawListStrokeCircle(arg0 *DrawList, center Vec2, radius float32, arg3 Color, segs uint32, thickness float32) {
	carg0, _ := (*C.struct_nk_draw_list)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	ccenter, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&center)), cgoAllocsUnknown
	cradius, _ := (C.float)(radius), cgoAllocsUnknown
	carg3, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg3)), cgoAllocsUnknown
	csegs, _ := (C.uint)(segs), cgoAllocsUnknown
	cthickness, _ := (C.float)(thickness), cgoAllocsUnknown
	C.nk_draw_list_stroke_circle(carg0, ccenter, cradius, carg3, csegs, cthickness)
}

// NkDrawListStrokeCurve function as declared in nk/nuklear.h:2152
func NkDrawListStrokeCurve(arg0 *DrawList, p0 Vec2, cp0 Vec2, cp1 Vec2, p1 Vec2, arg5 Color, segments uint32, thickness float32) {
	carg0, _       := (*C.struct_nk_draw_list)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cp0, _          = Vec2(*(*C.struct_nk_vec2)(unsafe.Pointer(&p0))), cgoAllocsUnknown
	ccp0, _        := *(*C.struct_nk_vec2)(unsafe.Pointer(&cp0)), cgoAllocsUnknown
	ccp1, _        := *(*C.struct_nk_vec2)(unsafe.Pointer(&cp1)), cgoAllocsUnknown
	cp1, _          = Vec2(*(*C.struct_nk_vec2)(unsafe.Pointer(&p1))), cgoAllocsUnknown
	carg5, _       := *(*C.struct_nk_color)(unsafe.Pointer(&arg5)), cgoAllocsUnknown
	csegments, _   := (C.uint)(segments), cgoAllocsUnknown
	cthickness, _  := (C.float)(thickness), cgoAllocsUnknown
	C.nk_draw_list_stroke_curve(carg0, C.struct_nk_vec2(cp0), ccp0, ccp1,C.struct_nk_vec2( cp1), carg5, csegments, cthickness)
}

// NkDrawListStrokePolyLine function as declared in nk/nuklear.h:2153
func NkDrawListStrokePolyLine(arg0 *DrawList, pnts *Vec2, cnt uint32, arg3 Color, arg4 DrawListStroke, thickness float32, arg6 AntiAliasing) {
	carg0, _ := (*C.struct_nk_draw_list)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cpnts, _ := (*C.struct_nk_vec2)(unsafe.Pointer(pnts)), cgoAllocsUnknown
	ccnt, _ := (C.uint)(cnt), cgoAllocsUnknown
	carg3, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg3)), cgoAllocsUnknown
	carg4, _ := (C.enum_nk_draw_list_stroke)(arg4), cgoAllocsUnknown
	cthickness, _ := (C.float)(thickness), cgoAllocsUnknown
	carg6, _ := (C.enum_nk_anti_aliasing)(arg6), cgoAllocsUnknown
	C.nk_draw_list_stroke_poly_line(carg0, cpnts, ccnt, carg3, carg4, cthickness, carg6)
}

// NkDrawListFillRect function as declared in nk/nuklear.h:2156
func NkDrawListFillRect(arg0 *DrawList, rect Rect, arg2 Color, rounding float32) {
	carg0, _ := (*C.struct_nk_draw_list)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	crect, _ := *(*C.struct_nk_rect)(unsafe.Pointer(&rect)), cgoAllocsUnknown
	carg2, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg2)), cgoAllocsUnknown
	crounding, _ := (C.float)(rounding), cgoAllocsUnknown
	C.nk_draw_list_fill_rect(carg0, crect, carg2, crounding)
}

// NkDrawListFillRectMultiColor function as declared in nk/nuklear.h:2157
func NkDrawListFillRectMultiColor(arg0 *DrawList, rect Rect, left Color, top Color, right Color, bottom Color) {
	carg0, _ := (*C.struct_nk_draw_list)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	crect, _ := *(*C.struct_nk_rect)(unsafe.Pointer(&rect)), cgoAllocsUnknown
	cleft, _ := *(*C.struct_nk_color)(unsafe.Pointer(&left)), cgoAllocsUnknown
	ctop, _ := *(*C.struct_nk_color)(unsafe.Pointer(&top)), cgoAllocsUnknown
	cright, _ := *(*C.struct_nk_color)(unsafe.Pointer(&right)), cgoAllocsUnknown
	cbottom, _ := *(*C.struct_nk_color)(unsafe.Pointer(&bottom)), cgoAllocsUnknown
	C.nk_draw_list_fill_rect_multi_color(carg0, crect, cleft, ctop, cright, cbottom)
}

// NkDrawListFillTriangle function as declared in nk/nuklear.h:2158
func NkDrawListFillTriangle(arg0 *DrawList, a Vec2, b Vec2, c Vec2, arg4 Color) {
	carg0, _ := (*C.struct_nk_draw_list)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	ca, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&a)), cgoAllocsUnknown
	cb, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&b)), cgoAllocsUnknown
	cc, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&c)), cgoAllocsUnknown
	carg4, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg4)), cgoAllocsUnknown
	C.nk_draw_list_fill_triangle(carg0, ca, cb, cc, carg4)
}

// NkDrawListFillCircle function as declared in nk/nuklear.h:2159
func NkDrawListFillCircle(arg0 *DrawList, center Vec2, radius float32, col Color, segs uint32) {
	carg0, _ := (*C.struct_nk_draw_list)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	ccenter, _ := *(*C.struct_nk_vec2)(unsafe.Pointer(&center)), cgoAllocsUnknown
	cradius, _ := (C.float)(radius), cgoAllocsUnknown
	ccol, _ := *(*C.struct_nk_color)(unsafe.Pointer(&col)), cgoAllocsUnknown
	csegs, _ := (C.uint)(segs), cgoAllocsUnknown
	C.nk_draw_list_fill_circle(carg0, ccenter, cradius, ccol, csegs)
}

// NkDrawListFillPolyConvex function as declared in nk/nuklear.h:2160
func NkDrawListFillPolyConvex(arg0 *DrawList, points *Vec2, count uint32, arg3 Color, arg4 AntiAliasing) {
	carg0, _ := (*C.struct_nk_draw_list)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	cpoints, _ := (*C.struct_nk_vec2)(unsafe.Pointer(points)), cgoAllocsUnknown
	ccount, _ := (C.uint)(count), cgoAllocsUnknown
	carg3, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg3)), cgoAllocsUnknown
	carg4, _ := (C.enum_nk_anti_aliasing)(arg4), cgoAllocsUnknown
	C.nk_draw_list_fill_poly_convex(carg0, cpoints, ccount, carg3, carg4)
}

// NkDrawListAddImage function as declared in nk/nuklear.h:2163
func NkDrawListAddImage(arg0 *DrawList, texture Image, rect Rect, arg3 Color) {
	carg0, _ := (*C.struct_nk_draw_list)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	ctexture, _ := *(*C.struct_nk_image)(unsafe.Pointer(&texture)), cgoAllocsUnknown
	crect, _ := *(*C.struct_nk_rect)(unsafe.Pointer(&rect)), cgoAllocsUnknown
	carg3, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg3)), cgoAllocsUnknown
	C.nk_draw_list_add_image(carg0, ctexture, crect, carg3)
}

// NkDrawListAddText function as declared in nk/nuklear.h:2164
func NkDrawListAddText(arg0 *DrawList, arg1 *UserFont, arg2 Rect, text string, len int32, fontHeight float32, arg6 Color) {
	carg0, _ := (*C.struct_nk_draw_list)(unsafe.Pointer(arg0)), cgoAllocsUnknown
	carg1, _ := arg1.PassRef()
	carg2, _ := *(*C.struct_nk_rect)(unsafe.Pointer(&arg2)), cgoAllocsUnknown
	ctext, _ := unpackPCharString(text)
	clen, _ := (C.int)(len), cgoAllocsUnknown
	cfontHeight, _ := (C.float)(fontHeight), cgoAllocsUnknown
	carg6, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg6)), cgoAllocsUnknown
	C.nk_draw_list_add_text(carg0, carg1, carg2, ctext, clen, cfontHeight, carg6)
}

// NkStyleItemImage function as declared in nk/nuklear.h:2608
func NkStyleItemImage(img Image) StyleItem {
	cimg, _ := *(*C.struct_nk_image)(unsafe.Pointer(&img)), cgoAllocsUnknown
	__ret := C.nk_style_item_image(cimg)
	__v := *(*StyleItem)(unsafe.Pointer(&__ret))
	return __v
}

// NkStyleItemColor function as declared in nk/nuklear.h:2609
func NkStyleItemColor(arg0 Color) StyleItem {
	carg0, _ := *(*C.struct_nk_color)(unsafe.Pointer(&arg0)), cgoAllocsUnknown
	__ret := C.nk_style_item_color(carg0)
	__v := *(*StyleItem)(unsafe.Pointer(&__ret))
	return __v
}

// NkStyleItemHide function as declared in nk/nuklear.h:2610
func NkStyleItemHide() StyleItem {
	__ret := C.nk_style_item_hide()
	__v := *(*StyleItem)(unsafe.Pointer(&__ret))
	return __v
}
