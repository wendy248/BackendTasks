# BackendTasks

Program BackendTasks dibuat tanpa menggunakan docker, sehingga diperlukan untuk menginstall beberapa aplikasi sebelum menjalankan program. 

Berikut adalah aplikasi yang digunakan dalam proses pembuatan :
    1. Visual Code Studio
    2. XAMPP    (Program ini menggunakan database MySQL)
    3. Postman
    4. Browser

Berikut adalah tahapan yang perlu dilakukan untuk menjalankan server yang telah dibuat pada program :
    1. Untuk menjalankan program ini, diperlukan untuk membuka XAMPP dan menjalankan Service "Apache" dan "MySQL".
    2. Buka url "http://localhost/phpmyadmin" dan buat database baru dengan nama "dbbackend".
    3. Jalankan program/server pada main.go dengan mengetik "go run main.go" pada terminal program.

Berikut adalah tahapan untuk melakukan test API menggunakan Postman :
    1. Untuk GET ALL request,  