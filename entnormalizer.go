package main

func stdType2EntType(stdType string) string {
	switch stdType {
	case "String":
		return "String"
	case "Int":
		return "Int"
	case "Float":
		return "Float"
	case "Map":
		return "JSON"
	case "File":
		return "String"
	default:
		return ""
	}
}
