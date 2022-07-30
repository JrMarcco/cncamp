package framework

import (
	"errors"
	"strings"
)

type trieTree struct {
	root *trieTreeNode
}

type trieTreeNode struct {
	parent     *trieTreeNode
	isLeafNode bool
	segment    string
	handlers   HandlerChain
	children   []*trieTreeNode
}

func newNode() *trieTreeNode {
	return &trieTreeNode{
		isLeafNode: false,
		children:   []*trieTreeNode{},
	}
}

func NewTrieTree() *trieTree {
	return &trieTree{root: newNode()}
}

func (tree *trieTree) AddRouter(url string, handlers HandlerChain) error {
	url = strings.TrimLeft(url, "/")

	// 判断是否路径已经注册
	rootNode := tree.root
	if rootNode.findMatchNode(url) != nil {
		return errors.New("router exist: " + url)
	}

	segments := strings.Split(url, "/")
	for index, segment := range segments {
		var objNode *trieTreeNode
		isLast := index == len(segments)-1

		children := rootNode.filterChildren(segment)
		if len(children) > 0 {
			// 按照分段查找子节点
			for _, child := range children {
				if child.segment == segment {
					objNode = child
					break
				}
			}
		}

		if objNode == nil {
			// 子节点不存在则构建新的子节点
			child := newNode()
			child.segment = segment
			if isLast {
				// 当前为叶子
				child.isLeafNode = true
				child.handlers = handlers
			}
			child.parent = rootNode
			rootNode.children = append(rootNode.children, child)
			objNode = child
		}
		// 以子节点为根继续构建子树
		rootNode = objNode
	}
	return nil
}

// filterChildren 过滤符合条件的备选节点
func (node *trieTreeNode) filterChildren(segment string) []*trieTreeNode {
	if len(node.children) == 0 {
		return nil
	}

	// 当前节点有通配符则其子节点则为备选节点
	if isWildSegment(segment) {
		return node.children
	}

	nodes := make([]*trieTreeNode, 0, len(node.children))
	for _, child := range node.children {
		// 节点有通配符节点或节点全文匹配
		if isWildSegment(child.segment) || child.segment == segment {
			nodes = append(nodes, child)
		}
	}

	return nodes
}

// findMatchNode 查找匹配节点
func (node *trieTreeNode) findMatchNode(url string) *trieTreeNode {
	url = strings.TrimLeft(url, "/")

	// 按 "/" 分割为两段
	// 分两段是为了分级依次递归查找匹配的 node
	segments := strings.SplitN(url, "/", 2)

	// 查找第一层匹配节点
	children := node.filterChildren(segments[0])
	if children == nil || len(children) == 0 {
		return nil
	}

	// 当前路径剩下最后一段
	if len(segments) == 1 {
		for _, child := range children {
			// 匹配的节点必须为叶子节点
			if child.isLeafNode {
				return child
			}
		}
		return nil
	}

	for _, child := range children {
		// 将剩余段交给子节点递归查找
		matchNode := child.findMatchNode(segments[1])
		if matchNode != nil {
			return matchNode
		}
	}
	return nil
}

func (node *trieTreeNode) parseParamsFromEndNode(url string) map[string]string {
	url = strings.TrimLeft(url, "/")
	params := map[string]string{}

	segments := strings.Split(url, "/")
	if len(segments) > 0 {
		currentNode := node
		for i := len(segments); i >= 0; i-- {
			if currentNode.segment == "" {
				break
			}

			if isWildSegment(currentNode.segment) {
				params[currentNode.segment[1:]] = segments[i]
			}
			currentNode = currentNode.parent
		}
	}
	return params
}

// isWildSegment 判断是否为通配 segment
func isWildSegment(segment string) bool {
	return strings.HasPrefix(segment, ":")
}
