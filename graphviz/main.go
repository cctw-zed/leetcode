package main

import (
	"fmt"
	"log"
	"os"
)

// 主入口
func main() {
    // 创建二叉树
    root := &TreeNode{Value: "A"}
    root.Left = &TreeNode{Value: "B"}
    root.Right = &TreeNode{Value: "C"}
    root.Left.Left = &TreeNode{Value: "D"}
    root.Left.Right = &TreeNode{Value: "E"}
    root.Right.Left = &TreeNode{Value: "F"}

    tree := &BinaryTree{Root: root}

    // 生成 DOT 格式的字符串
    dotCode := tree.GenerateDot()

    // 将 DOT 格式的字符串写入到文件
    err := writeDotToFile("binary_tree.dot", dotCode)
    if err != nil {
        log.Fatalf("Failed to write DOT file: %v", err)
    }
}

// 将生成的 DOT 格式的字符串写入文件
func writeDotToFile(filename, dotCode string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(dotCode)
	if err != nil {
		return err
	}

	fmt.Printf("DOT code has been written to %s\n", filename)
	return nil
}
