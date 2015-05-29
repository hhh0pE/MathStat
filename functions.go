package main
import (
    "fmt"
    "math"
)

type Class struct {
    Min, Max float64
    Frequency int
    FuncNum int
}


func classesFullValues (classes []Class) string {
    str := ""
    var counter int
    for i, class := range classes {
        counter += class.Frequency
        str += fmt.Sprintf("%d", counter)
        if i<len(classes)-1 {
            str += ", "
        }
    }

    return str
}
func classesValues (classes []Class) string {
    str := ""
    for i, class := range classes {
        str += fmt.Sprintf("%d", class.Frequency)
        if i<len(classes)-1 {
            str += ", "
        }
    }

    return str
}

func classesIntervals (classes []Class) string {
    str := ``
    for i, class := range classes {
        str += fmt.Sprintf("\"[%6.2f;%6.2f )\"", class.Min, class.Max)
        if i < len(classes)-1 {
            str += ", "
        }
    }

    return str
}

func DataToString(data []float64) string {
    str := ""
    for i, num := range data {
        str += fmt.Sprintf("%f", num)
        if i<(len(data)-1) {
            str += ", "
        }
    }
    return str
}

func Average(numbers []float64) float64 {
    var sum float64
    for _, num := range numbers {
        sum += num
    }

    return sum/float64(len(numbers))
}

func S(numbers []float64) float64 {
    av := Average(numbers)
    N := len(numbers)

    var s float64
    for _, num := range numbers {
        s += math.Pow(num, 2) - math.Pow(av, 2)
    }

    s = math.Sqrt((1/float64(N-1))*s)
    return s
}

func E(numbers []float64) float64 {
    av := Average(numbers)
    N := len(numbers)

    var e float64
    for _, num := range numbers {
        e += math.Pow(num-av, 4)
    }

    s4 := math.Pow(S(numbers), 4)
    e = e/(float64(N)*s4)

    return e
}

func contrE(numbers []float64) float64 {
    E := E(numbers)
    if E < 0 {
        E = -1*E
    }

    return 1/math.Sqrt(E)
}

func W(numbers []float64) float64 {
    return S(numbers)/Average(numbers)
}

func A(numbers []float64) float64 {
    var a float64
    av := Average(numbers)
    N := float64(len(numbers))
    S := S(numbers)
    for _, num := range numbers {
        a += math.Pow(num-av, 3)
    }

    a = a/(float64(len(numbers))*math.Pow(S, 3))
    return (math.Sqrt(N*(N-1))/(N-2)) * a
}

func Mediana(numbers []float64) float64 {
    N := len(numbers)

    if N%2==0 {
        return 0.5*float64(numbers[N/2]+numbers[N/2+1])
    } else {
        return numbers[N/2]
    }
}

func Module(num float64) float64 {
    if num < 0 {
        return -1*num
    } else {
        return num
    }
}

func IntervalMin(num float64, S_num float64) float64 {
    return num-1.67*S_num
}

func IntervalMax(num float64, S_num float64) float64 {
    return num+1.67*S_num
}