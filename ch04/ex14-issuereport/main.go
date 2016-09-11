package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const templ = `
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
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
<h2>Bugs</h2>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items | listBugReports}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
</table>
<h2>Users</h2>
<table>
<tr style='text-align: left'>
  <th>Avatar</th>
  <th>User</th>
</tr>
{{range .Items | listUsers}}
<tr>
  <td><img src='{{.AvatarURL}}' width='20' height='20' /></td>
  <td><a href='{{.HTMLURL}}'>{{.Login}}</a></td>
</tr>
{{end}}
</table>
<h2>Milestones</h2>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>Milestone</th>
</tr>
{{range .Items | listMilestone}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`

func listBugReports(issues []*Issue) []*Issue {
	var bugIssues []*Issue

	for _, issue := range issues {
		if issue.Labels != nil {
			isBug := false
			for _, l := range issue.Labels {
				if l.Name == "bug" {
					isBug = true
				}
			}
			if isBug {
				bugIssues = append(bugIssues, issue)
			}
		}
	}

	return bugIssues
}

func listUsers(issues []*Issue) []*User {
	users := make(map[string]*User)

	// remove duplication
	for _, issue := range issues {
		if issue.User != nil {
			users[issue.User.Login] = issue.User
		}
	}

	var list []*User
	for _, v := range users {
		list = append(list, v)
	}

	return list
}
func listMilestone(issues []*Issue) []*Milestone {
	milestones := make(map[string]*Milestone)

	// remove duplication
	for _, issue := range issues {
		if issue.Milestone != nil {
			milestones[issue.Milestone.Title] = issue.Milestone
		}
	}

	var list []*Milestone
	for _, v := range milestones {
		list = append(list, v)
	}

	return list
}

var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"listMilestone": listMilestone}).
	Funcs(template.FuncMap{"listUsers": listUsers}).
	Funcs(template.FuncMap{"listBugReports": listBugReports}).
	Parse(templ))

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}

		val, ok := r.Form["q"]
		if ok {
			var terms []string
			terms = append(terms, val[0])

			result, err := SearchIssues(terms)
			if err != nil {
				log.Fatal(err)
			}
			if err := report.Execute(w, result); err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Fprint(w, "invalid query")
		}
	})

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
