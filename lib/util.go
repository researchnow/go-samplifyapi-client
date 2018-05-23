package samplify

import "os"

func getEnvironment() string {
	env := os.Getenv("env")
	if len(env) == 0 {
		return "uat"
	}
	return env
}

func isProdEnv() bool {
	return getEnvironment() == "prod"
}
