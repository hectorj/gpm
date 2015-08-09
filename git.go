package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

var gitCommand string

func init() {
	var err error
	gitCommand, err = exec.LookPath("git")
	if err != nil {
		panic(err)
	}
}

func gitClone(targetPath, remoteURI string) ([]byte, error) {
	cmd := exec.Command(gitCommand, "clone", remoteURI, targetPath)
	return cmd.CombinedOutput()
}

func gitAddSubmodule(repoDir, remoteURI, targetPath string) ([]byte, error) {
	cmd := exec.Command(gitCommand, "submodule", "add", "-f", remoteURI, targetPath)
	cmd.Dir = repoDir
	return cmd.CombinedOutput()
}

func gitCheckoutCommit(repoDir, commitHash string) ([]byte, error) {
	cmd := exec.Command(gitCommand, "checkout", commitHash)
	cmd.Dir = repoDir
	return cmd.CombinedOutput()
}

var remoteExtractRegexp = regexp.MustCompile(`^([^\s]+)\s+([^\s]+) \(fetch\)`)

func gitGetRemoteURI(repoDir string, allowLocal bool) (string, error) {
	cmd := exec.Command(gitCommand, "remote", "-v")
	cmd.Dir = repoDir
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	matches := remoteExtractRegexp.FindStringSubmatch(string(output))
	if matches == nil {
		if allowLocal {
			// @TODO : Maybe. Vendoring local repo doesn't actually sound like a good idea, gotta see if there is
			// some real usecases
			panic("Getting local repo URI : not implemented yet")
		}
		err = fmt.Errorf("Could not extract remote URL from %q", repoDir)
		return "", err
	}
	return matches[2], nil
}

func gitGetCurrentCommitHash(repoDir string) (string, error) {
	cmd := exec.Command(gitCommand, "rev-parse", "--verify", "HEAD")
	cmd.Dir = repoDir

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.Trim(string(output), "\n"), nil
}
