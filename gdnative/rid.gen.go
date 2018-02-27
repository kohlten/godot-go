package gdnative

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "types.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

/*
#include "gdnative.gen.h"
#include <gdnative/rid.h>
// Include all headers for now. TODO: Look up all the required
// headers we need to import based on the method arguments and return types.
#include <gdnative/aabb.h>
#include <gdnative/array.h>
#include <gdnative/basis.h>
#include <gdnative/color.h>
#include <gdnative/dictionary.h>
#include <gdnative/gdnative.h>
#include <gdnative/node_path.h>
#include <gdnative/plane.h>
#include <gdnative/pool_arrays.h>
#include <gdnative/quat.h>
#include <gdnative/rect2.h>
#include <gdnative/rid.h>
#include <gdnative/string.h>
#include <gdnative/string_name.h>
#include <gdnative/transform.h>
#include <gdnative/transform2d.h>
#include <gdnative/variant.h>
#include <gdnative/vector2.h>
#include <gdnative/vector3.h>
#include <gdnative_api_struct.gen.h>
*/
import "C"
import "unsafe"

type Rid struct {
	base *C.godot_rid
}

func (gdt Rid) getBase() *C.godot_rid {
	return gdt.base
}

// NewRid godot_rid_new [[godot_rid * r_dest]] void
func NewRid() *Rid {
	var dest C.godot_rid
	C.go_godot_rid_new(GDNative.api, &dest)
	return &Rid{base: &dest}
}

// GetId godot_rid_get_id [[const godot_rid * p_self]] godot_int
func (gdt *Rid) GetId() Int {
	arg0 := gdt.getBase()

	ret := C.go_godot_rid_get_id(GDNative.api, arg0)

	return Int(ret)
}

// NewRidWithResource godot_rid_new_with_resource [[godot_rid * r_dest] [const godot_object * p_from]] void
func NewRidWithResource(from Object) *Rid {
	var dest C.godot_rid
	arg1 := unsafe.Pointer(from.getBase())
	C.go_godot_rid_new_with_resource(GDNative.api, &dest, arg1)
	return &Rid{base: &dest}
}

// OperatorEqual godot_rid_operator_equal [[const godot_rid * p_self] [const godot_rid * p_b]] godot_bool
func (gdt *Rid) OperatorEqual(b Rid) Bool {
	arg0 := gdt.getBase()
	arg1 := b.getBase()

	ret := C.go_godot_rid_operator_equal(GDNative.api, arg0, arg1)

	return Bool(ret)
}

// OperatorLess godot_rid_operator_less [[const godot_rid * p_self] [const godot_rid * p_b]] godot_bool
func (gdt *Rid) OperatorLess(b Rid) Bool {
	arg0 := gdt.getBase()
	arg1 := b.getBase()

	ret := C.go_godot_rid_operator_less(GDNative.api, arg0, arg1)

	return Bool(ret)
}
