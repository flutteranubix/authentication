http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        username := r.FormValue("username")
        password := r.FormValue("password")

        // Check if username and password are correct
        if username == "your_username" && password == "your_password" {
            // Authentication successful
            // Set a session or cookie to remember the user
            http.SetCookie(w, &http.Cookie{
                Name:  "auth",
                Value: "authenticated",
            })
            http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
            return
        }
    }

    // If authentication fails or it's a GET request, show the login form
    http.ServeFile(w, r, "login.html")
})
