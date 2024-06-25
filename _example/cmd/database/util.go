package database

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"strings"

	"github.com/canonical/rebac-admin-ui-handlers/v1/resources"
)

var customEncoder = base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!?")

// randomSource we use a static random seed to make the results reproducible for testing purposes.
var randomSource = rand.NewSource(0)

// NewRandomId returns a random 8-char string.
func NewRandomId() string {
	p := make([]byte, 3)
	_, _ = rand.New(randomSource).Read(p)
	return customEncoder.EncodeToString(p)
}

func GetMapValues[K comparable, V any](m map[K]V) []V {
	result := make([]V, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

func EntitlementToString(e resources.EntityEntitlement) string {
	// For example: "can_read::controller:foo"
	return fmt.Sprintf("%s::%s:%s", e.EntitlementType, e.EntityType, e.EntityName)
}

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
