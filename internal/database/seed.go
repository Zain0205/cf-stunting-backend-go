package database

import (
	"github.com/Zain0205/cf-stunting-backend-go/internal/models"
	"gorm.io/gorm"
)

func SeedAll(db *gorm.DB) error {
	if err := SeedCategories(db); err != nil {
		return err
	}
	if err := SeedDomains(db); err != nil {
		return err
	}
	if err := SeedQuestionsPrakonsepsi(db); err != nil {
		return err
	}
	if err := SeedQuestionsMelahirkan(db); err != nil {
		return err
	}
	if err := SeedQuestionsRemaja(db); err != nil {
		return err
	}
	if err := SeedAnswerMappings(db); err != nil {
		return err
	}
	return nil
}

func SeedCategories(db *gorm.DB) error {
	categories := []models.Category{
		{Code: "PRAKONSEPSI", Name: "Perempuan Prakonsepsi"},
		{Code: "PERNAH_MELAHIRKAN", Name: "Perempuan Pernah Melahirkan"},
		{Code: "REMAJA_19", Name: "Remaja 19 Tahun"},
	}

	for _, cat := range categories {
		if err := db.Where("code = ?", cat.Code).FirstOrCreate(&cat).Error; err != nil {
			return err
		}
	}
	return nil
}

func SeedDomains(db *gorm.DB) error {
	domains := []struct {
		CategoryCode string
		Code         string
		Name         string
	}{
		// PRAKONSEPSI
		{"PRAKONSEPSI", "A", "Gizi dan Suplementasi"},
		{"PRAKONSEPSI", "B", "Ketahanan Pangan"},
		{"PRAKONSEPSI", "C", "Lingkungan dan Perilaku Risiko"},
		{"PRAKONSEPSI", "D", "Psikososial"},

		// PERNAH_MELAHIRKAN
		{"PERNAH_MELAHIRKAN", "A", "Pemberian Makan"},
		{"PERNAH_MELAHIRKAN", "B", "Lingkungan Fisik"},
		{"PERNAH_MELAHIRKAN", "C", "Psikososial"},

		// REMAJA_19
		{"REMAJA_19", "A", "Biologis Intergenerasional"},
		{"REMAJA_19", "B", "Pola Makan"},
		{"REMAJA_19", "C", "Infeksi"},
		{"REMAJA_19", "D", "Sanitasi dan Perilaku"},
		{"REMAJA_19", "E", "Ketahanan Pangan"},
		{"REMAJA_19", "F", "Lingkungan Sosial"},
	}

	for _, d := range domains {
		var category models.Category
		if err := db.Where("code = ?", d.CategoryCode).First(&category).Error; err != nil {
			return err
		}

		domain := models.Domain{
			CategoryID: category.ID,
			Code:       d.Code,
			Name:       d.Name,
		}
		if err := db.Where("category_id = ? AND code = ?", category.ID, d.Code).
			FirstOrCreate(&domain).Error; err != nil {
			return err
		}
	}
	return nil
}

