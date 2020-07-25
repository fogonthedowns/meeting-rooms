package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	items := [][]int{
		{0, 30},
		{5, 10},
		{15, 20},
	}
	ans := minMeetingRooms(items)
	fmt.Println(ans)
}

func minMeetingRooms(items [][]int) int {
	if len(items) == 0 {
		return 0
	}

	// sort the meeting by start time
	// so the algorithm works
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] < items[j][0]
	})

	// make the PriorityQueue
	pq := make(PriorityQueue, 1)

	pq[0] = &Meeting{
		start: items[0][0],
		end:   items[0][1],
		index: 0,
	}
	heap.Init(&pq)

	for i := 1; i < len(items); i++ {
		start := items[i][0]
		end := items[i][1]
		m := &Meeting{
			start: start,
			end:   end,
			index: i,
		}
		top := pq[0]
		// room is empty at this time.
		// we need to keep the room count the same.
		// by keeping the heap the same size.
		// pop out the last meeting.
		// insert the new meeting.
		// this will resize the heap.
		if start >= top.end {
			heap.Pop(&pq)
			heap.Push(&pq, m)
		} else {
			// room is totally full
			// we need to add a room by expanding the heap
			heap.Push(&pq, m)
		}
	}

	return pq.Len()
}

// An Item is something we manage in a priority queue.
type Meeting struct {
	start int
	end   int // The priority of the item
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Meeting

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].end < pq[j].end
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Meeting))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(m *Meeting, start int, end int) {
	m.start = start
	m.end = end
	heap.Fix(pq, m.index)
}
