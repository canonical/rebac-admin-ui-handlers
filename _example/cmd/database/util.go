package database

import (
	"fmt"
	"strings"

	"github.com/canonical/rebac-admin-ui-handlers/v1/resources"
)

// EntitlementToString marshals a given entitlement as a string.
func EntitlementToString(e resources.EntityEntitlement) string {
	// For example: "can_read::controller:foo"
	return fmt.Sprintf("%s::%s:%s", e.EntitlementType, e.EntityType, e.EntityName)
}

// EntitlementFromString unmarshals a entitlement from a given string.
func EntitlementFromString(s string) resources.EntityEntitlement {
	parts := strings.SplitN(s, ":", 4)
	result := resources.EntityEntitlement{}
	if len(parts) > 0 {
		result.EntitlementType = parts[0]
	}
	if len(parts) > 1 {
		result.EntityType = parts[1]
	}
	if len(parts) > 2 {
		result.EntityName = parts[2]
	}
	return result
}

func getMapValues[K comparable, V any](m map[K]V) []V {
	result := make([]V, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}
	return result
}
