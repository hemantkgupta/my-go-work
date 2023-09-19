package main

func main() {
	// Both lang and year are newly declared.
	lang, year := "Go language", 2007

	// Only createdBy is a new declared variable.
	// The year variable has already been
	// declared before, so here its value is just
	// modified, or we can say it is redeclared.
	year, createdBy := 2009, "Google Research"

	// This is a pure assignment.
	lang, year = "Go", 2012

	print(lang, " is created by ", createdBy)
	println(", and released at year", year)
}
