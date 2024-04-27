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
Algoritma BFS sebagai berikut:
1. Pranala awal dimasukkan ke map dengan _key_ judul artikel dan _value_ "start".
2. Pranala akan ditelurusi kemudian diambil seluruh pranala yang mengarah ke laman Wikipedia lainnya.
3. Pranala yang ditemukan akan dimasukkan ke map dengan _key_ pranala yang ditemukan dan pranala yang<br>sedang ditelusuri sebagai _value_.
4. Jika ditemukan judul artikel tujuan ketika pranala ditelusuri, program akan mencari jalurnya dari map.
5. Jika judul artikel yang dituju sama dengan judul artikel yang ditelusuri, maka program akan mencari<br>jalurnya dan menghentikan penelusuran
<br>
Algoritma IDS sebagai berikut:
1. Algoritma IDS menggunakan DLS dengan pengecekan awal pada kedalaman 0, kemudian bertambah 1 di setiap iterasinya
2. Program akan mengecek link yang pertama ditemukan sampai batas kedalaman tertentu.
3. Program tidak melakukan _backtrack_, tetapi mengecek setiap link yang ditemukan secara rekursif
4. Jika ditemukan judul artikel tujuan ketika pranala ditelusuri, pencarian ke bawah akan dihentikan.
5. Jika judul artikel yang dituju sama dengan judul artikel yang ditelusuri, program berhenti mencari<br>
dan menyimpan rute hasil.

## Requirement
1. Golang
2. Gocolly
3. React.js
4. npm

## Setup and Usage
1. Pastikan Requirement di atas sudah terinstall dengan benar.
1. Clone repository ini dengan command
```
git clone 
```

## Authors
| NIM | Nama |
|-----|------|
| 13522004 | Eduardus Alvito Kristiadi |
| 13522033 | Bryan Cornelius Lauwrence |
| 13522049 | Vanson Kurnialim |