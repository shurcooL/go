package vcs

import (
	"os/exec"

	. "gist.github.com/5258650.git"
	. "gist.github.com/5892738.git"
)

func init() {
	if _, err := exec.LookPath("hg"); err == nil {
		addVcsProvider(func(path string) Vcs {
			if isRepo, rootPath := getHgRepoRoot(path); isRepo {
				return &hgVcs{commonVcs{rootPath: rootPath}}
			}
			return nil
		})
	}
}

func getHgRepoRoot(path string) (isHgRepo bool, rootPath string) {
	cmd := exec.Command("hg", "root")
	cmd.Dir = path

	if out, err := cmd.CombinedOutput(); err == nil {
		return true, TrimLastNewline(string(out))
	} else {
		return false, ""
	}
}

type hgVcs struct {
	commonVcs
}

func (this *hgVcs) Type() Type { return Hg }

func (this *hgVcs) GetStatus() string {
	cmd := exec.Command("hg", "status")
	cmd.Dir = this.rootPath

	if out, err := cmd.CombinedOutput(); err == nil {
		return string(out)
	} else {
		return ""
	}
}

func (this *hgVcs) GetDefaultBranch() string {
	return "default"
}

func (this *hgVcs) GetLocalBranch() string {
	cmd := exec.Command("hg", "branch")
	cmd.Dir = this.rootPath

	if out, err := cmd.CombinedOutput(); err == nil {
		return TrimLastNewline(string(out))
	} else {
		return ""
	}
}

// Length of a Mercurial revision hash.
const hgRevisionLength = 40

func (this *hgVcs) GetLocalRev() string {
	// Alternative: hg parent --template '{node}'
	cmd := exec.Command("hg", "--debug", "identify", "-i")
	cmd.Dir = this.rootPath

	if out, err := cmd.CombinedOutput(); err == nil && len(out) >= hgRevisionLength {
		return string(out[:hgRevisionLength])
	} else {
		return ""
	}
}

func (this *hgVcs) GetRemoteRev() string {
	// TODO: Make this more robust and proper, etc.
	cmd := exec.Command("hg", "--debug", "identify", "-i", "default")
	cmd.Dir = this.rootPath

	if out, err := cmd.CombinedOutput(); err == nil {
		// Get the last line of output
		if lines := GetLines(TrimLastNewline(string(out))); len(lines) > 0 {
			return lines[len(lines)-1]
		}
	}
	return ""
}