# 📚 CRUD Sederhana dengan Golang

Aplikasi CRUD (Create, Read, Update, Delete) sederhana untuk mengelola data mahasiswa menggunakan:

- **Backend**: Golang
- **Database**: MySQL
- **Frontend**: HTML + CSS
- **Admin Database**: phpMyAdmin

## 📁 Struktur Project

```
Learning-Golang/
├── main.go              # File utama aplikasi (entry point + struct)
├── database.go          # Koneksi dan setup database
├── routes.go            # Semua routing/URL aplikasi
├── handlers.go          # Handler functions untuk setiap route
├── go.mod              # File dependency Go
├── README.md           # File ini
├── templates/          # Folder template HTML
│   ├── index.html      # Halaman utama (list data)
│   ├── tambah.html     # Halaman tambah data
│   └── edit.html       # Halaman edit data
└── static/             # Folder file statis
    └── style.css       # File CSS untuk styling
```

## 🎯 Fitur Aplikasi

- ✅ **Tambah Data**: Menambah data mahasiswa baru
- ✅ **Lihat Data**: Menampilkan semua data mahasiswa dalam tabel
- ✅ **Edit Data**: Mengubah data mahasiswa yang sudah ada
- ✅ **Hapus Data**: Menghapus data mahasiswa
- ✅ **Responsive Design**: Tampilan yang bagus di desktop dan mobile
