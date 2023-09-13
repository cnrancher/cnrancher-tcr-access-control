package policystatus

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Status struct {
	Status   string           `json:"status" yaml:"status"`
	Reason   string           `json:"reason" yaml:"reason"`
	Policies []SecurityPolicy `json:"policies" yaml:"policies"`
}

type SecurityPolicy struct {
	Index       int64  `json:"index" yaml:"index"`
	CIDR        string `json:"cidr" yaml:"cidr"`
	Description string `json:"description" yaml:"description"`
}

func (s *Status) String() string {
	if s == nil {
		return "<nil>"
	}

	b := &strings.Builder{}
	fmt.Fprintf(b, "External Endpoint (公网访问入口) Status: %v\n", s.Status)
	if s.Reason != "" {
		fmt.Fprintf(b, "Reason: %v\n", s.Reason)
	}
	fmt.Fprintf(b, "\n")

	if len(s.Policies) == 0 {
		fmt.Fprintf(b, "Security Policy is empty.\n")
		return b.String()
	}

	fmt.Fprintf(b, "Security Policies:\n")
	fmt.Fprintf(b, "-------+--------------------+--------------\n")
	fmt.Fprintf(b, " INDEX |        CIDR        |  Description\n")
	fmt.Fprintf(b, "-------+--------------------+--------------\n")
	for _, p := range s.Policies {
		fmt.Fprintf(b, " %5d | %18s | %v \n",
			p.Index, p.CIDR, p.Description)
	}
	fmt.Fprintf(b, "-------+--------------------+--------------\n")

	return b.String()
}

func (s *Status) Json() string {
	d, _ := json.MarshalIndent(s, "", "  ")

	return string(d)
}
