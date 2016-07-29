package main

import (
	"Tests/Generated/report"
	"Tests/go-acceptance-tests/utils"
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	server, err := startServer()
	if err != nil {
		fmt.Printf("Error! %v\n", err)
	}
	allPass := true
	runTests(&allPass)
	getReport()
	fmt.Println("Kill server...")
	server.Kill()
	if !allPass {
		panic("Not all tests passed")
	}
}

func startServer() (*os.Process, error) {
	fmt.Println("Start server...")
	testServerPath := "../../../../../TestServer/server"
	install := exec.Command("npm", "install")
	install.Dir = testServerPath
	server := exec.Command("node", "startup/www")
	server.Dir = testServerPath
	if err := install.Run(); err != nil {
		return install.Process, err
	}
	if err := server.Start(); err != nil {
		return server.Process, err
	}
	return server.Process, nil
}

func runTests(allPass *bool) {
	fmt.Println("Run tests...")
	fmt.Println("========================")
	testSuites := []string{
		"arraygroup",
		"booleangroup",
		"bytegroup",
		"complexgroup",
		"dategroup",
		"datetimegroup",
		"datetimerfc1123group",
		"dictionarygroup",
		"durationgroup",
		"filegroup",
		"formdatagroup",
		"headergroup",
		"httpInfrastructuregroup",
		"integergroup",
		"modelflatteninggroup",
		"numbergroup",
		"optionalgroup",
		"stringgroup",
		"urlgroup",
		"validationgroup"}
	for _, suite := range testSuites {
		fmt.Printf("Run test (go test ./go-acceptance-tests/%v_test -v) ...\n", suite)
		tests := exec.Command("go", "test", fmt.Sprintf("./go-acceptance-tests/%v_test", suite), "-v")
		var stdout, stderr bytes.Buffer
		tests.Stdout, tests.Stderr = &stdout, &stderr
		err := tests.Run()
		fmt.Printf("STDOUT\n%v", stdout.String())
		fmt.Println("------------------------")
		fmt.Printf("STDERR\n%v", stderr.String())
		if err != nil {
			fmt.Printf("Error! %v\n", err)
			*allPass = false
		}
		fmt.Println("========================")
		if stderr.String()[:2] != "OK" {
			*allPass = false
		}
	}
}

func getReport() {
	fmt.Println("Get report...")
	// var reportClient = report.NewWithBaseURI(utils.GetBaseURI(), "jsklad")
	var reportClient = report.NewWithBaseURI(utils.GetBaseURI())
	res, err := reportClient.GetReport()
	if err != nil {
		fmt.Println("Error:", err)
	}

	count := 0
	for key, val := range *res.Value {
		if *val <= 0 /*strings.Contains(strings.ToUpper(key), "UUID") */ {
			fmt.Println(key, *val)
			count++
		}
	}
	total := len(*res.Value)
	fmt.Println("Total: ", total)
	fmt.Println("Left:", count)
}
