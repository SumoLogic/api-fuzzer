package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/AdityaVallabh/swagger_meqa/meqa/mqutil"
)

func TestMqgo(t *testing.T) {
	var baseURL, username, password, apitoken string
	wd, _ := os.Getwd()
	meqaPath := filepath.Join(wd, "../../testdata")
	swaggerPath := filepath.Join(meqaPath, "petstore_meqa.yml")
	planPath := filepath.Join(meqaPath, "simple.yml")
	resultPath := filepath.Join(meqaPath, "result.yml")
	testToRun := "all"
	baseURL, username, password, apitoken, dataset := "", "", "", "", ""
	fuzzType := "none"
	batchSize := 0
	repro, verbose := false, false

	mqutil.Logger = mqutil.NewFileLogger(filepath.Join(meqaPath, "mqgo.log"))
	runMeqa(&meqaPath, &swaggerPath, &planPath, &resultPath, &testToRun, &username, &password, &apitoken, &baseURL, &dataset, &fuzzType, &batchSize, &repro, &verbose)
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