// ==================== PRAKONSEPSI ====================
func SeedQuestionsPrakonsepsi(db *gorm.DB) error {
	var category models.Category
	if err := db.Where("code = ?", "PRAKONSEPSI").First(&category).Error; err != nil {
		return err
	}

	questions := []struct {
		DomainCode string
		Code       string
		Text       string
		CFPakar    float64
		IsReverse  bool
	}{
		// Domain A - Gizi dan Suplementasi
		{"A", "A1", "Dalam 6 bulan terakhir, seberapa sering anda melakukan pemeriksaan Hb (hemoglobin)?", 0.80, false},
		{"A", "A2", "Dalam 1 bulan terakhir, seberapa sering Anda minum suplemen asam folat sesuai anjuran?", 0.90, true},
		{"A", "A3", "Dalam 1 bulan terakhir, seberapa sering Anda minum tablet tambah darah (TTD)?", 0.90, true},
		{"A", "A4", "Seberapa sering Anda menggunakan garam beryodium (dengan logo beryodium) saat memasak di rumah?", 0.85, true},
		{"A", "A5", "Sejauh ini, sejauh mana Anda sudah memiliki rencana atau anjuran tertulis dari tenaga kesehatan tentang suplemen yang akan dikonsumsi saat hamil (asam folat, zat besi, kalsium, dsb.)?", 0.85, true},

		// Domain B - Ketahanan Pangan
		{"B", "B1", "Dalam 12 bulan terakhir, seberapa sering Anda khawatir persediaan makanan di rumah akan habis sebelum bisa membeli lagi?", 0.85, false},
		{"B", "B2", "Dalam 12 bulan terakhir, seberapa sering makanan di rumah benar-benar habis dan Anda tidak punya uang untuk membeli lagi?", 0.90, false},
		{"B", "B3", "Dalam 12 bulan terakhir, seberapa sering anggota keluarga mengurangi ukuran porsi makan karena alasan ekonomi?", 0.90, false},
		{"B", "B4", "Dalam 12 bulan terakhir, seberapa sering Anda mengurangi frekuensi makan per hari karena alasan ekonomi?", 1.00, false},
		{"B", "B5", "Dalam 12 bulan terakhir, seberapa sering Anda tidak makan seharian untuk menghemat makanan/karena tidak ada makanan?", 1.00, false},

		// Domain C - Lingkungan dan Perilaku Risiko
		{"C", "C1", "Dalam 3 bulan terakhir, seberapa sering Anda merokok?", 0.85, false},
		{"C", "C2", "Dalam 7 hari terakhir, berapa hari Anda terpapar asap rokok di rumah atau di tempat kerja?", 0.80, false},
		{"C", "C3", "Dalam 3 bulan terakhir, seberapa sering Anda mengonsumsi minuman beralkohol?", 0.90, false},
		{"C", "C4", "Sumber air minum utama di rumah Anda:", 0.75, false},
		{"C", "C5", "Kepemilikan jamban:", 0.75, false},
		{"C", "C6", "Seberapa konsisten Anda mencuci tangan dengan sabun sebelum menyiapkan dan menyantap makanan?", 0.85, false},

		// Domain D - Psikososial
		{"D", "D1", "Dalam 2 minggu terakhir, seberapa sering Anda merasa sedih atau putus asa?", 0.85, false},
		{"D", "D2", "Dalam 2 minggu terakhir, seberapa sering Anda kehilangan minat atau kenikmatan dalam aktivitas sehari-hari?", 0.80, false},
		{"D", "D3", "Dalam 12 bulan terakhir, seberapa sering Anda merasa tidak aman di rumah atau mengalami kekerasan (emosional/fisik/seksual) dari pasangan/anggota keluarga?", 0.90, false},
		{"D", "D4", "Seberapa besar keinginan Anda untuk mengikuti konseling prakonsepsi (gizi, penyakit kronis, KB/interval, suplementasi)?", 0.75, true},
	}

	for _, q := range questions {
		var domain models.Domain
		if err := db.Where("category_id = ? AND code = ?", category.ID, q.DomainCode).First(&domain).Error; err != nil {
			return err
		}

		question := models.Question{
			DomainID:  domain.ID,
			Code:      q.Code,
			Text:      q.Text,
			CFPakar:   q.CFPakar,
			IsReverse: q.IsReverse,
		}
		if err := db.Where("domain_id = ? AND code = ?", domain.ID, q.Code).
			FirstOrCreate(&question).Error; err != nil {
			return err
		}
	}

	return nil
}

// ==================== PERNAH MELAHIRKAN ====================
func SeedQuestionsMelahirkan(db *gorm.DB) error {
	var category models.Category
	if err := db.Where("code = ?", "PERNAH_MELAHIRKAN").First(&category).Error; err != nil {
		return err
	}

	questions := []struct {
		DomainCode string
		Code       string
		Text       string
		CFPakar    float64
		IsReverse  bool
	}{
		// Domain A - Pemberian Makan
		{"A", "A1", "Berapa lama anak terakhir Anda mendapatkan ASI eksklusif tanpa tambahan makanan/minuman lain?", 0.85, false},
		{"A", "A2", "Dalam 7 hari terakhir, berapa hari anak mendapat MP-ASI berkualitas (protein hewani, buah, sayur)?", 0.85, false},

		// Domain B - Lingkungan Fisik
		{"B", "B1", "Dalam 7 hari terakhir, berapa hari Anda atau anak terpapar asap rokok di rumah?", 0.80, false},
		{"B", "B2", "Dalam 7 hari terakhir, berapa hari rumah menggunakan bahan bakar selain gas elpiji/listrik?", 0.80, false},

		// Domain C - Psikososial
		{"C", "C1", "Dalam 2 minggu terakhir, seberapa sering Anda merasa sedih atau putus asa?", 0.90, false},
		{"C", "C2", "Dalam 12 bulan terakhir, seberapa sering Anda mengalami kekerasan emosional/fisik/seksual dari pasangan/keluarga?", 0.95, false},
	}

	for _, q := range questions {
		var domain models.Domain
		if err := db.Where("category_id = ? AND code = ?", category.ID, q.DomainCode).First(&domain).Error; err != nil {
			return err
		}

		question := models.Question{
			DomainID:  domain.ID,
			Code:      q.Code,
			Text:      q.Text,
			CFPakar:   q.CFPakar,
			IsReverse: q.IsReverse,
		}
		if err := db.Where("domain_id = ? AND code = ?", domain.ID, q.Code).
			FirstOrCreate(&question).Error; err != nil {
			return err
		}
	}

	return nil
}

