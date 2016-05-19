package main

import (
	"fmt"
	"bufio"
	"os"
	"bytes"
	"io/ioutil"
	"strings"
	"io"
)

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

func main() {
	fmt.Print("Enter domain: ")

	reader := bufio.NewReader(os.Stdin)
	domain, _ := reader.ReadString('\n')

	domain = strings.Replace(domain, "\n", "", -1)
	app := strings.Replace(domain, ("."), ("_"), -1)

	// Balancer
	os.Mkdir("balancer", 0777)
	nginx, _ := ioutil.ReadFile("src/balancer/nginx.conf.tmpl")
	nginx = bytes.Replace(nginx, []byte("#domain#"), []byte(domain), -1)
	nginx = bytes.Replace(nginx, []byte("#app#"), []byte(app), -1)
	ioutil.WriteFile("balancer/nginx.conf", nginx, 0666)

	CopyFile("balancer/Dockerfile", "src/balancer/Dockerfile")

	// App
	os.Mkdir(domain, 0777)
	os.Mkdir(domain + "/app", 0777)
	CopyFile(domain + "/app/default_server", "src/application/app/default_server")
	CopyFile(domain + "/app/Dockerfile.tmpl", "src/application/app/Dockerfile.tmpl")

	//docker-compose
	compose, _ := ioutil.ReadFile("src/docker-compose.yml")
	compose = bytes.Replace(compose, []byte("#app#"), []byte(app), -1)
	ioutil.WriteFile("docker-compose.yml", compose, 0666)

	//up
	up, _ := ioutil.ReadFile("src/up.sh")
	up = bytes.Replace(up, []byte("#domain#"), []byte(domain), -1)
	up = bytes.Replace(up, []byte("#app#"), []byte(app), -1)
	ioutil.WriteFile("up.sh", up, 0666)
}