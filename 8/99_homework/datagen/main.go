package main

import (
	"encoding/json"
	"fmt"
	"github.com/icrowley/fake"
	"os"
	"strconv"
)

func main() {

	files := 100
	linesPerFile := 20000

	for fileNum := 0; fileNum < files; fileNum++ {
		fmt.Println("generationg", fileNum, "of", files)

		file, err := os.Create("../searcher/data/logs" + strconv.Itoa(fileNum) + ".txt")
		if err != nil {
			panic(err)
		}

		for line := 0; line < linesPerFile; line++ {

			data := map[string]interface{}{
				"name":     fake.FirstName() + " " + fake.LastName(),
				"browsers": []string{fake.UserAgent(), fake.UserAgent(), fake.UserAgent(), fake.UserAgent()},
				"email":    fake.EmailAddress(),
				"country":  fake.Country(),
				"company":  fake.Company(),
				"phone":    fake.Phone(),
				"job":      fake.JobTitle(),
			}
			str, _ := json.Marshal(data)

			fmt.Fprintf(file, "%s\n", str)
		}

	}

}
