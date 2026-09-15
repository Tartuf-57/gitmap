package main


func main() {
	var folder string
	var email string
	flag.StringVar(&folder, "add","", "Add a new folder to scan for git contributions")
	flag.StringVar(&email, "email", "your@email.com", "The email to scan")
	flag.Parse()

	if folder != "" {
		scan(folder)
		return
	}
	stats(email)
}

