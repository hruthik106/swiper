package fileformatter

import (
	"encoding/csv"
	"encoding/json"
	"os"
)

func OutputAsJSON(data any) error {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(data)
}

func OutputAsCSV(data []string) error {
	writer := csv.NewWriter(os.Stdout)
	defer writer.Flush()

	for _, d := range data {
		if err := writer.Write([]string{d}); err != nil {
			return err
		}
	}
	return nil
}
