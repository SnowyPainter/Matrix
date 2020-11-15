# Matrix  
<img src="./matrixgo.png" style="width:300px;height:240px">  

This module implements some functions for matrix.  
Matrix implements functions with Rowing and Looping.

# Download

``` bash
go get -u github.com/snowypainter/Matrix
```

# Documentation

## 1. Declare & Set Matrix
---------------------------
``` Go
mat := matrix.NewMatrix(2, 3) // 2 rows, 3 columns
mat.Reset([][]float64{{1, 2}, {3, 4}, {5, 6}})

/* same above
    m.SetRow(0, []float64{1, 2})
	m.SetRow(1, []float64{3, 4})
	m.SetRow(2, []float64{5, 6})
*/
```

## 2. Calculate with Matrix
--------------------------
``` Go
/*
	WARN : variables are not a car!
*/
m := matrix.NewMatrix(3, 3)
m.Reset([][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
m2 := matrix.NewMatrix(3, 3)
m2.Reset([][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
```

### 1. Plus & Minus
``` Go
m3, err := m.Plus(m2)
fmt.Println("m + m2 = ", m3)

m4, err := m.Minus(m2)
fmt.Println("m - m2 = ", m4)
```

Plus & Minus use ***same core*** function.  
***add*** function which is private for developers and in /matrix/calculate.go.  
So it pass only +, - function internally. like :   
``` Go
func(a float64, b float64) float64 {
	return a + b // or a - b for Minus
}
 ```

### 2. Multiply
``` Go
m5, err := m.Multiply(m2)
fmt.Println("m * m2 = ", m5)
```

It does multiplying calculates as you know.

### 3. Scalar Multiply
``` Go
m6 := m.ScalarMultiply(5)
fmt.Println("|m| * |5| = ", m6)
```

Mutiply scalar value to matrix. no matter what matrix size is it.

### 4. Transpose
``` Go
fmt.Println("m.Transpose() = ", m.Transpose())
```

### 5. Results
Run main.go after cloning
``` Bash
go run main.go ./matrix
```
``` Bash
m + m2 =  [
  [2 4 6]
  [8 10 12]
  [14 16 18]
]
m - m2 =  [
  [0 0 0]
  [0 0 0]
  [0 0 0]
]
m * m2 =  [
  [30 36 42]
  [66 81 96]
  [102 126 150]
]
|m| * |5| =  [
  [5 10 15]
  [20 25 30]
  [35 40 45]
]
m.Transpose() =  [
  [1 4 7]
  [2 5 8]
  [3 6 9]
]
```

# Thank you for reading readme.md