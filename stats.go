package main

func stats(email string) {
	commits := processRepos(email)
	printCommitStats(commits)
}


