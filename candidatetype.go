package ice

// CandidateType represents the type of candidate
type CandidateType byte

// CandidateType enum
const (
	CandidateTypeUnspecified CandidateType = iota
	CandidateTypeHost
	CandidateTypeServerReflexive
	CandidateTypePeerReflexive
	CandidateTypeRelay
)

// String makes CandidateType printable
func (c CandidateType) String() string {
	switch c {
	case CandidateTypeHost:
		return "host"
	case CandidateTypeServerReflexive:
		return "srflx"
	case CandidateTypePeerReflexive:
		return "prflx"
	case CandidateTypeRelay:
		return "relay"
	}
	return "Unknown candidate type"
}

func ToCandidateType(value string) CandidateType {
	switch value {
	case "host":
		return CandidateTypeHost
	case "srflx":
		return CandidateTypeServerReflexive
	case "prflx":
		return CandidateTypePeerReflexive
	case "relay":
		return CandidateTypeRelay
	}
	return CandidateTypeUnspecified
}

// Preference returns the preference weight of a CandidateType
//
// 4.1.2.2.  Guidelines for Choosing Type and Local Preferences
// The RECOMMENDED values are 126 for host candidates, 100
// for server reflexive candidates, 110 for peer reflexive candidates,
// and 0 for relayed candidates.
func (c CandidateType) Preference() uint16 {
	switch c {
	case CandidateTypeHost:
		return 126
	case CandidateTypePeerReflexive:
		return 110
	case CandidateTypeServerReflexive:
		return 100
	}
	return 0
}

func containsCandidateType(candidateType CandidateType, candidateTypeList []CandidateType) bool {
	if candidateTypeList == nil {
		return false
	}
	for _, ct := range candidateTypeList {
		if ct == candidateType {
			return true
		}
	}
	return false
}
