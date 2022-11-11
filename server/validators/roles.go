package validators

import "github.com/ArkamFahry/GateGuardian/server/utils"

// IsValidRoles validates roles
func IsValidRoles(defaultRoles []string, allowedRoles []string) bool {
	valid := true
	for _, allowedRoles := range allowedRoles {
		if !utils.StringSliceContains(defaultRoles, allowedRoles) {
			valid = false
			break
		}
	}

	return valid
}
