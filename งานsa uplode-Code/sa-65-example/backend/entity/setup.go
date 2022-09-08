package entity

import (
	"gorm.io/gorm"

	"gorm.io/driver/sqlite"
)

var db *gorm.DB

func DB() *gorm.DB {

	return db

}

func SetupDatabase() {

	database, err := gorm.Open(sqlite.Open("sa-65.db"), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")

	}

	// Migrate the schema

	database.AutoMigrate(
		&User{},
		&Patient{},
		&Symptom{},
		&Covid{},
		&Screening{},
	)

	db = database
	db.Model(&User{}).Create(&User{Name: "Nawamin", Email: "nwmin14@gmail.com"})
	db.Model(&User{}).Create(&User{Name: "None", Email: "none@example.com"})
	var Nawamin User
	var None User
	db.Raw("SELECT * FROM users WHERE email = ?", "nwmin14@gmail.com").Scan(&Nawamin)
	db.Raw("SELECT * FROM users WHERE email = ?", "none@example.com").Scan(&None)

	db.Model(&Patient{}).Create(&Patient{Name: "Prayut", Email: "Prayut@godmail.com"})
	db.Model(&Patient{}).Create(&Patient{Name: "Prawit", Email: "Prawit@sleep.com"})
	db.Model(&Patient{}).Create(&Patient{Name: "Thanathon", Email: "Thanathon@marathon.com"})
	var Prayut Patient
	var Prawit Patient
	var Thanathon Patient
	db.Raw("SELECT * FROM patients WHERE email = ?", "Prayut@godmail.com").Scan(&Prayut)
	db.Raw("SELECT * FROM patients WHERE email = ?", "Prawit@sleep.com").Scan(&Prawit)
	db.Raw("SELECT * FROM patients WHERE email = ?", "Thanathon@marathon.com").Scan(&Thanathon)

	normal := Symptom{Name: "ปกติ"}
	db.Model(Symptom{}).Create(&normal)
	fever := Symptom{Name: "มีไข้"}
	db.Model(Symptom{}).Create(&fever)
	ill := Symptom{Name: "มีไข้ ไอ เจ็บคอ"}
	db.Model(Symptom{}).Create(&ill)

	never := Covid{Name: "ไม่เคย"}
	db.Model(Covid{}).Create(&never)
	aweek := Covid{Name: "1-7วัน"}
	db.Model(Covid{}).Create(&aweek)
	hmonth := Covid{Name: "7-14วัน"}
	db.Model(Covid{}).Create(&hmonth)

	db.Model(Screening{}).Create(&Screening{Patient: Prayut, Symptom: normal, Covid: never, User: Nawamin})
	db.Model(Screening{}).Create(&Screening{Patient: Prawit, Symptom: fever, Covid: aweek, User: Nawamin})
	db.Model(Screening{}).Create(&Screening{Patient: Thanathon, Symptom: ill, Covid: hmonth, User: None})


	//==== Query ====
	var target Patient
	db.Model(&Patient{}).Find(&target,db.Where("email=?","Prayut@godmail.com"))
	var whoTarget Patient
	db.Model(&Patient{}).Find(&whoTarget,db.Where("name = ? and patient_id = ?",target.ID))
	var outScreening []*Screening
	db.Model(&Screening{}).Joins("Symptom").Joins("Covid").Joins("User").Find(&outScreening, db.Where("patient_id = ?",whoTarget.ID))
	
	for, os := range outScreening {
		fmt.Printf("Out Screening: %v\n",os.ID)
		fmt.Printf("Out Screening: %v\n",os.ID)    // กลับมาแก้ด้วย
		fmt.Printf("Out Screening: %v\n",os.ID)
		fmt.Printf("Out Screening: %v\n",os.ID)
		fmt.Printf("Out Screening: %v\n",os.ID)
	}
}
