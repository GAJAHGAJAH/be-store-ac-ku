package main

import (
	"github.com/GAJAHGAJAH/be-store-ac-ku/config"
	"github.com/GAJAHGAJAH/be-store-ac-ku/models"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	godotenv.Load()
	config.InitDatabase()
	products := []models.Product{
		{
			Name:        "AC AQUA 1pk",
			Price:       75000,
			Category:    "PendinginRuangan",
			Stock:       50,
			Description: "AC AQUA 1pk kualitas terbaik sedunia",
			ImageURL:    "https://aquaelektronik.com/upload_files/1/56445b50c2-ac-split.png",
		},
		{
			Name:        "AC TRON PREMIUM 2pk",
			Price:       300000,
			Category:    "PendinginRuangan",
			Stock:       20,
			Description: "AC TRON PREMIUM 2pk kualitas terbaik sedunia",
			ImageURL:    "https://utamateknik.co.id/wp-content/uploads/2025/10/AC_Tron-scaled.webp",
		},
		{
			Name:        "AC Polytron 1pk",
			Price:       1200,
			Category:    "PendinginRuangan",
			Stock:       500,
			Description: "AC Polytron 1pk kualitas terbaik sedunia",
			ImageURL:    "https://cdn.polytron.co.id/public-assets/polytroncoid/2026/06/AC-neuva.webp",
		},
		{
			Name:        "AC SHARP 1pk",
			Price:       85000,
			Category:    "PendinginRuangan",
			Stock:       100,
			Description: "AC SHARP 1pk kualitas terbaik sedunia",
			ImageURL:    "https://www.static-src.com/wcsstore/Indraprastha/images/catalog/full/catalog-image/MTA-100454786/sharp_sharp_ah-a9zcy_ac_split_1_pk_standard_turbo_cool_new_-unit_only-_full05_k277wjzc.jpeg",
		},
		{
			Name:        "AC Daikin 1pk",
			Price:       95000,
			Category:    "PendinginRuangan",
			Stock:       40,
			Description: "AC Daikin 1pk kualitas terbaik sedunia",
			ImageURL:    "https://hakapolar.co.id/wp-content/uploads/2025/06/FTC25-1-pk-yv-hakapolar-768x768.jpg",
		},
		{
			Name:        "AC Daikin 2pk",
			Price:       125000,
			Category:    "PendinginRuangan",
			Stock:       25,
			Description: "AC Daikin 2pk kualitas terbaik sedunia",
			ImageURL:    "https://hakapolar.co.id/wp-content/uploads/2025/06/FTC25-1-pk-yv-hakapolar-768x768.jpg",
		},
	}
	for _, p := range products {
		config.DB.Create(&p)
	}
	log.Printf("Seed berhasil: %d produk ditambahkan", len(products))
}
