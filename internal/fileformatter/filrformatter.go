package fileformatter

import (
	"encoding/csv"
	"encoding/json"
	"io"
)

func OutputAsJSON(w io.Writer, data any) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(data)
}

func OutputAsCSV(w io.Writer, records []string) error {
	writer := csv.NewWriter(w)
	defer writer.Flush()

	for _, r := range records {
		if err := writer.Write([]string{r}); err != nil {
			return err
		}
	}
	return nil
}
