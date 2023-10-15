package todo

import "fmt"

type State int

const (
	StateOpen State = iota + 1
	StateClosed
)

var stateToString = map[State]string{
	StateOpen:   "open",
	StateClosed: "closed",
}

var stringToState = func() map[string]State {
	m := map[string]State{}
	for k, v := range stateToString {
		m[v] = k
	}
	return m
}()

func (s State) String() string {
	return stateToString[s]
}

func ParseState(s string) (State, error) {
	if v, ok := stringToState[s]; ok {
		return v, nil
	}
	return 0, fmt.Errorf("invalid state: %s", s)
}
