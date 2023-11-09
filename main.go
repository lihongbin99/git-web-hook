package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return
		}

		fmt.Print("\n\n", time.Now(), string(data), "\n\n")

		msg := GitMsg{}
		json.Unmarshal(data, &msg)

		if msg.Ref == "refs/heads/main" {

			cmd_path := "./" + msg.Repository.Name + ".sh"

			if len(cmd_path) > 0 {
				cmd := exec.Command(cmd_path)
				cmd.Stdout = os.Stdout
				if err := cmd.Run(); err != nil {
					fmt.Println(err)
				}
			}

		}
	})

	fmt.Println("启动成功...")
	http.ListenAndServe(":13520", nil)
}
