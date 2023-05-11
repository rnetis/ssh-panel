package main

import (
	"log"
	"fmt"
	"os/exec"
	"errors"
)

func sCommand(s string) (string, error) {
	out, err := exec.Command("sh", "-c", s).Output()
	if err != nil {
		return "", err
	}
	return string(out[:]), nil
}

func actionUser(action string, data ...interface{}) (bool, error) {

	switch action {
		case "new":
			//check Existanse on database also send notification to phone
			//add new user to database
			command := fmt.Sprintf("add %s %s %d", data[0]/*username*/, data[1]/*password*/, data[2]/*expiration date by number of days*/)
			args := strings.Split(command)
		case "delete":
			//check existance on database
			//delete user from database
			command := fmt.Sprintf("del %s", data[0]/*username*/)
			args := strings.Split(command)
		case "lock":
			command := fmt.Sprintf("lock %s", data[0]/*username*/)
			args := strings.Split(command)
		case "unlock":
			command := fmt.Sprintf("unlock %s", data[0]/*username*/)
			args := strings.Split(command)
		case "chpasswd":
			//check existance on database also update database
			command := fmt.Sprintf("passwd %s %s", data[1]/*New Password*/, data[0]/*Username*/)
			args := strings.Split(command)
		case "extend":
			//check for existance and update the database
			command := fmt.Sprintf("extendExp %d %s", data[1]/*Number of day for extending exiration*/, data[0]/*Username*/)
			args := strings.Split(command)
		default:
			return false, errors.New("Invalid Action")
	}
	_ , err := exec.Command("mngusers", args...).Output()
	if err != nil {
		return false, err
	}
	return true, nil
}