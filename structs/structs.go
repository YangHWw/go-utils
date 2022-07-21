package structs

import (
	"encoding/json"
	"fmt"
	"github.com/YangHWw/go-utils"
	"github.com/pkg/errors"
	"reflect"
)

func dataPrepare(base, update interface{}) (map[string]interface{}, map[string]interface{}, error) {
	jsonBase, err := json.Marshal(base)
	if err != nil {
		return nil, nil, errors.Wrap(err, "base data struct marshal")
	}
	jsonUpdate, err := json.Marshal(update)
	if err != nil {
		return nil, nil, errors.Wrap(err, "update data struct marshal")
	}
	var mapBase, mapUpdate map[string]interface{}

	err = json.Unmarshal(jsonBase, &mapBase)
	if err != nil {
		return nil, nil, errors.Wrap(err, "base data unmarshal to map")
	}
	err = json.Unmarshal(jsonUpdate, &mapUpdate)
	if err != nil {
		return nil, nil, errors.Wrap(err, "update data unmarshal to map")
	}

	return mapBase, mapUpdate, nil
}

func isMap(val interface{}) bool {
	value := utils.ValueOf(val)

	if value.Kind() == reflect.Map {
		return true
	}
	return false
}

// updateMap, update the keys value that both appear in updated and base dictionary
func updateMap(base, update map[string]interface{}) {
	for key, baseValue := range base {
		if _, ok := update[key]; !ok {
			continue
		}
		updateValue := update[key]
		baseOk := isMap(baseValue)
		updateOk := isMap(updateValue)

		if baseOk && updateOk {
			deepUpdate(baseValue.(map[string]interface{}), updateValue.(map[string]interface{}))
		} else {
			base[key] = updateValue
		}
	}
}

// addMap, add keys that appear only in the updated dictionary to the base dictionary
func addMap(base, update map[string]interface{}) map[string]interface{} {
	for key, updateValue := range update {
		if _, ok := base[key]; !ok {
			// the key just in update map, insert it directly
			base[key] = updateValue
		}
	}
	return base
}

func deepUpdate(base, update map[string]interface{}) {
	updateMap(base, update)
	addMap(base, update)
}

func DeepUpdateStruct(b, u, out interface{}) (retErr error) {
	defer func() {
		if e := recover(); e != nil {
			retErr = fmt.Errorf("deep update struct panic, %v", e)
			return
		}
	}()

	baseType := utils.TypeOf(b)
	updateType := utils.TypeOf(u)

	if baseType != updateType {
		return fmt.Errorf("baseType[%v] not equal with updateType[%v]", baseType, updateType)
	}

	mapBase, mapUpdate, err := dataPrepare(b, u)
	if err != nil {
		return errors.Wrap(err, "data prepare")
	}
	deepUpdate(mapBase, mapUpdate)

	mapStr, err := json.Marshal(mapBase)
	if err != nil {
		return fmt.Errorf("map marshal error, %v", err)
	}
	err = json.Unmarshal(mapStr, out)
	if err != nil {
		return fmt.Errorf("map str unmarshal to output data error, %v", err)
	}

	return nil
}
