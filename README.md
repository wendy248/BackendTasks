# BackendTasks

Program BackendTasks dibuat tanpa menggunakan docker, sehingga diperlukan untuk menginstall beberapa aplikasi sebelum menjalankan program. 

Aplikasi ini menggunakan database MySQL

Berikut adalah tahapan yang perlu dilakukan untuk menjalankan server yang telah dibuat pada program :
1. Untuk menjalankan program ini, diperlukan untuk membuka XAMPP dan menjalankan Service "Apache" dan "MySQL".
2. Buka url "http://localhost/phpmyadmin" dan buat database baru dengan nama "dbbackend". (tidak perlu membuat table)
3. Jalankan program/server pada main.go dengan mengetik "go run main.go" pada terminal program.

Berikut adalah URL untuk melakukan test API menggunakan Postman :
1. GET ALL      : "localhost:8080/v1/animal"
2. GET by ID    : "localhost:8080/v1/animal/{id}"
3. POST DATA    : "localhost:8080/v1/animal"
4. PUT DATA     : "localhost:8080/v1/animal"
5. DELETE DATA  : "localhost:8080/v1/animal/{animal name}"

Pada repository, juga terlampir Postman Collection yang dapat digunakan untuk pengujian API dengan format .JSON. File tersebut dapat digunakan untuk menguji API dengan menggunakan Postman 

Catatan :
PUT handler menggunakan animal name sebagai identifikasi data di database.