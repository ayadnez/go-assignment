package handler

import (
	"fmt"
	"go-backend/models"
	"io"
	"net/http"
	"os"

	"github.com/google/uuid"

	"path/filepath"

	"github.com/h2non/bimg"
)

func Downloadimg(urls []string, pid int) {
	db, err := models.InitDB()
	var success bool = false
	if err != nil {
		panic("connection failed" + err.Error())
	}
	var successfullDownloads []string
	for _, imageURL := range urls {
		path, err := downloadAndCompressImage(imageURL)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			success = true
			fmt.Println("Image downloaded and compressed successfully.")
			successfullDownloads = append(successfullDownloads, path)
		}
	}
	if success {
		db.Model(&models.DbProduct{}).Where("product_id = ?", pid).Update("compressed_product_images", successfullDownloads)
	}

}

func downloadAndCompressImage(imageURL string) (string, error) {
	response, err := http.Get(imageURL)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP Error: %s", response.Status)
	}

	// Read the image from the response body
	imageData, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	// Generate a unique filename using a UUID
	fileName := uuid.New().String() + ".jpg" // You can change the file extension as needed

	// Specify the local storage directory
	downloadDirectory := `./`

	// Check if the download directory exists, and create it if it doesn't
	if _, err := os.Stat(downloadDirectory); os.IsNotExist(err) {
		err = os.MkdirAll(downloadDirectory, 0755)
		if err != nil {
			return "", err
		}
	}

	// Create the complete path to save the file
	localPath := filepath.Join(downloadDirectory, fileName)

	// Compress the image using bimg
	processedImage, err := bimg.Resize(imageData, bimg.Options{
		Quality: 50, // Set the compression quality (JPEG)
	})
	if err != nil {
		return "", err
	}

	// Save the processed image to local storage
	err = os.WriteFile(localPath, processedImage, 0644)
	if err != nil {
		return "", err
	}

	fmt.Printf("Image saved as %s\n", localPath)

	return localPath, nil
}
