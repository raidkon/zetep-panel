package i18n

func mergeMaps(base map[string]string, over map[string]string) map[string]string {
	if len(over) == 0 {
		out := make(map[string]string, len(base))
		for k, v := range base {
			out[k] = v
		}
		return out
	}
	out := make(map[string]string, len(base)+len(over))
	for k, v := range base {
		out[k] = v
	}
	for k, v := range over {
		out[k] = v
	}
	return out
}