// ==================== REMAJA 19 TAHUN ====================
func SeedQuestionsRemaja(db *gorm.DB) error {
	var category models.Category
	if err := db.Where("code = ?", "REMAJA_19").First(&category).Error; err != nil {
		return err
	}

	questions := []struct {
		DomainCode string
		Code       string
		Text       string
		CFPakar    float64
		IsReverse  bool
	}{
		// Domain A - Biologis Intergenerasional
		{"A", "A1", "Apakah Anda pernah diberi tahu atau didiagnosis sebagai 'pendek untuk usia Anda' pada masa kanak-kanak atau remaja?", 0.90, false},
		{"A", "A2", "Apakah Anda lahir dengan berat badan kurang dari 2.500 gram?", 0.90, false},
		{"A", "A3", "Apakah jarak kelahiran Anda kurang dari 3 tahun dari kakak kandung terdekat?", 0.80, false},
		{"A", "A4", "Apakah tinggi badan ibu kandung Anda kurang dari 150 cm?", 0.85, false},

		// Domain B - Pola Makan
		{"B", "B1", "Dalam 7 hari terakhir, apakah Anda mengonsumsi ≥ 5 kelompok pangan berbeda (misalnya: nasi/karbohidrat, lauk hewani, lauk nabati, sayur, buah) pada ≥ 4 hari dalam seminggu?", 0.85, true},
		{"B", "B2", "Apakah frekuensi makan utama Anda kurang dari 3 kali per hari pada ≥ 4 hari dalam seminggu?", 0.85, false},
		{"B", "B3", "Dalam 7 hari terakhir, apakah Anda mengonsumsi pangan hewani (telur, ikan, daging, ayam) kurang dari 4 hari dalam seminggu?", 0.80, false},
		{"B", "B4", "Apakah Anda minum minuman berpemanis (teh manis, minuman berenergi, minuman kemasan manis) setiap hari?", 0.70, false},
		{"B", "B5", "Apakah Anda pernah didiagnosis anemia (kurang darah) atau pernah mendapatkan anjuran minum tablet tambah darah (Fe) oleh tenaga kesehatan?", 0.85, false},
		{"B", "B6", "Apakah Anda sering mengalami haid dengan perdarahan sangat banyak (misalnya durasi >7 hari atau sering harus mengganti pembalut karena penuh)?", 0.80, false},

		// Domain C - Infeksi
		{"C", "C1", "Dalam 2 minggu terakhir, apakah Anda mengalami diare ≥ 3 hari berturut-turut atau demam karena infeksi (misal ISPA, tifus, dll.)?", 0.85, false},
		{"C", "C2", "Dalam 6 bulan terakhir, apakah Anda pernah mengalami cacingan atau mendapatkan obat cacing?", 0.80, false},
		{"C", "C3", "Apakah Anda mendapatkan imunisasi dasar lengkap waktu kecil (BCG, DPT, Polio, Campak), sesuai buku KIA atau catatan imunisasi?", 0.75, true},

		// Domain D - Sanitasi dan Perilaku
		{"D", "D1", "Apakah sumber air minum utama di rumah Anda adalah air minum terlindung (PDAM, sumur terlindung, atau air kemasan)?", 0.70, true},
		{"D", "D2", "Apakah rumah Anda memiliki atau menggunakan jamban yang layak (tidak mencemari lingkungan, beralaskan semen, punya saluran pembuangan yang aman)?", 0.70, true},
		{"D", "D3", "Apakah Anda biasanya mencuci tangan dengan sabun pada lima momen penting (sebelum makan, sebelum menyiapkan makanan, setelah dari jamban, setelah membersihkan anak, setelah memegang hewan/kotoran)?", 0.85, false},

		// Domain E - Ketahanan Pangan
		{"E", "E1", "Dalam 12 bulan terakhir, seberapa sering Anda atau keluarga khawatir persediaan makanan akan habis sebelum punya uang untuk membeli lagi?", 0.80, false},
		{"E", "E2", "Dalam 12 bulan terakhir, seberapa sering Anda atau anggota keluarga lain mengurangi ukuran porsi makan karena alasan ekonomi?", 0.90, false},
		{"E", "E3", "Dalam 12 bulan terakhir, seberapa sering Anda atau anggota keluarga lain mengurangi jumlah frekuensi makan per hari karena alasan ekonomi?", 1.00, false},
		{"E", "E4", "Dalam 12 bulan terakhir, seberapa sering Anda atau anggota keluarga lain tidak makan seharian penuh karena tidak ada makanan/untuk menghemat uang?", 1.00, false},
		{"E", "E5", "Dalam 12 bulan terakhir, seberapa sering Anda atau anggota keluarga lain hanya makan makanan rendah mutu (misalnya hanya nasi/karbohidrat tanpa lauk) karena alasan ekonomi?", 0.70, false},

		// Domain F - Lingkungan Sosial
		{"F", "F1", "Apakah Anda tinggal di daerah pedesaan terpencil atau di lingkungan kumuh perkotaan?", 0.70, false},
		{"F", "F2", "Apakah ada anggota keluarga yang merokok di dalam rumah secara rutin?", 0.75, false},
		{"F", "F3", "Apakah Anda saat ini sedang hamil atau pernah hamil pada usia <20 tahun?", 0.80, false},
		{"F", "F4", "Apakah Anda saat ini sudah menikah?", 0.70, false},
	}

	for _, q := range questions {
		var domain models.Domain
		if err := db.Where("category_id = ? AND code = ?", category.ID, q.DomainCode).First(&domain).Error; err != nil {
			return err
		}

		question := models.Question{
			DomainID:  domain.ID,
			Code:      q.Code,
			Text:      q.Text,
			CFPakar:   q.CFPakar,
			IsReverse: q.IsReverse,
		}
		if err := db.Where("domain_id = ? AND code = ?", domain.ID, q.Code).
			FirstOrCreate(&question).Error; err != nil {
			return err
		}
	}

	return nil
}

