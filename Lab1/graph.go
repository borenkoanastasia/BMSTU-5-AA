package main 
import (
    //"bytes"
    "os"
    "github.com/wcharczuk/go-chart" //exposes "chart"
)



func make_graph(x[] float64, y[] float64){
    graph := chart.Chart{
        Series: []chart.Series{
            chart.ContinuousSeries{
                XValues: x,//[]float64{1.0, 2.0, 3.0, 4.0},
                YValues: y,//[]float64{1.0, 2.0, 3.0, 4.0},
            },
        },
    }

    //buffer := bytes.NewBuffer([]byte{})
    //err := graph.Render(chart.PNG, buffer)
    //fmt.Print( err)
    
	pngFile, err := os.Create("output.png")
	if err != nil {
		panic(err)
	}

	if err := graph.Render(chart.PNG, pngFile); err != nil {
		panic(err)
	}

	if err := pngFile.Close(); err != nil {
		panic(err)
	}
}


