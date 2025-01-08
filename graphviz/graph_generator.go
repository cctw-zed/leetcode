package main

import "github.com/awalterschulze/gographviz"

// 定义一个接口，所有生成图的结构体都要实现该接口
type GraphGenerator interface {
	GenerateDot() string
}

// 有向图生成器结构体
type DirectedGraph struct {
	Nodes []string
	Edges [][2]string
}

// 实现 GraphGenerator 接口的 GenerateDot 方法
func (g *DirectedGraph) GenerateDot() string {
	graph := gographviz.NewGraph()
	graph.SetName("G")
	graph.SetDir(true) // 设置为有向图

	// 添加节点
	for _, node := range g.Nodes {
		graph.AddNode("G", node, nil)
	}

	// 添加边
	for _, edge := range g.Edges {
		graph.AddEdge(edge[0], edge[1], true, nil)
	}

	return graph.String() // 返回生成的 DOT 格式字符串
}

// 二叉树节点结构体
type TreeNode struct {
    Value       string
    Left, Right *TreeNode
}

// 二叉树生成器结构体
type BinaryTree struct {
    Root *TreeNode
}

// 实现 GraphGenerator 接口的 GenerateDot 方法
func (t *BinaryTree) GenerateDot() string {
    graph := gographviz.NewGraph()
    graph.SetName("BinaryTree")
    graph.SetDir(true) // 设置为有向图

    // 从根节点开始递归地添加节点和边
    addTreeNodesAndEdges(graph, t.Root)

    return graph.String() // 返回生成的 DOT 格式字符串
}

// 辅助函数，递归地添加树的节点和边
func addTreeNodesAndEdges(graph *gographviz.Graph, node *TreeNode) {
    if node == nil {
        return
    }

    // 添加当前节点
    graph.AddNode("BinaryTree", node.Value, nil)

    // 如果有左子树，添加左子树节点和边
    if node.Left != nil {
        graph.AddEdge(node.Value, node.Left.Value, true, nil)
        addTreeNodesAndEdges(graph, node.Left)
    }

    // 如果有右子树，添加右子树节点和边
    if node.Right != nil {
        graph.AddEdge(node.Value, node.Right.Value, true, nil)
        addTreeNodesAndEdges(graph, node.Right)
    }
}
