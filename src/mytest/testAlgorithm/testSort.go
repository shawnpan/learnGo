package testAlgorithm

import (
	"fmt"
	"math/rand"
	_ "sort"
	"time"
)

func TestSort() {
	size := 10
	arr := mkArr(size)
	start := time.Now().Unix()
	testBubbleSort1(arr)
	fmt.Println("time use : ", (time.Now().Unix() - start), "s")
	// checkArr(arr)

	arr = mkArr(size)
	start = time.Now().Unix()
	testBubbleSort2(arr)
	fmt.Println("time use : ", (time.Now().Unix() - start), "s")
	// checkArr(arr)

	// size = 10
	arr = mkArr(size)
	start = time.Now().Unix()
	testQuickSort(arr)
	fmt.Println("time use : ", (time.Now().Unix() - start), "s")
	// checkArr(arr)

	// size = 100
	arr = mkArr(size)
	start = time.Now().Unix()
	testInsertSort(arr)
	fmt.Println("time use : ", (time.Now().Unix() - start), "s")
	// checkArr(arr)

	size = 10000
	arr = mkArr(size)
	start = time.Now().Unix()
	testSleepSort(arr)
	fmt.Println("time use : ", (time.Now().Unix() - start), "s")
	checkArr(arr)
}
func mkArr(size int) []int {
	arr := make([]int, size, size)
	for i, _ := range arr {
		arr[i] = rand.Intn(1000)
	}
	return arr
}
func checkArr(arr []int) bool {
	res := true
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			res = false

		}
	}
	fmt.Println("check arry : ", res)
	return res
}

func testBubbleSort1(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
func testBubbleSort2(arr []int) {
	length := len(arr)
	for i := 0; i < length; i++ {
		if length%2 == 0 && 2*i == length {
			break
		}
		if length%2 == 1 && 2*i > length {
			break
		}
		max, min := arr[i], arr[i]
		maxIndex, minIndex := i, i
		for j := i; j < length-i; j++ {
			if arr[j] > max {
				max = arr[j]
				maxIndex = j
			}
			if arr[j] < min {
				min = arr[j]
				minIndex = j
			}
		}
		arr[minIndex] = arr[i]
		arr[i] = min
		if i == maxIndex {
			maxIndex = minIndex
		}
		arr[maxIndex] = arr[length-i-1]
		arr[length-i-1] = max
	}
}

func testQuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

// 一趟快速排序，递归算法
func quickSort(arr []int, low int, high int) {
	// low、high指定序列的下界和上界
	if low >= high { // 序列有效
		return
	}
	i, j := low, high
	vot := arr[i] // 第一个值作为基准值
	for i != j {  // 一趟排序
		for i < j && vot <= arr[j] { // 从后向前寻找较小值
			j--
		}
		if i < j {
			arr[i] = arr[j] // 较小元素向前移动
			i++
		}
		for i < j && arr[i] < vot { // 从前向后寻找较大值
			i++
		}
		if i < j { // 较大元素向后移动
			arr[j] = arr[i]
			j--
		}
	}
	arr[i] = vot              // 基准值的最终位置
	quickSort(arr, low, j-1)  // 前端子序列再排序
	quickSort(arr, i+1, high) // 后端子序列再排序
}

// 插入排序
func testInsertSort(arr []int) {
	length := len(arr)
	var temp int
	for i := 1; i < length; i++ {
		for j := i - 1; j >= 0; j-- {
			if arr[j+1] < arr[j] {
				temp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			} else {
				break
			}
		}
	}
}

// 睡眠排序
//睡眠排序算法是一个天才程序员发明的，想法很简单，就是针对数组里的不同的数开多个线程，
//每个线程根据数的大小睡眠，自然睡的时间越长的，数越大，哈哈，搞笑吧，这种算法看起来很荒唐，
//但实际上很天才，它可以充分利用多核cpu进行计算
func testSleepSort(arr []int) {

	ch := make(chan int)
	for _, value := range arr {
		go func(val int) {
			for i := 0; i < val; i++ {
				time.Sleep(10000 * time.Nanosecond) // sleep方法只能接受time.Duration，但不能乘以val，晕
			}
			ch <- val
		}(value)
	}
	res := []int{}
	for i := 0; i < len(arr); i++ {
		val := <-ch
		res = append(res, val)
	}
	for i, val := range res {
		arr[i] = val
	}
}
