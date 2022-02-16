package main

import (
	"io/ioutil"
	"os/exec"
	"strings"
)

const rootDirPathDefault = "/home/bitrix/www"
const rootDirPath = "/home/bitrix/ext_www/bitrix.rzdba.ru"

func main() {

	replaceStringInFile("/etc/httpd/bx/conf/default.conf", rootDirPathDefault, rootDirPath)
	replaceStringInFile("/etc/nginx/bx/site_avaliable/ssl.s1.conf", rootDirPathDefault, rootDirPath)

	err := exec.Command("/usr/sbin/nginx", "-s", "reload").Run()
	if err != nil {
		panic(err)
	}
	err = exec.Command("/usr/sbin/httpd", "-k", "graceful").Run()
	if err != nil {
		panic(err)
	}
}

func replaceStringInFile(filePath string, old string, new string) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")

	for i, line := range lines {
		if strings.Contains(line, old) {
			lines[i] = strings.Replace(line, old, new, -1)
		}
	}

	output := strings.Join(lines, "\n")

	err = ioutil.WriteFile(filePath, []byte(output), 644)
	if err != nil {
		panic(err)
	}
}
