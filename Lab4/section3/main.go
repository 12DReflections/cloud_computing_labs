package ED_EXAMPLE

import (
		"fmt"
        "net/http"
        "math"
        "math/rand"
        "sort"
        "time"
)

func init() {
        http.HandleFunc("/", root)
}

func root(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, 
		"<html>",
		  "<head>",
		    "<title>Point set</title>",
		    "<style>",
		        "table, th, td {",
		            "border: 1px solid black;",
		            "border-collapse: collapse;",
		        "}",
		    "</style>",
		  "</head>",
		  "<body>")

		t0 := time.Now()
        test1(w)
        t1 := time.Now()

		fmt.Fprint(w, "<pre>The call took ", t1.Sub(t0), " to run.\n</pre>")

        fmt.Fprint(w,
			"</body>",
		"</html>")
}

//To test showing the distance of two points
func test1(w http.ResponseWriter) {
	    pt1 := Point {
        	x: 1,
        	y: 1,
        }

        pt2 := Point {
        	x: 2,
        	y: 2,
        }

        showResult(w, pt1, pt2)
}

//to create 10 points randomly
func test2(w http.ResponseWriter) {
		num := 10
		ptSet := randomDataset(num)

		showTable(w, ptSet, num)
}

//to test searching the nearest point
func test3(w http.ResponseWriter) {
	num := 10
	ptSet := randomDataset(num)

	startPt := 0

	npt := search1(ptSet[startPt], startPt, ptSet)

	showTable(w, ptSet, num)
	fmt.Fprint(w, "<pre>The nearest point: [", npt.x , npt.y, "]</pre>")
	showResult(w, npt, ptSet[startPt])
}


//To test searching n nearest points
func test4(w http.ResponseWriter) {
	num := 10
	ptSet := randomDataset(num)

	startPt := 0

	count := 3

	npt := search3(ptSet[startPt], startPt, count, ptSet)

	showTable(w, ptSet, num)

	for i:= 0; i < count; i++ {
		fmt.Fprint(w, "<pre>The nearest point ", i, ": [", npt[i].x , npt[i].y, "]</pre>")
		showResult(w, npt[i], ptSet[startPt])

		fmt.Fprint(w, "<pre>----------------------------------------------------------</pre>")
	}
}


//To test sorting
func test5(w http.ResponseWriter) {
	num := 10
	ptSet := randomDataset(num)

	startPt := 0

	count := 3

	result := search2(ptSet[startPt], ptSet)

	showTable(w, ptSet, num)

	for i:= 1; i <= count; i++ {
		pt := ptSet[result[i].idx]
		fmt.Fprint(w, "<pre>The nearest point ", i, ": [", pt.x , pt.y, "]</pre>")
		showResult(w, pt, ptSet[startPt])

		fmt.Fprint(w, "<pre>----------------------------------------------------------</pre>")
	}
}

func showResult(w http.ResponseWriter, pt1 Point, pt2 Point) {
	    fmt.Fprint(w, "<pre>Point 1: [", pt1.x, pt1.y, "]\n</pre>")
        fmt.Fprint(w, "<pre>Point 2: [", pt2.x, pt2.y, "]\n</pre>")

        fmt.Fprint(w, "<pre>Distance: ", distance(pt1, pt2), "</pre>")
}

func showTable(w http.ResponseWriter, ptSet []Point, num int) {
		fmt.Fprint(w, 
		"<table>",
		    "<tr>",
		        "<th>X</th>",
		        "<th>Y</th>",
		    "</tr>")
		

		for i := 0; i < num; i++ {
			fmt.Fprint(w,
				"<tr> <td>", ptSet[i].x, "</td> <td>", ptSet[i].y, "</td> </tr>")
		}

		fmt.Fprint(w, "</table>")
}

type Point struct {
        x           float64
        y           float64
}

func distance(a Point, b Point) float64 {
    x := a.x - b.x
    y := a.y - b.y
    return math.Sqrt(x*x + y*y)
}

func randomDataset(num int) []Point {
    var ds []Point
    for i := 0; i < num; i++ { 
        x := rand.Float64()
        y := rand.Float64()
        ds = append(ds,
        	Point{
        		x:      x,
            	y:      y,
        	})
    }
    return ds
}


//To search for the nearest point
func search1(c Point, idx int, dataSet []Point) Point {
    minDist := math.MaxFloat64
    minPt := dataSet[0]

    for i := 0; i < len(dataSet); i++ {
        dis := distance(c, dataSet[i])

        if idx != i && minDist > dis {
            minDist = dis
            minPt = dataSet[i]
        }
    }

    return minPt
}


//To search for n nearest points
func search3(c Point, idx int, n int, dataSet []Point) []Point {
    result := make([]Point, n)
    dList := make([]float64, n)

    for j := 0; j < n; j++ {
        dList[j] = math.MaxFloat64
    }

    for i := 0; i < len(dataSet); i++ {
    	if idx != i {
	        dis := distance(c, dataSet[i])

	  		// To add the new shorest distance and sort
	        for j := 0; j < n; j++ {
	            if dList[j] > dis {
	                for k := n-1; k < j; k-- {
	                    dList[k] = dList[k-1]
	                }

	                dList[j] = dis
	                result[j] = dataSet[i]
	                break
	            }
	        }
    	}     
    }

    return result
}


//To sort whole distances
type Item struct {
    dis         float64
    idx         int
}

type ByDist []Item

func (a ByDist) Len() int { return len(a) }
func (a ByDist) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByDist) Less(i, j int) bool { return (a[i].dis < a[j].dis) }

func search2(c Point, dataSet []Point) []Item {
    dataLength := len(dataSet)
    dList := make([]Item, dataLength)

    for i := 0; i < len(dataSet); i++ {
        dList[i] = Item {
            dis:    distance(c, dataSet[i]),
            idx:    i,
        }
    }

    sort.Sort(ByDist(dList))

    return dList
}