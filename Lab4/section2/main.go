package mycalc

import (
        "html/template"
        "net/http"
        "strconv"

        "appengine"
        "appengine/datastore"
)

type RESULT struct {
        P           int64
        Q           int64		
        Div     	float64
}

func init() {
        http.HandleFunc("/", root)
        http.HandleFunc("/sign", sign)
}


func resultKey(c appengine.Context) *datastore.Key {
        return datastore.NewKey(c, "MyResult", "my_result", 0, nil)
}

func root(w http.ResponseWriter, r *http.Request) {
        c := appengine.NewContext(r)
        q := datastore.NewQuery("RESULT").Ancestor(resultKey(c)).Limit(10)
        results := make([]RESULT, 0, 10)
        if _, err := q.GetAll(c, &results); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }

        if err := mainpageTemplate.Execute(w, results); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
        }
}

var mainpageTemplate = template.Must(template.New("mainpage").Parse(`
<html>
  <head>
    <title>Lab 4</title>
    <style>
        table, th, td {
            border: 1px solid black;
            border-collapse: collapse;
        }
    </style>
  </head>
  <body>
    <table style="width:300px;margin:10px;">
    <tr>
        <th>P</th>
        <th>Q</th> 		
        <th>Div</th>        
    </tr>
    {{range .}}
    <tr>
        <td>{{.P}}</td>
        <td>{{.Q}}</td> 
        <td>{{.Div | printf "%.2f"}}</td>
    </tr>
    {{end}}
    </table>
    <form action="/sign" method="post">
      <div>P<input name="value_p" type="number"></input></div>
      <div>Q<input name="value_q" type="number"></input></div>
      <div><input type="submit" value="Submit"></div>
    </form>
  </body>
</html>
`))

func sign(w http.ResponseWriter, r *http.Request) {
        c := appengine.NewContext(r)
        p, _ := strconv.ParseInt(r.FormValue("value_p"), 10, 64)
        q, _ := strconv.ParseInt(r.FormValue("value_q"), 10, 64)		
        d := float64(p)/float64(q)        
        g := RESULT{
                P:          p,
                Q:          q,                
                Div:        d,                
        }
        key := datastore.NewIncompleteKey(c, "RESULT", resultKey(c))
        _, err := datastore.Put(c, key, &g)
        if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }
        http.Redirect(w, r, "/", http.StatusFound)
}