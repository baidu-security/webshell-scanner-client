package scanner

import (
	"encoding/json"
	"github.com/olekukonko/tablewriter"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

func Result(url string) (resultResponse *[]ResultResponse, err error) {
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &resultResponse)
	return
}

func PrintResult(resultResponse ResultResponse) {
	table := tablewriter.NewWriter(os.Stdout)
	// table.SetBorder(false)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeader([]string{"Filename", "Result"})
	for _, data := range resultResponse.Data {
		if len(data.Descr) == 0 {
			data.Descr = "-"
		}

		table.Append([]string{
			data.Path,
			data.Descr,
		})
	}

	table.SetFooter([]string{
		"",
		"Scanned:     " + strconv.Itoa(resultResponse.Scanned) + "\nDetected:    " + strconv.Itoa(resultResponse.Detected) + "\nTotal files: " + strconv.Itoa(resultResponse.Total)})
	table.Render()
}
