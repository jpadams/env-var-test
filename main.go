package main

import (
	"dagger/env/internal/dagger"
)

type EnvVarTest struct{}

// Returns a container with a bunch of env vars
func (m *EnvVarTest) Test(
	env1 string,
	env2 string,
	env3 string,
	env4 string,
	env5 string,
	env6 string,
	env7 string,
	env8 string,
) *dagger.Container {
	return dag.
		Container().
		From("alpine:latest").
		WithEnvVariable("ENV1", env1).
		WithEnvVariable("ENV2", env2).
		WithEnvVariable("ENV3", env3).
		WithEnvVariable("ENV4", env4).
		WithEnvVariable("ENV5", env5).
		WithEnvVariable("ENV6", env6).
		WithEnvVariable("ENV7", env7).
		WithEnvVariable("ENV8", env8)
}
