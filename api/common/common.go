package common

import "reflect"

func JsonSuccess(message string, metadata interface{}, data ...interface{}) interface{} {
	return jsonSuccess(message, metadata, data)
}

func jsonSuccess(message string, metadata interface{}, data []interface{}) interface{} {
	obj := map[string]interface{}{
		"success": true,
		"code":    200,
		"message": message,
	}
	if len(data) > 0 && data[0] != nil {
		if reflect.TypeOf(data[0]).Kind() == reflect.Slice {
			if reflect.ValueOf(data[0]).IsNil() {
				obj["data"] = []interface{}{}
			} else {
				obj["data"] = data[0]
			}
		} else {
			obj["data"] = []interface{}{data[0]}
		}
	} else {
		obj["data"] = []string{}
	}
	if metadata != nil {
		obj["metadata"] = metadata
	}

	return obj
}

func JsonError(code int, message string) interface{} {
	return jsonError(code, []string{message})
}

func jsonError(code int, messages []string) interface{} {
	obj := map[string]interface{}{
		"success": false,
		"code":    code,
		"errors":  messages,
	}

	return obj
}
