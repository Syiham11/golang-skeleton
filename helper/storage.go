package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

const serviceAccountURL = "eventori-google.json"
const BucketName = "st-core"
const TemporaryDirectory = "public/temp/"
const BaseStorageURL = "https://storage.googleapis.com"

type (
	ServiceAccount struct {
		Email      string `json:"client_email"`
		PrivateKey string `json:"private_key"`
	}
)

func GetServiceAccountData(credentialFile string) ServiceAccount {
	b, _ := ioutil.ReadFile(credentialFile)
	c := ServiceAccount{}
	json.Unmarshal(b, &c)

	return c
}

func GetObjectList(prefix string) []string {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithServiceAccountFile(serviceAccountURL))
	if err != nil {
		log.Fatal(err)
	}

	bkt := client.Bucket(BucketName)

	query := &storage.Query{Prefix: prefix}
	var names []string
	it := bkt.Objects(ctx, query)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		names = append(names, attrs.Name)
	}

	return names
}

func MoveObject(src, dst string) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithServiceAccountFile(serviceAccountURL))
	if err != nil {
		return fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	srcObj := client.Bucket(BucketName).Object(src)
	dstObj := client.Bucket(BucketName).Object(dst)

	if _, err := dstObj.CopierFrom(srcObj).Run(ctx); err != nil {
		return fmt.Errorf("Object(%q).CopierFrom(%q).Run: %v", dst, src, err)
	}
	if err := srcObj.Delete(ctx); err != nil {
		return fmt.Errorf("Object(%q).Delete: %v", src, err)
	}
	fmt.Println("Blob %v moved to %v.\n", src, dst)
	return nil
}

func DeleteObject(object string) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithServiceAccountFile(serviceAccountURL))
	if err != nil {
		return fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	obj := client.Bucket(BucketName).Object(object)
	if err := obj.Delete(ctx); err != nil {
		return fmt.Errorf("Object(%q).Delete: %v", object, err)
	}

	return nil
}

func UploadFile(c echo.Context, prefix string, fieldname string) ([]string, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithServiceAccountFile(serviceAccountURL))
	if err != nil {
		return []string{}, fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	form, err := c.MultipartForm()
	if err != nil {
		return []string{}, err
	}

	files := form.File[fieldname]
	var filenames []string

	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			return filenames, err
		}
		defer src.Close()

		filename := fmt.Sprintf("%s-%s", strconv.FormatInt(time.Now().Unix(), 10), file.Filename)
		wc := client.Bucket(BucketName).Object(prefix + filename).NewWriter(ctx)
		if _, err = io.Copy(wc, src); err != nil {
			return filenames, fmt.Errorf("io.Copy: %v", err)
		}
		if err := wc.Close(); err != nil {
			return filenames, fmt.Errorf("Writer.Close: %v", err)
		}
		fmt.Println(fmt.Sprintf("Blob %v uploaded.\n", filename))

		filenames = append(filenames, filename)
	}

	return filenames, nil
}

func GenerateSignedURL(object string) (string, error) {
	serviceAccountData := GetServiceAccountData(serviceAccountURL)
	fmt.Println(serviceAccountData)

	url, err := storage.SignedURL(BucketName, object, &storage.SignedURLOptions{
		GoogleAccessID: serviceAccountData.Email,
		PrivateKey:     []byte(serviceAccountData.PrivateKey),
		Method:         "GET",
		Expires:        time.Now().Add(48 * time.Hour),
	})

	if err != nil {
		return "", err
	}

	return url, nil
}

func GetDirectoryAndFilename(absoluteURL string) (string, string) {
	baseBucketLength := len(BaseStorageURL) + len(BucketName) + 2
	if baseBucketLength < len(absoluteURL) {
		segments := strings.Split(absoluteURL[baseBucketLength:], "/")
		dir := strings.Join(segments[:len(segments)-1], "/")
		filename := segments[len(segments)-1]
		return dir, filename
	}
	return "", ""
}
