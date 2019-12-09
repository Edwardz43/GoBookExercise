package main

import (
	"html/template"
	"log"
	"os"
	"time"

	"github.com/Edwardz43/GoBookExercise/book/ch4/github"
)

var issuesList = template.Must(template.New("issueslist").Parse(`
<h1>{{.TotalCount}} issues</h1> 
<table> 
<tr style='text-align: left'>
	<th>#</th>
	<th>State</th>
	<th>User</th>
	<th>Title</th> 
</tr> 
{{range .Items}} 
<tr>
	<td><a href='{{.HTMLURL}}'>{{.Number}}</td>
	<td>{{.State}}</td>
	<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
	<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td> 
</tr> 
{{end}} 
</table> 
`))

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func main() {

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := issuesList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
