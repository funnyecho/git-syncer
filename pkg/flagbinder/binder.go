package flagbinder

import (
	"flag"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

func MustBind(f interface{}, fs *flag.FlagSet) {
	if err := bindFlags(f, fs); err != nil {
		panic(err)
	}
}

func Bind(f interface{}, fs *flag.FlagSet) error {
	return bindFlags(f, fs)
}

func bindFlags(f interface{}, fs *flag.FlagSet) error {
	if fs == nil {
		fs = flag.CommandLine
	}

	v := reflect.ValueOf(f).Elem()

	return bindFlagsFromValue("", v, fs)
}

func bindFlagsFromValue(prefix string, v reflect.Value, flagset *flag.FlagSet) error {
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		vField := v.Field(i)
		tField := t.Field(i)

		fieldKind := vField.Kind()
		fieldAddr := vField.UnsafeAddr()

		rawNames := tField.Tag.Get("flag")
		dValue := tField.Tag.Get("value")
		usage := tField.Tag.Get("usage")

		if fieldKind == reflect.Struct {
			if rawNames != "" {
				for _, name := range strings.Split(rawNames, ",") {
					var p string
					if prefix == "" {
						p = name
					} else {
						p = fmt.Sprintf("%s--%s", prefix, name)
					}
					if err := bindFlagsFromValue(p, vField, flagset); err != nil {
						return err
					}
				}
			} else {
				if err := bindFlagsFromValue(prefix, vField, flagset); err != nil {
					return err
				}
			}

			continue
		}

		if rawNames == "" {
			continue
		}

		names := strings.Split(rawNames, ",")
		for _, name := range names {
			if prefix != "" {
				name = fmt.Sprintf("%s--%s", prefix, name)
			}
			switch fieldKind {
			case reflect.Bool:
				value, _ := strconv.ParseBool(dValue)
				flagset.BoolVar((*bool)(unsafe.Pointer(fieldAddr)), name, value, usage)
			case reflect.String:
				flagset.StringVar((*string)(unsafe.Pointer(fieldAddr)), name, dValue, usage)
			case reflect.Int:
				value, _ := strconv.Atoi(dValue)
				flagset.IntVar((*int)(unsafe.Pointer(fieldAddr)), name, value, usage)
			default:
				return fmt.Errorf("type of field %s:%s do not support", tField.Name, tField.Type.String())
			}
		}
	}

	return nil
}
