package keycodes

import (
	"fmt"
	"strconv"
)

const (
	KeycodeStatusUnknown      = 0
	KeycodeStatusNoKeycode    = 1
	KeycodeStatusInvalid      = 2
	KeycodeStatusExpired      = 3
	KeycodeStatusNotActivated = 4
	KeycodeStatusValid        = 5

	MessageCPUCoreExceed = "Cluster CPU cores (%s) exceed the licensed number of CPU cores (%s)"
)

var KeycodeStatusName = map[int]string{
	0: "Unknown",
	1: "No Keycode",
	2: "Invalid",
	3: "Expired",
	4: "Not Activated",
	5: "Valid",
}

var KeycodeStatusMessage = map[int]string{
	0: "Unknown keycode is detected",
	1: "No keycode is applied",
	2: "Invalid keycode is detected",
	3: "Keycode is expired",
	4: "Keycode is not activated",
	5: "A valid keycode is applied",
}

type KeycodeStatusObject struct {
	reason string
}

func NewKeycodeStatusObject() *KeycodeStatusObject {
	keycodeStatusObject := KeycodeStatusObject{}
	return &keycodeStatusObject
}

func (c *KeycodeStatusObject) GetStatus() int {
	c.reason = ""
	if c.isNoKeycode() {
		return KeycodeStatusNoKeycode
	}
	if c.isInvalid() {
		return KeycodeStatusInvalid
	}
	if c.isExpired() {
		return KeycodeStatusExpired
	}
	if c.isCapacityCPUCoreExceed() {
		return KeycodeStatusInvalid
	}
	if c.isNotActivated() {
		return KeycodeStatusNotActivated
	}
	if c.isValid() {
		return KeycodeStatusValid
	}
	return KeycodeStatusUnknown
}

func (c *KeycodeStatusObject) GetReason() string {
	return c.reason
}

func (c *KeycodeStatusObject) isNoKeycode() bool {
	if KeycodeList == nil || len(KeycodeList) == 0 {
		return true
	}
	return false
}

func (c *KeycodeStatusObject) isInvalid() bool {
	if KeycodeSummary != nil {
		if KeycodeSummary.LicenseState == "Invalid" {
			return true
		}
	}
	return false
}

func (c *KeycodeStatusObject) isExpired() bool {
	if KeycodeSummary != nil {
		if KeycodeSummary.LicenseState == "Expired" {
			return true
		}
	}
	return false
}

func (c *KeycodeStatusObject) isNotActivated() bool {
	if KeycodeSummary != nil {
		if KeycodeSummary.KeycodeType == "Regular" && KeycodeSummary.LicenseState == "Valid" {
			if KeycodeSummary.Registered == false {
				return true
			}
		}
	}
	return false
}

func (c *KeycodeStatusObject) isValid() bool {
	if KeycodeSummary != nil {
		if KeycodeSummary.LicenseState == "Valid" {
			return true
		}
	}
	return false
}

func (c *KeycodeStatusObject) isCapacityCPUCoreExceed() bool {
	if KeycodeSummary.Capacity.CPUs >= 0 {
		if KeycodeSummary.Capacity.CPUs < ClusterCPUCores {
			c.reason = fmt.Sprintf(MessageCPUCoreExceed, strconv.Itoa(ClusterCPUCores),
				strconv.Itoa(KeycodeSummary.Capacity.CPUs))
			scope.Error(c.reason)
			return true
		}
	}
	return false
}
