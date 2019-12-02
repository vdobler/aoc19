package aoc

import (
	"io/ioutil"
	"log"
	"reflect"
	"strconv"
	"strings"
)

var InputFile = "input.txt"

// ReadInput reads the file "input.txt" and fills v with the data.
//
// The argument v must be a pointer to a slice. Each line in input.txt results
// in an element of the slice. Each field in a line is converted to field in
// the slice's element by splitting the line on elements of ws.
func ReadInput(v interface{}, ws string) {
	vt := reflect.TypeOf(v)
	if vt.Kind() != reflect.Ptr {
		log.Fatalf("Non pointer type %T", v)
	}

	all, err := ioutil.ReadFile(InputFile)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(strings.TrimSpace(string(all)), "\n")
	N := len(lines)

	st := vt.Elem()
	et := st.Elem()
	slice := reflect.MakeSlice(st, N, N)

	wsr := []rune(ws)
	fff := func(x rune) bool {
		for _, r := range wsr {
			if x == r {
				return true
			}
		}
		return false
	}

	for i, line := range lines {
		cur := slice.Index(i)
		fields := strings.FieldsFunc(line, fff)
		for f := 0; f < et.NumField(); f++ {
			target := cur.Field(f)
			if cur.Type().Field(f).Name == "_" {
				continue
			}
			switch et.Field(f).Type.Kind() {
			case reflect.String:
				target.SetString(fields[f])
			case reflect.Int:
				val, err := strconv.Atoi(fields[f])
				if err != nil {
					log.Fatalf("Line %d, fields %s %d: %q is not an int: %s",
						i+1, et.Field(f).Name, f+1, fields[f], err)
				}

				target.SetInt(int64(val))
			default:
				log.Fatalf("Unsupported type %s of %s for col %d",
					et.Field(f).Type, et.Field(f).Name, f+1)
			}
		}
	}
	reflect.ValueOf(v).Elem().Set(slice)
}
