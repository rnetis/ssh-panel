package main

import (
	"log"
	"os/exec"
)

func sCommand(s string) (string, error) {
	out, err := exec.Command("sh", "-c", s).Output()
	if err != nil {
		return "", err
	}
	return string(out[:]), nil
}

func userCreate(name string, password string, duration int) bool {
	command := `useradd -M -s /usr/sbin/nologin -p $(echo ` + password + ` | openssl passwd -1 -stdin) `
	if duration != 0 {
		command += name
	} else {
		command += `-e $(date --date="+` + string(duration) + ` month" +"%F") ` + name
	}
	out, err := sCommand(command)
	if err != nil {
		log.Println("Error on Creating User:", err)
		return false
	}
	log.Println(string(out[:]))
	return true
}

func userDelete(name string) bool {
	command := `deluser ` + name
	_, err := sCommand(command)
	if err != nil {
		log.Println("Error on Deleting User:", err)
		return false
	}
	return true
}

func userRead(name string) (string, error) {
	out, err := sCommand(`chage -l ` + name + ` | grep 'Account expires'`)
	if err != nil {
		log.Println("Error on Reading User:", err)
		return "", err
	}
	lastLogin, err := sCommand(`cat /var/log/auth.log | awk '/Accepted password/ {print}' | grep ` + name + ` | tail -n 1`)
	if err != nil {
		log.Println("Error on Reading User:", err)
		return "", err
	}
	return out + "\n" + lastLogin, nil
}

func main() {
	userCreate("goTest", "1029384756", 1)
	getUser, err := userRead("goTest")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(getUser)
	userDelete("goTest")
}
