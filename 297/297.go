// UVa 297 - Quadtrees

package main

import (
	"fmt"
	"os"
	"strings"
)

const total_weight = 1024

func preOrder(tree string, weight int, image *string) string {
	if weight == total_weight {
		if tree[0] == 'p' {
			return preOrder(tree[1:], weight/4, image)
		}
		if len(tree) == 1 {
			*image += strings.Repeat(tree, weight)
			return ""
		}
	}
	for i := 0; i < 4; i++ {
		if tree[0] == 'p' {
			tree = preOrder(tree[1:], weight/4, image)
		} else {
			*image += strings.Repeat(string(tree[0]), weight)
			tree = tree[1:]
		}
	}
	return tree
}

func add(tree1, tree2 string) int {
	var image1, image2 string
	preOrder(tree1, total_weight, &image1)
	preOrder(tree2, total_weight, &image2)

	total := 0
	for i := range image1 {
		if image1[i] == 'f' || image2[i] == 'f' {
			total++
		}
	}
	return total
}

func main() {
	in, _ := os.Open("297.in")
	defer in.Close()
	out, _ := os.Create("297.out")
	defer out.Close()

	var kase int
	var tree1, tree2 string
	fmt.Fscanf(in, "%d", &kase)
	for kase > 0 {
		fmt.Fscanf(in, "%s", &tree1)
		fmt.Fscanf(in, "%s", &tree2)
		fmt.Fprintf(out, "There are %d black pixels.\n", add(tree1, tree2))
		kase--
	}
}
