package workspaces

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/gabriel-vasile/mimetype"
	"github.com/tus/tusd/pkg/filestore"
	tusd "github.com/tus/tusd/pkg/handler"
)

func FileActionCreate(
	dto *FileEntity, query QueryDSL,
) (*FileEntity, *IError) {
	return FileActionCreateFn(dto, query)
}

func FileActionUpdate(
	query QueryDSL,
	fields *FileEntity,
) (*FileEntity, *IError) {
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
		model.UniqueId = UUID()
	}
	return GetDbRef().Create(&model).Error
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

func afterTusUploadedOnDisk(event *tusd.HookEvent, q *QueryDSL, ctx *FileUploadContext) (*FileEntity, error) {
	fname := event.Upload.MetaData["filename"]
	fpath := event.Upload.MetaData["path"]
	fsize := event.Upload.Size
	ftype := event.Upload.MetaData["filetype"]
	diskPath := event.Upload.ID
	entity := &FileEntity{
		Name:        &fname,
		VirtualPath: &fpath,
		DiskPath:    &diskPath,
		UniqueId:    event.Upload.ID,
		Size:        &fsize,
		Type:        &ftype,
		WorkspaceId: &q.WorkspaceId,
		UserId:      &q.UserId,
	}

	fmt.Println("----", ctx)

	if ctx != nil {

		fmt.Println("After creation:", len(ctx.AfterCreatedHooks))
		for _, item := range ctx.AfterCreatedHooks {
			item(&event.Upload, entity)
		}
	}

	CreateFile(entity)

	return entity, nil
}

var GlobalTusFileUploadContext *FileUploadContext

func LiftTusServer() {

	if config.Storage == "" {
		return
	}

	store := filestore.FileStore{
		Path: config.Storage,
	}

	os.Mkdir(config.Storage, 0777)

	composer := tusd.NewStoreComposer()
	store.UseIn(composer)

	handler, err := tusd.NewHandler(tusd.Config{
		BasePath:              "/files/",
		StoreComposer:         composer,
		NotifyCompleteUploads: true,
	})

	if err != nil {
		panic(fmt.Errorf("unable to create handler: %s", err))
	}

	go func() {
		for {
			event := <-handler.CompleteUploads
			var result *AuthResultDto

			if os.Getenv("BYPASS_WORKSPACES") == "YES" {
				result = &AuthResultDto{
					WorkspaceId: &WORKSPACE_SYSTEM,
					UserId:      &WORKSPACE_SYSTEM,
				}
			} else {
				wi := event.HTTPRequest.Header.Get("workspace-id")
				tk := event.HTTPRequest.Header.Get("authorization")

				result, err = WithAuthorizationPure(&AuthContextDto{
					WorkspaceId:  &wi,
					Token:        &tk,
					Capabilities: []PermissionInfo{},
				})

				if result != nil {
					q := QueryDSL{
						WorkspaceId: *result.WorkspaceId,
						UserId:      *result.UserId,
					}

					afterTusUploadedOnDisk(&event, &q, GlobalTusFileUploadContext)
				}
			}
		}
	}()

	fmt.Println("TUS is listenning on", ":"+config.TusPort)
	if os.Getenv("BYPASS_WORKSPACES") == "YES" {
		http.Handle("/files/", http.StripPrefix("/files/", handler))
	} else {
		http.Handle("/files/",
			WithAuthorizationHttp(http.StripPrefix("/files/", handler), true),
		)
		http.Handle("/files-inline/", http.StripPrefix("/files-inline/", http.FileServer(http.Dir(config.TusPort))))
	}
	err = http.ListenAndServe(":"+config.TusPort, nil)
	if err != nil {
		panic(fmt.Errorf("Unable to listen: %s", err))
	}
}

func copyFile(src string, dst string) {
	// Read all content of src to data, may cause OOM for a large file.
	data, _ := ioutil.ReadFile(src)

	ioutil.WriteFile(dst, data, 0644)

}

func UploadFromDisk(filePath string) (*FileEntity, string, error) {
	fi, _ := os.Stat(filePath)
	fmt.Printf("The file is %d bytes long", fi.Size())
	fmt.Println("Source:", filePath)

	mtype, _ := mimetype.DetectFile(filePath)
	fmt.Println("Type:", mtype.String())

	file := tusd.FileInfo{
		ID: UUID_Long(),
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

	entity, err := afterTusUploadedOnDisk(&event, &QueryDSL{
		WorkspaceId: "system",
		UserId:      "system",
	}, GlobalTusFileUploadContext)

	if err != nil {
		return nil, "", err
	}

	return entity, file.ID, nil
}

func UploadFromFs(fs *embed.FS, filePath string) (*FileEntity, string, error) {

	sourceFile, _ := fs.ReadFile(filePath)
	var fileSize int = len(sourceFile)

	fmt.Printf("The file is %d bytes long", fileSize)
	fmt.Println("Source:", filePath)

	mimetype := ""

	fmt.Println("Type:", mimetype)

	file := tusd.FileInfo{
		ID: UUID_Long(),
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
	os.MkdirAll(config.Storage, os.ModePerm)
	ioutil.WriteFile(fileTarget, sourceFile, 0644)
	os.WriteFile(path.Join(config.Storage, file.ID+".info"), dicJson, 0644)

	entity, err := afterTusUploadedOnDisk(&event, &QueryDSL{
		WorkspaceId: "system",
		UserId:      "system",
	}, GlobalTusFileUploadContext)

	if err != nil {
		return nil, "", err
	}

	return entity, file.ID, nil
}
