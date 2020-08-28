package main

import "os/exec"

func Samba() {
	cmd := exec.Command("mkdir", "/home/data")
	err := cmd.Run()
	if err != nil {
		println(err.Error())
	}
	cmd = exec.Command("mount", "-t", "cifs", "-v", "-o", "username=zapsi,password=Jahoda123,domain=ROMPACZ", "//10.60.1.9/Interface/Zapsi2LN", "/home/data")
	err = cmd.Run()
	if err != nil {
		println(err.Error())
	}
}
