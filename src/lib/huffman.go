package lib

import (
	"container/heap"
)

var KeyCodeMap = make(map[int]string)

type HuffmanTree interface {
	Freq() int
}

type HuffmanLeaf struct {
	freq int
	value int
}

type HuffmanNode struct {
	freq int
	left, right HuffmanTree
}

func (self HuffmanLeaf) Freq() int {
	return self.freq
}

func (self HuffmanNode) Freq() int {
	return self.freq
}

type treeHeap []HuffmanTree

func (th treeHeap) Len() int { return len(th) }
func (th treeHeap) Less(i, j int) bool {
	return th[i].Freq() < th[j].Freq()
}
func (th *treeHeap) Push(ele interface{}) {
	*th = append(*th, ele.(HuffmanTree))
}
func (th *treeHeap) Pop() (popped interface{}) {
	popped = (*th)[len(*th)-1]
	*th = (*th)[:len(*th)-1]
	return
}
func (th treeHeap) Swap(i, j int) { th[i], th[j] = th[j], th[i] }

func buildTree(symFreqs map[int]int) HuffmanTree {
	var trees treeHeap
	for c, f := range symFreqs {
		trees = append(trees, HuffmanLeaf{f, c})
	}
	heap.Init(&trees)
	for trees.Len() > 1 {
		a := heap.Pop(&trees).(HuffmanTree)
		b := heap.Pop(&trees).(HuffmanTree)
		heap.Push(&trees, HuffmanNode{a.Freq() + b.Freq(), a, b})
	}
	return heap.Pop(&trees).(HuffmanTree)
}

func setMap(tree HuffmanTree, prefix []byte) {
	switch i := tree.(type) {
	case HuffmanLeaf:
		KeyCodeMap[i.value]=string(prefix)
	case HuffmanNode:
		prefix = append(prefix, '0')
		setMap(i.left, prefix)
		prefix = prefix[:len(prefix)-1]
		prefix = append(prefix, '1')
		setMap(i.right, prefix)
		prefix = prefix[:len(prefix)-1]
	}
}

func GetHuffmanMap(array [] int) map[int]string {
	symFreqs := make(map[int]int)
	for _, c := range array {
		symFreqs[c]++
	}
	exampleTree := buildTree(symFreqs)
	setMap(exampleTree, []byte{})
	return KeyCodeMap
}
