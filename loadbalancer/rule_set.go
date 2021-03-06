// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// RuleSet A named set of rules associated with a load balancer. Rules are objects that represent actions to apply to a listener,
// such as adding, altering, or removing HTTP headers.
type RuleSet struct {

	// The name for this set of rules. It must be unique and it cannot be changed. Avoid entering
	// confidential information.
	// Example: `example_rule_set`
	Name *string `mandatory:"true" json:"name"`

	// An array of rules that compose the rule set.
	Items []Rule `mandatory:"true" json:"items"`
}

func (m RuleSet) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *RuleSet) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Name  *string `json:"name"`
		Items []rule  `json:"items"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	m.Name = model.Name
	m.Items = make([]Rule, len(model.Items))
	for i, n := range model.Items {
		nn, err := n.UnmarshalPolymorphicJSON(n.JsonData)
		if err != nil {
			return err
		}
		if nn != nil {
			m.Items[i] = nn.(Rule)
		} else {
			m.Items[i] = nil
		}
	}
	return
}
