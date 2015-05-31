package main

import (
	"fmt"
	"math"
    "sort"
)

type Class struct {
	Min, Max  float64
	Frequency int
	FuncNum   int
}

func Classes(data []float64) []Class {
    M := global_M
    h := h(data)
    data_i := 0
    frequency := Frequency(data)

    min := data[0]

    classes := make([]Class, M)

    for class_i, _ := range classes {
        classes[class_i].Min = min + float64(class_i)*h
        classes[class_i].Max = min + (float64(class_i)+1)*h

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

    return classes
}

func classesFullValues(classes []Class) string {
	str := ""
	var counter int
	for i, class := range classes {
		counter += class.Frequency
		str += fmt.Sprintf("%d", counter)
		if i < len(classes)-1 {
			str += ", "
		}
	}

	return str
}
func classesValues(classes []Class) string {
	str := ""
	for i, class := range classes {
		str += fmt.Sprintf("%d", class.Frequency)
		if i < len(classes)-1 {
			str += ", "
		}
	}

	return str
}

func classesIntervals(classes []Class) string {
	str := ``
	for i, class := range classes {
		str += fmt.Sprintf("\"[%6.2f;%6.2f )\"", class.Min, class.Max)
		if i < len(classes)-1 {
			str += ", "
		}
	}

	return str
}

func IntDataToString(data []int) string {
    str := ""
    for i, num := range data {
        str += fmt.Sprintf("%d", num)
        if i < (len(data) - 1) {
            str += ", "
        }
    }
    return str
}

func FloatDataToString(data []float64) string {
	str := ""
	for i, num := range data {
		str += fmt.Sprintf("%f", num)
		if i < (len(data) - 1) {
			str += ", "
		}
	}
	return str
}

func DataPToString(data []float64) string {
    str := ""
    freq := Frequency(data)
    for i, num := range data {
        str += fmt.Sprintf("%d  ", freq[num])
        if i < (len(data) - 1) {
            str += ", "
        }
    }
    return str
}

func Frequency (data []float64) map[float64]int {
    frequency := make(map[float64]int, len(data))

    for _, number := range data {
        frequency[number]++
    }

    return frequency
}

// среднее арифметическое
func Average(numbers []float64) float64 {
	var sum float64
	for _, num := range numbers {
		sum += num
	}

	return sum / float64(len(numbers))
}

// среднее арифметическое от квадратов
func Average2(numbers []float64) float64 {
    var sum float64
    for _, num := range numbers {
        sum += math.Pow(num, 2)
    }

    return sum / float64(len(numbers))
}

// Среднее квадратическое отклонение
func S(numbers []float64) float64 {
	av := Average(numbers)
	N := len(numbers)

	var s float64
	for _, num := range numbers {
		s += math.Pow(num, 2) - math.Pow(av, 2)
	}

	s = math.Sqrt((1 / float64(N-1)) * s)
	return s
}

// коефициент Эксцесса
func E(numbers []float64) float64 {
	av := Average(numbers)
	N := len(numbers)

	var e float64
	for _, num := range numbers {
		e += math.Pow(num-av, 4)
	}

	s4 := math.Pow(S(numbers), 4)
	e = e / (float64(N) * s4)

	return e
}

// контрэксцесс
func contrE(numbers []float64) float64 {
	E := E(numbers)
	if E < 0 {
		E = -1 * E
	}

	return 1 / math.Sqrt(E)
}

// вариация Пирсона
func W(numbers []float64) float64 {
	return S(numbers) / Average(numbers)
}

// коефициент ассиметрии
func A(numbers []float64) float64 {
	var a float64
	av := Average(numbers)
	N := float64(len(numbers))
	S := S(numbers)
	for _, num := range numbers {
		a += math.Pow(num-av, 3)
	}

	a = a / (float64(len(numbers)) * math.Pow(S, 3))
	return (math.Sqrt(N*(N-1)) / (N - 2)) * a
}

func h(numbers []float64) float64 {
    sort.Float64s(numbers)

    return (numbers[len(numbers)-1] - numbers[0])/ float64(global_M)
}

// lab4 a
func a(numbers []float64) float64 {
    av := Average(numbers)
    av2 := Average2(numbers)

    return av - math.Sqrt(3*(av2-math.Pow(av, 2)))
}

// lab4 b
func b(numbers []float64) float64 {
    av := Average(numbers)
    av2 := Average2(numbers)

    return av + math.Sqrt(3*(av2-math.Pow(av, 2)))
}

// медиана
func Mediana(numbers []float64) float64 {
	N := len(numbers)

	if N%2 == 0 {
		return 0.5 * float64(numbers[N/2]+numbers[N/2+1])
	} else {
		return numbers[N/2]
	}
}

// модуль от числа
func Module(num float64) float64 {
	if num < 0 {
		return -1 * num
	} else {
		return num
	}
}

func IntervalMin(num float64, S_num float64) float64 {
	return num - 1.67*S_num
}

func IntervalMax(num float64, S_num float64) float64 {
	return num + 1.67*S_num
}

func Kolmagorov(z, N float64) float64 {
    var result float64

    for i:=1; i<=3; i++ {
        k := float64(i)
        f1 := math.Pow(k, 2) - 0.5*(1-math.Pow(-1, k))
        f2 := 5*math.Pow(k,2)+22-7.5*(1-math.Pow(-1, k))


        result +=
        math.Pow(-1, k)*math.Exp(  -2 * math.Pow(k,2) * math.Pow(z,2)  )*
        (1 - (2*math.Pow(k, 2)*z)/3*math.Sqrt(N) - 1/(18*N)*(
        (f1-4*(f1+3))*math.Pow(k,2)*math.Pow(z,2) + 8*math.Pow(k, 4)*math.Pow(z, 4)) + (math.Pow(k, 2)*z)/(27*math.Sqrt(math.Pow(N, 3))) *
        ((math.Pow(f2, 2)/5) - (4*(f2+45)*math.Pow(k, 2)*math.Pow(z, 2))/(15) + 8*math.Pow(k,4)*math.Pow(z,4) )  )
    }

    return 1.0+2*result
}