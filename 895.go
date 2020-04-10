package main

import "container/list"

type FreqStack struct {
	freq    map[int]int
	group   map[int]*list.List
	maxFreq int
}

func Constructor() FreqStack {
	return FreqStack{
		freq:  make(map[int]int),
		group: make(map[int]*list.List),
	}
}

func (this *FreqStack) Push(x int) {
	this.freq[x] = this.freq[x] + 1
	if this.freq[x] > this.maxFreq {
		this.maxFreq = this.freq[x]
	}
	cur := this.freq[x]
	if lis, exist := this.group[cur]; exist {
		lis.PushBack(x)
	} else {
		lis = list.New()
		lis.PushBack(x)
		this.group[cur] = lis
	}
}

func (this *FreqStack) Pop() int {
	lis := this.group[this.maxFreq]
	v := lis.Remove(lis.Back())
	if lis.Len() == 0 {
		this.maxFreq--
	}
	result := v.(int)
	this.freq[result]--
	return result
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 */
