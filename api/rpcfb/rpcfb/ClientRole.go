// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package rpcfb

import "strconv"

type ClientRole int8

const (
	ClientRoleCLIENT_ROLE_UNKNOWN   ClientRole = 0
	ClientRoleCLIENT_ROLE_PM        ClientRole = 1
	ClientRoleCLIENT_ROLE_DATA_NODE ClientRole = 2
	ClientRoleCLIENT_ROLE_CUSTOMER  ClientRole = 3
)

var EnumNamesClientRole = map[ClientRole]string{
	ClientRoleCLIENT_ROLE_UNKNOWN:   "CLIENT_ROLE_UNKNOWN",
	ClientRoleCLIENT_ROLE_PM:        "CLIENT_ROLE_PM",
	ClientRoleCLIENT_ROLE_DATA_NODE: "CLIENT_ROLE_DATA_NODE",
	ClientRoleCLIENT_ROLE_CUSTOMER:  "CLIENT_ROLE_CUSTOMER",
}

var EnumValuesClientRole = map[string]ClientRole{
	"CLIENT_ROLE_UNKNOWN":   ClientRoleCLIENT_ROLE_UNKNOWN,
	"CLIENT_ROLE_PM":        ClientRoleCLIENT_ROLE_PM,
	"CLIENT_ROLE_DATA_NODE": ClientRoleCLIENT_ROLE_DATA_NODE,
	"CLIENT_ROLE_CUSTOMER":  ClientRoleCLIENT_ROLE_CUSTOMER,
}

func (v ClientRole) String() string {
	if s, ok := EnumNamesClientRole[v]; ok {
		return s
	}
	return "ClientRole(" + strconv.FormatInt(int64(v), 10) + ")"
}
