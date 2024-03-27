package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var IRRMap = make(map[string]map[string]float64)
var IRRMapGlobal = make(map[string]float64)

func updateIRR(req []string) {

	var tenant = req[0]
	var product = req[1]

	if _, mapped := IRRMap[tenant]; !mapped {
		IRRMap[tenant] = make(map[string]float64)
	}

	if _, found := IRRMap[tenant][product]; !found{
		IRRMap[tenant][product] = 0
	}else{
		IRRMap[tenant][product]++
	}
}

func calculateIRR() {
	var count float64
	var req float64
	for tenant, products := range IRRMap{
		count = 0
		req = 0
		for _, quantity := range products {
			count += quantity
			req += quantity + 1
		}
		IRR := count/req
		fmt.Printf("%s,%f\n", tenant, IRR)	
	}
}

func updateIRRGlobal(req []string) {

	var product = req[1]

	if _, mapped := IRRMapGlobal[product]; !mapped {
		IRRMapGlobal[product] = 0
	}else{
		IRRMapGlobal[product]++
	}
}

func calculateIRRGlobal() {
	var count float64
	var req float64
	for _, quantity := range IRRMapGlobal{
		count += quantity
		req += quantity + 1	
	}
	IRR := count/req
	fmt.Printf("%f\n", IRR)
}

func readTrace() {
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		_ = scanner.Text()
	}

	for scanner.Scan() {
		line := scanner.Text()
		result := strings.Split(line, ",")

		var resultList []string
		for _, element := range result {
			resultList = append(resultList, element)
		}

		updateIRR(resultList) //Em caso de calculo de IRR por tenant
		updateIRRGlobal(resultList) //Em caso de calculo de IRR Global
	}

	calculateIRR() //Em caso de calculo de IRR por tenant
	calculateIRRGlobal() //Em caso de calculo de IRR Global

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Erro ao ler a entrada:", err)
	}
}

var group string
func getArguments(){
	flag.Parse()
}

var trace = make([][]string, 65000000)
var traceSize int

func main() {
	getArguments()
	readTrace()
	//getIRR()
}

