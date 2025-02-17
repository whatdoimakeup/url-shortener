package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/whatdoimakeup/url-shortener/internal/database"
	urlprocessor "github.com/whatdoimakeup/url-shortener/internal/url-processor"
	"gorm.io/gorm"
)



func Greet(c *fiber.Ctx) error {
	return c.SendString("Hello!")
}

func ShortenUrl(c * fiber.Ctx) error {
	url := c.Query("url")
	db := database.DB

	shortUrl := database.ShortUrl{Original: url}

	err := db.Transaction(func(tx *gorm.DB) error {
        if err := tx.Where("original = ?", url).
            FirstOrCreate(&shortUrl, database.ShortUrl{Original: url}).Error; err != nil {
            return err
        }
        if shortUrl.Short == "" {
            b58url := urlprocessor.Encode(int(shortUrl.ID))
            if err := tx.Model(&shortUrl).Update("Short", b58url).Error; err != nil {
                return err
            }
            shortUrl.Short = b58url
        }
        return nil
    })
    if err != nil {
        log.Fatal(err)
    }
    return c.SendString(shortUrl.Short)

}

func RedirectToOriginal(c * fiber.Ctx) error {
	url := c.Params("short")
	db := database.DB

	origUrl := database.ShortUrl{Short: url}
	result := db.Where("short = ?", url).Limit(1).Find(&origUrl)
	if result.Error != nil {
		return c.Redirect("/")
	}

	return c.Redirect(origUrl.Original)
}