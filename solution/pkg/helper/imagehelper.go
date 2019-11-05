package helper

import (
	"crypto/rand"
	"fmt"
	"io"
	"os"

	"github.com/pdfcrowd/pdfcrowd-go"
)

//ImageHelper will hold the ImageHelper configuration
type ImageHelper struct {
	APIUser string
	APIkey  string
	Path    string
}

//NewImageHelper will fill the Image Helper object
func (imgHelper *ImageHelper) NewImageHelper(parImageHelper *ImageHelper) *ImageHelper {
	pvImageHelper := new(ImageHelper)
	pvImageHelper.APIUser = parImageHelper.APIUser
	pvImageHelper.APIkey = parImageHelper.APIkey
	pvImageHelper.Path = parImageHelper.Path
	return pvImageHelper

}

//StartTakeScreenShot to add the screenshot images to the folder path
func (imgHelper *ImageHelper) StartTakeScreenShot(parImgHelper *ImageHelper, parUrls []string) {
	go imgHelper.urlToImages(parImgHelper, parUrls)
}

func (imgHelper *ImageHelper) urlToImages(parImgHelper *ImageHelper, parUrls []string) {

	pvClient := pdfcrowd.NewHtmlToImageClient(parImgHelper.APIUser, parImgHelper.APIkey)
	pvClient.SetOutputFormat("png")

	go func() {
		for _, urlsData := range parUrls {
			pvUID, pvError := newUUID()
			if pvError != nil {
				panic(pvError)
			}
			pvOutPutStream, pvError := os.Create(parImgHelper.Path + fmt.Sprint(pvUID) + ".png")
			if pvError != nil {
				handleError(pvError)

			}
			pvError = pvClient.ConvertUrlToStream(urlsData, pvOutPutStream)
			if pvError != nil {
				handleError(pvError)
			}
			defer pvOutPutStream.Close()
		}
	}()
}

// newUUID generates a random UUID according to RFC 4122
func newUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
func handleError(parErr error) {
	if parErr != nil {

		pvErrorMsg, pvOk := parErr.(pdfcrowd.Error)
		if pvOk {
			os.Stderr.WriteString(fmt.Sprintf("Pdfcrowd Error: %s\n", pvErrorMsg))
		} else {
			os.Stderr.WriteString(fmt.Sprintf("Generic Error: %s\n", parErr))
		}

		panic(parErr.Error())
	}
}
