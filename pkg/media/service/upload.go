package media_service

import (
	"io"
	"mime"
	"mime/multipart"
	"net/textproto"
	"os"
	"strings"

	"gitlab.com/mlc-d/ff/pkg/hash"
)

// mediaTypeOrDefault extracts the MIME Type of the file
func mediaTypeOrDefault(header textproto.MIMEHeader) string {
	mediaType, _, err := mime.ParseMediaType(header.Get("Content-Type"))
	if err != nil {
		return "application/octet-stream"
	}
	return mediaType
}

// buildFileName returns the join between "name" and "ext"
// e.g name: "landscape", ext: "png" --> "landscape.png"
func buildFileName(name, ext string) string {
	return name + "." + ext
}

// UploadFile saves file to disk. If the file already exists, it skips that task and just return
// the hash of the file
func (ms *mediaService) uploadFile(file *multipart.FileHeader) (filename string, err error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	md5sum, err := hash.Md5Sum(src)
	if err != nil {
		return "", err
	}

	isBlacklisted, err := ms.repo.IsBlacklisted(md5sum)
	if err != nil || isBlacklisted {
		return "", err
	}

	// build file name
	_, ext, _ := strings.Cut(mediaTypeOrDefault(file.Header), "/")
	filename = buildFileName(md5sum, ext)

	// with this flags, we try to create the file. If it already exists,
	// an error is returned
	dst, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// set position back to start.
	if _, err := src.Seek(0, 0); err != nil {
		return "", err
	}

	// copy to disk
	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}
	return filename, nil
}
