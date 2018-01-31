package goap

type State map[string]bool

func (s *State) Contains(otherState State) bool {
	for key, value := range otherState {
		ourValue, ok := (*s)[key]
		if ok == false || value != ourValue {
			return false
		}
	}
	return true
}

func (s *State) Update(otherState State) {
	for key, value := range otherState {
		(*s)[key] = value
	}
}
