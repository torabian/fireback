/*
Contains a general csv operation code which can be used on different type of structs in golang
*/
package workspaces

import (
	"encoding/json"
	"os"

	"github.com/gocarina/gocsv"
	"gopkg.in/yaml.v2"
)

type ProgressUpdate struct {
	ItemsProcessed int
	Message        string
	Complete       bool
	Error          error
}

// CSV2Exporter exports data from a channel to a CSV format using the provided io.Writer.
func CSV2ExporterWriter[T any](source chan []*T, fp string) (chan ProgressUpdate, error) {
	progress := make(chan ProgressUpdate) // Channel to send progress updates
	value, err := gocsv.MarshalBytes([]T{})
	if err != nil {
		return nil, err
	}

	writer, err := os.Create(fp)
	if err != nil {
		return nil, err
	}

	if _, err := writer.Write(value); err != nil {
		return nil, err
	}

	go func() {
		defer writer.Close()

		batchSize := 100

		var batch []*T

		for record := range source {

			batch = append(batch, record...)

			if len(batch) >= batchSize {

				if res, err := gocsv.MarshalStringWithoutHeaders(batch); err == nil {
					if _, err := writer.Write([]byte(res)); err != nil {
						// return err
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
			if res, err := gocsv.MarshalStringWithoutHeaders(batch); err == nil {
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
func YamlExporterWriter[T any](source chan []*T, fp string) (chan ProgressUpdate, error) {
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

				if res, err := yaml.Marshal(batch); err == nil {
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
			if res, err := yaml.Marshal(batch); err == nil {
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
