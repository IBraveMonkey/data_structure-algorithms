# 🔙 Поиск с возвратом (Backtracking)

**Описание**: 
Поиск с возвратом (Backtracking) — это алгоритмическая стратегия, основанная на методе "проб и ошибок". Она используется для решения задач, где нужно перебрать множество вариантов, чтобы найти один или все подходящие под условия.

- **Как это устроено внутри**: Алгоритм строит решение по шагам. Если на каком-то этапе становится ясно, что текущий путь ведет в тупик (не удовлетворяет правилам), мы делаем "шаг назад" (unde) в предыдущее состояние и пробуем другой вариант. Это реализуется через рекурсию: мы погружаемся вглубь дерева решений и всплываем обратно при неудаче.
- **Аналогия**: Представьте, что вы находитесь в лабиринте. Вы доходите до развилки и выбираете левый путь. Если вы упираетесь в стену, вы возвращаетесь к последней развилке и пробуете правый путь. Если и там тупик — возвращаетесь еще дальше. Вы продолжаете так, пока не найдете выход.


### Преимущества и недостатки
✅ **Плюсы**:
1. **Универсальность**: Позволяет решать сложнейшие комбинаторные задачи (Судоку, расстановка ферзей, генерация всех подмножеств).
2. **Экономия времени**: Благодаря "отсечению ветвей" (pruning), алгоритм не проверяет варианты, которые заведомо не приведут к успеху, что быстрее полного перебора.

❌ **Минусы**:
1. **Огромная сложность**: В худшем случае время работы растет экспоненциально (O(2^n) или O(n!)). На больших данных алгоритм будет работать вечно.
2. **Память**: Из-за глубокой рекурсии может возникнуть переполнение стека (Stack Overflow).

---


### Сложность

| Метрика | Сложность (O) |
|:---|:---:|
| Время | O(b^d) / O(N!) |
| Память | O(d) |

\*b — коэффициент ветвления (branching factor), d — глубина дерева решений.
\*\*Сложность сильно зависит от конкретной задачи и эффективности отсечения ветвей.


## 🚀 Примеры


### Задача 1: Генерация всех перестановок
Нужно найти все возможные способы расставить $N$ чисел.

```go
func Permutations(nums []int) [][]int {
    var result [][]int
    backtrack(nums, 0, &result)
    return result
}

func backtrack(nums []int, start int, result *[][]int) {
    if start == len(nums) {
        temp := make([]int, len(nums))
        copy(temp, nums)
        *result = append(*result, temp)
        return
    }

    for i := start; i < len(nums); i++ {
        nums[start], nums[i] = nums[i], nums[start]
        backtrack(nums, start+1, result)
        nums[start], nums[i] = nums[i], nums[start]
    }
}
```
```javascript
function permutations(nums) {
  const result = [];
  
  const backtrack = (start) => {
    if (start === nums.length) {
      result.push([...nums]);
      return;
    }
    
    for (let i = start; i < nums.length; i++) {
      [nums[start], nums[i]] = [nums[i], nums[start]];
      backtrack(start + 1);
      [nums[start], nums[i]] = [nums[i], nums[start]];
    }
  };

  backtrack(0);
  return result;
}
```


### Задача 2: N Ферзей
Классическая задача: расставить $N$ ферзей на шахматной доске $N \times N$ так, чтобы ни один не бил другого.

```go
func SolveNQueens(n int) [][]string {
    var result [][]string
    board := make([][]string, n)
    for i := range board {
        board[i] = make([]string, n)
        for j := range board[i] { board[i][j] = "." }
    }
    backtrack(board, 0, n, &result)
    return result
}

func backtrack(board [][]string, row, n int, result *[][]string) {
    if row == n {
        *result = append(*result, construct(board))
        return
    }
    for col := 0; col < n; col++ {
        if isValid(board, row, col, n) {
            board[row][col] = "Q"
            backtrack(board, row+1, n, result)
            board[row][col] = "."
        }
    }
}

func isValid(board [][]string, row, col, n int) bool {
    for i := 0; i < row; i++ {
        if board[i][col] == "Q" { return false }
    }
    for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
        if board[i][j] == "Q" { return false }
    }
    for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
        if board[i][j] == "Q" { return false }
    }
    return true
}

func construct(board [][]string) []string {
    var res []string
    for _, row := range board {
        line := ""
        for _, char := range row { line += char }
        res = append(res, line)
    }
    return res
}
```
```javascript
function solveNQueens(n) {
  const result = [];
  const board = Array.from({ length: n }, () => Array(n).fill('.'));

  const isValid = (row, col) => {
    for (let i = 0; i < row; i++) {
      if (board[i][col] === 'Q') return false;
    }
    for (let i = row - 1, j = col - 1; i >= 0 && j >= 0; i--, j--) {
      if (board[i][j] === 'Q') return false;
    }
    for (let i = row - 1, j = col + 1; i >= 0 && j < n; i--, j++) {
      if (board[i][j] === 'Q') return false;
    }
    return true;
  };

  const backtrack = (row) => {
    if (row === n) {
      result.push(board.map(r => r.join('')));
      return;
    }
    for (let col = 0; col < n; col++) {
      if (isValid(row, col)) {
        board[row][col] = 'Q';
        backtrack(row + 1);
        board[row][col] = '.';
      }
    }
  };

  backtrack(0);
  return result;
}
```

<!-- QUIZ_START 
[
    {
        "question": "По какому принципу работает алгоритм поиска с возвратом (Backtracking)?",
        "options": ["Жадный выбор лучшего варианта", "Метод проб и ошибок: построение решения по шагам и отмена шага при попадании в тупик", "Случайный перебор всех элементов", "Сортировка данных перед поиском"],
        "correctIndex": 1
    },
    {
        "question": "Что такое 'отсечение ветвей' (pruning) в контексте Backtracking?",
        "options": ["Удаление части кода", "Пропуск вариантов, которые заведомо не приведут к решению, для экономии времени", "Остановка программы при ошибке", "Использование только половины массива"],
        "correctIndex": 1
    },
    {
        "question": "Какова типичная временная сложность задач, решаемых с помощью Backtracking (например, задача о ферзях или перестановках)?",
        "options": ["Линейная O(n)", "Логарифмическая O(log n)", "Экспоненциальная (O(2^n)) или факториальная (O(n!))", "Константная O(1)"],
        "correctIndex": 2
    }
]
QUIZ_END -->

