package main

import (
	"fmt"
)

func indexAction() string {
	var data_value string
	if len(global_data) > 0 {
		for i, num := range global_data {
			if i < len(global_data)-1 {
				data_value += fmt.Sprintf("%6.4f\n", num)
			} else {
				data_value += fmt.Sprintf("%6.4f", num)
			}

		}
	}

	content := fmt.Sprintf(`
    <h1>Математическая статистика. Ввод данных</h1>
    <form method="POST" action="/analys/">
                    <textarea name="data" placeholder="Введите данные сюда. Каждое значение с новой строки.">%s</textarea> <br />
                    <label>Количество классов: </label><input type="number" value="%d" min=0 name="class_count" /><br /><div>0 - определить автоматически</div> <br />
                    <input class="btn btn-primary" type="submit" value="Обработать!" />
                </form>`, data_value, global_M)

	return content
}
