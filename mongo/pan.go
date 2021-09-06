package mongo

import (
	"gateway/tools/oss"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

const panDocumentName = "pan"

type PanDocument struct {
	UID uint `json:"uid"`

	// Type string `json:"type"`
	Name string `json:"name"`
	Hash string `json:"hash"`

	Children []PanChild `json:"children"`

	CreateAt int64 `json:"createAt"`
	UpdateAt int64 `json:"updateAt"`
}

type PanChild struct {
	Type        string `json:"type"`
	Name        string `json:"name"`
	Hash        string `json:"etag"`
	Size        int64  `json:"size"`
	ContentType string `json:"contentType"`
	// LastMotified int64  `json:"lastMotified"`
}

// Insert 插入文件/文件夹
func (p *PanDocument) Insert(uid uint, dirHash string, child *PanChild) error {
	if child == nil {
		panic("child cannot be nil")
	}

	err := p.Query(uid, dirHash)
	if err != nil {
		return err
	}

	p.Children = append(p.Children, *child)
	tt := time.Now().Unix()
	p.UpdateAt = tt

	_, err = p.Update(uid, dirHash)
	if err != nil {
		return err
	}

	if child.Type != oss.Tree {
		return nil
	}

	document := &PanDocument{
		UID:      p.UID,
		Name:     child.Name,
		Hash:     child.Hash,
		Children: []PanChild{},
		CreateAt: tt,
	}
	return document.Create()
}

// Create 创建新文档
func (p *PanDocument) Create() error {
	return client.Insert(panDocumentName, p)
}

// Update 更新文档
func (p *PanDocument) Update(uid uint, hash string) (int64, error) {
	filter := bson.D{{Key: "uid", Value: uid}, {Key: "hash", Value: hash}}

	return client.Update(panDocumentName, filter, p)
}

// func (p *PanDocument) Update(uid uint, hash string, children []PanChild) (int64, error) {
// 	filter := bson.D{{Key: "uid", Value: uid}, {Key: "hash", Value: hash}}
// 	update := bson.D{{Key: "$set", Value: map[string]interface{}{"children": children}}}

// 	return client.Update(panDocumentName, filter, update)
// }

// Query 查询详细信息
func (p *PanDocument) Query(uid uint, hash string) error {
	res := client.FindOne(panDocumentName, bson.D{{Key: "uid", Value: uid}, {Key: "hash", Value: hash}})
	if err := res.Err(); err != nil {
		return err
	}

	return res.Decode(p)
}


// Delete 删除文件/文件夹
func (p *PanDocument) Delete(uid uint, dirHash, thisHash string) error {
	err := p.Query(uid, dirHash)
	if err != nil {
		return err
	}

	for idx, child := range p.Children {
		if child.Hash == thisHash {
			if child.Type != oss.Tree {
				if idx+1 < len(p.Children) {
					p.Children = append(p.Children[0:idx], p.Children[idx+1:]...)
				} else {
					p.Children = p.Children[0 : len(p.Children)-1]
				}
				_, err = p.Update(uid, dirHash)
			} else {
				var hashList []string
				err = p.del(uid, thisHash, &hashList)
				if err != nil {
					return err
				}
				err = p.removeMany(uid, hashList)
			}
			break
		}
	}

	return err
}

func (p *PanDocument) del(uid uint, dirHash string, hashList *[]string) error {
	err := p.Query(uid, dirHash)
	if err != nil {
		return err
	}

	for _, child := range p.Children {
		if child.Type == oss.Tree {
			*hashList = append(*hashList, child.Hash)
			err := p.del(uid, child.Hash, hashList)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (p *PanDocument) removeMany(uid uint, hashList []string) error {
	filter := bson.D{{Key: "uid", Value: uid}, {Key: "$in", Value: bson.A{hashList}}}
	return client.DeleteMany(panDocumentName, filter)
}
