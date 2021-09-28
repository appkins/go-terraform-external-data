package terraform

import (
	"encoding/json"
	"log"
	"os"
)

type QueryData map[string]string

func check(e error) {
	if e != nil {
		log.Panic(e)
	}
}

func getQueryData() (d QueryData, e error) {
	e = json.NewDecoder(os.Stdin).Decode(&d)
	check(e)

	return d, e
}

func writeResponseRaw(res []byte) (err error) {
	if err != nil {
		log.Fatal(err)
	}
	_, err = os.Stdout.Write(res)
	check(err)
	return err
}

func writeResponse(res interface{}) (err error) {
	err = json.NewEncoder(os.Stdout).Encode(&res)
	check(err)
	return err
}

func ExternalDataRaw(fn func(QueryData) (bytes []byte, e error)) (err error) {
	input, err := getQueryData()
	check(err)
	if res, err := fn(input); err == nil {
		return writeResponseRaw(res)
	}
	return err
}

func ExternalData(fn func(QueryData) (r interface{}, e error)) (err error) {
	input, err := getQueryData()
	check(err)
	if res, err := fn(input); err == nil {
		return writeResponse(res)
	}
	return err
}
