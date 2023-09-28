package utils

import "encoding/json"

func MergeIntoMap(obj any, target *map[string]interface{}) error {
	objMap, err := objectToMap(obj)
	if err != nil {
		return err
	}
	for k, v := range objMap {
		(*target)[k] = v
	}
	return nil
}

func FetchFromMap(source *map[string]interface{}, target any) error {
	err := mapToObject(source, target)
	if err == nil {
		targetMap, err := objectToMap(target)
		if err == nil {
			for k := range targetMap {
				delete(*source, k)
			}
		}
	}
	return err
}

func objectToMap(obj any) (map[string]interface{}, error) {
	jsonObj, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]interface{})
	err = json.Unmarshal(jsonObj, &objMap)
	return objMap, err
}

func mapToObject(source *map[string]interface{}, target any) error {
	jsonObj, err := json.Marshal(source)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonObj, target)
	return err
}
