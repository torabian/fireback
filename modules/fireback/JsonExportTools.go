package fireback

import (
	"encoding/json"
	"os"
)

func JsonExporterWriter[T any](source chan []*T, fp string) (chan ProgressUpdate, error) {
	progress := make(chan ProgressUpdate) // Channel to send progress updates

	writer, err := os.Create(fp)
	if err != nil {
		return nil, err
	}

	go func() {
		defer writer.Close()

		batchSize := 100

		var batch []*T

		for record := range source {

			batch = append(batch, record...)

			if len(batch) >= batchSize {

				if res, err := json.MarshalIndent(batch, "", "  "); err == nil {
					if _, err := writer.Write([]byte(res)); err != nil {
						progress <- ProgressUpdate{
							ItemsProcessed: 0,
							Error:          err,
						}
					} else {

						progress <- ProgressUpdate{
							ItemsProcessed: len(batch),
						}
					}
				}
				batch = nil
			}
		}

		if len(batch) > 0 {
			if res, err := json.MarshalIndent(batch, "", "  "); err == nil {
				if _, err := writer.Write([]byte(res)); err != nil {
					progress <- ProgressUpdate{
						ItemsProcessed: 0,
						Error:          err,
					}
				} else {
					progress <- ProgressUpdate{
						ItemsProcessed: len(batch),
					}
				}
			}
		}

		close(progress)

	}()

	return progress, nil
}
