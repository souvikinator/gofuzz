package data

//TODO: rename the file, sounds weird
type FuzzData struct {
	AsciiFuzzData []string //from -a
	CharFuzzData  []string //from -c
	NumFuzzData   []string //from -n options
	InputFuzzData []string //from file input
}

//check if we have some data to fuzz?
func (m FuzzData) IsEmpty() bool {
	isEmpty := 0
	if len(m.AsciiFuzzData) == 0 {
		isEmpty += 1
	}
	if len(m.CharFuzzData) == 0 {
		isEmpty += 1
	}
	if len(m.NumFuzzData) == 0 {
		isEmpty += 1
	}
	if len(m.InputFuzzData) == 0 {
		isEmpty += 1
	}
	//=> all are empty
	if isEmpty == 4 {
		return true
	}
	//=> not all are empty
	return false
}
