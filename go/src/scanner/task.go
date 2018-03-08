package scanner

import (
	"fmt"
	"log"
	"time"
)

func ProcessFile(filename string) {
	if !FileExists(filename) {
		log.Printf("No such file: %s\n", filename)
		return
	}

	fmt.Println("")
	log.Printf("Submitting %s ..\n", filename)

	apiResponse, err := Enqueue(filename)
	if err != nil {
		log.Println("Enqueue API error:", err)
	} else if apiResponse.Status != 0 {
		log.Println("Enqueue API error:", apiResponse.Descr)
	} else {
		log.Println("Success. MD5 is", apiResponse.Md5)

		// query for result every few seconds
		for {
			resultResponseList, err := Result(apiResponse.Url)
			if err != nil {
				log.Println("Result API error:", err)
				break
			} else {
				// parse first object
				resultResponse := (*resultResponseList)[0]
				if resultResponse.Status == "pending" {
					log.Printf("Task %s pending\n", apiResponse.Md5)
				} else if resultResponse.Status == "done" {
					log.Printf("Task %s completed\n", apiResponse.Md5)
					PrintResult(resultResponse)
					break
				} else {
					log.Printf("Task %s running: %d / %d processed\n",
						apiResponse.Md5,
						resultResponse.Scanned,
						resultResponse.Total,
					)
				}
			}

			time.Sleep(2000 * time.Millisecond)
		}
	}
}
