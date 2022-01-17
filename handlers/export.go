package handlers

import (
	"fmt"
	"os/exec"
	"strconv"
)

type ReqExport struct {
	RawFile string `json:"RawFile"`
	Ranges  []struct {
		Start int
		End   int
	} `json:"Ranges"`
}

func exportCsv(r ReqExport) error {
	args := []string{}
	for _, _r := range r.Ranges {
		args = append(args, "-r", strconv.Itoa(_r.Start), strconv.Itoa(_r.End))
	}

	cmd := exec.Command("./scripts/exporter.py", args...)
	stdout, err := cmd.Output()

	fmt.Print(string(stdout[:]))

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// json.Unmarshal(stdout, &resultPath)

	return err
}
