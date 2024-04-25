package git

import (
    "bytes"
    "os"
    "os/exec"
    "strings"
)

func GetUserName() (string, error) {
    cmd := exec.Command("git", "config", "--get", "user.name")
    var out bytes.Buffer
    cmd.Stdout = &out
    err := cmd.Run()
    if err != nil {
        return "", err
    }
    return strings.TrimSpace(out.String()), nil
}

func SaveUsername(filename string) error {
    username, err := GetUserName()
    if err != nil {
        return err
    }
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    _, err = file.WriteString(username)
    if err != nil {
        return err
    }
    return nil
}
