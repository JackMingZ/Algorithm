package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var s []int
	for i := 0; i < 1000000; i++ {
		r := rand.Intn(100000)
		//append给m自动扩容
		s = append(s, r)
	}
	t1 := time.Now()
	QuickSort(0, len(s)-1, s)
	elapsed := time.Since(t1)
	fmt.Println("正常时间=", elapsed)
	t2 := time.Now()
	var d []int
	for i := 0; i < 1000000; i++ {
		r := rand.Intn(100000)
		//append给m自动扩容
		d = append(d, r)
	}
	elapsed = time.Since(t2)
	fmt.Print("goroutine时间=", elapsed)
}

func QuickSort(left int, right int, arr []int) {
	//left始终指向数组最左边，不会动。right始终指向数组最右边，不会动
	//l一开始指向数组最左边，后面会逐渐右移。r一开始指向数组最右边，后面会逐渐左移
	l := left
	r := right

	middle := arr[(left+right)/2]

	//for循环的目标是将比middle小的数放到左边;比middle大的数放到右边
	for l < r {
		//从middle的左边找到大于等于middle的值
		for arr[l] < middle {
			l++
		}
		//从middle的右边边找到小于等于middle的值
		for arr[r] > middle {
			r--
		}
		// 1 >= r 表明本次分解任务完成, break
		if l >= r {
			break
		}
		//将middle左边大于middle的值 与 middle右边小于middle的值 交换
		arr[l], arr[r] = arr[r], arr[l]
		//优化
		if arr[l] == middle {
			r--
		}
		if arr[r] == middle {
			l++
		}
	}
	// 如果  1== r, 再移动下,为了跳出循环，
	//同时也是为了下面递归时，找到左边集合数组中最右边的数r，和找到右边集合数组中最左边的数l
	if l == r {
		l++
		r--
	}
	// 向左递归
	if left < r {
		QuickSort(left, r, arr)
	}
	// 向右递归
	if right > l {
		QuickSort(l, right, arr)
	}

}

func QuickSortOnGoroutine(data []int) []int {

	if len(data) <= 1 {
		return data
	}

	var wg sync.WaitGroup

	c := data[0]

	var s1, s2 []int

	for k, v := range data {
		if k == 0 {
			continue
		}
		if c < v {
			s2 = append(s2, v) //s2中值总是比data[0]大
		} else {
			s1 = append(s1, v) //s1中值总是比data[0]小or等
		}
	}

	wg.Add(2)

	go func() {
		s1 = QuickSortOnGoroutine(s1)
		wg.Done()
	}()

	go func() {
		s2 = QuickSortOnGoroutine(s2)
		wg.Done()
	}()

	wg.Wait()

	data = []int{c}
	if len(s1) > 0 {
		data = append(s1, data...)
	}
	if len(s2) > 0 {
		data = append(data, s2...)
	}
	return data
}
