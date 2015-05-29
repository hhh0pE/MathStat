package main

import (
	"fmt"
	"math"
    "sort"
)

func analysAction() string {

    data := global_data

    sort.Float64s(data)

	size := len(data)

	frequency := make(map[float64]int, len(data))

	for _, number := range data {
		frequency[number]++
	}

	content := `<a name="variations"></a>`
	content += "<h1>Вариационный ряд</h1>"

	content += `
        <table class="table">
            <tr>
                <td>#</td>
                <td>Варианта</td>
                <td>Частота</td>
                <td>Относительная частота</td>
                <td>Значение эмпирической фунции распределения</td>
            </tr>`

	counter := 0
	a := 0
	for i, val := range data {
		if i > 0 && data[i-1] == val {
			continue
		}
		a++
		counter += frequency[val]
		content += fmt.Sprintf(`
        <tr>
            <td>%d</td>
            <td>%6.3f</td>
            <td>%d</td>
            <td>%d/%d</td>
            <td>%d/%d</td>
        </tr>`, a, val, frequency[val], frequency[val], size, counter, size)
	}

	min, max := data[0], data[len(data)-1]

    if global_M == 0 {
        global_M = int(math.Floor(math.Sqrt(float64(size))))
    }
    M := global_M
	h := (max - min) / float64(M)

	content += `</table>
    <div id="chart1"></div>
    <div id="chart3"></div>
    <a name="classes"></a>`

	content += `<h1>Классы</h1>`

	content += fmt.Sprintf(`
    <div class="panel-default">
        <div>N = %d</div>
        <div>Max = %6.2f</div>
        <div>Min = %6.2f</div>
        <div>M = %d</div>
        <div>h = %6.3f</div>
    </div>`, size, min, max, M, h)

	content += fmt.Sprintf(`
        <script>
            c3.generate({
                bindto: "#chart1",
                data: {
                    columns: [
                        ['', %s]
                    ]
                },
                axis: {
                    x: {
                        tick: {
                            format: function(d) { return d+1; }
                        }
                    }
                }
            })
        </script>
    `, DataToString(data))

	classes := make([]Class, M)

	data_i := 0
	for class_i, _ := range classes {
		classes[class_i].Min = min + float64(class_i)*h
		classes[class_i].Max = min + (float64(class_i)+1)*h

		//        fmt.Printf("Min(%6.3f) Max(%6.3f)\n", classes[class_i].Min, classes[class_i].Max)

		for ; data[data_i] <= classes[class_i].Max; data_i++ {
			if data_i > 0 && data[data_i] == data[data_i-1] {
				continue
			}

			classes[class_i].Frequency += frequency[data[data_i]]

			if data_i == len(data)-1 {
				break
			}
		}

		if class_i > 0 {
			classes[class_i].FuncNum = classes[class_i-1].FuncNum + classes[class_i].Frequency
		} else {
			classes[0].FuncNum = classes[0].Frequency
		}
	}

	content += `
        <table class="table">
            <tr>
                <td>#</td>
                <td>Границы</td>
                <td>Частота</td>
                <td>Относительная частота</td>
                <td>Значение эмпирической фунции распределения</td>
            </tr>
            `

	for i, class := range classes {
		content += fmt.Sprintf(`
            <tr>
                <td>%d</td>
                <td>[%6.2f;%6.2f )</td>
                <td>%d</td>
                <td>%d/%d</td>
                <td>%d/%d</td>
            </tr>
        `, i+1, class.Min, class.Max, class.Frequency, class.Frequency, size, class.FuncNum, size)
	}
	content += `</table>`

	content += fmt.Sprintf(`<div id="chart2"></div>
    <div id="chart4"></div>
    <script>
        c3.generate({
                bindto: "#chart2",
                data: {
                    columns: [
                        ['Classes', %s]
                    ],
                    type: "bar"
                },
                axis: {
                    x: {
                        type: 'category',
                        categories: [%s]
                    },
                    y: {
                        tick: {
                            format: function(d) { return d+"/%d";}
                        }
                    }
                },
                bar: {
                    width: {
                        ratio: 1
                    }
                }
            })

            c3.generate({
                bindto: "#chart3",
                data: {
                    columns: [
                        ['', %s]
                    ],
                    type: "bar"
                },
                axis: {
                    x: {
                        tick: {
                            format: function(d) { return d+1; }
                        }
                    }
                },
                bar: {
                    width: {
                        ratio: 1
                    }
                }
            })

            c3.generate({
                bindto: "#chart4",
                data: {
                    columns: [
                        ['Classes', %s]
                    ],
                    type: "bar"
                },
                axis: {
                    x: {
                        type: 'category',
                        categories: [%s]
                    },
                    y: {
                        tick: {
                            format: function(d) { return d+"/%d";}
                        }
                    }
                },
                bar: {
                    width: {
                        ratio: 1
                    }
                }
            })
    </script>
    `, classesValues(classes), classesIntervals(classes), size, DataToString(data), classesFullValues(classes), classesIntervals(classes), size)

	content += "<h1>Количественные характеристики выборки</h1>"

	N := float64(len(data))

	average := Average(data)
	S := S(data)

	E := E(data)
	cE := contrE(data)

	A := A(data)

	W := W(data)

	S_average := average / math.Sqrt(N)
	S_S := S / math.Sqrt(2*N)
	S_A := math.Sqrt((6.0 * (N - 2.0)) / ((N + 1.0) * (N + 3.0)))
	S_E := (24 * N * (N - 2) * (N - 3)) / (math.Pow(N+1, 2) * (N + 3) * (N + 5))

	S_cE := math.Sqrt(Module(cE)/29*N) * math.Pow(Module(math.Pow(cE*cE-1, 3)), 1/4)

	S_W := W * math.Sqrt(math.Pow(1+2*W, 2)/(2*N))

	content += `
        <table class="table">
            <tr>
                <td></td>
                <td>Значение</td>
                <td>Среднеквадратическое отклонение</td>
                <td>Доверительный интервал</td>
            </tr>`
	content += fmt.Sprintf(`
        <tr>
            <td>Среднее арифметическое</td>
            <td>%6.4f</td>
            <td>%6.4f</td>
            <td>[ %6.4f; %6.4f ]</td>
        </tr>`, average, S_average, IntervalMin(average, S_average), IntervalMax(average, S_average))
	content += fmt.Sprintf(`
        <tr>
            <td>Медиана</td>
            <td>%6.4f</td>
            <td>-</td>
            <td>-</td>
        </tr>`, Mediana(data))
	content += fmt.Sprintf(`
        <tr>
            <td>Среднеквадратическое</td>
            <td>%6.4f</td>
            <td>%6.4f</td>
            <td>[ %6.4f; %6.4f ]</td>
        </tr>
        `, S, S_S, IntervalMin(S, S_S), IntervalMax(S, S_S))
	content += fmt.Sprintf(`<tr>
            <td>Коефициент ассиметрии</td>
            <td>%6.4f</td>
            <td>%6.4f</td>
            <td>[ %6.4f; %6.4f ]</td>
        </tr>`, A, S_A, IntervalMin(A, S_A), IntervalMax(A, S_A))
	content += fmt.Sprintf(`<tr>
            <td>Коефициент эксцесса</td>
            <td>%6.4f</td>
            <td>%6.4f</td>
            <td>[ %6.4f; %6.4f ]</td>
        </tr>`, E, S_E, IntervalMin(E, S_E), IntervalMax(E, S_E))
	content += fmt.Sprintf(`
        <tr>
            <td>Коефициент контрэксцесса</td>
            <td>%6.4f</td>
            <td>%6.4f</td>
            <td>[ %6.4f; %6.4f ]</td>
        </tr>`, cE, S_cE, IntervalMin(cE, S_cE), IntervalMax(cE, S_cE))
	content += fmt.Sprintf(`
        <tr>
            <td>Коефициент вариации</td>
            <td>%6.4f</td>
            <td>%6.4f</td>
            <td>[ %6.4f; %6.4f ]</td>
        </tr>
    `, W, S_W, IntervalMin(W, S_W), IntervalMax(W, S_W))

	return content
}
