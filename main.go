package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v1"
	"io/ioutil"
	"math"
	"os"
	"text/tabwriter"
)

var (
	pertfile = flag.String("file", "pert.yml", "YAML PERT Input file")
)

func pert(o float64, m float64, p float64) (float64, float64) {
	// Estimated time
	e := (o + (4 * m) + p) / 6
	// Variance
	v := math.Pow(p-o, 2) / math.Pow(6, 2)
	// Standard Deviation
	sd := math.Sqrt(v)

	return e, sd
}

type TaskContainer struct {
	Tasks []Task
}

type Task struct {
	Name         string
	Optimal      float64
	Nominal      float64
	Pessimistic  float64
	Dependencies []Task
}

func main() {

	fc, _ := ioutil.ReadFile(*pertfile)

	tc := TaskContainer{}

	yaml.Unmarshal(fc, &tc)

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)

	fmt.Fprintf(w, "Task\tO\tM\tP\t|Est\tSD\t\n")

	for _, value := range tc.Tasks {
		e, sd := pert(value.Optimal, value.Nominal, value.Pessimistic)
		fmt.Fprintf(w, "%v\t%v\t%.2f\t%.2f\t|%.2f\t%.2f\t\n", value.Name, value.Optimal, value.Nominal, value.Pessimistic, e, sd)
	}

	w.Flush()

}
