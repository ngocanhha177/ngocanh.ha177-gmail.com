package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    // Phản hồi văn bản khi có người truy cập trang web
    fmt.Fprintf(w, "Chào mừng! Tôi là Hà Ngọc Anh 177")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Server đang chạy tại http://localhost:8080 ...")
    
    // Bắt đầu lắng nghe kết nối
    http.ListenAndServe(":8080", nil)
}
