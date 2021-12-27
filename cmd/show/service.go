/*MIT License

Copyright (c)  2020-2021 YPSI SAS
Centctl is developped by : Mélissa Bertin

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
package show

import (
	"bytes"
	"centctl/display"
	"centctl/request"
	"centctl/resources/service"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
)

// serviceCmd represents the service command
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Show one service's details",
	Long:  `Show one service's details of the Centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		output, _ := cmd.Flags().GetString("output")
		metrics, _ := cmd.Flags().GetBool("metrics")
		pathGraph, _ := cmd.Flags().GetString("pathGraph")
		err := ShowService(name, description, debugV, output, metrics, pathGraph)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ShowService permits to display the details of one service
func ShowService(name string, description string, debugV bool, output string, metrics bool, pathGraph string) error {
	output = strings.ToLower(output)

	values := name + ";" + description
	err, body := request.GeneriqueCommandV1Post("show", "service", values, "show service", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the services contain into the response body
	services := service.DetailResult{}
	json.Unmarshal(body, &services)

	//Permits to find the good service in the array
	var ServiceFind service.DetailService
	for _, v := range services.DetailServices {
		if strings.ToLower(v.HostName) == strings.ToLower(name) && strings.ToLower(v.Description) == strings.ToLower(description) {
			ServiceFind = v
		}
	}

	var server service.DetailServer
	if ServiceFind.Description != "" {
		//Organization of data
		server = service.DetailServer{
			Server: service.DetailInformations{
				Name:    os.Getenv("SERVER"),
				Service: &ServiceFind,
			},
		}
	} else {
		server = service.DetailServer{
			Server: service.DetailInformations{
				Name:    os.Getenv("SERVER"),
				Service: nil,
			},
		}
	}

	//Display details of the service
	displayService, err := display.DetailService(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayService)

	if metrics {
		hostID, _ := strconv.Atoi(ServiceFind.HostID)
		serviceID, _ := strconv.Atoi(ServiceFind.ID)
		displayMetrics(hostID, serviceID, debugV, description, pathGraph)
	}

	return nil
}

type Graph struct {
	Global  map[string]string `json:"global"`
	Metrics []Metric          `json:"metrics"`
	Times   []string          `json:"times"`
}

type Metric struct {
	Legend       string            `json:"legend"`
	DsData       map[string]string `json:"ds_data"`
	Data         []float64         `json:"data"`
	Unit         string            `json:"unit"`
	MetricLegend string            `json:"metric_legend"`
	MinimumValue float64           `json:"minimum_value"`
	MaximumValue float64           `json:"maximum_value"`
}

func xvalues(rawx []string) []time.Time {

	var dates []time.Time
	for _, ts := range rawx {
		parsed, _ := time.Parse(time.RFC3339, ts)
		dates = append(dates, parsed)
	}
	return dates
}

func deleteNullValues(data []float64, hours []time.Time) ([]float64, []time.Time) {
	for data[len(data)-1] == 0 {
		if len(data) == 1 {
			return []float64{}, []time.Time{}
		}
		data = data[:len(data)-1]
		hours = hours[:len(hours)-1]
	}
	return data, hours
}

func getMinMaxValue(data []float64) (float64, float64) {
	var max float64 = data[0]
	var min float64 = data[0]
	for _, value := range data {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func displayMetrics(hostID int, serviceID int, debugV bool, description string, pathGraph string) {
	urlCentreon := "/monitoring/hosts/" + strconv.Itoa(hostID) + "/services/" + strconv.Itoa(serviceID) + "/metrics/performance"
	_, body := request.GeneriqueCommandV2Get(urlCentreon, "show service", debugV)

	graphe := Graph{}
	_ = json.Unmarshal(body, &graphe)

	xv := xvalues(graphe.Times)

	if len(graphe.Metrics) == 0 {
		fmt.Println("no metrics")
		os.Exit(1)
	}

	for _, m := range graphe.Metrics {
		datas, hours := deleteNullValues(m.Data, xv)
		if len(datas) == 0 {
			continue
		}
		min, max := getMinMaxValue(datas)
		if min == max {
			min = min - 1.0
			max = max + 1.0
		}
		ts1 := chart.TimeSeries{ //TimeSeries{
			Name: m.Legend,
			Style: chart.Style{
				Show:        true,
				StrokeColor: drawing.ColorFromHex(strings.ReplaceAll(m.DsData["ds_color_line"], "#", "")),
				FillColor:   drawing.ColorFromHex(strings.ReplaceAll(m.DsData["ds_color_line"], "#", "")).WithAlpha(64),
			},
			XValues: hours,
			YValues: datas,
		}
		graph := chart.Chart{
			Title:      graphe.Global["title"],
			TitleStyle: chart.Style{Show: true},
			Background: chart.Style{
				Padding: chart.Box{
					Top:    50,
					Left:   25,
					Right:  25,
					Bottom: 10,
				},
				FillColor: drawing.ColorFromHex("efefef"),
			},

			XAxis: chart.XAxis{
				Style:          chart.Style{Show: true},
				ValueFormatter: chart.TimeHourValueFormatter,
			},

			YAxis: chart.YAxis{
				Name:           m.Unit,
				NameStyle:      chart.StyleShow(),
				Style:          chart.Style{Show: true},
				ValueFormatter: chart.FloatValueFormatter,
				Range: &chart.ContinuousRange{
					Min: min,
					Max: max,
				},
			},

			Series: []chart.Series{ts1},
		}
		graph.Elements = []chart.Renderable{
			chart.LegendLeft(&graph),
		}

		buffer := bytes.NewBuffer([]byte{})
		err := graph.Render(chart.PNG, buffer)
		if err != nil {
			log.Fatal(err)
		}

		_ = os.Mkdir(pathGraph, os.ModePerm)
		fo, err := os.Create(pathGraph + "/" + m.MetricLegend + ".png")
		if err != nil {
			panic(err)
		}

		if _, err := fo.Write(buffer.Bytes()); err != nil {
			panic(err)
		}
	}

}

func init() {
	serviceCmd.Flags().StringP("name", "n", "", "Host's name")
	serviceCmd.MarkFlagRequired("name")
	serviceCmd.RegisterFlagCompletionFunc("name", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var values []string
		if request.InitAuthentification(cmd) {
			values = request.GetHostNames()
		}
		return values, cobra.ShellCompDirectiveDefault
	})
	serviceCmd.Flags().StringP("description", "d", "", "Service's description")
	serviceCmd.MarkFlagRequired("description")
	serviceCmd.RegisterFlagCompletionFunc("description", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var values []string
		if serviceCmd.Flag("name").Value.String() != "" {
			if request.InitAuthentification(cmd) {
				values = request.GetServiceDescriptions(serviceCmd.Flag("name").Value.String())
			}
		}

		return values, cobra.ShellCompDirectiveDefault
	})
	serviceCmd.Flags().Bool("metrics", false, "To display metrics of the service")
	serviceCmd.Flags().StringP("pathGraph", "p", "graphs", "Path directory for metric's graph (default in current directory)")
}
