package play_go

import (
	"github.com/stretchr/testify/assert"
	"os/exec"
	"testing"
)

func TestIsValidDockerCommand(t *testing.T) {
	assert := assert.New(t)
	//IsValidDockerCommand(exec.Command("wooooo"))
	assert.False(IsValidDockerCommand(exec.Command("docker", "attack")), "docker attack")
	assert.True(IsValidDockerCommand(exec.Command("docker", "attach")), "docker attach")
	assert.True(IsValidDockerCommand(exec.Command("docker", "exec")), "docker exec")
}
