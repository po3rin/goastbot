package gendoc

import (
	"os/exec"
	"strings"
)

// GenDoc generate doc.
func GenDoc(arg string) (string, error) {
	s := []string{"doc"}
	args := strings.Split(arg, " ")
	for _, arg := range args {
		s = append(s, arg)
	}
	got, err := exec.Command("go", s...).Output()
	if err != nil {
		return "", err
	}
	out := string(got)
	return out, nil
}
