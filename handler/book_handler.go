package book_handler

import (
	"encoding/json" //karna kita mau mengencode data ke json
	"net/http" //untuk menangani request dan response HTTP
	"sis-pustaka/model"	//ambil model yang kita buat sebelumnya	
	"sis-pustaka/storage" //ambil storage yang kita buat sebelumnya
	"github.com/go-chi/chi/v5" //import chi untuk routing
)


