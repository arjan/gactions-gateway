package main

import (
	"io/ioutil"
	"os/exec"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type update_actions struct {
	Creds   string  `json:"creds"`
	Actions string `json:"actions"`
	Project string `json:"project"`
}

func gactionsUpdateHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		d := json.NewDecoder(r.Body)
		req := &update_actions{}
		err := d.Decode(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpdir, err := ioutil.TempDir("", "bsqd")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Println(tmpdir)

		ioutil.WriteFile(tmpdir + "/creds.data", []byte(req.Creds), 0644)
		ioutil.WriteFile(tmpdir + "/action.json", []byte(req.Actions), 0644)

		cmd := exec.Command("gactions", "update", "--action_package", "action.json", "--project", req.Project)
		cmd.Dir = tmpdir

		stdoutStderr, e := cmd.CombinedOutput()
		if e != nil {
			log.Printf("Command finished with error: %v", e)
			log.Println(string(stdoutStderr[:]))
			http.Error(w, e.Error(), http.StatusBadRequest)
			return
		}

		log.Println("-----")
		log.Println(string(stdoutStderr[:]))
		log.Println("Updated project " + req.Project)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "OK")

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "I can't do that.")
	}
}

func main() {
	http.HandleFunc("/update", gactionsUpdateHandler)

	log.Println("Go!")
	http.ListenAndServe(":8080", nil)
}
