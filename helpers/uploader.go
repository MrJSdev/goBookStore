package helpers

import (
	"path/filepath"

	"github.com/kataras/iris/v12"
)

const maxSize = 8 * iris.MB

func FileUploader(ctx iris.Context) {
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	ctx.SetMaxRequestBodySize(maxSize)
	// OR
	// app.Use(iris.LimitRequestBodySize(maxSize))
	// OR
	// OR iris.WithPostMaxMemory(maxSize)

	// single file
	_, fileHeader, err := ctx.FormFile("file")
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	// Upload the file to specific destination.
	dest := filepath.Join("./uploads", fileHeader.Filename)
	ctx.SaveFormFile(fileHeader, dest)

	ctx.Writef("File: %s uploaded!", fileHeader.Filename)

}
