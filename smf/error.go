package smf

//MidiError error type for midi package //就是定义一个自己的midierror
type MidiError struct {
	errorString string
}

// Error implements standart error interface
func (m *MidiError) Error() string {
	return "Midi error: " + m.errorString
}
