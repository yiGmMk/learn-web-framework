package gee

type node struct {
	pattern  string
	part     string
	children []*node
	isWild   bool
}

// 插入一条路径
func (n *node) insert(pattern string, parts []string, height int) {

}

// 根据pattern查找节点
func (n *node) search(parts []string, height int) *node {

	return nil
}
