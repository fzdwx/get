package git

import (
	"fmt"
	"github.com/fzdwx/get/pkg/execx"
	"os"
	"os/exec"
)

const cloneUrl = "https://github.com/%s.git"

func Clone(repo string) {
	err := execx.ClassicCmd(exec.Command("git", "clone", fmt.Sprintf(cloneUrl, repo))).Run()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s", err.Error())
	}
}
