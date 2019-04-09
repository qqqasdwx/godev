package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}
	switch os.Args[1] {
	case "init":
		devInit()
	case "up":
		up()
	case "down":
		down()
	case "cli":
		cli()
	default:
		usage()
	}
}

func cli() {
	cmd := exec.Command("docker-compose", "exec", "workspace", "sh")
	cmd.Dir = fmt.Sprintf("%s/%s", getCurrentAbsPath(), "workspace")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func down() {
	cmd := exec.Command("docker-compose", "down")
	cmd.Dir = fmt.Sprintf("%s/%s", getCurrentAbsPath(), "workspace")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	// stdout, err := cmd.StdoutPipe()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// cmd.Start()

	// reader := bufio.NewReader(stdout)
	// for {
	// 	line, err := reader.ReadString('\n')
	// 	if err != nil || io.EOF == err {
	// 		break
	// 	}
	// 	fmt.Print(line)
	// }

	// cmd.Wait()
}

func up() {
	cmd := exec.Command("docker-compose", "up", "-d")
	cmd.Dir = fmt.Sprintf("%s/%s", getCurrentAbsPath(), "workspace")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)

	}

	cmd.Start()

	reader := bufio.NewReader(stdout)
	for {
		line, err := reader.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		fmt.Print(line)
	}

	cmd.Wait()
}

func usage() {
	fmt.Println("Usage: godev [init] [up] [down] [cli]")
}

func devInit() {
	if _, err := os.Stat(path.Join(".", "workspace")); !os.IsNotExist(err) {
		fmt.Println("please remove your workspace dir!")
		return
	}
	if err := os.MkdirAll(path.Join(".", "workspace"), 0755); err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("./workspace/.env", []byte(strings.Replace(httpGet("file/env"), "$YOUR_PROJECT_NAME", getCurrentPath(), -1)), 0755); err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("./workspace/Dockerfile", []byte(httpGet("file/Dockerfile")), 0755); err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("./workspace/docker-compose.yml", []byte(httpGet("file/docker-compose.yml")), 0755); err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("./workspace/bashrc", []byte(httpGet("file/bashrc")), 0755); err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("./Makefile", []byte(strings.Replace(
		httpGet("Makefile"),
		"$YOUR_PROJECT_PATH",
		fmt.Sprintf("%s/%s", strings.Split(getCurrentAbsPath(), "/")[len(strings.Split(getCurrentAbsPath(), "/"))-2], getCurrentPath()),
		-1)), 0755); err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("./Dockerfile", []byte(httpGet("Dockerfile")), 0755); err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("./.gitignore", []byte(getCurrentPath()), 0755); err != nil {
		log.Fatal(err)
	}
}

func getCurrentAbsPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func getCurrentPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return path.Base(dir)
}

func httpGet(path string) string {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	resp, err := client.Get(fmt.Sprintf("https://raw.githubusercontent.com/qqqasdwx/godev/master/%s", path))
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body_bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(body_bytes)
}
