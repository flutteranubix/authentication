http.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
    authCookie, err := r.Cookie("auth")
    if err != nil || authCookie.Value != "authenticated" {
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }

    // Render the protected dashboard
    // ...
})
