package main

import (
	"embed"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/qianlifeng/tshare/entity"
	"github.com/qianlifeng/tshare/util"
	"github.com/struCoder/pidusage"
	"github.com/teris-io/shortid"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"time"
)

//go:embed template
var template embed.FS
var startTime time.Time

func main() {
	startTime = time.Now()

	app := fiber.New(fiber.Config{
		Views:     html.NewFileSystem(http.FS(template), ".html"),
		BodyLimit: 1024 * 1024 * 1024, // file limit up to 1024M
	})

	app.Get("/", home)
	app.Post("/", upload)
	app.Get("/:id/:name", download)

	postAppStart()

	port := 3000
	if len(os.Args) > 1 {
		userPort, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic("invalid http port")
		}
		port = userPort
	}

	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}

func home(c *fiber.Ctx) error {
	var totalShares int64
	util.GetDB().Model(&entity.Upload{}).Count(&totalShares)

	var memory float64
	if sysInfo, err := pidusage.GetStat(os.Getpid()); err == nil {
		memory = sysInfo.Memory / 1024 / 1024
	}

	return c.Render("template/index", fiber.Map{
		"total":   totalShares,
		"runDays": fmt.Sprintf("%.0f", time.Now().Sub(startTime).Hours()/24),
		"memory":  fmt.Sprintf("%.1f", memory),
	})
}

func upload(c *fiber.Ctx) error {
	log.Printf("start upload file")
	file, err := c.FormFile("file")
	if err != nil {
		log.Printf("upload error: %s", err.Error())
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
