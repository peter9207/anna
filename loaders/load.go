package loaders

import "os"
import "encoding/csv"
import "io"
import "log"

func FromCSV(filename string, callback func([]string)) {
	var file io.Reader
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	r := csv.NewReader(file)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		callback(record)
	}
	return
}
