Migrasi dari Laravel (PHP) ke Go bukan sekadar "ganti bahasa". Ini adalah perubahan paradigma dari bahasa yang Dynamic/Interpreted (PHP) ke Static/Compiled & Concurrent (Go).

Risiko Utama Migrasi Laravel ke Go:

API Response Mismatch: Laravel Eloquent memiliki format JSON standar (misal: pembungkus data, format pagination meta/links). Jika Backend Go tidak meniru struktur ini 100%, Frontend Vue akan blank atau undefined.
Type Safety vs Dynamic Types: PHP membolehkan tipe data yang longgar (string "1" == int 1). Go sangat ketat. Frontend Vue mungkin mengirim string, Go mengharapkan integer -> Error 400 Bad Request.
Date/Time Handling: Laravel menggunakan Carbon. Go menggunakan time.Time. Format tanggal sering kali menjadi sumber bug utama saat migrasi (Timezone mismatch, format string berbeda).
Berikut adalah Master Test Plan spesifik untuk migrasi taraNote (Laravel -> Go).

🗓 Master Test Plan: Migration Protocol (Laravel to Go)
🏁 Definition of Done (DoD) - Migrasi
Zero Frontend Changes: Codebase Vue tidak boleh diubah hanya untuk mengakomodasi Backend baru (kecuali fitur baru).
Data Consistency: Data yang di-import dari DB lama (jika ada) terbaca sempurna di sistem baru.
Performance: Response time Go harus minimal 2x lebih cepat dari Laravel untuk operasi Read.
🏃 Sprint 1: API Contract & Auth Parity (The "Handshake")
Fokus: Memastikan Frontend Vue bisa "Login" dan menerima struktur data yang dikenali tanpa crash.

Goal: User bisa Login, Logout, dan melihat profil user. Format JSON response Go identik dengan Laravel.
Risiko Teknis: Format Token (Sanctum/JWT) di Laravel berbeda implementasinya di Go.
Checklist Pengujian:

| ID | Skenario Tes | Tipe | Prioritas |
| :-- | :--- | :-- | :-- |
| S1-01 | Payload Structure Check: Bandingkan JSON response Login dari Laravel vs Go. Pastikan key-nya sama (misal: apakah access_token atau token? Ada field user di dalamnya?). | Contract | Critical |
| S1-02 | Type Strictness: Kirim data login dengan tipe data "salah" (misal: password berupa angka, bukan string). PHP mungkin memproses, Go mungkin panic/error. Pastikan Go handle error dengan graceful (pesan error terbaca). | Negative | High |
| S1-03 | Auth Middleware: Akses endpoint /api/notes tanpa token. Pastikan Go mengembalikan 401 Unauthorized (bukan 500 panic). | Security | Critical |
| S1-04 | Date Format: Cek field created_at pada user profile. Apakah formatnya ISO 8601 yang sama dengan output Carbon Laravel? | Data | High |
| S1-05 | CORS Handling: Akses API Go dari port Frontend Vue (misal localhost:3000 ke localhost:8080). Pastikan headers CORS dikonfigurasi benar di Go. | Config | High |

🏃 Sprint 2: CRUD & Logic Translation (The "Clone")
Fokus: Memindahkan logika bisnis "Note" dari PHP ke Go.

Goal: User bisa Create, Read, Update, Delete (CRUD) catatan. Fitur Pagination dan Search bekerja persis seperti versi Laravel.
Risiko Teknis: Logic Query Scope di Laravel (seperti whereUser()) harus ditulis ulang manual di SQL/ORM Go.
Checklist Pengujian:

| ID | Skenario Tes | Tipe | Prioritas |
| :-- | :--- | :-- | :-- |
| S2-01 | Pagination Parity: Request GET /notes?page=1. Cek apakah Go mengembalikan metadata pagination (total, per_page, last_page) yang strukturnya dimengerti komponen Pagination Vue. | UI/Funct | Critical |
| S2-02 | Validation Messages: Coba simpan Note kosong. Laravel biasanya mengembalikan error bag {"errors": {"title": ["Required"]}}. Apakah Go mengembalikan struktur JSON error yang sama? Jika beda, Vue tidak akan menampilkan pesan error di form. | UI/UX | High |
| S2-03 | Null Handling: Buat note tanpa konten (hanya judul). Cek di SQLite: apakah tersimpan sebagai NULL atau empty string ""? Pastikan konsisten dengan logika lama. | Data | Medium |
| S2-04 | Search Query: Lakukan pencarian note. Test sensitivitas huruf (Case Sensitive vs Insensitive). SQLite LIKE kadang case-sensitive defaultnya, beda dengan MySQL/Postgres di Laravel. | Funct | Medium |
| S2-05 | Tags/Categories (Relation): Jika ada relasi (Note has many Tags), pastikan JSON response menyertakan nested object tags. | Data | High |

🏃 Sprint 3: SQLite Concurrency & Stability (The "Stress")
Fokus: Karena pindah ke Go (Concurrent) + SQLite (File based), kita harus pastikan database tidak terkunci.

Goal: Aplikasi tidak error database is locked saat diakses banyak request cepat.
Checklist Pengujian:

| ID | Skenario Tes | Tipe | Prioritas |
| :-- | :--- | :-- | :-- |
| S3-01 | Parallel Writes: Gunakan tool (seperti JMeter/K6/Apache Benchmark), kirim 50 request "Create Note" dalam 1 detik. Pastikan semua masuk DB tanpa error. | Stress | Critical |
| S3-02 | Transaction Safety: Simulasikan error di tengah proses simpan (misal: matikan DB connection manual saat process). Pastikan tidak ada data "setengah jadi" yang tersimpan (Atomicity). | Integrity | High |
| S3-03 | Memory Leak Check: Monitor penggunaan RAM aplikasi Go saat dibiarkan berjalan 24 jam atau dipukul load test. Go tidak boleh memakan RAM terus menerus tanpa release (Garbage Collection). | Perf | Medium |

🏃 Sprint 4: The "DocFlow" Upgrade (New Features)
Fokus: Setelah fitur lama stabil, barulah kita test fitur baru (Versioning & WebSocket) yang Anda rencanakan.

Goal: Fitur Versioning berjalan di atas stack baru ini.
Checklist Pengujian:

| ID | Skenario Tes | Tipe | Prioritas |
| :-- | :--- | :-- | :-- |
| S4-01 | Versioning Trigger: Update note yang sudah ada. Pastikan logic Go tidak menimpa data lama, tapi membuat entry baru di tabel versions (sesuai rencana DB baru). | Logic | Critical |
| S4-02 | Diff Logic: Bandingkan Versi 1 dan Versi 2. Pastikan tidak ada karakter aneh/encoding error akibat perbedaan cara handle string antara PHP dan Go. | Data | High |
| S4-03 | WebSocket Auth: Pastikan user harus mengirim token valid saat connect ke WebSocket. Jangan sampai user anonim bisa listen ke perubahan dokumen. | Security | Critical |

💡 Rekomendasi QA Tools untuk Project Ini
Postman / Insomnia: Wajib. Buat "Collection" berisi respons asli dari backend Laravel, lalu jalankan collection yang sama ke backend Go. Hasilnya harus All Green.
K6 (Load Testing): Tool berbasis Go/JS untuk stress test. Sangat cocok untuk test concurrency SQLite.
DB Browser for SQLite: Untuk cek manual integritas data di file .db.
