package main

import (
	"os"
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	url := "https://github.com/xingyizhou/CenterNet.git"
	ramp_up_score(url)
	fmt.Println("TESTING")
}

func Test2(t *testing.T) {
	var personal_token string
	personal_token = os.Getenv("GITHUB_TOKEN")
	owner := "xingyizhou"
	repo := "CenterNet"
	correctness_score(personal_token, owner, repo)
	fmt.Println("TESTING")
}

func Test3(t *testing.T) {
	var personal_token string
	personal_token = os.Getenv("GITHUB_TOKEN")
	owner := "xingyizhou"
	repo := "CenterNet"
	responseviness_score(personal_token, owner, repo)
	fmt.Println("TESTING")
}

func Test4(t *testing.T) {
	var personal_token string
	personal_token = os.Getenv("GITHUB_TOKEN")
	owner := "xingyizhou"
	repo := "CenterNet"
	bus_factor_score(personal_token, owner, repo)
	fmt.Println("TESTING")
}

func Test5(t *testing.T) {
	var personal_token string
	personal_token = os.Getenv("GITHUB_TOKEN")
	owner := "xingyizhou"
	repo := "CenterNet"
	license_score(personal_token, owner, repo)
	fmt.Println("TESTING")
}

func Test6(t *testing.T) {
	get_git_url("https://www.npmjs.com/package/express")
	fmt.Println("TESTING")
}

func Test7(t *testing.T) {
	analyze_git("https://github.com/xingyizhou/CenterNet.git", "https://github.com/xingyizhou/CenterNet.git")
	fmt.Println("TESTING")
}

func Test8(t *testing.T) {
	url := "https://github.com/hwholiday/short_url.git"
	ramp_up_score(url)
	fmt.Println("TESTING")
}

func Test9(t *testing.T) {
	var personal_token string
	personal_token = os.Getenv("GITHUB_TOKEN")
	owner := "hwholiday"
	repo := "short_url"
	correctness_score(personal_token, owner, repo)
	fmt.Println("TESTING")
}

func Test10(t *testing.T) {
	var personal_token string
	personal_token = os.Getenv("GITHUB_TOKEN")
	owner := "hwholiday"
	repo := "short_url"
	responseviness_score(personal_token, owner, repo)
	fmt.Println("TESTING")
}

func Test11(t *testing.T) {
	var personal_token string
	personal_token = os.Getenv("GITHUB_TOKEN")
	owner := "hwholiday"
	repo := "short_url"
	bus_factor_score(personal_token, owner, repo)
	fmt.Println("TESTING")
}

func Test12(t *testing.T) {
	var personal_token string
	personal_token = os.Getenv("GITHUB_TOKEN")
	owner := "hwholiday"
	repo := "short_url"
	license_score(personal_token, owner, repo)
	fmt.Println("TESTING")
}

func Test13(t *testing.T) {
	url := "https://github.com/Caturra000/co.git"
	ramp_up_score(url)
	fmt.Println("TESTING")
}

func Test14(t *testing.T) {
	var personal_token string
	personal_token = os.Getenv("GITHUB_TOKEN")
	owner := "Caturra000"
	repo := "co"
	correctness_score(personal_token, owner, repo)
	fmt.Println("TESTING")
}

func Test15(t *testing.T) {
	var personal_token string
	personal_token = os.Getenv("GITHUB_TOKEN")
	owner := "Caturra000"
	repo := "co"
	responseviness_score(personal_token, owner, repo)
	fmt.Println("TESTING")
}

func Test16(t *testing.T) {
	var personal_token string
	personal_token = os.Getenv("GITHUB_TOKEN")
	owner := "Caturra000"
	repo := "co"
	bus_factor_score(personal_token, owner, repo)
	fmt.Println("TESTING")
}

func Test17(t *testing.T) {
	var personal_token string
	personal_token = os.Getenv("GITHUB_TOKEN")
	owner := "Caturra000"
	repo := "co"
	license_score(personal_token, owner, repo)
	fmt.Println("TESTING")
}

func Test18(t *testing.T) {
	get_git_url("https://www.npmjs.com/package/browserify")
	fmt.Println("TESTING")
}

func Test19(t *testing.T) {
	analyze_git("https://github.com/Caturra000/co.git", "https://github.com/Caturra000/co.git")
	fmt.Println("TESTING")
}

func Test20(t *testing.T) {
	calc_score("./test_file.txt")
	fmt.Println("TESTING")
}