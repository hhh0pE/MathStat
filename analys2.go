package main
import "fmt"


func analys2Action() (content string) {

    data := global_data

    a, b := a(data), b(data)
    data2 := make([]float64, len(data))

    for i, num := range data {
        if num < a {
            data2[i] = 0
            continue
        }
        if num > b {
            data2[i] = 1
            continue
        }

        data2[i] = (num-a)/(b-num)
    }

    content += fmt.Sprintf(`
    <div><var>a</var> = <var>%6.3f</var></div>
    <div><var>b</var> = <var>%6.3f</var></div>
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
                        ['Варианты', %s]
                    ],

                    types: {
                        'Вероятностная сетка': "line",
                        'Варианты': "area-step"
                    }
                },
                axis: {
                    x: {
                        tick: {
                            format: function(d) { return d+1; }
                        }
                    }
                },

            })
        </script>
    `, DataToString(data2), DataToString(data))

    return
}
