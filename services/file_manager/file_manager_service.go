package file_manager

import (
	"fmt"
	"infolelang/lib"
	models "infolelang/models/file_manager"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"gitlab.com/golang-package-library/logger"
	"gitlab.com/golang-package-library/minio"
)

type FileManagerDefinition interface {
	MakeUpload(request models.FileManagerRequest) (responses models.FileManagerResponse, err error)
	GetFile(request models.FileManagerRequest) (response models.FileManagerResponseUrl, err error)
	RemoveObject(request models.FileManagerRequest) (response bool, err error)
}
type FileManagerService struct {
	logger logger.Logger
	minio  minio.Minio
}

func NewFileManagerService(
	minio minio.Minio,
	logger logger.Logger) FileManagerDefinition {
	return FileManagerService{
		logger: logger,
		minio:  minio,
	}
}

// GetFile implements FileManagerDefinition
func (fm FileManagerService) MakeUpload(request models.FileManagerRequest) (response models.FileManagerResponse, err error) {
	var minioPath string
	bucketName := os.Getenv("BUCKET_NAME")

	timeNow := lib.GetTimeNow("date")
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
	if _, err := os.Stat(dir + "/storage/uploads/" + bucketName); os.IsNotExist(err) {
		err = os.MkdirAll(dir+"/storage/uploads/"+bucketName, os.ModePerm)
		if err != nil {
			fm.logger.Zap.Error(err)
		}
	}

	fileLocation := filepath.Join(dir, "storage/uploads/"+bucketName+"/", filename)
	dst, err := os.Create(fileLocation)

	if err != nil {
		fm.logger.Zap.Error(err)
	}
	defer dst.Close()
	if _, err = io.Copy(dst, src); err != nil {
		fm.logger.Zap.Error(err)
	}

	pathFile := "storage/uploads/" + bucketName + "/" + filename

	// check bucket exist
	// if true do upload else create bucket and do upload
	bucketExist := fm.minio.BucketExist(fm.minio.Client(), bucketName)

	uuid := uuid.New()
	minioPath = "tmp/" + request.Subdir + lib.GetTimeNow("year") + "/" + lib.GetTimeNow("month") + "/" + lib.GetTimeNow("day") + "/" + uuid.String() + "/" + filename

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

		_, err = fm.minio.UploadObject(fm.minio.Client(), bucketName, minioPath, pathFile, contentType)
		if err != nil {
			fm.logger.Zap.Error(err)
		}

	} else {
		fm.minio.MakeBucket(fm.minio.Client(), bucketName, "")

		// Get Content Type
		dataFile, err := os.Open(fileLocation)
		if err != nil {
			fmt.Println(err)
		}

		contentType, err := GetFileContentType(dataFile)
		if err != nil {
			fmt.Println(err.Error())
		}

		_, err = fm.minio.UploadObject(fm.minio.Client(), bucketName, minioPath, pathFile, contentType)
		if err != nil {
			fm.logger.Zap.Error(err)
		}

	}

	err = os.Remove(fileLocation)

	if err != nil {
		fmt.Println(err)
	}

	fileResponse := models.FileManagerResponse{
		Subdir:   minioPath,
		Size:     fmt.Sprint(request.File.Size),
		Filename: request.File.Filename,
	}

	return fileResponse, err

}

// MakeUpload implements FileManagerDefinition
func (fm FileManagerService) GetFile(request models.FileManagerRequest) (responses models.FileManagerResponseUrl, err error) {
	bucket := os.Getenv("BUCKET_NAME")
	subdir := request.Subdir
	filename := request.Filename
	// strings.Split(subdir, "/")
	// fmt.Println(filename[3])

	// Minio Init
	var minioPath string
	// minioClient, err := minio.Init()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// check bucket exist
	// if true do upload else create bucket and do upload
	bucketExist := fm.minio.BucketExist(fm.minio.Client(), bucket)
	if bucketExist {
		preSign := fm.minio.SignUrl(fm.minio.Client(), bucket, subdir, request.Filename)
		minioPath = fmt.Sprint(preSign)
		fmt.Println(filename)
		fmt.Println("presign url", preSign)

		responses := models.FileManagerResponseUrl{
			MinioPath:  minioPath,
			PreSignUrl: preSign,
		}

		return responses, err
	} else {
		return responses, err
	}
}

// RemoveObject implements FileManagerDefinition
func (fm FileManagerService) RemoveObject(request models.FileManagerRequest) (response bool, err error) {
	bucket := os.Getenv("BUCKET_NAME")
	objectName := request.ObjectName

	// Minio Init
	// var minioPath string
	// minioClient, err := minio.Init()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// check bucket exist
	// if true do upload else create bucket and do upload
	bucketExist := fm.minio.BucketExist(fm.minio.Client(), bucket)
	if bucketExist {
		remove := fm.minio.RemoveObject(fm.minio.Client(), bucket, objectName)
		if remove {
			return true, err
		} else {
			return false, err
		}

	} else {
		return false, err
	}
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
