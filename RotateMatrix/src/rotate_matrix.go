/*
@author : Anchal Verma
*/

package main

import "fmt"

// Function to rotate matrix
func rotate_matrix(matrix [][]int, size int) [][]int {
	// choosing start and the end index
	start, end := 0, size - 1
	// Traverse through the matrix and swap/rotate the start, end rows until the middle row is reached
	for start < end {
		// Storing the "start" row in temp variable
		temp := matrix[start]
		// Swapping start and end row of the matrix
		matrix[start] = matrix[end]
		matrix[end] = temp
		// Move to the next set of (start, end) rows of matrix
		start++
		end--
	}

	// Swap/rotate the column of the matrix
	for i:= 0; i < len(matrix); i++ {
		// for each row, swap the i(row) element with j(column) element 
		// except the ith column(i.e. starting from i+1 to the size of column)
		for j := i+1; j < len(matrix[i]); j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = tmp
		}
	}
	return matrix
}

// main function
func main(){

	// Taking matrix input from user
	fmt.Println("Enter the matrix size : ")
	size := 0
	fmt.Scanln(&size)

	fmt.Println("Enter the matrix : ")
	matrix := make([][]int, 0, size)
	for i:= 0; i<size; i++ {
		row := make([]int, size)
		for j, _ := range row {
			fmt.Scan(&row[j])
		}
		matrix = append(matrix, row)
	}
	fmt.Println("The given matrix is : ")
	fmt.Println(matrix)

	//  Calling function to rotate the matrix
	matrix = rotate_matrix(matrix, size)
	fmt.Println("The matrix after rotation is : ")
	fmt.Println(matrix)
}
