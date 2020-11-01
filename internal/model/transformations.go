package models

func WtsToString(wts []WasteType) string {
	var result string = "[\n"
	for _, item := range wts {
		result += "\t{\"ID\": \"" + item.ID.String + "\";" +
			" \"Name\": \"" + item.Name.String + "\""
		if item.Description.Valid {
			result += "; \"Description\": \"" + item.Description.String + "\""
		}
		result += "}\n"
	}
	result += "]"
	return result
}
