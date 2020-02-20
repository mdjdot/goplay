package main 
import ( 
    "fmt" 
    "net/http" 
) 
func handler(w http.ResponseWriter, r *http.Request) { 
    fmt.Fprintln(w, "正在通过自己创建的多路复用器处理你的请求") 
} 
func main() { 
    mux := http.NewServeMux() 
    mux.HandleFunc("/myMux", handler) 
    http.ListenAndServe(":8080", mux) 
} 