/*
Contains a general csv operation code which can be used on different type of structs in golang
*/
package workspaces

import (
	"io"
	"os"

	"github.com/gocarina/gocsv"
)

type ProgressUpdate struct {
	ItemsProcessed int
	Message        string
}

// CSV2Exporter exports data from a channel to a CSV format using the provided io.Writer.
func CSV2ExporterWriter[T any](source chan []*T, writer io.Writer) (chan ProgressUpdate, error) {
	progress := make(chan ProgressUpdate) // Channel to send progress updates
	value, err := gocsv.MarshalBytes([]T{})
	if err != nil {
		return nil, err
	}

	if _, err := writer.Write(value); err != nil {
		return nil, err
	}

	go func() {
		batchSize := 100

		var batch []*T

		for record := range source {

			batch = append(batch, record...)

			if len(batch) >= batchSize {

				if res, err := gocsv.MarshalStringWithoutHeaders(batch); err == nil {
					if _, err := writer.Write([]byte(res)); err != nil {
						// return err

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
			if res, err := gocsv.MarshalStringWithoutHeaders(batch); err == nil {
				if _, err := writer.Write([]byte(res)); err != nil {
					progress <- ProgressUpdate{
						ItemsProcessed: len(batch),
					}
				} else {
					progress <- ProgressUpdate{
						ItemsProcessed: len(batch),
					}
				}
			}
		}
	}()

	return progress, nil
}

func CSV2ExporterToFile[T any](source chan []*T, fp string) (chan ProgressUpdate, error) {
	// Example using an os.File as the writer
	file, err := os.Create(fp)
	if err != nil {
		return nil, err
	}
	// defer file.Close()

	return CSV2ExporterWriter(source, file)
}