// ==================== ANSWER MAPPINGS ====================
func SeedAnswerMappings(db *gorm.DB) error {
	// Helper function to get question
	getQuestion := func(categoryCode, questionCode string) (*models.Question, error) {
		var question models.Question
		err := db.Joins("JOIN domains ON domains.id = questions.domain_id").
			Joins("JOIN categories ON categories.id = domains.category_id").
			Where("categories.code = ? AND questions.code = ?", categoryCode, questionCode).
			First(&question).Error
		return &question, err
	}

	// ========== PRAKONSEPSI ==========

	// A1 - Frekuensi pemeriksaan Hb
	if q, err := getQuestion("PRAKONSEPSI", "A1"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "0", CFEvidence: 0.0},
			{QuestionID: q.ID, AnswerKey: "1", CFEvidence: 0.4},
			{QuestionID: q.ID, AnswerKey: "2", CFEvidence: 0.7},
			{QuestionID: q.ID, AnswerKey: "3", CFEvidence: 1.0},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// A2 - Frekuensi minum asam folat (REVERSED)
	if q, err := getQuestion("PRAKONSEPSI", "A2"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "0", CFEvidence: 1.0}, // Reversed
			{QuestionID: q.ID, AnswerKey: "1", CFEvidence: 0.7},
			{QuestionID: q.ID, AnswerKey: "2", CFEvidence: 0.4},
			{QuestionID: q.ID, AnswerKey: "3", CFEvidence: 0.0},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// A3 - Frekuensi minum TTD (REVERSED)
	if q, err := getQuestion("PRAKONSEPSI", "A3"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "0", CFEvidence: 1.0}, // Reversed
			{QuestionID: q.ID, AnswerKey: "1", CFEvidence: 0.7},
			{QuestionID: q.ID, AnswerKey: "2", CFEvidence: 0.4},
			{QuestionID: q.ID, AnswerKey: "3", CFEvidence: 0.0},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// A4 - Penggunaan garam beryodium (REVERSED)
	if q, err := getQuestion("PRAKONSEPSI", "A4"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "0", CFEvidence: 1.0}, // Reversed
			{QuestionID: q.ID, AnswerKey: "1", CFEvidence: 0.7},
			{QuestionID: q.ID, AnswerKey: "2", CFEvidence: 0.4},
			{QuestionID: q.ID, AnswerKey: "3", CFEvidence: 0.0},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// A5 - Rencana suplementasi hamil (REVERSED)
	if q, err := getQuestion("PRAKONSEPSI", "A5"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "0", CFEvidence: 1.0}, // Reversed
			{QuestionID: q.ID, AnswerKey: "1", CFEvidence: 0.7},
			{QuestionID: q.ID, AnswerKey: "2", CFEvidence: 0.4},
			{QuestionID: q.ID, AnswerKey: "3", CFEvidence: 0.0},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// B1 - Kekhawatiran pangan habis
	if q, err := getQuestion("PRAKONSEPSI", "B1"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "0", CFEvidence: 0.0},
			{QuestionID: q.ID, AnswerKey: "1", CFEvidence: 0.4},
			{QuestionID: q.ID, AnswerKey: "2", CFEvidence: 0.7},
			{QuestionID: q.ID, AnswerKey: "3", CFEvidence: 1.0},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// B2 - Makanan benar-benar habis
	if q, err := getQuestion("PRAKONSEPSI", "B2"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "0", CFEvidence: 0.0},
			{QuestionID: q.ID, AnswerKey: "1", CFEvidence: 0.4},
			{QuestionID: q.ID, AnswerKey: "2", CFEvidence: 0.7},
			{QuestionID: q.ID, AnswerKey: "3", CFEvidence: 1.0},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// B3 - Mengurangi porsi makan
	if q, err := getQuestion("PRAKONSEPSI", "B3"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "0", CFEvidence: 0.0},
			{QuestionID: q.ID, AnswerKey: "1", CFEvidence: 0.4},
			{QuestionID: q.ID, AnswerKey: "2", CFEvidence: 0.7},
			{QuestionID: q.ID, AnswerKey: "3", CFEvidence: 1.0},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// B4 - Mengurangi frekuensi makan
	if q, err := getQuestion("PRAKONSEPSI", "B4"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "0", CFEvidence: 0.0},
			{QuestionID: q.ID, AnswerKey: "1", CFEvidence: 0.4},
			{QuestionID: q.ID, AnswerKey: "2", CFEvidence: 0.7},
			{QuestionID: q.ID, AnswerKey: "3", CFEvidence: 1.0},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// B5 - Tidak makan seharian
	if q, err := getQuestion("PRAKONSEPSI", "B5"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "0", CFEvidence: 0.0},
			{QuestionID: q.ID, AnswerKey: "1", CFEvidence: 0.4},
			{QuestionID: q.ID, AnswerKey: "2", CFEvidence: 0.7},
			{QuestionID: q.ID, AnswerKey: "3", CFEvidence: 1.0},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// C1 - Frekuensi merokok
	if q, err := getQuestion("PRAKONSEPSI", "C1"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "0", CFEvidence: 0.0},
			{QuestionID: q.ID, AnswerKey: "1", CFEvidence: 0.4},
			{QuestionID: q.ID, AnswerKey: "2", CFEvidence: 0.7},
			{QuestionID: q.ID, AnswerKey: "3", CFEvidence: 1.0},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// C2 - Hari terpapar asap rokok
	if q, err := getQuestion("PRAKONSEPSI", "C2"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "0", CFEvidence: 0.0},
			{QuestionID: q.ID, AnswerKey: "1", CFEvidence: 0.4},
			{QuestionID: q.ID, AnswerKey: "2", CFEvidence: 0.7},
			{QuestionID: q.ID, AnswerKey: "3", CFEvidence: 1.0},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// C3 - Frekuensi konsumsi alkohol
	if q, err := getQuestion("PRAKONSEPSI", "C3"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "0", CFEvidence: 0.0},
			{QuestionID: q.ID, AnswerKey: "1", CFEvidence: 0.4},
			{QuestionID: q.ID, AnswerKey: "2", CFEvidence: 0.7},
			{QuestionID: q.ID, AnswerKey: "3", CFEvidence: 1.0},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// C4 - Sumber air minum (Binary: 0=aman, 2=risiko)
	if q, err := getQuestion("PRAKONSEPSI", "C4"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "0", CFEvidence: 0.0},
			{QuestionID: q.ID, AnswerKey: "2", CFEvidence: 1.0},
		}

		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// C5 - Kepemilikan jamban
	if q, err := getQuestion("PRAKONSEPSI", "C5"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "0", CFEvidence: 0.0},
			{QuestionID: q.ID, AnswerKey: "1", CFEvidence: 0.4},
			{QuestionID: q.ID, AnswerKey: "2", CFEvidence: 1.0},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// C6 - Kebiasaan cuci tangan
	if q, err := getQuestion("PRAKONSEPSI", "C6"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "0", CFEvidence: 0.0},
			{QuestionID: q.ID, AnswerKey: "1", CFEvidence: 0.4},
			{QuestionID: q.ID, AnswerKey: "2", CFEvidence: 0.7},
			{QuestionID: q.ID, AnswerKey: "3", CFEvidence: 1.0},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// D1 - Perasaan sedih
	if q, err := getQuestion("PRAKONSEPSI", "D1"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "0", CFEvidence: 0.0},
			{QuestionID: q.ID, AnswerKey: "1", CFEvidence: 0.4},
			{QuestionID: q.ID, AnswerKey: "2", CFEvidence: 0.7},
			{QuestionID: q.ID, AnswerKey: "3", CFEvidence: 1.0},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// D2 - Kehilangan minat
	if q, err := getQuestion("PRAKONSEPSI", "D2"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "0", CFEvidence: 0.0},
			{QuestionID: q.ID, AnswerKey: "1", CFEvidence: 0.4},
			{QuestionID: q.ID, AnswerKey: "2", CFEvidence: 0.7},
			{QuestionID: q.ID, AnswerKey: "3", CFEvidence: 1.0},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// D3 - Merasa tidak aman/kekerasan
	if q, err := getQuestion("PRAKONSEPSI", "D3"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "0", CFEvidence: 0.0},
			{QuestionID: q.ID, AnswerKey: "1", CFEvidence: 0.4},
			{QuestionID: q.ID, AnswerKey: "2", CFEvidence: 0.7},
			{QuestionID: q.ID, AnswerKey: "3", CFEvidence: 1.0},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// D4 - Minat konseling (REVERSED - Protektif)
	if q, err := getQuestion("PRAKONSEPSI", "D4"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "0", CFEvidence: 0.0}, // Sangat berminat = tidak berisiko
			{QuestionID: q.ID, AnswerKey: "1", CFEvidence: 0.4},
			{QuestionID: q.ID, AnswerKey: "2", CFEvidence: 0.7},
			{QuestionID: q.ID, AnswerKey: "3", CFEvidence: 1.0},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// ========== PERNAH MELAHIRKAN ==========

	// A1 - Lama ASI eksklusif
	if q, err := getQuestion("PERNAH_MELAHIRKAN", "A1"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "0", CFEvidence: 0.0},
			{QuestionID: q.ID, AnswerKey: "1", CFEvidence: 0.4},
			{QuestionID: q.ID, AnswerKey: "2", CFEvidence: 0.7},
			{QuestionID: q.ID, AnswerKey: "3", CFEvidence: 1.0},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// A2 - Frekuensi MP-ASI berkualitas
	if q, err := getQuestion("PERNAH_MELAHIRKAN", "A2"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "0", CFEvidence: 0.0},
			{QuestionID: q.ID, AnswerKey: "1", CFEvidence: 0.4},
			{QuestionID: q.ID, AnswerKey: "2", CFEvidence: 0.7},
			{QuestionID: q.ID, AnswerKey: "3", CFEvidence: 1.0},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// B1 - Paparan asap rokok
	if q, err := getQuestion("PERNAH_MELAHIRKAN", "B1"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "0", CFEvidence: 0.0},
			{QuestionID: q.ID, AnswerKey: "1", CFEvidence: 0.4},
			{QuestionID: q.ID, AnswerKey: "2", CFEvidence: 0.7},
			{QuestionID: q.ID, AnswerKey: "3", CFEvidence: 1.0},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// B2 - Bahan bakar selain LPG
	if q, err := getQuestion("PERNAH_MELAHIRKAN", "B2"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "0", CFEvidence: 0.0},
			{QuestionID: q.ID, AnswerKey: "1", CFEvidence: 0.4},
			{QuestionID: q.ID, AnswerKey: "2", CFEvidence: 0.7},
			{QuestionID: q.ID, AnswerKey: "3", CFEvidence: 1.0},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// C1 - Perasaan sedih/putus asa
	if q, err := getQuestion("PERNAH_MELAHIRKAN", "C1"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "0", CFEvidence: 0.0},
			{QuestionID: q.ID, AnswerKey: "1", CFEvidence: 0.4},
			{QuestionID: q.ID, AnswerKey: "2", CFEvidence: 0.7},
			{QuestionID: q.ID, AnswerKey: "3", CFEvidence: 1.0},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// C2 - Kekerasan
	if q, err := getQuestion("PERNAH_MELAHIRKAN", "C2"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "0", CFEvidence: 0.0},
			{QuestionID: q.ID, AnswerKey: "1", CFEvidence: 0.4},
			{QuestionID: q.ID, AnswerKey: "2", CFEvidence: 0.7},
			{QuestionID: q.ID, AnswerKey: "3", CFEvidence: 1.0},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// ========== REMAJA 19 TAHUN ==========

	// A1 - Pernah dibilang pendek (Ya/Tidak/Tidak Tahu)
	if q, err := getQuestion("REMAJA_19", "A1"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "YA", CFEvidence: 1.0},
			{QuestionID: q.ID, AnswerKey: "TIDAK", CFEvidence: 0.0},
			{QuestionID: q.ID, AnswerKey: "TIDAK_TAHU", CFEvidence: 0.4},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// A2 - Lahir BBLR
	if q, err := getQuestion("REMAJA_19", "A2"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "YA", CFEvidence: 1.0},
			{QuestionID: q.ID, AnswerKey: "TIDAK", CFEvidence: 0.0},
			{QuestionID: q.ID, AnswerKey: "TIDAK_TAHU", CFEvidence: 0.4},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// A3 - Jarak kelahiran <3 tahun
	if q, err := getQuestion("REMAJA_19", "A3"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "YA", CFEvidence: 1.0},
			{QuestionID: q.ID, AnswerKey: "TIDAK", CFEvidence: 0.0},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// A4 - Tinggi ibu <150cm
	if q, err := getQuestion("REMAJA_19", "A4"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "YA", CFEvidence: 1.0},
			{QuestionID: q.ID, AnswerKey: "TIDAK", CFEvidence: 0.0},
			{QuestionID: q.ID, AnswerKey: "TIDAK_TAHU", CFEvidence: 0.4},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// B1 - ≥5 kelompok pangan (REVERSED)
	if q, err := getQuestion("REMAJA_19", "B1"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "YA", CFEvidence: 0.0},          // Baik
			{QuestionID: q.ID, AnswerKey: "TIDAK", CFEvidence: 0.8},       // Risiko
			{QuestionID: q.ID, AnswerKey: "TIDAK_INGAT", CFEvidence: 0.4}, // Cukup berisiko
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// B2 - Frekuensi makan <3x
	if q, err := getQuestion("REMAJA_19", "B2"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "YA", CFEvidence: 0.8},
			{QuestionID: q.ID, AnswerKey: "TIDAK", CFEvidence: 0.0},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// B3 - Konsumsi hewani <4 hari
	if q, err := getQuestion("REMAJA_19", "B3"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "YA", CFEvidence: 0.8},
			{QuestionID: q.ID, AnswerKey: "TIDAK", CFEvidence: 0.0},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// B4 - Minuman manis setiap hari
	if q, err := getQuestion("REMAJA_19", "B4"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "YA", CFEvidence: 0.8},
			{QuestionID: q.ID, AnswerKey: "TIDAK", CFEvidence: 0.0},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// B5 - Riwayat anemia
	if q, err := getQuestion("REMAJA_19", "B5"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "YA", CFEvidence: 0.8},
			{QuestionID: q.ID, AnswerKey: "TIDAK", CFEvidence: 0.0},
			{QuestionID: q.ID, AnswerKey: "TIDAK_TAHU", CFEvidence: 0.4},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// B6 - Perdarahan banyak
	if q, err := getQuestion("REMAJA_19", "B6"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "YA", CFEvidence: 0.8},
			{QuestionID: q.ID, AnswerKey: "TIDAK", CFEvidence: 0.0},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// C1 - Riwayat diare/demam
	if q, err := getQuestion("REMAJA_19", "C1"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "YA", CFEvidence: 0.8},
			{QuestionID: q.ID, AnswerKey: "TIDAK", CFEvidence: 0.0},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// C2 - Riwayat cacingan
	if q, err := getQuestion("REMAJA_19", "C2"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "YA", CFEvidence: 0.8},
			{QuestionID: q.ID, AnswerKey: "TIDAK", CFEvidence: 0.0},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// C3 - Imunisasi dasar lengkap (REVERSED)
	if q, err := getQuestion("REMAJA_19", "C3"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "YA", CFEvidence: 0.0},
			{QuestionID: q.ID, AnswerKey: "TIDAK", CFEvidence: 0.8},
			{QuestionID: q.ID, AnswerKey: "TIDAK_TAHU", CFEvidence: 0.4},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// D1 - Air minum terlindung (REVERSED)
	if q, err := getQuestion("REMAJA_19", "D1"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "YA", CFEvidence: 0.0},
			{QuestionID: q.ID, AnswerKey: "TIDAK", CFEvidence: 0.8},
			{QuestionID: q.ID, AnswerKey: "TIDAK_TAHU", CFEvidence: 0.4},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// D2 - Jamban layak (REVERSED)
	if q, err := getQuestion("REMAJA_19", "D2"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "YA", CFEvidence: 0.0},
			{QuestionID: q.ID, AnswerKey: "TIDAK", CFEvidence: 0.8},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// D3 - Cuci tangan dengan sabun
	if q, err := getQuestion("REMAJA_19", "D3"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "0", CFEvidence: 0.0}, // Selalu
			{QuestionID: q.ID, AnswerKey: "1", CFEvidence: 0.2}, // Sering
			{QuestionID: q.ID, AnswerKey: "2", CFEvidence: 0.5}, // Kadang
			{QuestionID: q.ID, AnswerKey: "3", CFEvidence: 0.8}, // Jarang/Tidak
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// E1 - Kekhawatiran persediaan makanan
	if q, err := getQuestion("REMAJA_19", "E1"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "0", CFEvidence: 0.0},
			{QuestionID: q.ID, AnswerKey: "1", CFEvidence: 0.4},
			{QuestionID: q.ID, AnswerKey: "2", CFEvidence: 0.7},
			{QuestionID: q.ID, AnswerKey: "3", CFEvidence: 1.0},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// E2 - Mengurangi porsi makan
	if q, err := getQuestion("REMAJA_19", "E2"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "0", CFEvidence: 0.0},
			{QuestionID: q.ID, AnswerKey: "1", CFEvidence: 0.4},
			{QuestionID: q.ID, AnswerKey: "2", CFEvidence: 0.7},
			{QuestionID: q.ID, AnswerKey: "3", CFEvidence: 1.0},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// E3 - Mengurangi frekuensi makan
	if q, err := getQuestion("REMAJA_19", "E3"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "0", CFEvidence: 0.0},
			{QuestionID: q.ID, AnswerKey: "1", CFEvidence: 0.4},
			{QuestionID: q.ID, AnswerKey: "2", CFEvidence: 0.7},
			{QuestionID: q.ID, AnswerKey: "3", CFEvidence: 1.0},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// E4 - Tidak makan seharian
	if q, err := getQuestion("REMAJA_19", "E4"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "0", CFEvidence: 0.0},
			{QuestionID: q.ID, AnswerKey: "1", CFEvidence: 0.4},
			{QuestionID: q.ID, AnswerKey: "2", CFEvidence: 0.7},
			{QuestionID: q.ID, AnswerKey: "3", CFEvidence: 1.0},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// E5 - Makanan rendah mutu
	if q, err := getQuestion("REMAJA_19", "E5"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "0", CFEvidence: 0.0},
			{QuestionID: q.ID, AnswerKey: "1", CFEvidence: 0.4},
			{QuestionID: q.ID, AnswerKey: "2", CFEvidence: 0.7},
			{QuestionID: q.ID, AnswerKey: "3", CFEvidence: 1.0},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// F1 - Tinggal di daerah terpencil/kumuh
	if q, err := getQuestion("REMAJA_19", "F1"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "YA", CFEvidence: 0.8},
			{QuestionID: q.ID, AnswerKey: "TIDAK", CFEvidence: 0.0},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// F2 - Anggota keluarga merokok di dalam rumah
	if q, err := getQuestion("REMAJA_19", "F2"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "YA", CFEvidence: 0.8},
			{QuestionID: q.ID, AnswerKey: "TIDAK", CFEvidence: 0.0},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// F3 - Riwayat hamil <20 tahun
	if q, err := getQuestion("REMAJA_19", "F3"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "YA", CFEvidence: 0.8},
			{QuestionID: q.ID, AnswerKey: "TIDAK", CFEvidence: 0.0},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	// F4 - Sudah menikah
	if q, err := getQuestion("REMAJA_19", "F4"); err == nil {
		mappings := []models.AnswerMapping{
			{QuestionID: q.ID, AnswerKey: "YA", CFEvidence: 0.8},
			{QuestionID: q.ID, AnswerKey: "TIDAK", CFEvidence: 0.0},
		}
		for _, m := range mappings {
			db.Where("question_id = ? AND answer_key = ?", m.QuestionID, m.AnswerKey).FirstOrCreate(&m)
		}
	}

	return nil
}
