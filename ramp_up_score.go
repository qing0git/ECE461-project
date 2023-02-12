package main

import (
    "fmt"
	"regexp"
	"strconv"
	"os"
    "os/exec"
    "log"
    "github.com/go-git/go-git/v5"
)

func ramp_up_score(url string) (float64, string, string) {
    // create a temp directory and clone the repository to local
	path := "./tmp"
    _, err := os.Stat(path)
    if err == nil {
        err = os.RemoveAll(path)
        if err != nil {
            log.Println(err)
            return 0, "", ""
        }
    }
    _, err = exec.Command("mkdir", path).Output()
    if err != nil {
        log.Println(err)
        return 0, "", ""
    }
    _, err = git.PlainClone(path, false, &git.CloneOptions{
        URL:               url,
        RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
    })
    if err != nil {
        log.Println(err)
        return 0, "", ""
    }

    // Use cloc to find the total lines of code
    command_output, err := exec.Command("./cloc", path).Output()
    if err != nil {
        log.Println(err)
        return 0, "", ""
    }
	re_sum, _ := regexp.Compile("SUM:\\s*\\d+\\s*\\d+\\s*\\d+\\s*\\d+")
    match_sum := re_sum.FindString(string(command_output))
	re_num_code, _ := regexp.Compile("\\d+")
	num_code_comment := re_num_code.FindAllString(match_sum, -1)
    sugar_logger.Debugf("Number of sloc: " + num_code_comment[len(num_code_comment) - 1])
    sugar_logger.Debugf("Number of lines of comments: " + num_code_comment[len(num_code_comment) - 2])
	num_code, _ := strconv.ParseFloat(num_code_comment[len(num_code_comment) - 1], 64)
	num_comment, _ := strconv.ParseFloat(num_code_comment[len(num_code_comment) - 2], 64)
	code_comment_ratio := 4.0 / ((num_code + 1) / (num_comment + 1))
    sugar_logger.Debugf("Ratio of code to comments: %.1f", code_comment_ratio)
	score, err := 1.0, nil
	if code_comment_ratio < 1 {
		score, _ = strconv.ParseFloat(fmt.Sprintf("%.1f", code_comment_ratio), 64)
	}
	err = os.RemoveAll(path)
    if err != nil {
        log.Println(err)
        return 0, "", ""
    }
	// parse user and repo
	re_user_repo, _ := regexp.Compile("/\\w+/\\w+.git")
    user_repo_raw := re_user_repo.FindString(url)
	re_user_repo, _ = regexp.Compile("\\w+")
	user_repo := re_user_repo.FindAllString(user_repo_raw, -1)
	return score, user_repo[0], user_repo[1]
}