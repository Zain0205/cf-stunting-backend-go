package database

import (
	"github.com/Zain0205/cf-stunting-backend-go/internal/models"
	"gorm.io/gorm"
)

func SeedQuestions(db *gorm.DB) error {
	type qSeed struct {
		Category string
		Domain   string
		Code     string
		Text     string
		CFPakar  float64
		Reverse  bool
	}

	questions := []qSeed{
		// ================= PRAKONSEPSI – DOMAIN A (Gizi dan Suplementasi) =================
		{"PRAKONSEPSI", "A", "A1", "Dalam 6 bulan terakhir, seberapa sering anda melakukan pemeriksaan Hb (hemoglobin)", 0.80, false},
		{"PRAKONSEPSI", "A", "A2", "Dalam 1 bulan terakhir, seberapa sering Anda minum suplemen asam folat sesuai anjuran?", 0.90, true},
		{"PRAKONSEPSI", "A", "A3", "Dalam 1 bulan terakhir, seberapa sering Anda minum tablet tambah darah (TTD)?", 0.90, true},
		{"PRAKONSEPSI", "A", "A4", "Seberapa sering Anda menggunakan garam beryodium (dengan logo beryodium) saat memasak di rumah?", 0.85, true},
		{"PRAKONSEPSI", "A", "A5", "Sejauh ini, sejauh mana Anda sudah memiliki rencana atau anjuran tertulis dari tenaga kesehatan tentang suplemen yang akan dikonsumsi saat hamil (asam folat, zat besi, kalsium, dsb.)?", 0.85, true},

		// ================= PRAKONSEPSI – DOMAIN B (Ketahanan Pangan) =================
		{"PRAKONSEPSI", "B", "B1", "Dalam 12 bulan terakhir, seberapa sering Anda khawatir persediaan makanan di rumah akan habis sebelum bisa membeli lagi?", 0.85, false},
		{"PRAKONSEPSI", "B", "B2", "Dalam 12 bulan terakhir, seberapa sering makanan di rumah benar-benar habis dan Anda tidak punya uang untuk membeli lagi?", 0.90, false},
		{"PRAKONSEPSI", "B", "B3", "Dalam 12 bulan terakhir, seberapa sering anggota keluarga mengurangi ukuran porsi makan karena alasan ekonomi?", 0.90, false},
		{"PRAKONSEPSI", "B", "B4", "Dalam 12 bulan terakhir, seberapa sering anggota keluarga mengurangi frekuensi makan per hari karena alasan ekonomi?", 1.00, false},
		{"PRAKONSEPSI", "B", "B5", "Dalam 12 bulan terakhir, seberapa sering terjadi tidak makan seharian untuk menghemat makanan/karena tidak ada makanan?", 1.00, false},

		// ================= PRAKONSEPSI – DOMAIN C (Lingkungan dan Perilaku Risiko) =================
		{"PRAKONSEPSI", "C", "C1", "Dalam 3 bulan terakhir, seberapa sering Anda merokok?", 0.85, false},
		{"PRAKONSEPSI", "C", "C2", "Dalam 7 hari terakhir, berapa hari Anda terpapar asap rokok di rumah atau di tempat kerja?", 0.80, false},
		{"PRAKONSEPSI", "C", "C3", "Dalam 3 bulan terakhir, seberapa sering Anda mengonsumsi minuman beralkohol?", 0.90, false},
		{"PRAKONSEPSI", "C", "C4", "Sumber air minum utama di rumah Anda", 0.75, false},
		{"PRAKONSEPSI", "C", "C5", "Kepemilikan jamban", 0.75, false},
		{"PRAKONSEPSI", "C", "C6", "Seberapa konsisten Anda mencuci tangan dengan sabun sebelum menyiapkan dan menyantap makanan?", 0.85, false},

		// ================= PRAKONSEPSI – DOMAIN D (Psikososial) =================
		{"PRAKONSEPSI", "D", "D1", "Dalam 2 minggu terakhir, seberapa sering Anda merasa sedih atau putus asa?", 0.85, false},
		{"PRAKONSEPSI", "D", "D2", "Dalam 2 minggu terakhir, seberapa sering Anda kehilangan minat atau kenikmatan dalam aktivitas sehari-hari?", 0.80, false},
		{"PRAKONSEPSI", "D", "D3", "Dalam 12 bulan terakhir, seberapa sering Anda merasa tidak aman di rumah atau mengalami kekerasan (emosional/fisik/seksual) dari pasangan/anggota keluarga?", 0.90, false},
		{"PRAKONSEPSI", "D", "D4", "Seberapa besar keinginan Anda untuk mengikuti konseling prakonsepsi (gizi, penyakit kronis, KB/interval, suplementasi)?", 0.75, true},

		// ================= PERNAH MELAHIRKAN – DOMAIN A (Pemberian Makan) =================
		{"PERNAH_MELAHIRKAN", "A", "A1", "Berapa lama anak terakhir Anda mendapatkan ASI eksklusif tanpa tambahan makanan/minuman lain?", 0.85, true},
		{"PERNAH_MELAHIRKAN", "A", "A2", "Dalam 7 hari terakhir, berapa hari anak mendapat MP-ASI berkualitas (protein hewani, buah, sayur)?", 0.85, true},

		// ================= PERNAH MELAHIRKAN – DOMAIN B (Lingkungan Fisik) =================
		{"PERNAH_MELAHIRKAN", "B", "B1", "Dalam 7 hari terakhir, berapa hari Anda atau anak terpapar asap rokok di rumah?", 0.80, false},
		{"PERNAH_MELAHIRKAN", "B", "B2", "Dalam 7 hari terakhir, berapa hari rumah menggunakan bahan bakar selain gas elpiji/listrik?", 0.80, false},

		// ================= PERNAH MELAHIRKAN – DOMAIN C (Psikososial) =================
		{"PERNAH_MELAHIRKAN", "C", "C1", "Dalam 2 minggu terakhir, seberapa sering Anda merasa sedih atau putus asa?", 0.90, false},
		{"PERNAH_MELAHIRKAN", "C", "C2", "Dalam 12 bulan terakhir, seberapa sering Anda mengalami kekerasan emosional/fisik/seksual dari pasangan/keluarga?", 0.95, false},

		// ================= REMAJA 19 – DOMAIN A (Biologis Intergenerasional) =================
		{"REMAJA_19", "A", "A2", "Apakah Anda lahir dengan berat badan kurang dari 2.500 gram?", 0.90, false},
		{"REMAJA_19", "A", "A3", "Apakah jarak kelahiran Anda kurang dari 3 tahun dari kakak kandung terdekat?", 0.80, false},
		{"REMAJA_19", "A", "A4", "Apakah tinggi badan ibu kandung Anda kurang dari 150 cm?", 0.85, false},

		// ================= REMAJA 19 – DOMAIN B (Pola Makan) =================
		{"REMAJA_19", "B", "B1", "Dalam 7 hari terakhir, apakah Anda mengonsumsi ≥ 5 kelompok pangan berbeda (misalnya: nasi/karbohidrat, lauk hewani, lauk nabati, sayur, buah) pada ≥ 4 hari dalam seminggu?", 0.85, true},
		{"REMAJA_19", "B", "B2", "Apakah frekuensi makan utama Anda kurang dari 3 kali per hari pada ≥ 4 hari dalam seminggu?", 0.85, false},
		{"REMAJA_19", "B", "B3", "Dalam 7 hari terakhir, apakah Anda mengonsumsi pangan hewani (telur, ikan, daging, ayam) kurang dari 4 hari dalam seminggu?", 0.80, false},
		{"REMAJA_19", "B", "B4", "Apakah Anda minum minuman berpemanis (teh manis, minuman berenergi, minuman kemasan manis) setiap hari?", 0.70, false},
		{"REMAJA_19", "B", "B5", "Apakah Anda pernah didiagnosis anemia (kurang darah) atau pernah mendapatkan anjuran minum tablet tambah darah (Fe) oleh tenaga kesehatan?", 0.85, false},
		{"REMAJA_19", "B", "B6", "Apakah Anda sering mengalami haid dengan perdarahan sangat banyak (misalnya durasi >7 hari atau sering harus mengganti pembalut karena penuh)?", 0.80, false},

		// ================= REMAJA 19 – DOMAIN C (Infeksi) =================
		{"REMAJA_19", "C", "C1", "Dalam 2 minggu terakhir, apakah Anda mengalami diare ≥ 3 hari berturut-turut atau demam karena infeksi (misal ISPA, tifus, dll.)?", 0.85, false},
		{"REMAJA_19", "C", "C2", "Dalam 6 bulan terakhir, apakah Anda pernah mengalami cacingan atau mendapatkan obat cacing?", 0.80, false},
		{"REMAJA_19", "C", "C3", "Apakah Anda mendapatkan imunisasi dasar lengkap waktu kecil (BCG, DPT, Polio, Campak), sesuai buku KIA atau catatan imunisasi?", 0.75, true},

		// ================= REMAJA 19 – DOMAIN D (Sanitasi dan Perilaku) =================
		{"REMAJA_19", "D", "D1", "Apakah sumber air minum utama di rumah Anda adalah air minum terlindung (PDAM, sumur terlindung, atau air kemasan)?", 0.70, true},
		{"REMAJA_19", "D", "D2", "Apakah rumah Anda memiliki atau menggunakan jamban yang layak (tidak mencemari lingkungan, beralaskan semen, punya saluran pembuangan yang aman)?", 0.70, true},
		{"REMAJA_19", "D", "D3", "Apakah Anda biasanya mencuci tangan dengan sabun pada lima momen penting (sebelum makan, sebelum menyiapkan makanan, setelah dari jamban, setelah membersihkan anak, setelah memegang hewan/kotoran)?", 0.85, true},

		// ================= REMAJA 19 – DOMAIN E (Ketahanan Pangan) =================
		{"REMAJA_19", "E", "E1", "Apakah Anda atau keluarga pernah khawatir persediaan makanan akan habis sebelum punya uang untuk membeli lagi? (12 bulan terakhir)", 0.80, false},
		{"REMAJA_19", "E", "E2", "Apakah pernah terjadi bahwa Anda atau anggota keluarga lain mengurangi ukuran porsi makan karena alasan ekonomi? (12 bulan terakhir)", 0.90, false},
		{"REMAJA_19", "E", "E3", "Apakah pernah terjadi bahwa Anda atau anggota keluarga lain mengurangi jumlah frekuensi makan per hari karena alasan ekonomi? (12 bulan terakhir)", 1.00, false},
		{"REMAJA_19", "E", "E4", "Apakah pernah terjadi bahwa Anda atau anggota keluarga lain tidak makan seharian penuh karena tidak ada makanan/untuk menghemat uang? (12 bulan terakhir)", 1.00, false},
		{"REMAJA_19", "E", "E5", "Apakah pernah terjadi bahwa Anda atau anggota keluarga lain hanya makan makanan rendah mutu (misalnya hanya nasi/karbohidrat tanpa lauk) karena alasan ekonomi? (12 bulan terakhir)", 0.70, false},

		// ================= REMAJA 19 – DOMAIN F (Lingkungan Sosial) =================
		{"REMAJA_19", "F", "F1", "Apakah Anda tinggal di daerah pedesaan terpencil atau di lingkungan kumuh perkotaan?", 0.70, false},
		{"REMAJA_19", "F", "F2", "Apakah ada anggota keluarga yang merokok di dalam rumah secara rutin?", 0.75, false},
		{"REMAJA_19", "F", "F3", "Apakah Anda saat ini sedang hamil atau pernah hamil pada usia <20 tahun?", 0.80, false},
		{"REMAJA_19", "F", "F4", "Apakah Anda saat ini sudah menikah?", 0.70, false},
	}

	for _, q := range questions {
		var domain models.Domain

		db.Joins("JOIN categories ON categories.id = domains.category_id").
			Where("categories.code = ? AND domains.code = ?", q.Category, q.Domain).
			First(&domain)

		question := models.Question{
			DomainID:  domain.ID,
			Code:      q.Code,
			Text:      q.Text,
			CFPakar:   q.CFPakar,
			IsReverse: q.Reverse,
		}

		db.Where("domain_id = ? AND code = ?", domain.ID, q.Code).
			FirstOrCreate(&question)
	}

	return nil
}
