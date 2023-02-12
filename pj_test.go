package main

import (
	"os"
	"fmt"
	"testing"
)

func Test_ramp_up_score1(t *testing.T) {
	url := "https://github.com/xingyizhou/CenterNet.git"
	ramp_up_score(url)
	fmt.Println("TESTING")
}

func Test_ramp_up_score2(t *testing.T) {
	url := "https://github.com/hwholiday/short_url.git"
	ramp_up_score(url)
	fmt.Println("TESTING")
}

func Test_ramp_up_score3(t *testing.T) {
	url := "https://github.com/Caturra000/co.git"
	ramp_up_score(url)
	fmt.Println("TESTING")
}

func Test_correctness_score1(t *testing.T) {
	var personal_token string
	personal_token = os.Getenv("GITHUB_TOKEN")
	owner := "xingyizhou"
	repo := "CenterNet"
	correctness_score(personal_token, owner, repo)
	fmt.Println("TESTING")
}

func Test_correctness_score2(t *testing.T) {
	var personal_token string
	personal_token = os.Getenv("GITHUB_TOKEN")
	owner := "hwholiday"
	repo := "short_url"
	correctness_score(personal_token, owner, repo)
	fmt.Println("TESTING")
}

func Test_correctness_score3(t *testing.T) {
	var personal_token string
	personal_token = os.Getenv("GITHUB_TOKEN")
	owner := "Caturra000"
	repo := "co"
	correctness_score(personal_token, owner, repo)
	fmt.Println("TESTING")
}

func Test_responseviness_score1(t *testing.T) {
	var personal_token string
	personal_token = os.Getenv("GITHUB_TOKEN")
	owner := "xingyizhou"
	repo := "CenterNet"
	responseviness_score(personal_token, owner, repo)
	fmt.Println("TESTING")
}

func Test_responseviness_score2(t *testing.T) {
	var personal_token string
	personal_token = os.Getenv("GITHUB_TOKEN")
	owner := "hwholiday"
	repo := "short_url"
	responseviness_score(personal_token, owner, repo)
	fmt.Println("TESTING")
}

func Test_responseviness_score3(t *testing.T) {
	var personal_token string
	personal_token = os.Getenv("GITHUB_TOKEN")
	owner := "Caturra000"
	repo := "co"
	responseviness_score(personal_token, owner, repo)
	fmt.Println("TESTING")
}

func Test_bus_factor_score1(t *testing.T) {
	var personal_token string
	personal_token = os.Getenv("GITHUB_TOKEN")
	owner := "xingyizhou"
	repo := "CenterNet"
	bus_factor_score(personal_token, owner, repo)
	fmt.Println("TESTING")
}

func Test_bus_factor_score2(t *testing.T) {
	var personal_token string
	personal_token = os.Getenv("GITHUB_TOKEN")
	owner := "hwholiday"
	repo := "short_url"
	bus_factor_score(personal_token, owner, repo)
	fmt.Println("TESTING")
}

func Test_bus_factor_score3(t *testing.T) {
	var personal_token string
	personal_token = os.Getenv("GITHUB_TOKEN")
	owner := "Caturra000"
	repo := "co"
	bus_factor_score(personal_token, owner, repo)
	fmt.Println("TESTING")
}

func Test_license_score1(t *testing.T) {
	var personal_token string
	personal_token = os.Getenv("GITHUB_TOKEN")
	owner := "xingyizhou"
	repo := "CenterNet"
	license_score(personal_token, owner, repo)
	fmt.Println("TESTING")
}

func Test_license_score2(t *testing.T) {
	var personal_token string
	personal_token = os.Getenv("GITHUB_TOKEN")
	owner := "hwholiday"
	repo := "short_url"
	license_score(personal_token, owner, repo)
	fmt.Println("TESTING")
}

func Test_license_score3(t *testing.T) {
	var personal_token string
	personal_token = os.Getenv("GITHUB_TOKEN")
	owner := "Caturra000"
	repo := "co"
	license_score(personal_token, owner, repo)
	fmt.Println("TESTING")
}

func Test_get_git_url_empty(t *testing.T) {
	get_git_url("")
	fmt.Println("TESTING")
}

func Test_get_git_url1(t *testing.T) {
	get_git_url("https://www.npmjs.com/package/express")
	fmt.Println("TESTING")
}

func Test_analyze_git_empty(t *testing.T) {
	analyze_git("", "")
	fmt.Println("TESTING")
}

func Test_analyze_git1(t *testing.T) {
	analyze_git("https://github.com/xingyizhou/CenterNet.git", "https://github.com/xingyizhou/CenterNet.git")
	fmt.Println("TESTING")
}

func Test_entire_program(t *testing.T) {
	calc_score("./test_file.txt")
	fmt.Println("TESTING")
}