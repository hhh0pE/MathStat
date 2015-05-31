package main
import "fmt"


func analys2Action() (content string) {

    data := global_data

    a, b := a(data), b(data)
    data_F := make([]float64, len(data))

    iteration_data := make([]float64, len(data))

    f := (1/(b-a))*h(data)*60

    classes := Classes(data)

    classes_dencity := make([]float64, len(classes))

    for i, _ := range classes {
        classes_dencity[i] = f
    }

    for i, num := range data {
        iteration_data[i] = float64(i)/float64(len(data))

        if num < a {
            data_F[i] = 0
            continue
        }
        if num > b {
            data_F[i] = 1
            continue
        }

        data_F[i] = (num-a)/(b-a)

    }

    content += fmt.Sprintf(`
    <div><var>a</var> = <var>%6.3f</var></div>
    <div><var>b</var> = <var>%6.3f</var></div>
    <a name="chart1"></a>
    <h1>Вероятностная сетка</h1>
    <div id="chart1"></div>
    `, a, b)

    content += fmt.Sprintf(`
        <script>
            c3.generate({
                bindto: "#chart1",
                data: {
                    columns: [
                        ['Вероятностная сетка', %s],
                        ['Линия', %s]
                    ],

                    types: {
                        'Вероятностная сетка': "scatter",
                        'Линия': "line"
                    }
                },
                axis: {
                    x: {
                        tick: {
                            format: function(d) { return d+1; }
                        }
                    }
                },
                point: {
                    show: false
                }

            })
        </script>
    `, FloatDataToString(data_F), FloatDataToString(iteration_data))

    content += `
    <a name="chart2"></a>
    <h1>Функция плотности</h1>
    <div id="chart2"></div>
    <a name="chart3"></a>
    <h1>Функция распределения</h1>
    <div id="chart3"></div>
    `

    content += fmt.Sprintf(`
    <script>
        c3.generate({
                bindto: "#chart2",
                data: {
                    columns: [
                        ['Классы', %s],
                        ['Функция плотности', %s]
                    ],
                    types: {
                        'Классы': "bar",
                        'Функция плотности': "line"
                    }
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
    `,
    classesValues(classes), FloatDataToString(classes_dencity), classesIntervals(classes), len(data))

    content += fmt.Sprintf(`
     <script>
        c3.generate({
                bindto: "#chart3",
                data: {
                    columns: [
                        ['Варианты', %s],
                        ['Функция распределения', %s]
                    ],
                    types: {
                        'Варианты': "bar",
                        'Функция распределения': "line"
                    }
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
                },
                point: {
                    show: false
                }

            })
    </script>
    `, FloatDataToString(data), FloatDataToString(data_F))

    content += fmt.Sprintf(`
    <a name="kolmagorov"></a>
        <h1>Критерий Колмагорова</h1>
        <div><var>K</var> = %6.4f<var></var></div>
    `,
    Kolmagorov(1, float64(len(data))))

    return
}
