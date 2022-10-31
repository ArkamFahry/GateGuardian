package validators

import "github.com/ArkamFahry/GateGuardian/server/utils"

func IsValidRoles(userRoles []string, roles []string) bool {
	valid := true
	for _, userRole := range userRoles {
		if !utils.StringSliceContains(roles, userRole) {
			valid = false
			break
		}
	}

	return valid
}
