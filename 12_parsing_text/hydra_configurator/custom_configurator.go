package hydra_configurator

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type ConfigFields map[string]reflect.Value

func (f ConfigFields) Add(name, value, t string) error {
	switch t {
	case "STRING":
		f[name] = reflect.ValueOf(value)
	case "INTEGER":
		i, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		f[name] = reflect.ValueOf(i)
	case "FLOAT":
		fl, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		f[name] = reflect.ValueOf(fl)
	case "BOOL":
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		f[name] = reflect.ValueOf(b)
	}
	return nil
}

func MarshalCustomConfig(v reflect.Value, filename string) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic Recovered", r)
		}
	}()

	if !v.CanSet() {
		return errors.New("Value passed not settable")
	}

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var fields = make(ConfigFields)
	var scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		var line = scanner.Text()
		fmt.Println("Processing line", line)
		var args = strings.Split(line, "|")
		var valueType = strings.Split(args[1], ";")
		name, value, vtype := strings.TrimSpace(args[0]), strings.TrimSpace(valueType[0]), strings.TrimSpace(valueType[1])
		fields.Add(name, value, vtype)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	var vt = v.Type()

	for i := 0; i < v.NumField(); i++ {
		var fieldType = vt.Field(i)
		var fieldVal = v.Field(i)

		name := fieldType.Tag.Get("name")
		if name == "" {
			name = fieldType.Name
		}

		if v, ok := fields[name]; ok {
			fieldVal.Set(v)
		}
	}

	return nil
}