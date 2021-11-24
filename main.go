package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/qianlifeng/cli-share-file/entity"
	"github.com/qianlifeng/cli-share-file/util"
	"github.com/teris-io/shortid"
	"log"
	"net/url"
	"path"
	"time"
)

func main() {
	app := fiber.New(fiber.Config{
		BodyLimit: 1024 * 1024 * 1024, // file limit up to 1024M
	})

	app.Post("/", upload)
	app.Get("/:id/:name", download)

	postAppStart()

	log.Fatal(app.Listen(":3000"))
}

func upload(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	uploadId := shortid.MustGenerate()
	upload := &entity.Upload{
		Id:        uploadId,
		Location:  path.Join(util.GetUploadFolder(), uploadId),
		FileName:  file.Filename,
		ExpiredAt: time.Now().Add(time.Hour * 24 * 7),
		CreatedAt: time.Now(),
	}

	saveErr := c.SaveFile(file, upload.Location)
	if saveErr != nil {
		return saveErr
	}

	util.GetDB().Save(upload)

	log.Printf("upload file [%s]: %s", upload.Id, upload.FileName)

	return c.SendString(fmt.Sprintf("%s://%s/%s/%s\n", c.Request().URI().Scheme(), c.Request().URI().Host(), upload.Id, url.QueryEscape(upload.FileName)))
}

func download(c *fiber.Ctx) error {
	var upload entity.Upload
	util.GetDB().First(&upload, "id = ?", c.Params("id"))
	if upload.Id == "" {
		return c.SendStatus(404)
	}

	log.Printf("download file [%s]: %s", upload.Id, upload.FileName)

	return c.Download(upload.Location, upload.FileName)
}

func postAppStart() {
	util.EnsureAppFolderExist()
	migrateDB()

	cleanExpiredUploads()
}

func migrateDB() {
	db := util.GetDB()
	db.AutoMigrate(&entity.Upload{})
}

func cleanExpiredUploads() {
	util.GetDB().Where("expired_at < ?", time.Now()).Delete(entity.Upload{})
}
