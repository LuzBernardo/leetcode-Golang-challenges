package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	visited := make([]*Node, 101)
	newNode := &Node{
		Val: node.Val,
	}

	var dfs func(*Node, *Node) *Node
	dfs = func(node *Node, newNode *Node) *Node {
		if visited[node.Val] != nil {
			return visited[node.Val]
		}

		visited[node.Val] = newNode

		for _, neighbor := range node.Neighbors {
			newNeighbor := &Node{
				Val: neighbor.Val,
			}

			newNode.Neighbors = append(newNode.Neighbors, dfs(neighbor, newNeighbor))
		}
		return newNode
	}

	return dfs(node, newNode)
}
