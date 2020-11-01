package models

// WtsToString transforms struct with sql.NullString values and returns formatted string
// seems like should be replaced by custom marshaling/unmarshaling methods
func (wtl WasteTypeList) WtsToString() string {
	var result string = "[\n"
	for _, item := range wtl {
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
