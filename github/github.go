package github

import (
	"fmt"

	"janjas/atomatiki/config"

	"log"
)


func Login(ghToken string) {

	bashCmd := loginBash(ghToken)

	cmd, err, f := config.BashCmdExec(bashCmd)
	if err != nil {
		log.Fatal(err)
	}
	if f != nil {
		config.ClearTempFile(f)
	}

	output, err := cmd.Output()
	if err != nil  {
		log.Fatal(err)
	}

	log.Print(string(output))
}


func startBash() (string) {
	// if runtime.GOOS == "windows" {
	return "#!/bin/sh"
}


func loginBash(ghToken string) (string) {
	return fmt.Sprintf(`%s
echo %s > .token && gh auth login --with-token < "./.token" && rm ./.token
	`, startBash(), ghToken)
}


func CreateRepoBash(repoName string, ghToken string) string {
	return fmt.Sprintf(`%s
gh repo create "%s" --private --confirm`, loginBash(ghToken), repoName)
}


func CreateRepoSecretsBash(secrets map[string]string ,repoName string, ghToken string) string {
	
	cmdString := ""
	for key, value := range secrets {
		cmdString += fmt.Sprintf(`gh secret set %s --body '%s' --repo %s \n`, key, value,  repoName)
	}

	return fmt.Sprintf(`%s\n%s`, loginBash(ghToken), cmdString)
	
}


func InitializePushGitBash(localRepoPath string, repoName string, ghToken string) string {
	return fmt.Sprintf(`
%s
cd %s && \
echo %s > README.md && \
git init && git add README.md && git add . && \
git commit -m "first commit" && \
git branch -M main && \
git remote add origin https://github.com/%s.git && \
git push -u origin main && \
git checkout -b develop main && \
git push -u origin develop
	`, loginBash(ghToken), localRepoPath, repoName, repoName)
}



