package models

import (
	"os"
	"fmt"
	"log"
	"encoding/json"
	"io/ioutil"
	"github.com/jinzhu/gorm"
)

func InitVoca(db *gorm.DB) {
	db.Unscoped().Where("text LIKE ?", "%").Delete(Voca{})
	db.Exec("ALTER TABLE vocas AUTO_INCREMENT=1")

	for i := 1; i <= 6; i++ {
		for j := 1; j <= 10; j++ {
			if lessonData, err := os.Open(fmt.Sprintf("static/data/vocabularies/lesson-%d-%d.json", i, j)); err != nil {
				log.Fatal(err.Error())
			} else {
				defer lessonData.Close()
				byteValue, _ := ioutil.ReadAll(lessonData)
				var vocas []map[string]interface{}
				json.Unmarshal([]byte(byteValue), &vocas)
				for k := 0; k < len(vocas); k++ {
					voca := Voca {
						Text: vocas[k]["text"].(string),
						Read: vocas[k]["read"].(string),
						Meaning: vocas[k]["meaning"].(string),
						ExampleText: vocas[k]["example-text"].(string),
						ExampleRead: vocas[k]["example-read"].(string),
						ExampleMeaning: vocas[k]["example-meaning"].(string),
						Lesson: (i - 1) * 10 + j,
					}

					if err := db.Create(&voca).Error; err != nil {
						log.Fatal(err.Error())
					}
				}
				
			}
		}
 	}
}