package transfer

import (
	"encoding/json"
	"fmt"
)

type NodeItem struct {
	Id     uint64
	Label  string
	Pid    uint64
	Status bool
}

type NodeList []NodeItem

func (l NodeList) ToMap() map[uint64]NodeItem {
	m := make(map[uint64]NodeItem)
	for _, v := range l {
		m[v.Id] = v
	}
	return m
}

type NodeTree struct {
	NodeItem
	Children []*NodeTree
}

// type NodeTreeMap map[uint64]*NodeTree

func BuildTree(nodes []NodeItem) []*NodeTree {
	treeList := make([]*NodeTree, 0)
	treeMap := make(map[uint64]*NodeTree)
	for _, node := range nodes {
		treeItem := NodeTree{
			NodeItem: NodeItem{
				Id:    node.Id,
				Pid:   node.Pid,
				Label: node.Label,
			},
			Children: make([]*NodeTree, 0),
		}
		treeMap[node.Id] = &treeItem
		// 根节点收集
		if node.Pid == 0 {
			treeList = append(treeList, &treeItem)
		}
	}
	for _, node := range treeMap {
		if node.Pid != 0 {
			if parent, ok := treeMap[node.Pid]; ok {
				parent.Children = append(parent.Children, node)
			}
		}
	}
	jsonRes, _ := json.Marshal(treeList)
	fmt.Println(string(jsonRes))
	return treeList
}

func BuildValidTree(nodes []NodeItem) []*NodeTree {
	nodeMap := make(map[uint64]*NodeTree)
	var roots []*NodeTree

	// 创建所有节点
	for i := range nodes {
		nodeMap[nodes[i].Id] = &NodeTree{
			NodeItem: nodes[i],
			Children: make([]*NodeTree, 0),
		}
	}

	// 构建树
	for i := range nodes {
		node := nodeMap[nodes[i].Id]
		if nodes[i].Pid == 0 {
			roots = append(roots, node)
		} else {
			parentNode := nodeMap[nodes[i].Pid]
			parentNode.Children = append(parentNode.Children, node)
		}
	}

	// 删除状态为false的叶子节点
	var prune func(node *NodeTree) bool
	prune = func(node *NodeTree) bool {
		for i := 0; i < len(node.Children); i++ {
			if prune(node.Children[i]) {
				node.Children = append(node.Children[:i], node.Children[i+1:]...)
				i--
			}
		}
		return len(node.Children) == 0 && !node.Status
	}
	for i := 0; i < len(roots); i++ {
		if prune(roots[i]) {
			roots = append(roots[:i], roots[i+1:]...)
			i--
		}
	}
	jsonRes, _ := json.Marshal(roots)
	fmt.Println(string(jsonRes))
	return roots
}
