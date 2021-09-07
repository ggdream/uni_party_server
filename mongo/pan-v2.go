package mongo

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type PanDocumentV2 struct {
	UID uint `json:"uid"`

	Tree *PanChildV2 `json:"tree"`

	CreateAt int64 `json:"createAt"`
	UpdateAt int64 `json:"updateAt"`
}

type PanChildV2 struct {
	Type        string        `json:"type"`        // 文件或文件夹
	Name        string        `json:"name"`        // 项目名称
	Hash        string        `json:"hash"`        // 哈希标识
	Size        int64         `json:"size"`        // 如果是文件，则返回文件的大小；如果是目录，则返回0
	Counter     int64         `json:"counter"`     // 如果是目录，则为包含的子项目数；如果是文件，则返回0
	ContentType string        `json:"contentType"` // MIME类型
	Children    []*PanChildV2 `json:"children"`

	CreateAt int64 `json:"createAt"`
	UpdateAt int64 `json:"updateAt"`
}

// findNodeByHash 通过hash值找对应节点
func (p *PanDocumentV2) findNodeByHash(hash string) *PanChildV2 {
	return p._findNodeByHash(p.Tree, hash)
}

// _findNodeByHash 通过hash值找对应节点
func (p *PanDocumentV2) _findNodeByHash(node *PanChildV2, hash string) *PanChildV2 {
	if node == nil || node.Hash == hash {
		return node
	}

	for _, child := range node.Children {
		res := p._findNodeByHash(child, hash)
		if node != nil {
			return res
		}
	}

	return nil
}

// unmarshal 将数据库的数据反序列化到结构体上
func (p *PanDocumentV2) unmarshal(uid uint) (err error) {
	res := client.FindOne(panDocumentName, bson.D{{Key: "uid", Value: uid}})
	if err = res.Err(); err != nil {
		return
	}

	err = res.Decode(p)
	return
}

// update 更新提交
func (p *PanDocumentV2) update() error {
	filter := bson.D{{Key: "uid", Value: p.UID}}
	_, err := client.Update(panDocumentName, filter, p)

	return err
}

// Insert 在某个目录插入新文件/文件夹
func (p *PanDocumentV2) Insert(uid uint, hash string, child *PanChildV2) error {
	if err := p.unmarshal(uid); err != nil {
		return err
	}

	node := p.findNodeByHash(hash)
	node.UpdateAt = time.Now().Unix()
	node.Children = append(node.Children, child)

	return p.update()
}

// Delete 删除文件/文件夹
func (p *PanDocumentV2) Delete(uid uint, dirHash, hash string) error {
	if err := p.unmarshal(uid); err != nil {
		return err
	}

	node := p.findNodeByHash(dirHash)
	idx := -1
	for i, child := range node.Children {
		if child.Hash == hash {
			idx = i
			break
		}
	}

	// 兼容了node的孩子为空的情况
	if idx == -1 {
		return errors.New("cannot find the file/folder")
	}

	length := len(node.Children)
	if idx == length-1 {
		node.Children = node.Children[0 : length-1]
	} else if idx < length-1 {
		node.Children = append(node.Children[0:idx], node.Children[idx+1:]...)
	}
	node.UpdateAt = time.Now().Unix()

	return p.update()
}

// Rename 重命名文件/文件夹
func (p *PanDocumentV2) Rename(uid uint, name, hash string) error {
	if err := p.unmarshal(uid); err != nil {
		return err
	}

	node := p.findNodeByHash(hash)
	node.UpdateAt = time.Now().Unix()
	node.Name = name

	return p.update()
}

// Query 获取文件夹下数据
func (p *PanDocumentV2) Query(uid uint, hash string, dest *PanChildV2) error {
	if err := p.unmarshal(uid); err != nil {
		return err
	}

	*dest = *p.findNodeByHash(hash)

	return nil
}

// QueryByPath 获取文件夹下数据
func (p *PanDocumentV2) QueryByPath(path []string, dest *PanChildV2) error {
	return nil
}
