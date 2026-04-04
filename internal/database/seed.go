package database

import (
	"github.com/Zain0205/cf-stunting-backend-go/internal/models"

	"gorm.io/gorm"
)

func SeedDatabase(db *gorm.DB) error {
	if err := seedCategories(db); err != nil {
		return err
	}
	if err := seedDomains(db); err != nil {
		return err
	}
	if err := seedQuestions(db); err != nil {
		return err
	}
	if err := seedAnswerMappings(db); err != nil {
		return err
	}
	return nil
}

func seedCategories(db *gorm.DB) error {
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

func seedDomains(db *gorm.DB) error {
	domains := []struct {
		CategoryCode string
		Code         string
		Name         string
	}{
		// Prakonsepsi
		{"PRAKONSEPSI", "A", "Gizi dan Suplementasi"},
		{"PRAKONSEPSI", "B", "Ketahanan Pangan"},
		{"PRAKONSEPSI", "C", "Lingkungan dan Perilaku Risiko"},
		{"PRAKONSEPSI", "D", "Psikososial"},

		// Pernah Melahirkan
		{"PERNAH_MELAHIRKAN", "A", "Pemberian Makan"},
		{"PERNAH_MELAHIRKAN", "B", "Lingkungan Fisik"},
		{"PERNAH_MELAHIRKAN", "C", "Psikososial"},

		// Remaja 19 Tahun
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
		if err := db.Where("category_id = ? AND code = ?", category.ID, d.Code).FirstOrCreate(&domain).Error; err != nil {
			return err
		}
	}
	return nil
}

func seedQuestions(db *gorm.DB) error {
	// CFPakar sesuai Tabel 2 paper.
	// IsReverse = true hanya sebagai penanda UI/display bahwa skala jawaban
	// bersifat protektif (jawaban terbaik = key "0"). CFEvidence sudah
	// mencerminkan nilai risiko yang benar; engine TIDAK perlu membalik
	// evidence secara otomatis berdasarkan flag ini.
	questions := []struct {
		CategoryCode string
		DomainCode   string
		Code         string
		Text         string
		CFPakar      float64
		IsReverse    bool
	}{
		// ===== PRAKONSEPSI - DOMAIN A =====
		{"PRAKONSEPSI", "A", "A1", "Dalam 6 bulan terakhir, seberapa sering anda melakukan pemeriksaan Hb (hemoglobin)?", 0.80, false},
		{"PRAKONSEPSI", "A", "A2", "Dalam 1 bulan terakhir, seberapa sering Anda minum suplemen asam folat sesuai anjuran?", 0.90, true},
		{"PRAKONSEPSI", "A", "A3", "Dalam 1 bulan terakhir, seberapa sering Anda minum tablet tambah darah (TTD)?", 0.90, true},
		{"PRAKONSEPSI", "A", "A4", "Seberapa sering Anda menggunakan garam beryodium (dengan logo beryodium) saat memasak di rumah?", 0.85, true},
		{"PRAKONSEPSI", "A", "A5", "Sejauh ini, sejauh mana Anda sudah memiliki rencana atau anjuran tertulis dari tenaga kesehatan tentang suplemen yang akan dikonsumsi saat hamil?", 0.85, true},

		// ===== PRAKONSEPSI - DOMAIN B =====
		{"PRAKONSEPSI", "B", "B1", "Dalam 12 bulan terakhir, seberapa sering Anda khawatir persediaan makanan di rumah akan habis sebelum bisa membeli lagi?", 0.85, false},
		{"PRAKONSEPSI", "B", "B2", "Dalam 12 bulan terakhir, seberapa sering makanan di rumah benar-benar habis dan Anda tidak punya uang untuk membeli lagi?", 0.90, false},
		{"PRAKONSEPSI", "B", "B3", "Dalam 12 bulan terakhir, seberapa sering anggota keluarga mengurangi ukuran porsi makan karena alasan ekonomi?", 0.90, false},
		{"PRAKONSEPSI", "B", "B4", "Dalam 12 bulan terakhir, seberapa sering anggota keluarga mengurangi frekuensi makan per hari karena ekonomi?", 1.00, false},
		{"PRAKONSEPSI", "B", "B5", "Dalam 12 bulan terakhir, seberapa sering tidak makan seharian untuk menghemat makanan/karena tidak ada makanan?", 1.00, false},

		// ===== PRAKONSEPSI - DOMAIN C =====
		{"PRAKONSEPSI", "C", "C1", "Dalam 3 bulan terakhir, seberapa sering Anda merokok?", 0.85, false},
		{"PRAKONSEPSI", "C", "C2", "Dalam 7 hari terakhir, berapa hari Anda terpapar asap rokok di rumah atau di tempat kerja?", 0.80, false},
		{"PRAKONSEPSI", "C", "C3", "Dalam 3 bulan terakhir, seberapa sering Anda mengonsumsi minuman beralkohol?", 0.90, false},
		{"PRAKONSEPSI", "C", "C4", "Sumber air minum utama di rumah Anda:", 0.75, false},
		{"PRAKONSEPSI", "C", "C5", "Kepemilikan jamban:", 0.75, false},
		// C6: skala risiko — Selalu=0.0 (baik/tidak berisiko), Jarang=1.0 (sangat berisiko)
		// IsReverse=false karena CFEvidence sudah mencerminkan nilai risiko langsung
		{"PRAKONSEPSI", "C", "C6", "Seberapa konsisten Anda mencuci tangan dengan sabun sebelum menyiapkan dan menyantap makanan?", 0.85, false},

		// ===== PRAKONSEPSI - DOMAIN D =====
		{"PRAKONSEPSI", "D", "D1", "Dalam 2 minggu terakhir, seberapa sering Anda merasa sedih atau putus asa?", 0.85, false},
		{"PRAKONSEPSI", "D", "D2", "Dalam 2 minggu terakhir, seberapa sering Anda kehilangan minat atau kenikmatan dalam aktivitas sehari-hari?", 0.80, false},
		{"PRAKONSEPSI", "D", "D3", "Dalam 12 bulan terakhir, seberapa sering Anda merasa tidak aman di rumah atau mengalami kekerasan?", 0.90, false},
		// D4: protektif/dibalik — Sangat berminat = risiko rendah (0.0), Tidak berminat = risiko tinggi (1.0)
		{"PRAKONSEPSI", "D", "D4", "Seberapa besar keinginan Anda untuk mengikuti konseling prakonsepsi?", 0.75, true},

		// ===== PERNAH MELAHIRKAN - DOMAIN A =====
		// A1 & A2: kondisi baik (ASI eksklusif lama / MP-ASI sering) = risiko rendah (0.0)
		{"PERNAH_MELAHIRKAN", "A", "A1", "Berapa lama anak terakhir Anda mendapatkan ASI eksklusif tanpa tambahan makanan/minuman lain?", 0.85, true},
		{"PERNAH_MELAHIRKAN", "A", "A2", "Dalam 7 hari terakhir, berapa hari anak mendapat MP-ASI berkualitas (protein hewani, buah, sayur)?", 0.85, true},

		// ===== PERNAH MELAHIRKAN - DOMAIN B =====
		{"PERNAH_MELAHIRKAN", "B", "B1", "Dalam 7 hari terakhir, berapa hari Anda atau anak terpapar asap rokok di rumah?", 0.80, false},
		{"PERNAH_MELAHIRKAN", "B", "B2", "Dalam 7 hari terakhir, berapa hari rumah menggunakan bahan bakar selain gas elpiji/listrik?", 0.80, false},

		// ===== PERNAH MELAHIRKAN - DOMAIN C =====
		{"PERNAH_MELAHIRKAN", "C", "C1", "Dalam 2 minggu terakhir, seberapa sering Anda merasa sedih atau putus asa?", 0.90, false},
		{"PERNAH_MELAHIRKAN", "C", "C2", "Dalam 12 bulan terakhir, seberapa sering Anda mengalami kekerasan emosional/fisik/seksual dari pasangan/keluarga?", 0.95, false},

		// ===== REMAJA 19 - DOMAIN A =====
		{"REMAJA_19", "A", "A1", "Apakah Anda pernah diberi tahu atau didiagnosis sebagai 'pendek untuk usia Anda' pada masa kanak-kanak atau remaja?", 0.90, false},
		{"REMAJA_19", "A", "A2", "Apakah Anda lahir dengan berat badan kurang dari 2.500 gram?", 0.90, false},
		{"REMAJA_19", "A", "A3", "Apakah jarak kelahiran Anda kurang dari 3 tahun dari kakak kandung terdekat?", 0.80, false},
		{"REMAJA_19", "A", "A4", "Apakah tinggi badan ibu kandung Anda kurang dari 150 cm?", 0.85, false},

		// ===== REMAJA 19 - DOMAIN B =====
		// B1: kondisi baik (konsumsi beragam pangan) = riPRAKONSEPSIsiko rendah (Ya=0.0)
		{"REMAJA_19", "B", "B1", "Dalam 7 hari terakhir, apakah Anda mengonsumsi ≥ 5 kelompok pangan berbeda pada ≥ 4 hari dalam seminggu?", 0.85, true},
		{"REMAJA_19", "B", "B2", "Apakah frekuensi makan utama Anda kurang dari 3 kali per hari pada ≥ 4 hari dalam seminggu?", 0.85, false},
		{"REMAJA_19", "B", "B3", "Dalam 7 hari terakhir, apakah Anda mengonsumsi pangan hewani kurang dari 4 hari dalam seminggu?", 0.80, false},
		{"REMAJA_19", "B", "B4", "Apakah Anda minum minuman berpemanis setiap hari?", 0.70, false},
		{"REMAJA_19", "B", "B5", "Apakah Anda pernah didiagnosis anemia atau pernah mendapatkan anjuran minum tablet tambah darah?", 0.85, false},
		{"REMAJA_19", "B", "B6", "Apakah Anda sering mengalami haid dengan perdarahan sangat banyak?", 0.80, false},

		// ===== REMAJA 19 - DOMAIN C =====
		{"REMAJA_19", "C", "C1", "Dalam 2 minggu terakhir, apakah Anda mengalami diare ≥ 3 hari berturut-turut atau demam karena infeksi?", 0.85, false},
		{"REMAJA_19", "C", "C2", "Dalam 6 bulan terakhir, apakah Anda pernah mengalami cacingan atau mendapatkan obat cacing?", 0.80, false},
		// C3: kondisi baik (imunisasi lengkap) = risiko rendah (Ya=0.0)
		{"REMAJA_19", "C", "C3", "Apakah Anda mendapatkan imunisasi dasar lengkap waktu kecil?", 0.75, true},

		// ===== REMAJA 19 - DOMAIN D =====
		// D1 & D2: kondisi baik (air terlindung/jamban layak) = risiko rendah (Ya=0.0)
		{"REMAJA_19", "D", "D1", "Apakah sumber air minum utama di rumah Anda adalah air minum terlindung?", 0.70, true},
		{"REMAJA_19", "D", "D2", "Apakah rumah Anda memiliki atau menggunakan jamban yang layak?", 0.70, true},
		// D3: kondisi baik (selalu cuci tangan) = risiko rendah (Selalu=0.0)
		{"REMAJA_19", "D", "D3", "Apakah Anda biasanya mencuci tangan dengan sabun pada lima momen penting?", 0.85, true},

		// ===== REMAJA 19 - DOMAIN E =====
		{"REMAJA_19", "E", "E1", "Apakah Anda atau keluarga pernah khawatir persediaan makanan akan habis sebelum punya uang untuk membeli lagi?", 0.80, false},
		{"REMAJA_19", "E", "E2", "Apakah pernah terjadi bahwa Anda atau anggota keluarga lain mengurangi ukuran porsi makan karena alasan ekonomi?", 0.90, false},
		{"REMAJA_19", "E", "E3", "Apakah pernah terjadi bahwa Anda atau anggota keluarga lain mengurangi jumlah frekuensi makan per hari karena alasan ekonomi?", 1.00, false},
		{"REMAJA_19", "E", "E4", "Apakah pernah terjadi bahwa Anda atau anggota keluarga lain tidak makan seharian penuh karena tidak ada makanan?", 1.00, false},
		{"REMAJA_19", "E", "E5", "Apakah pernah terjadi bahwa Anda atau anggota keluarga lain hanya makan makanan rendah mutu karena alasan ekonomi?", 0.70, false},

		// ===== REMAJA 19 - DOMAIN F =====
		{"REMAJA_19", "F", "F1", "Apakah Anda tinggal di daerah pedesaan terpencil atau di lingkungan kumuh perkotaan?", 0.70, false},
		{"REMAJA_19", "F", "F2", "Apakah ada anggota keluarga yang merokok di dalam rumah secara rutin?", 0.75, false},
		{"REMAJA_19", "F", "F3", "Apakah Anda saat ini sedang hamil atau pernah hamil pada usia <20 tahun?", 0.80, false},
		{"REMAJA_19", "F", "F4", "Apakah Anda saat ini sudah menikah?", 0.70, false},
	}

	for _, q := range questions {
		var category models.Category
		if err := db.Where("code = ?", q.CategoryCode).First(&category).Error; err != nil {
			return err
		}
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
		if err := db.Where("domain_id = ? AND code = ?", domain.ID, q.Code).FirstOrCreate(&question).Error; err != nil {
			return err
		}
	}
	return nil
}

func seedAnswerMappings(db *gorm.DB) error {
	// Semua CFEvidence di sini mencerminkan nilai RISIKO langsung (0.0 = tidak berisiko, 1.0 = risiko tertinggi).
	// Sumber: Tabel 2 Rulebase Certainty Factor (paper hal. 5–9).
	// Engine CF menggunakan nilai ini AS-IS; tidak ada pembalikan otomatis.
	answerMappings := []struct {
		CategoryCode string
		QuestionCode string
		AnswerKey    string
		Label        string
		CFEvidence   float64
	}{
		// ========================================================
		// PRAKONSEPSI
		// ========================================================

		// A1 — Frekuensi pemeriksaan Hb (semakin jarang = semakin berisiko)
		{"PRAKONSEPSI", "A1", "0", "Tidak pernah", 0.0},
		{"PRAKONSEPSI", "A1", "1", "1 kali", 0.4},
		{"PRAKONSEPSI", "A1", "2", "2 kali", 0.7},
		{"PRAKONSEPSI", "A1", "3", "≥3 kali", 1.0},

		// A2 — Frekuensi minum asam folat (DIBALIK: setiap hari = baik = risiko rendah)
		// Paper: 0=Setiap hari → CFEvidence 1.0 (DIBALIK), dst.
		// Namun makna risiko: setiap hari = tidak berisiko.
		// Sesuai kolom "Mapping Jawaban – CF Evidence" Tabel 2:
		//   0 → 1.0 (dibalik, ini adalah CF evidence setelah inversi label)
		//   1 → 0.7
		//   2 → 0.4
		//   3 → 0.0
		// CATATAN: paper menampilkan mapping setelah dibalik sehingga key "0" (jawaban terbaik)
		// memiliki CFEvidence 1.0 di kolom tabel. Ini berarti engine HARUS menggunakan
		// logika: CFitem = CFPakar × (1 - CFEvidence) untuk item IsReverse=true,
		// ATAU CFEvidence sudah disesuaikan sedemikian rupa sehingga key "0" (setiap hari)
		// menghasilkan CF_item rendah.
		//
		// Dari contoh perhitungan section 2.3:
		//   A2: CF_evidence = 1.0 (Jawaban ke-3, yaitu "Tidak pernah")
		//   CF_A2 = 0.90 × 1.0 = 0.90  → kontribusi risiko TINGGI
		// Artinya "Tidak pernah" (key "3") → CFEvidence = 1.0 (risiko tinggi). ✓
		// Dan "Setiap hari" (key "0") → CFEvidence = 0.0? Atau 1.0 sesuai tabel?
		//
		// Kesimpulan dari contoh hitung: paper menyimpan CFEvidence sebagai nilai RISIKO murni.
		// Key "0" (Setiap hari) = kondisi BAIK = CFEvidence risiko = 0.0
		// Key "3" (Tidak pernah) = kondisi BURUK = CFEvidence risiko = 1.0
		// Kolom tabel Tabel 2 yang menampilkan "0 → 1.0" adalah tampilan DIBALIK untuk
		// keperluan kuesioner (urutan skala). Nilai yang disimpan ke DB = nilai risiko murni.
		{"PRAKONSEPSI", "A2", "0", "Setiap hari", 0.0},
		{"PRAKONSEPSI", "A2", "1", "4–6 hari/minggu", 0.4},
		{"PRAKONSEPSI", "A2", "2", "1–3 hari/minggu", 0.7},
		{"PRAKONSEPSI", "A2", "3", "Tidak pernah", 1.0},

		// A3 — Frekuensi minum TTD (DIBALIK: ≥4 tablet/minggu = baik = risiko rendah)
		{"PRAKONSEPSI", "A3", "0", "≥4 tablet/minggu", 0.0},
		{"PRAKONSEPSI", "A3", "1", "2–3 tablet/minggu", 0.4},
		{"PRAKONSEPSI", "A3", "2", "1 tablet/minggu", 0.7},
		{"PRAKONSEPSI", "A3", "3", "Tidak pernah", 1.0},

		// A4 — Penggunaan garam beryodium (DIBALIK: Selalu = baik = risiko rendah)
		{"PRAKONSEPSI", "A4", "0", "Selalu", 0.0},
		{"PRAKONSEPSI", "A4", "1", "Sering", 0.4},
		{"PRAKONSEPSI", "A4", "2", "Jarang", 0.7},
		{"PRAKONSEPSI", "A4", "3", "Tidak pernah", 1.0},

		// A5 — Rencana suplementasi hamil (DIBALIK: Ada rencana jelas = baik = risiko rendah)
		{"PRAKONSEPSI", "A5", "0", "Ada rencana jelas dan tertulis dari tenaga kesehatan", 0.0},
		{"PRAKONSEPSI", "A5", "1", "Ada rencana tetapi belum tertulis", 0.4},
		{"PRAKONSEPSI", "A5", "2", "Hanya pernah mendengar tetapi belum punya rencana", 0.7},
		{"PRAKONSEPSI", "A5", "3", "Belum pernah membahas / belum punya rencana sama sekali", 1.0},

		// B1 — Kekhawatiran pangan habis
		{"PRAKONSEPSI", "B1", "0", "Tidak pernah", 0.0},
		{"PRAKONSEPSI", "B1", "1", "Jarang (1–2 kali)", 0.4},
		{"PRAKONSEPSI", "B1", "2", "Kadang-kadang (3–10 kali)", 0.7},
		{"PRAKONSEPSI", "B1", "3", "Sering (>10 kali)", 1.0},

		// B2 — Makanan benar-benar habis
		{"PRAKONSEPSI", "B2", "0", "Tidak pernah", 0.0},
		{"PRAKONSEPSI", "B2", "1", "Jarang (1–2 kali)", 0.4},
		{"PRAKONSEPSI", "B2", "2", "Kadang-kadang (3–10 kali)", 0.7},
		{"PRAKONSEPSI", "B2", "3", "Sering (>10 kali)", 1.0},

		// B3 — Mengurangi porsi makan
		{"PRAKONSEPSI", "B3", "0", "Tidak pernah", 0.0},
		{"PRAKONSEPSI", "B3", "1", "Jarang (1–2 kali)", 0.4},
		{"PRAKONSEPSI", "B3", "2", "Kadang-kadang (3–10 kali)", 0.7},
		{"PRAKONSEPSI", "B3", "3", "Sering (>10 kali)", 1.0},

		// B4 — Mengurangi frekuensi makan per hari (CFPakar = 1.00)
		{"PRAKONSEPSI", "B4", "0", "Tidak pernah", 0.0},
		{"PRAKONSEPSI", "B4", "1", "Jarang (1–2 kali)", 0.4},
		{"PRAKONSEPSI", "B4", "2", "Kadang-kadang (3–10 kali)", 0.7},
		{"PRAKONSEPSI", "B4", "3", "Sering (>10 kali)", 1.0},

		// B5 — Tidak makan seharian (CFPakar = 1.00)
		{"PRAKONSEPSI", "B5", "0", "Tidak pernah", 0.0},
		{"PRAKONSEPSI", "B5", "1", "Jarang (1–2 kali)", 0.4},
		{"PRAKONSEPSI", "B5", "2", "Kadang-kadang (3–10 kali)", 0.7},
		{"PRAKONSEPSI", "B5", "3", "Sering (>10 kali)", 1.0},

		// C1 — Frekuensi merokok
		{"PRAKONSEPSI", "C1", "0", "Tidak pernah", 0.0},
		{"PRAKONSEPSI", "C1", "1", "Kadang-kadang (≤1 batang/minggu)", 0.4},
		{"PRAKONSEPSI", "C1", "2", "Sering (2–6 batang/minggu)", 0.7},
		{"PRAKONSEPSI", "C1", "3", "Setiap hari", 1.0},

		// C2 — Paparan asap rokok (hari dalam 7 hari terakhir)
		{"PRAKONSEPSI", "C2", "0", "0 hari", 0.0},
		{"PRAKONSEPSI", "C2", "1", "1–2 hari", 0.4},
		{"PRAKONSEPSI", "C2", "2", "3–5 hari", 0.7},
		{"PRAKONSEPSI", "C2", "3", "6–7 hari", 1.0},

		// C3 — Frekuensi konsumsi alkohol
		{"PRAKONSEPSI", "C3", "0", "Tidak pernah", 0.0},
		{"PRAKONSEPSI", "C3", "1", "Pernah, tetapi <1 kali/bulan", 0.4},
		{"PRAKONSEPSI", "C3", "2", "1–3 kali/bulan", 0.7},
		{"PRAKONSEPSI", "C3", "3", "≥1 kali/minggu", 1.0},

		// C4 — Sumber air minum (hanya 2 opsi, key "1" tidak ada)
		{"PRAKONSEPSI", "C4", "0", "Air kemasan/PDAM/sumur terlindung", 0.0},
		{"PRAKONSEPSI", "C4", "2", "Sumur tidak terlindung / air permukaan yang belum diolah", 1.0},

		// C5 — Kepemilikan jamban
		{"PRAKONSEPSI", "C5", "0", "Jamban pribadi layak", 0.0},
		{"PRAKONSEPSI", "C5", "1", "Jamban komunal layak", 0.4},
		{"PRAKONSEPSI", "C5", "2", "Tidak memiliki jamban / jamban tidak layak", 1.0},

		// C6 — Cuci tangan (Selalu = perilaku sangat baik = risiko rendah 0.0)
		// Paper Tabel 2: "0=Selalu → 0.0 (perilaku sangat baik), 1→0.4, 2→0.7, 3→1.0"
		{"PRAKONSEPSI", "C6", "0", "Selalu", 0.0},
		{"PRAKONSEPSI", "C6", "1", "Sering", 0.4},
		{"PRAKONSEPSI", "C6", "2", "Kadang-kadang", 0.7},
		{"PRAKONSEPSI", "C6", "3", "Jarang / Tidak pernah", 1.0},

		// D1 — Perasaan sedih/putus asa
		{"PRAKONSEPSI", "D1", "0", "Tidak pernah", 0.0},
		{"PRAKONSEPSI", "D1", "1", "Beberapa hari", 0.4},
		{"PRAKONSEPSI", "D1", "2", "Lebih dari separuh hari", 0.7},
		{"PRAKONSEPSI", "D1", "3", "Hampir setiap hari", 1.0},

		// D2 — Kehilangan minat/kenikmatan
		{"PRAKONSEPSI", "D2", "0", "Tidak pernah", 0.0},
		{"PRAKONSEPSI", "D2", "1", "Beberapa hari", 0.4},
		{"PRAKONSEPSI", "D2", "2", "Lebih dari separuh hari", 0.7},
		{"PRAKONSEPSI", "D2", "3", "Hampir setiap hari", 1.0},

		// D3 — Merasa tidak aman / kekerasan
		{"PRAKONSEPSI", "D3", "0", "Tidak pernah", 0.0},
		{"PRAKONSEPSI", "D3", "1", "Jarang (1–2 kali)", 0.4},
		{"PRAKONSEPSI", "D3", "2", "Kadang-kadang (3–10 kali)", 0.7},
		{"PRAKONSEPSI", "D3", "3", "Sering (>10 kali)", 1.0},

		// D4 — Minat konseling prakonsepsi (PROTEKTIF/DIBALIK)
		// Paper: "Sangat berminat → 0.0 (risiko rendah), Tidak berminat → 1.0 (risiko tinggi)"
		{"PRAKONSEPSI", "D4", "0", "Sangat berminat", 0.0},
		{"PRAKONSEPSI", "D4", "1", "Berminat", 0.4},
		{"PRAKONSEPSI", "D4", "2", "Kurang berminat", 0.7},
		{"PRAKONSEPSI", "D4", "3", "Tidak berminat sama sekali", 1.0},

		// ========================================================
		// PERNAH MELAHIRKAN
		// ========================================================

		// A1 — Lama ASI eksklusif (DIBALIK: ≥6 bulan = baik = risiko rendah 0.0)
		// Paper Tabel 2: "0=≥6 bulan → 0.0, 1=4–5 bulan → 0.4, 2=1–3 bulan → 0.7, 3=Tidak ASI → 1.0"
		{"PERNAH_MELAHIRKAN", "A1", "0", "≥6 bulan", 0.0},
		{"PERNAH_MELAHIRKAN", "A1", "1", "4–5 bulan", 0.4},
		{"PERNAH_MELAHIRKAN", "A1", "2", "1–3 bulan", 0.7},
		{"PERNAH_MELAHIRKAN", "A1", "3", "Tidak ASI eksklusif", 1.0},

		// A2 — Frekuensi MP-ASI berkualitas (DIBALIK: ≥6 hari = baik = risiko rendah 0.0)
		// Paper Tabel 2: "0=≥6 hari → 0.0, 1=4–5 hari → 0.4, 2=2–3 hari → 0.7, 3=0–1 hari → 1.0"
		{"PERNAH_MELAHIRKAN", "A2", "0", "≥6 hari", 0.0},
		{"PERNAH_MELAHIRKAN", "A2", "1", "4–5 hari", 0.4},
		{"PERNAH_MELAHIRKAN", "A2", "2", "2–3 hari", 0.7},
		{"PERNAH_MELAHIRKAN", "A2", "3", "0–1 hari", 1.0},

		// B1 — Paparan asap rokok (hari dalam 7 hari terakhir)
		{"PERNAH_MELAHIRKAN", "B1", "0", "0 hari", 0.0},
		{"PERNAH_MELAHIRKAN", "B1", "1", "1–2 hari", 0.4},
		{"PERNAH_MELAHIRKAN", "B1", "2", "3–5 hari", 0.7},
		{"PERNAH_MELAHIRKAN", "B1", "3", "6–7 hari", 1.0},

		// B2 — Penggunaan bahan bakar selain LPG/listrik
		{"PERNAH_MELAHIRKAN", "B2", "0", "0 hari", 0.0},
		{"PERNAH_MELAHIRKAN", "B2", "1", "1–2 hari", 0.4},
		{"PERNAH_MELAHIRKAN", "B2", "2", "3–5 hari", 0.7},
		{"PERNAH_MELAHIRKAN", "B2", "3", "6–7 hari", 1.0},

		// C1 — Perasaan sedih/putus asa
		{"PERNAH_MELAHIRKAN", "C1", "0", "Tidak pernah", 0.0},
		{"PERNAH_MELAHIRKAN", "C1", "1", "Beberapa hari", 0.4},
		{"PERNAH_MELAHIRKAN", "C1", "2", "Lebih dari separuh hari", 0.7},
		{"PERNAH_MELAHIRKAN", "C1", "3", "Hampir setiap hari", 1.0},

		// C2 — Kekerasan emosional/fisik/seksual (CFPakar = 0.95)
		{"PERNAH_MELAHIRKAN", "C2", "0", "Tidak pernah", 0.0},
		{"PERNAH_MELAHIRKAN", "C2", "1", "Jarang (1–2 kali/tahun)", 0.4},
		{"PERNAH_MELAHIRKAN", "C2", "2", "Kadang (3–10 kali)", 0.7},
		{"PERNAH_MELAHIRKAN", "C2", "3", "Sering (>10 kali)", 1.0},

		// ========================================================
		// REMAJA 19 TAHUN
		// ========================================================

		// A1 — Pernah didiagnosis pendek (Ya/Tidak/Tidak tahu)
		{"REMAJA_19", "A1", "0", "Ya", 1.0},
		{"REMAJA_19", "A1", "1", "Tidak", 0.0},
		{"REMAJA_19", "A1", "2", "Tidak tahu", 0.4},

		// A2 — Lahir BB < 2500 g
		{"REMAJA_19", "A2", "0", "Ya", 1.0},
		{"REMAJA_19", "A2", "1", "Tidak", 0.0},
		{"REMAJA_19", "A2", "2", "Tidak tahu", 0.4},

		// A3 — Jarak kelahiran < 3 tahun (hanya Ya/Tidak)
		{"REMAJA_19", "A3", "0", "Ya", 1.0},
		{"REMAJA_19", "A3", "1", "Tidak", 0.0},

		// A4 — Tinggi ibu < 150 cm
		{"REMAJA_19", "A4", "0", "Ya", 1.0},
		{"REMAJA_19", "A4", "1", "Tidak", 0.0},
		{"REMAJA_19", "A4", "2", "Tidak tahu", 0.4},

		// B1 — ≥5 kelompok pangan (DIBALIK: Ya = kondisi baik = risiko rendah 0.0)
		// Paper Tabel 2: "Ya = 0.0 (baik), Tidak = 0.8 (risiko), Tidak ingat = 0.4"
		{"REMAJA_19", "B1", "0", "Ya", 0.0},
		{"REMAJA_19", "B1", "1", "Tidak", 0.8},
		{"REMAJA_19", "B1", "2", "Tidak ingat", 0.4},

		// B2 — Frekuensi makan < 3 kali/hari
		{"REMAJA_19", "B2", "0", "Ya", 0.8},
		{"REMAJA_19", "B2", "1", "Tidak", 0.0},

		// B3 — Konsumsi hewani < 4 hari/minggu
		{"REMAJA_19", "B3", "0", "Ya", 0.8},
		{"REMAJA_19", "B3", "1", "Tidak", 0.0},

		// B4 — Minuman berpemanis setiap hari
		{"REMAJA_19", "B4", "0", "Ya", 0.8},
		{"REMAJA_19", "B4", "1", "Tidak", 0.0},

		// B5 — Riwayat anemia/Fe
		{"REMAJA_19", "B5", "0", "Ya", 0.8},
		{"REMAJA_19", "B5", "1", "Tidak", 0.0},
		{"REMAJA_19", "B5", "2", "Tidak tahu", 0.4},

		// B6 — Perdarahan haid banyak
		{"REMAJA_19", "B6", "0", "Ya", 0.8},
		{"REMAJA_19", "B6", "1", "Tidak", 0.0},

		// C1 — Riwayat diare/demam infeksi
		{"REMAJA_19", "C1", "0", "Ya", 0.8},
		{"REMAJA_19", "C1", "1", "Tidak", 0.0},

		// C2 — Riwayat cacingan / obat cacing
		{"REMAJA_19", "C2", "0", "Ya", 0.8},
		{"REMAJA_19", "C2", "1", "Tidak", 0.0},

		// C3 — Status imunisasi dasar lengkap (DIBALIK: Ya = kondisi baik = risiko rendah 0.0)
		// Paper Tabel 2: "Ya → 0.0, Tidak → 0.8, TT → 0.4"
		{"REMAJA_19", "C3", "0", "Ya", 0.0},
		{"REMAJA_19", "C3", "1", "Tidak", 0.8},
		{"REMAJA_19", "C3", "2", "Tidak tahu", 0.4},

		// D1 — Sumber air minum terlindung (DIBALIK: Ya = baik = risiko rendah 0.0)
		// Paper Tabel 2: "Ya → 0.0, Tidak → 0.8, TT → 0.4"
		{"REMAJA_19", "D1", "0", "Ya", 0.0},
		{"REMAJA_19", "D1", "1", "Tidak", 0.8},
		{"REMAJA_19", "D1", "2", "Tidak tahu", 0.4},

		// D2 — Kepemilikan jamban layak (DIBALIK: Ya = baik = risiko rendah 0.0)
		// Paper Tabel 2: "Ya → 0.0, Tidak → 0.8"
		{"REMAJA_19", "D2", "0", "Ya", 0.0},
		{"REMAJA_19", "D2", "1", "Tidak", 0.8},

		// D3 — Cuci tangan dengan sabun pada 5 momen penting (DIBALIK: Selalu = risiko rendah 0.0)
		// Paper Tabel 2: "0=Selalu → 0.0, 1=Sering → 0.2, 2=Kadang → 0.5, 3=Jarang/Tidak → 0.8"
		{"REMAJA_19", "D3", "0", "Selalu", 0.0},
		{"REMAJA_19", "D3", "1", "Sering", 0.2},
		{"REMAJA_19", "D3", "2", "Kadang", 0.5},
		{"REMAJA_19", "D3", "3", "Jarang / Tidak pernah", 0.8},

		// E1 — Kekhawatiran persediaan pangan habis
		{"REMAJA_19", "E1", "0", "Tidak pernah", 0.0},
		{"REMAJA_19", "E1", "1", "Ya, 1–2 kali", 0.4},
		{"REMAJA_19", "E1", "2", "Ya, 3–10 kali", 0.7},
		{"REMAJA_19", "E1", "3", "Ya, lebih dari 10 kali", 1.0},

		// E2 — Mengurangi porsi makan karena ekonomi
		{"REMAJA_19", "E2", "0", "Tidak pernah", 0.0},
		{"REMAJA_19", "E2", "1", "Ya, 1–2 kali", 0.4},
		{"REMAJA_19", "E2", "2", "Ya, 3–10 kali", 0.7},
		{"REMAJA_19", "E2", "3", "Ya, lebih dari 10 kali", 1.0},

		// E3 — Mengurangi frekuensi makan karena ekonomi (CFPakar = 1.00)
		{"REMAJA_19", "E3", "0", "Tidak pernah", 0.0},
		{"REMAJA_19", "E3", "1", "Ya, 1–2 kali", 0.4},
		{"REMAJA_19", "E3", "2", "Ya, 3–10 kali", 0.7},
		{"REMAJA_19", "E3", "3", "Ya, lebih dari 10 kali", 1.0},

		// E4 — Tidak makan seharian penuh (CFPakar = 1.00)
		{"REMAJA_19", "E4", "0", "Tidak pernah", 0.0},
		{"REMAJA_19", "E4", "1", "Ya, 1–2 kali", 0.4},
		{"REMAJA_19", "E4", "2", "Ya, 3–10 kali", 0.7},
		{"REMAJA_19", "E4", "3", "Ya, lebih dari 10 kali", 1.0},

		// E5 — Makanan rendah mutu karena ekonomi (CFPakar = 0.70)
		{"REMAJA_19", "E5", "0", "Tidak pernah", 0.0},
		{"REMAJA_19", "E5", "1", "Ya, 1–2 kali", 0.4},
		{"REMAJA_19", "E5", "2", "Ya, 3–10 kali", 0.7},
		{"REMAJA_19", "E5", "3", "Ya, lebih dari 10 kali", 1.0},

		// F1 — Tinggal di daerah terpencil/kumuh
		{"REMAJA_19", "F1", "0", "Ya", 0.8},
		{"REMAJA_19", "F1", "1", "Tidak", 0.0},

		// F2 — Anggota keluarga merokok di dalam rumah
		{"REMAJA_19", "F2", "0", "Ya", 0.8},
		{"REMAJA_19", "F2", "1", "Tidak", 0.0},

		// F3 — Riwayat hamil < 20 tahun
		{"REMAJA_19", "F3", "0", "Ya", 0.8},
		{"REMAJA_19", "F3", "1", "Tidak", 0.0},

		// F4 — Sudah menikah
		{"REMAJA_19", "F4", "0", "Ya", 0.8},
		{"REMAJA_19", "F4", "1", "Tidak", 0.0},
	}

	for _, am := range answerMappings {
		var category models.Category
		if err := db.Where("code = ?", am.CategoryCode).First(&category).Error; err != nil {
			return err
		}

		questionCode := am.QuestionCode
		domainCode := string(questionCode[0])

		var domain models.Domain
		if err := db.Where("category_id = ? AND code = ?", category.ID, domainCode).First(&domain).Error; err != nil {
			return err
		}

		var question models.Question
		if err := db.Where("domain_id = ? AND code = ?", domain.ID, am.QuestionCode).First(&question).Error; err != nil {
			return err
		}

		answerMapping := models.AnswerMapping{
			QuestionID: question.ID,
			AnswerKey:  am.AnswerKey,
			Label:      am.Label,
			CFEvidence: am.CFEvidence,
		}
		if err := db.Where("question_id = ? AND answer_key = ?", question.ID, am.AnswerKey).FirstOrCreate(&answerMapping).Error; err != nil {
			return err
		}
	}

	return nil
}
