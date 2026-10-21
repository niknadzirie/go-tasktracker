package model

import "encoding/json"

type Status int

const (
	StatusUnknown Status = iota
	StatusPending
	StatusInProgress
	StatusCompleted
)

var statusStrings = map[Status]string{
	StatusPending:    "pending",
	StatusInProgress: "in-progress",
	StatusCompleted:  "completed",
}

var statusValues = map[string]Status{
	"pending":     StatusPending,
	"in-progress": StatusInProgress,
	"completed":   StatusCompleted,
}

func (s Status) String() string {
	if str, ok := statusStrings[s]; ok {
		return str
	}
	return "unknown"
}

func (s Status) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s *Status) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}

	val, ok := statusValues[str]
	if !ok {
		*s = StatusUnknown
		return nil
	}
	*s = val
	return nil
}
