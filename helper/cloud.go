package helper

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"strings"
	"time"
	"toko/config"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type Uploads interface {
	Upload(file *multipart.FileHeader) (string, error)
	Destroy(publicID string) error
}
type cloudUpload struct {
	clds *cloudinary.Cloudinary
}

func NewCloud(cfg *config.AppConfig) Uploads {
	clds, err := cloudinary.NewFromParams(cfg.CLOUDINARY_CLOUD_NAME, cfg.CLOUDINARY_API_KEY, cfg.CLOUDINARY_API_SECRET)

	if err != nil {
		log.Println("init cloudinary gagal", err)
		return nil
	}

	return &cloudUpload{clds: clds}
}

func (cl *cloudUpload) Upload(file *multipart.FileHeader) (string, error) {
	src, _ := file.Open()
	defer src.Close()

	publicID := fmt.Sprintf("%d-%s", int(file.Size), time.Now().Format("20060102-150405")) // Format  "file_size-(YY-MM-DD)-(hh-mm-ss)""

	uploadResult, err := cl.clds.Upload.Upload(
		context.Background(),
		src,
		uploader.UploadParams{
			PublicID: publicID,
			Folder:   os.Getenv("CLOUDINARY_UPLOAD_FOLDER"),
		})
	if err != nil {
		return "", err
	}
	fmt.Println(uploadResult.SecureURL)
	return uploadResult.SecureURL, nil

}
func (cl *cloudUpload) Destroy(publicID string) error {
	_, err := cl.clds.Upload.Destroy(
		context.Background(),
		uploader.DestroyParams{
			PublicID: publicID,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
func GetPublicID(secureURL string) string {
	// Proses filter Public ID dari SecureURL(avatar)
	urls := strings.Split(secureURL, "/")
	urls = urls[len(urls)-2:]                               // array [file, random_name.extension]
	noExtension := strings.Split(urls[len(urls)-1], ".")[0] // remove ".extension", result "random_name"
	urls = append(urls[:1], noExtension)                    // new array [file, random_name]
	publicID := strings.Join(urls, "/")                     // "file/random_name"

	return publicID
}
