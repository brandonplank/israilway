package israilway

import "os"

func IsRailway() bool {
	vars := []string{"RAILWAY_STATIC_URL", "RAILWAY_ENVIRONMENT", "RAILWAY_HEALTHCHECK_TIMEOUT_SEC", "RAILWAY_GIT_COMMIT_SHA", "RAILWAY_GIT_AUTHOR", "RAILWAY_GIT_BRANCH", "RAILWAY_GIT_REPO_NAME", "RAILWAY_GIT_REPO_OWNER", "RAILWAY_GIT_COMMIT_MESSAGE"}
	for _, env := range vars {
		_, hasVar := os.LookupEnv(env)
		if hasVar {
			return true
		}
	}
	return false
}
