package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func GetUserEmail() (string, error) {
	cmd := exec.Command("git", "config", "user.email")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("GetUserEmail: Git user.emailni olishda xatolik: %v", err)
	}
	return string(output), nil
}

func WriteFile(filename string, data []byte, perm os.FileMode) error {
	err := ioutil.WriteFile("file1.txt", data, perm)
	if err != nil {
		return fmt.Errorf("WriteFile: Faylga yozishda xatolik: %v", err)
	}
	return nil
}

func main() {
	email, err := GetUserEmail()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = WriteFile("userEmail.txt", []byte(email), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Foydalanuvchi email faylga yozildi.")
}
