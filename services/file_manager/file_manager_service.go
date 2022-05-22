package file_manager

import (
	"fmt"
	"infolelang/lib"
	models "infolelang/models/file_manager"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	minio "gitlab.com/golang-package-library/minio"
)

type FileManagerDefinition interface {
	MakeUpload(request models.FileManagerRequest) (responses []models.FileManagerResponse, err error)
	GetFile(request models.FileManagerRequest) (response models.FileManagerResponse, err error)
	RemoveObject(request models.FileManagerRequest) (bool, error)
}
type FileManagerService struct {
	minio       minio.MinioDefinition
	minioClient minio.Minio
	logger      lib.Logger
}

func NewFileManagerService(
	minio minio.MinioDefinition,
	minioClient minio.Minio,
	logger lib.Logger) FileManagerDefinition {
	return FileManagerService{
		minio:       minio,
		minioClient: minioClient,
		logger:      logger,
	}
}

// GetFile implements FileManagerDefinition
func (fm FileManagerService) GetFile(request models.FileManagerRequest) (response models.FileManagerResponse, err error) {
	var minioPath string
	currentTime := time.Now()
	time.LoadLocation("Asia/Jakarta")
	timeNow := currentTime.Format("01-02-2006")
	filename := timeNow + "-" + request.File.Filename

	src, err := request.File.Open()
	if err != nil {
		fm.logger.Zap.Error(err)
	}
	defer src.Close()

	dir, err := os.Getwd()

	if err != nil {
		fm.logger.Zap.Error(err)
	}

	// mkdir bucket if not exist
	if _, err := os.Stat(dir + "/storage/uploads/" + request.BucketName); os.IsNotExist(err) {
		err = os.MkdirAll(dir+"/storage/uploads/"+request.BucketName, os.ModePerm)
		if err != nil {
			fm.logger.Zap.Error(err)
		}
	}

	fileLocation := filepath.Join(dir, "storage/uploads/"+request.BucketName+"/", filename)
	dst, err := os.Create(fileLocation)

	if err != nil {
		fm.logger.Zap.Error(err)
	}
	defer dst.Close()
	if _, err = io.Copy(dst, src); err != nil {
		fm.logger.Zap.Error(err)
	}

	pathFile := "storage/uploads/" + request.BucketName + "/" + filename

	// check bucket exist
	// if true do upload else create bucket and do upload
	bucketExist := fm.minio.BucketExist(fm.minioClient.MinioClient, request.BucketName)

	uuid := uuid.New()
	if bucketExist {
		// Get Content Type
		dataFile, err := os.Open(fileLocation)
		if err != nil {
			fmt.Println(err)
		}

		contentType, err := GetFileContentType(dataFile)
		if err != nil {
			fmt.Println(err.Error())
		}

		objectMinioPath := "tmp/" + uuid.String() + "/" + request.Subdir + "/" + filename
		_, err = fm.minio.UploadObject(fm.minioClient.MinioClient, request.BucketName, objectMinioPath, pathFile, contentType)
		if err != nil {
			fm.logger.Zap.Error(err)
		}

	} else {
		fm.minio.MakeBucket(fm.minioClient.MinioClient, request.BucketName, "")

		// Get Content Type
		dataFile, err := os.Open(fileLocation)
		if err != nil {
			fmt.Println(err)
		}

		contentType, err := GetFileContentType(dataFile)
		if err != nil {
			fmt.Println(err.Error())
		}

		objectMinioPath := "tmp/" + uuid.String() + "/" + request.Subdir + "/" + filename

		_, err = fm.minio.UploadObject(fm.minioClient.MinioClient, request.BucketName, objectMinioPath, pathFile, contentType)
		if err != nil {
			fm.logger.Zap.Error(err)
		}

	}

	err = os.Remove(fileLocation)

	if err != nil {
		fmt.Println(err)
	}

	minioPath = "tmp/" + uuid.String() + "/" + request.Subdir + "/" + filename
	fileResponse := models.FileManagerResponse{
		Subdir:   minioPath,
		Size:     fmt.Sprint(request.File.Size),
		Filename: request.File.Filename,
	}

	return fileResponse, err

}

// MakeUpload implements FileManagerDefinition
func (fm FileManagerService) MakeUpload(request models.FileManagerRequest) (responses []models.FileManagerResponse, err error) {
	panic("unimplemented")
}

// RemoveObject implements FileManagerDefinition
func (fm FileManagerService) RemoveObject(request models.FileManagerRequest) (bool, error) {
	panic("unimplemented")
}

func GetFileContentType(out *os.File) (string, error) {
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}
