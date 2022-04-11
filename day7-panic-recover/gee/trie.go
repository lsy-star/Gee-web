package gee

import (
	"fmt"
	"strings"
)

type node struct {
	pattern  string  //待匹配路由，例如 /p/:lang/abc
	part     string  //路由中的一部分	/:lang
	children []*node //子节点
	isWild   bool    //是否精确匹配，part部分为":"、"*"时为true
}

func (n *node) String() string {
	return fmt.Sprintf("node{pattern=%s, part=%s, isWild=%t}", n.pattern, n.part, n.isWild)
}

func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

func (n *node) matchChildren(part string) (res []*node) {
	for _, child := range n.children {
		if child.isWild || child.part == part {
			res = append(res, child)
		}
	}
	return
}

func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{
			part:   part,
			isWild: part[0] == ':' || part[0] == '*',
		}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)
	for _, child := range children {
		if result := child.search(parts, height+1); result != nil {
			return result
		}
	}
	return nil
}
