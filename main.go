package main

import (
	"github.com/qqqasdwx/godev/cmd"
)

func main() {
	cmd.Execute()
}

// func usage() {
// 	fmt.Println("Usage: godev [init] [up] [down] [cli]")
// }

// func devInit() {
// 	if _, err := os.Stat(path.Join(".", "workspace")); !os.IsNotExist(err) {
// 		fmt.Println("please remove your workspace dir!")
// 		return
// 	}
// 	if err := os.MkdirAll(path.Join(".", "workspace"), 0755); err != nil {
// 		log.Fatal(err)
// 	}
// 	if err := ioutil.WriteFile("./workspace/.env", []byte(strings.Replace(httpGet("file/env"), "$YOUR_PROJECT_NAME", getCurrentPath(), -1)), 0755); err != nil {
// 		log.Fatal(err)
// 	}
// 	if err := ioutil.WriteFile("./workspace/Dockerfile", []byte(httpGet("file/Dockerfile")), 0755); err != nil {
// 		log.Fatal(err)
// 	}
// 	if err := ioutil.WriteFile("./workspace/docker-compose.yml", []byte(httpGet("file/docker-compose.yml")), 0755); err != nil {
// 		log.Fatal(err)
// 	}
// 	if err := ioutil.WriteFile("./workspace/bashrc", []byte(httpGet("file/bashrc")), 0755); err != nil {
// 		log.Fatal(err)
// 	}

// 	if err := ioutil.WriteFile("./Makefile", []byte(strings.Replace(
// 		httpGet("Makefile"),
// 		"$YOUR_PROJECT_PATH",
// 		fmt.Sprintf("%s/%s", strings.Split(getCurrentAbsPath(), "/")[len(strings.Split(getCurrentAbsPath(), "/"))-2], getCurrentPath()),
// 		-1)), 0755); err != nil {
// 		log.Fatal(err)
// 	}
// 	if err := ioutil.WriteFile("./Dockerfile", []byte(httpGet("Dockerfile")), 0755); err != nil {
// 		log.Fatal(err)
// 	}
// 	if err := ioutil.WriteFile("./.gitignore", []byte(getCurrentPath()), 0755); err != nil {
// 		log.Fatal(err)
// 	}
// }
