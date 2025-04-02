package main

type Element struct {
	Timestamp int
	Value     string
}

type TimeMap struct {
	Elements map[string][]Element
}

func binarySearch(list []Element, target int) int {
	left, right := 0, len(list)

	for left < right {
		mid := left + (right-left)/2
		if target > list[mid].Timestamp {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

func Constructor() TimeMap {
	return TimeMap{
		Elements: make(map[string][]Element, 0),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if _, exists := this.Elements[key]; !exists {
		this.Elements[key] = make([]Element, 0)
	}

	this.Elements[key] = append(this.Elements[key], Element{
		Value:     value,
		Timestamp: timestamp,
	})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if _, exists := this.Elements[key]; !exists {
		return ""
	}

	idx := binarySearch(this.Elements[key], timestamp)
	idx -= 1
	if idx == -1 {
		return ""
	}

	return this.Elements[key][idx].Value
}
