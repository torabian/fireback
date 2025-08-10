package abac

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/tus/tusd/pkg/filestore"
	tusd "github.com/tus/tusd/pkg/handler"
)

func FileActionCreate(
	dto *FileEntity, query fireback.QueryDSL,
) (*FileEntity, *fireback.IError) {
	return FileActionCreateFn(dto, query)
}

func FileActionUpdate(
	query fireback.QueryDSL,
	fields *FileEntity,
) (*FileEntity, *fireback.IError) {
	return FileActionUpdateFn(query, fields)
}

type FileInfo struct {
	Name      string    `json:"name"`
	Size      int64     `json:"size"`
	Directory string    `json:"directory"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	IsDir     bool      `json:"isDir"`
}

type Directory struct {
	Parent uint   `json:"parent"`
	Name   string `json:"name"`
}

func CreateFile(model *FileEntity) error {

	if model.UniqueId == "" {
		model.UniqueId = fireback.UUID()
	}
	return fireback.GetDbRef().Create(&model).Error
}

/*
Use this to define different actions, maybe based on the file type
after upload is completed. For example, you might want to create
a hook that crops images, or creates a preview out of pdf files.
*/
type UploadEventHook func(*tusd.FileInfo, *FileEntity) error

type FileUploadContext struct {
	AfterCreatedHooks []UploadEventHook
}

func afterTusUploadedOnDisk(event *tusd.HookEvent, q *fireback.QueryDSL, ctx *FileUploadContext) (*FileEntity, error) {
	fname := event.Upload.MetaData["filename"]
	fpath := event.Upload.MetaData["path"]
	fsize := event.Upload.Size
	ftype := event.Upload.MetaData["filetype"]
	diskPath := event.Upload.ID
	entity := &FileEntity{
		Name:        fname,
		VirtualPath: fpath,
		DiskPath:    diskPath,
		UniqueId:    event.Upload.ID,
		Size:        fsize,
		Type:        ftype,
		WorkspaceId: fireback.NewString(q.WorkspaceId),
		UserId:      fireback.NewString(q.UserId),
	}

	if ctx != nil {
		for _, item := range ctx.AfterCreatedHooks {
			item(&event.Upload, entity)
		}
	}

	err2 := CreateFile(entity)
	if err2 != nil {
		return nil, err2
	}

	return entity, nil
}

var GlobalTusFileUploadContext *FileUploadContext

// func LiftTusServer() {
// 	config := fireback.GetConfig()

// 	if config.Storage == "" {
// 		return
// 	}

// 	store := filestore.FileStore{
// 		Path: config.Storage,
// 	}

// 	os.Mkdir(config.Storage, 0777)

// 	composer := tusd.NewStoreComposer()
// 	store.UseIn(composer)

// 	handler, err := tusd.NewHandler(tusd.Config{
// 		BasePath:              "/files/",
// 		StoreComposer:         composer,
// 		NotifyCompleteUploads: true,
// 	})

// 	if err != nil {
// 		panic(fmt.Errorf("unable to create handler: %s", err))
// 	}

// 	go func() {
// 		for {
// 			event := <-handler.CompleteUploads
// 			var result *fireback.AuthResultDto

// 			wi := event.HTTPRequest.Header.Get("workspace-id")
// 			tk := event.HTTPRequest.Header.Get("authorization")

// 			result, err = fireback.WithAuthorizationPure(&fireback.AuthContextDto{
// 				WorkspaceId:  wi,
// 				Token:        tk,
// 				Capabilities: []fireback.PermissionInfo{},
// 			})

// 			if result != nil {
// 				q := fireback.QueryDSL{
// 					WorkspaceId: wi,
// 					UserId:      result.UserId.String,
// 				}

// 				afterTusUploadedOnDisk(&event, &q, GlobalTusFileUploadContext)
// 			}

// 		}
// 	}()

// 	fireback.LOG.Info("TUS file uploader separate port is listening:", zap.String("port", config.TusPort))

// 	http.Handle("/files/",
// 		WithAuthorizationHttp(http.StripPrefix("/files/", handler), true),
// 	)
// 	http.Handle("/files-inline/", http.StripPrefix("/files-inline/", http.FileServer(http.Dir(config.TusPort))))

// 	err = http.ListenAndServe(":"+config.TusPort, nil)
// 	if err != nil {
// 		panic(fmt.Errorf("Unable to listen: %s", err))
// 	}
// }

// TUS file uploading systems can be running directly into a gin server
// This is done because, the golang routines are capable of nonblcking io.

func LiftTusServerInHttp(app *gin.Engine) {
	config := fireback.GetConfig()
	if config.Storage == "" {
		fireback.LOG.Info("File server has been skipped, because there is no directory.")
		return
	}

	store := filestore.FileStore{
		Path: config.Storage,
	}

	os.Mkdir(config.Storage, 0777)

	composer := tusd.NewStoreComposer()
	store.UseIn(composer)

	handler, err := tusd.NewUnroutedHandler(tusd.Config{
		BasePath:              "/tus/",
		StoreComposer:         composer,
		NotifyCompleteUploads: true,
	})

	if err != nil {
		panic(fmt.Errorf("unable to create handler: %s", err))
	}

	go func() {
		for {
			event := <-handler.CompleteUploads
			var result *fireback.AuthResultDto

			wi := event.HTTPRequest.Header.Get("workspace-id")
			tk := event.HTTPRequest.Header.Get("authorization")

			result, err = fireback.WithAuthorizationPure(&fireback.AuthContextDto{
				WorkspaceId:  wi,
				Token:        tk,
				Capabilities: []fireback.PermissionInfo{},
			})

			if result != nil {
				q := fireback.QueryDSL{
					WorkspaceId: wi,
					UserId:      result.UserId.String,
				}

				afterTusUploadedOnDisk(&event, &q, GlobalTusFileUploadContext)
			}

		}
	}()

	app.POST("/tus", gin.WrapF(handler.PostFile))
	app.HEAD("/tus/:id", gin.WrapF(handler.HeadFile))
	app.PATCH("/tus/:id", gin.WrapF(handler.PatchFile))
	app.GET("/files-inline/:id", gin.WrapF(handler.GetFile))
}

func copyFile(src string, dst string) {
	// Read all content of src to data, may cause OOM for a large file.
	data, _ := ioutil.ReadFile(src)

	ioutil.WriteFile(dst, data, 0644)

}

func UploadFromDisk(filePath string) (*FileEntity, string, error) {
	config := fireback.GetConfig()
	fi, _ := os.Stat(filePath)

	mtype, _ := mimetype.DetectFile(filePath)

	file := tusd.FileInfo{
		ID: fireback.UUID_Long(),
		MetaData: tusd.MetaData{
			"filename": filepath.Base(filePath),
			"filetype": mtype.String(),
		},
		Size: fi.Size(),
	}

	event := tusd.HookEvent{
		Upload: file,
	}

	dicJson, _ := json.MarshalIndent(file, "", "  ")

	fileTarget := path.Join(config.Storage, file.ID)
	copyFile(filePath, fileTarget)
	os.WriteFile(path.Join(config.Storage, file.ID+".info"), dicJson, 0644)

	entity, err := afterTusUploadedOnDisk(&event, &fireback.QueryDSL{
		WorkspaceId: "system",
		UserId:      "system",
	}, GlobalTusFileUploadContext)

	if err != nil {
		return nil, "", err
	}

	return entity, file.ID, nil
}

func UploadFromFs(fs *embed.FS, filePath string) (*FileEntity, string, error) {
	config := fireback.GetConfig()
	sourceFile, err := fs.ReadFile(filePath)

	if err != nil {
		return nil, "", err
	}

	var fileSize int = len(sourceFile)

	if fileSize == 0 {
		log.Default().Printf("its strange that the file %s on embed resource is 0 bytes, are you sure the address of it is correct?", filePath)
	}

	mimetype := ""

	file := tusd.FileInfo{
		ID: fireback.UUID_Long(),
		MetaData: tusd.MetaData{
			"filename": filepath.Base(filePath),
			"filetype": mimetype,
		},
		Size: int64(fileSize),
	}

	event := tusd.HookEvent{
		Upload: file,
	}

	dicJson, _ := json.MarshalIndent(file, "", "  ")

	fileTarget := path.Join(config.Storage, file.ID)
	err = os.MkdirAll(config.Storage, os.ModePerm)
	if err != nil {
		log.Default().Printf("storage directory creation error: %w", err)
		return nil, "", err
	}
	err = os.WriteFile(fileTarget, sourceFile, 0644)
	if err != nil {
		log.Default().Printf("writing file error on upload from fs: %w", err)
		return nil, "", err
	}
	err = os.WriteFile(path.Join(config.Storage, file.ID+".info"), dicJson, 0644)
	if err != nil {
		log.Default().Printf("writing tus meta data error: %w", err)
		return nil, "", err
	}

	entity, err := afterTusUploadedOnDisk(&event, &fireback.QueryDSL{
		WorkspaceId: "system",
		UserId:      "system",
	}, GlobalTusFileUploadContext)

	if err != nil {
		return nil, "", err
	}

	return entity, file.ID, nil
}

func ImportYamlFromFsResources(fs *embed.FS, filePath string) []fireback.ResourceMap {
	result := []fireback.ResourceMap{}
	var resources fireback.ContentImport[any]
	err := fireback.ReadYamlFileEmbed(fs, filePath, &resources)

	if err != nil {
		log.Fatalln("Error importing content:", err, filePath)
	}

	for _, resource := range resources.Resources {
		actualPath := path.Join(filepath.Dir(filePath), resource.Path)

		uniqueId := ""
		fileId := ""
		var fileBytes []byte = nil

		if resource.Blob {

			r, err := fs.ReadFile(actualPath)
			if err != nil {
				log.Default().Printf("could not read the: %v as blob because: %w", actualPath, err)
				continue
			}

			fileBytes = r
		} else {
			entity, fileId2, err := UploadFromFs(fs, actualPath)
			uniqueId = entity.UniqueId
			fileId = fileId2
			if err != nil {
				log.Default().Printf("uploading failed: %w", err)

				continue
			}
		}

		result = append(result, fireback.ResourceMap{
			DriveId:  uniqueId,
			FileId:   fileId,
			Key:      resource.Key,
			DiskPath: actualPath,
			Blob:     fileBytes,
		})
	}

	return result
}
