package handlers

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type ResFiltered struct {
	Result string `json:"Result"`
	CyGt   string `json:"CyGt"`
	CyLt   string `json:"CyLt"`
	CyRt   string `json:"CyRt"`
	CyDb   string `json:"CyDb"`
}

func getFilteredData(csvFile, outDir string) (ResFiltered, error) {
	cmd := exec.Command("./scripts/filter.py", "-f", csvFile, "-s", outDir)
	stdout, err := cmd.Output()

	resultPath := ResFiltered{}
	if err != nil {
		fmt.Println(err.Error())
		return resultPath, err
	}

	json.Unmarshal(stdout, &resultPath)

	return resultPath, err
}