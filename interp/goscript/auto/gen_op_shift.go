package main

import (
	"fmt"
	"strings"
	"text/template"
)

var opShiftList = map[string]opValues{
	"SHL": {
		op:     "<<",
		values: removeDefaultValuesAttributes("String", "Float32", "Float64", "Complex64", "Complex128"),
	},
	"SHR": {
		op:     ">>",
		values: removeDefaultValuesAttributes("String", "Float32", "Float64", "Complex64", "Complex128"),
	},
}

var templateShiftFuncBody = `
	ir := pfn.regIndex(instr)
	ix, kx, vx := pfn.regIndex3(instr.X)
	iy, ky, vy := pfn.regIndex3(instr.Y)
	xtyp := pfn.Interp.preToType(instr.X.Type())
	ytyp := pfn.Interp.preToType(instr.Y.Type())
	xkind := xtyp.Kind()
	ykind := ytyp.Kind()
	{{ $op := .op  }}
	{{ $types := .types  }}
	if kx == kindConst && ky == kindConst {
		t := xtype.TypeOfType(xtyp)
		switch xkind {
		{{ range $key, $value := $types }}
		case reflect.{{ $key }}:
			x := xtype.{{ $key }}(vx)
			switch ykind {
			{{ range $ykey, $yvalue := $types }}
			case reflect.{{ $ykey }}:
				v := xtype.Make(t, x {{ $op }} xtype.{{ $ykey }}(vy))
				return func(vm *goVm) { vm.setReg(ir, v) }
			{{ end }}
			}
		{{ end }}
		}
	}
	if xtyp.PkgPath() == "" {
		switch xkind {
		{{ range $key, $value := $types }}
		case reflect.{{ $key }}:
			if kx == kindConst {
				x := vx.({{ $value }})
				switch ykind {
				{{ range $ykey, $yvalue := $types }}
				case reflect.{{ $ykey }}:
					return func(vm *goVm) { vm.setReg(ir, x {{ $op }} vm.{{ $yvalue }}(iy)) }
				{{ end }}
				}
			} else if ky == kindConst {
				switch ykind {
				{{ range $ykey, $yvalue := $types }}
				case reflect.{{ $ykey }}:
					y := xtype.{{ $ykey }}(vy)
					return func(vm *goVm) { vm.setReg(ir, vm.reg(ix).({{ $value }}) {{ $op }} y) }
				{{ end }}
				}
			} else {
				switch ykind {
				{{ range $ykey, $yvalue := $types }}
				case reflect.{{ $ykey }}:
					return func(vm *goVm) { vm.setReg(ir, vm.reg(ix).({{ $value }}) {{ $op }} vm.{{ $yvalue }}(iy)) }
				{{ end }}
				}
			}
		{{ end }}
		}
	} else {
		t := xtype.TypeOfType(xtyp)
		switch xkind {
		{{ range $key, $value := $types }}
		case reflect.{{ $key }}:
			if kx == kindConst {
				x := xtype.{{ $key }}(vx)
				switch ykind {
				{{ range $ykey, $yvalue := $types }}
				case reflect.{{ $ykey }}:
					return func(vm *goVm) { vm.setReg(ir, xtype.Make(t, x {{ $op }} vm.{{ $yvalue }}(iy))) }
				{{ end }}
				}
			} else if ky == kindConst {
				switch ykind {
				{{ range $ykey, $yvalue := $types }}
				case reflect.{{ $ykey }}:
					y := xtype.{{ $ykey }}(vy)
					return func(vm *goVm) { vm.setReg(ir, xtype.Make(t, vm.{{ $value }}(ix) {{ $op }} y)) }
				{{ end }}
				}
			} else {
				switch ykind {
				{{ range $ykey, $yvalue := $types }}
				case reflect.{{ $ykey }}:
					return func(vm *goVm) { vm.setReg(ir, xtype.Make(t, vm.{{ $value }}(ix) {{ $op }} vm.{{ $yvalue }}(iy))) }
				{{ end }}
				}
			}
		{{ end }}
		}
	}
	panic("unreachable kind: " + fmt.Sprintf("%v", xkind))
`

func generateOpShift() {
	var resultStr string
	for name, op := range opShiftList {
		templateStr := new2Line + "func makeBinOp" + name + "(pfn *function, instr *ssa.BinOp) func(vm *goVm) {" + templateShiftFuncBody + "}"
		op.values["op"] = op.op
		tx := template.Must(template.New(name).Parse(templateStr))
		result := new(strings.Builder)
		if err := tx.Execute(result, op.values); err != nil {
			panic(err)
		} else {
			resultStr += result.String()
		}
	}
	resultStr = generateImportWithSSA + resultStr
	fmt.Println(resultStr)
}