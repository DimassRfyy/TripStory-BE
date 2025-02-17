# TripStory Backend

TripStory Backend adalah aplikasi backend untuk platform TripStory yang dibangun menggunakan arsitektur hexagonal. Arsitektur ini memungkinkan pemisahan yang jelas antara logika bisnis dan infrastruktur, sehingga memudahkan pengembangan, pemeliharaan, dan pengujian.

## Struktur Proyek

- **cmd/**: Berisi kode untuk menjalankan aplikasi.
- **config/**: Berisi konfigurasi aplikasi, termasuk konfigurasi database.
- **database/**: Berisi skema dan migrasi database.
- **internal/**: Berisi kode internal aplikasi yang diorganisir berdasarkan lapisan arsitektur hexagonal.
  - **adapter/**: Berisi adapter untuk berinteraksi dengan infrastruktur eksternal seperti database, API eksternal, dll.
  - **app/**: Berisi kode aplikasi utama.
  - **core/**: Berisi logika bisnis utama.
    - **domain/**: Berisi entitas dan objek nilai.
    - **service/**: Berisi layanan aplikasi.
- **lib/**: Berisi pustaka dan utilitas yang digunakan oleh aplikasi.
- **main.go**: Titik masuk utama aplikasi.