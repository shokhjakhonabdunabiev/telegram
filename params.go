package telegram

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type Params map[string]string

func (p Params) AddFirstValid(key string, args ...any) error {
	for _, arg := range args {
		switch v := arg.(type) {
		case int:
			if v != 0 {
				p[key] = strconv.Itoa(v)
				return nil
			}
		case int64:
			if v != 0 {
				p[key] = strconv.FormatInt(v, 10)
				return nil
			}
		case string:
			if v != "" {
				p[key] = v
				return nil
			}
		case nil:
		default:
			b, err := json.Marshal(arg)
			if err != nil {
				return err
			}
			p[key] = string(b)
			return nil
		}
	}

	return nil
}

func structToQuery(params any) (url.Values, error) {
	values := url.Values{}

	if params == nil {
		return values, nil
	}

	b, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	var m map[string]any
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}

	for k, v := range m {
		switch val := v.(type) {
		case string:
			values.Set(k, val)
		case float64:
			values.Set(k, fmt.Sprintf("%.0f", val))
		case bool:
			values.Set(k, fmt.Sprintf("%t", val))
		default:
			nested, err := json.Marshal(val)
			if err != nil {
				return nil, err
			}
			values.Set(k, string(nested))
		}
	}

	return values, nil
}
