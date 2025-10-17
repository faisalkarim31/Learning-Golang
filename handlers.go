package main

import (
	"html/template"
	"net/http"
	"strconv"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query("SELECT id, nama, nim, email FROM mahasiswa ORDER BY id DESC")
	if err != nil {
		http.Error(w, "Gagal mengambil data: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var mahasiswaList []Mahasiswa
	for rows.Next() {
		var m Mahasiswa
		err := rows.Scan(&m.ID, &m.Nama, &m.NIM, &m.Email)
		if err != nil {
			http.Error(w, "Gagal membaca data: "+err.Error(), http.StatusInternalServerError)
			return
		}
		mahasiswaList = append(mahasiswaList, m)
	}

	tmpl, err := template.New("index.html").Funcs(template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
	}).ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Gagal load template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, mahasiswaList)
	if err != nil {
		http.Error(w, "Gagal render template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func TambahHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/tambah.html")
	if err != nil {
		http.Error(w, "Gagal load template: "+err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func SimpanHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	nama := r.FormValue("nama")
	nim := r.FormValue("nim")
	email := r.FormValue("email")

	_, err := DB.Exec("INSERT INTO mahasiswa (nama, nim, email) VALUES (?, ?, ?)", nama, nim, email)
	if err != nil {
		http.Error(w, "Gagal menyimpan data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func EditHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/edit/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID tidak valid", http.StatusBadRequest)
		return
	}

	var m Mahasiswa
	err = DB.QueryRow("SELECT id, nama, nim, email FROM mahasiswa WHERE id = ?", id).Scan(&m.ID, &m.Nama, &m.NIM, &m.Email)
	if err != nil {
		http.Error(w, "Data tidak ditemukan: "+err.Error(), http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("templates/edit.html")
	if err != nil {
		http.Error(w, "Gagal load template: "+err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, m)
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		http.Error(w, "ID tidak valid", http.StatusBadRequest)
		return
	}

	nama := r.FormValue("nama")
	nim := r.FormValue("nim")
	email := r.FormValue("email")

	_, err = DB.Exec("UPDATE mahasiswa SET nama = ?, nim = ?, email = ? WHERE id = ?", nama, nim, email, id)
	if err != nil {
		http.Error(w, "Gagal update data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func HapusHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/hapus/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID tidak valid", http.StatusBadRequest)
		return
	}

	_, err = DB.Exec("DELETE FROM mahasiswa WHERE id = ?", id)
	if err != nil {
		http.Error(w, "Gagal menghapus data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
