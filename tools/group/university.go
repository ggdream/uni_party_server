package group

import (
	"fmt"
)

type SingleStruct struct {
	Campus  string `json:"campus"`
	College string `json:"college"`
	Grade   uint8  `json:"grade"`
	Major   string `json:"major"`
	Class   string `json:"class"`

	UID uint `json:"uid"`
}

// Head 头结点
type Head struct {
	Rank    int8
	Include *[]*Node
}

// Node 指针节点
type Node struct {
	UID      uint
	Mark     string
	Children *[]*Node

	Campus  string
	College string
	Grade   uint8
	Major   string
	Class   string
}

// NewUniversity 实例化一个大学结构
// data：校区、学院、年级、专业、班级、uid
// department: [2]interface{}: [0]为uid, [1]为名称
func NewUniversity(name string, uid uint, departments [][2]interface{}, data map[int][]SingleStruct) *[6]*Head {
	var result [6]*Head

	// 学校
	university := &Node{
		UID:  uid,
		Mark: name,
	}
	result[0] = &Head{
		Rank:    0,
		Include: &[]*Node{university},
	}

	// 部门
	departmentsNode := make([]*Node, len(departments))
	for i, v := range departments {
		departmentsNode[i] = &Node{
			UID:      v[0].(uint),
			Mark:     v[1].(string),
			Children: &[]*Node{},
		}
	}
	university.Children = &departmentsNode
	result[1] = &Head{
		Rank:    1,
		Include: &departmentsNode,
	}

	// 校区
	nodeList0 := make([]*Node, len(data[0]))
	for i, v := range data[0] {
		temp := &Node{
			UID:      0,
			College:  v.Campus,
			Mark:     v.Campus,
			Children: &[]*Node{},
		}
		nodeList0[i] = temp

		for _, node := range departmentsNode {
			*node.Children = append(*node.Children, temp)
		}
	}
	result[2] = &Head{
		Rank:    2,
		Include: &nodeList0,
	}

	// 学院
	nodeList1 := make([]*Node, len(data[1]))
	for i, v := range data[1] {
		temp := &Node{
			UID:      0,
			College:  v.College,
			Mark:     fmt.Sprintf("%s:%s", v.Campus, v.College),
			Children: &[]*Node{},
		}
		nodeList1[i] = temp

		for _, node := range nodeList0 {
			if v.Campus == node.Mark {
				*node.Children = append(*node.Children, temp)
			}
		}
	}
	result[3] = &Head{
		Rank:    3,
		Include: &nodeList1,
	}

	// 年级
	nodeList2 := make([]*Node, len(data[2]))
	for i, v := range data[2] {
		temp := &Node{
			UID:      v.UID,
			Grade:    v.Grade,
			Mark:     fmt.Sprintf("%s:%s:%d", v.Campus, v.College, v.Grade),
			Children: &[]*Node{},
		}
		nodeList2[i] = temp

		cMark := fmt.Sprintf("%s:%s", v.Campus, v.College)
		for _, node := range nodeList1 {
			if cMark == node.Mark {
				*node.Children = append(*node.Children, temp)
			}
		}
	}
	result[4] = &Head{
		Rank:    4,
		Include: &nodeList2,
	}

	// 专业
	nodeList3 := make([]*Node, len(data[3]))
	for i, v := range data[3] {
		temp := &Node{
			UID:      v.UID,
			Major:    v.Major,
			Mark:     fmt.Sprintf("%s:%s:%d:%s", v.Campus, v.College, v.Grade, v.Major),
			Children: &[]*Node{},
		}
		nodeList3[i] = temp

		mMark := fmt.Sprintf("%s:%s:%d", v.Campus, v.College, v.Grade)
		for _, node := range nodeList2 {
			if mMark == node.Mark {
				*node.Children = append(*node.Children, temp)
			}
		}
	}
	result[5] = &Head{
		Rank:    5,
		Include: &nodeList3,
	}

	// 班级
	for _, v := range data[4] {
		sMark := fmt.Sprintf("%s:%s:%d:%s", v.Campus, v.College, v.Grade, v.Major)
		for _, node := range nodeList3 {
			if sMark == node.Mark {
				temp := &Node{
					UID:   v.UID,
					Class: v.Class,
					Mark:  fmt.Sprintf("%s:%s:%d:%s:%s", v.Campus, v.College, v.Grade, v.Major, v.Class),
				}
				*node.Children = append(*node.Children, temp)
			}
		}
	}

	return &result
}
