package workspaces

import (
	"fmt"
	"log"
	"path"

	"github.com/tus/tusd/pkg/handler"
)

func ImageToWebpFormatHook(sizes []ImageCropSize) UploadEventHook {
	return func(fi *handler.FileInfo, fe *FileEntity) error {
		// Implemenet the image to webp format.
		config := GetAppConfig()
		source := path.Join(config.Drive.Storage, fi.ID)

		for _, size := range sizes {
			fmt.Println("Convering image size: ", size)
			id := UUID_Long()
			dist := path.Join(config.Drive.Storage, id)
			err := ConvertToWebp(source, dist, size)
			if err != nil {
				log.Default().Panicln("Error:", err)
			}

			m := fmt.Sprintf("%v_%v", size.Width, size.Height)

			if size.Name != "" {
				m = size.Name
			}
			fe.Variations = append(fe.Variations, &FileVariations{
				UniqueId: id,
				Name:     &m,
			})
		}

		return nil
	}
}
