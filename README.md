# Tugas Besar 2 IF2211 Strategi Algoritma<br>Semester II tahun 2023/2024<br>_WikiRace Solver_ Menggunakan Algortima IDS dan BFS
> by Dominus Vobiscum

## Table of Contents
* [General Information](#general-information)
* [Requirement](#requirement)
* [Setup and Usage](#setup-and-usage)
* [Authors](#authors)

## General Information
Pencarian jalur dari satu laman Wikipedia ke laman Wikipedia lainnya dengan algoritma BFS dan IDS<br>
<br>
Algoritma BFS sebagai berikut:<br>
a. Pranala awal dimasukkan ke map dengan _key_ judul artikel dan _value_ "start".<br>
b. Pranala akan ditelurusi kemudian diambil seluruh pranala yang mengarah ke laman Wikipedia lainnya.<br>
c. Pranala yang ditemukan akan dimasukkan ke map dengan _key_ pranala yang ditemukan dan pranala yang<br>sedang ditelusuri sebagai _value_.<br>
d. Jika ditemukan judul artikel tujuan ketika pranala ditelusuri, program akan mencari jalurnya dari map.<br>
e. Jika judul artikel yang dituju sama dengan judul artikel yang ditelusuri, maka program akan mencari<br>jalurnya dan menghentikan penelusuran<br>

<br>
Algoritma IDS sebagai berikut:<br>
a. Algoritma IDS menggunakan DLS dengan pengecekan awal pada kedalaman 0, kemudian bertambah 1 di setiap iterasinya<br>
b. Program akan mengecek link yang pertama ditemukan sampai batas kedalaman tertentu.<br>
c. Program tidak melakukan _backtrack_, tetapi mengecek setiap link yang ditemukan secara rekursif<br>
d. Jika ditemukan judul artikel tujuan ketika pranala ditelusuri, pencarian ke bawah akan dihentikan.<br>
e. Jika judul artikel yang dituju sama dengan judul artikel yang ditelusuri, program berhenti mencari<br>
dan menyimpan rute hasil.<br>

## Requirement
1. Golang
2. Gocolly
3. React.js
4. npm

## Setup and Usage
1. Pastikan Requirement di atas sudah terinstall dengan benar.
2. Clone repository ini dengan command
```
git clone https://github.com/Edvardesu/Tubes2_Dominus-Vobiscum.git
```
3. Masuk ke folder website dengan perintah
```
cd website
```
4. Jalankan website dengan command
```
npm run dev
```
5. Masuk ke folder backend dengan command
```
cd src/backend
```
6. Jalankan command berikut
```
go run scrap.go BFS.go IDS.go
```
7. Buka website pada localhost
8. Masukkan artikel awal pada input box di atas dan artikel tujuan pada input box bawah
9. Pilih algoritma yang ingin dijalankan (BFS atau IDS)
10. Jika ingin mencari hanya 1 path, centang kotak `Single Path`
11. Tekan tombol `Sikat!!!`
12. Tunggu hingga hasil ditampilkan

## Authors
| NIM | Nama |
|-----|------|
| 13522004 | Eduardus Alvito Kristiadi |
| 13522033 | Bryan Cornelius Lauwrence |
| 13522049 | Vanson Kurnialim |