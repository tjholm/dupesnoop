package analyzer

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/shirou/gopsutil/v4/disk"
	"github.com/spf13/afero"
)

type Analyzer struct {
}

type DiskInfo struct {
	Device string `json:"device"`
	Mount  string `json:"mount"`
	FsType string `json:"fs"`
	Used   uint64 `json:"used"`
	Total  uint64 `json:"total"`
}

// Node is a struct for the disk usage sunburst chart.
type Node struct {
	Name     string  `json:"name"`
	Color    string  `json:"color,omitempty"`
	Loc      int64   `json:"loc,omitempty"`
	Children []*Node `json:"children,omitempty"`
}

func (n *Node) FindNode(dir string) *Node {
	if n.Name == dir {
		return n
	}

	for _, child := range n.Children {
		if parent := child.FindNode(dir); parent != nil {
			return parent
		}
	}

	return nil
}

func (n *Node) FindNearestNode(dir string) *Node {
	for _, child := range n.Children {
		if parent := child.FindNearestNode(dir); parent != nil {
			return parent
		}
	}

	if strings.HasPrefix(dir, n.Name) {
		return n
	}

	return nil
}

func (a *Analyzer) ListDisks() ([]DiskInfo, error) {
	partitions, err := disk.Partitions(false)
	if err != nil {
		return nil, err
	}

	disks := make([]DiskInfo, 0, len(partitions))
	for _, partition := range partitions {
		usageStat, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			return nil, err
		}

		disks = append(disks, DiskInfo{
			Device: partition.Device,
			Mount:  partition.Mountpoint,
			FsType: partition.Fstype,
			Used:   usageStat.Used,
			Total:  usageStat.Total,
		})
	}

	return disks, nil
}

func analyzeUsage(fs afero.Fs, dir string, depth int) (*Node, error) {
	// Create the root node
	var tree *Node = nil

	err := afero.Walk(fs, dir, func(path string, info os.FileInfo, err error) error {
		// Add the root node
		fmt.Println(path)
		if tree == nil {
			tree = &Node{
				Name: path,
				// Color:    "hsl(240, 70%, 50%)",
				Children: []*Node{},
			}
		} else {
			relativePath, err := filepath.Rel(tree.Name, path)
			if err != nil {
				return err
			}

			fmt.Println("relative path depth", len(strings.Split(relativePath, "/")))

			if !(len(strings.Split(relativePath, "/")) > depth) {
				// Skip files that are too deep
				parentNode := tree.FindNode(filepath.Dir(path))
				if parentNode == nil {
					fmt.Println("parent node not found for", path)
					return fmt.Errorf("parent node not found for %s", path)
				}

				parentNode.Children = append(parentNode.Children, &Node{
					Name:     path,
					Loc:      info.Size(),
					Children: []*Node{},
				})
			} else {
				// add to nearest parent
				parentNode := tree.FindNearestNode(filepath.Dir(path))
				if parentNode == nil {
					fmt.Println("parent node not found for", path)
					parentNode = tree
				}

				fmt.Println("adding to nearest parent", parentNode.Name)
				parentNode.Loc = parentNode.Loc + info.Size()
			}
		}

		return nil
	})
	if err != nil {
		fmt.Println("returning error", err)
		return nil, err
	}

	json, _ := json.Marshal(tree)

	fmt.Printf("returning tree %s\n", string(json))
	return tree, nil
}

func (a *Analyzer) GetDiskUsage(path string) (*Node, error) {
	// Walk the file system from the given path and build the sunburst chart.
	return analyzeUsage(afero.NewOsFs(), path, 3)
	// root := &Node{
	// 	Name:  "nivo",
	// 	Color: "hsl(240, 70%, 50%)",
	// 	Children: []*Node{
	// 		{
	// 			Name:  "viz",
	// 			Color: "hsl(117, 70%, 50%)",
	// 			Children: []*Node{
	// 				{
	// 					Name:  "stack",
	// 					Color: "hsl(279, 70%, 50%)",
	// 					Children: []*Node{
	// 						{Name: "cchart", Color: "hsl(280, 70%, 50%)", Loc: 80314},
	// 						{Name: "xAxis", Color: "hsl(315, 70%, 50%)", Loc: 111007},
	// 						{Name: "yAxis", Color: "hsl(51, 70%, 50%)", Loc: 9214},
	// 						{Name: "layers", Color: "hsl(197, 70%, 50%)", Loc: 37677},
	// 					},
	// 				},
	// 				{
	// 					Name:  "ppie",
	// 					Color: "hsl(146, 70%, 50%)",
	// 					Children: []*Node{
	// 						{
	// 							Name:  "chart",
	// 							Color: "hsl(239, 70%, 50%)",
	// 							Children: []*Node{
	// 								{Name: "pie", Color: "hsl(213, 70%, 50%)", Children: []*Node{
	// 									{Name: "outline", Color: "hsl(189, 70%, 50%)", Loc: 152961},
	// 									{Name: "slices", Color: "hsl(56, 70%, 50%)", Loc: 38624},
	// 									{Name: "bbox", Color: "hsl(171, 70%, 50%)", Loc: 100170},
	// 								}},
	// 								{Name: "donut", Color: "hsl(174, 70%, 50%)", Loc: 74520},
	// 								{Name: "gauge", Color: "hsl(148, 70%, 50%)", Loc: 103242},
	// 							},
	// 						},
	// 						{Name: "legends", Color: "hsl(301, 70%, 50%)", Loc: 132401},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		{
	// 			Name:  "colors",
	// 			Color: "hsl(181, 70%, 50%)",
	// 			Children: []*Node{
	// 				{Name: "rgb", Color: "hsl(106, 70%, 50%)", Loc: 172323},
	// 				{Name: "hsl", Color: "hsl(109, 70%, 50%)", Loc: 96178},
	// 			},
	// 		},
	// 		{
	// 			Name:  "utils",
	// 			Color: "hsl(174, 70%, 50%)",
	// 			Children: []*Node{
	// 				{Name: "randomize", Color: "hsl(155, 70%, 50%)", Loc: 65701},
	// 				{Name: "resetClock", Color: "hsl(267, 70%, 50%)", Loc: 126719},
	// 				{Name: "noop", Color: "hsl(87, 70%, 50%)", Loc: 69300},
	// 				{Name: "tick", Color: "hsl(285, 70%, 50%)", Loc: 83076},
	// 				{Name: "forceGC", Color: "hsl(318, 70%, 50%)", Loc: 192733},
	// 				{Name: "stackTrace", Color: "hsl(226, 70%, 50%)", Loc: 139715},
	// 				{Name: "dbg", Color: "hsl(314, 70%, 50%)", Loc: 84768},
	// 			},
	// 		},
	// 	},
	// }

	// return root, nil
}

func NewAnalyzer() *Analyzer {
	return &Analyzer{}
}
