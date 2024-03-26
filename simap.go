package ext

// SIKey comprises keys used to access common information from a SIMap.
type SIKey string

// SIMap - String Interface Map.
type SIMap map[SIKey]interface{}

func Pair(fields ...interface{}) SIMap {
	results := make(SIMap)

	end := 0
	if fields != nil {
		end = len(fields)
	}

	for idx := 0; idx < end; {
		var key SIKey

		if k, ok := fields[idx].(SIKey); ok {
			key = k
		} else {
			idx += 2

			continue
		}

		var value interface{}

		if (idx + 1) < end {
			value = fields[idx+1]
		}

		if valueType, ok := value.([]interface{}); ok {
			value = Pair(valueType...)
		}

		results[key] = value

		idx += 2
	}

	return results
}
