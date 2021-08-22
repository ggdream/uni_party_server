package group

import (
	"encoding/json"
	"unsafe"
)

type uStruct = map[string]map[string]map[string]map[string][]string

// Group 用map描述大学内部结构tree
// 实现了:
// 1.解析高校发送过来的学生信息表，解析到 Group 的data字段和info字段
// 2.将data字段的内容序列化为JSON，方便持久化该树状结构
// 3.info字段提供入MySQL的行数据片段
type Group struct {
	// 校区、学院、年级、专业、班级
	data uStruct
	// 为了入数据库，创建组级别账户
	info map[int][][5]string
}

// NewGroup 实例化一个 Group
func NewGroup() *Group {
	return &Group{
		data: uStruct{},
		info: map[int][][5]string{},
	}
}

// GetInfo 获取info字段
func (d *Group) GetInfo() map[int][][5]string {
	return d.info
}

// Add 添加
func (d Group) Add(campus, college, grade, major, class string) {
	// 校区
	if d.data[campus] == nil {
		d.data[campus] = map[string]map[string]map[string][]string{}
		if d.info[6] == nil {
			d.info[6] = make([][5]string, 0)
		}
		d.info[6] = append(d.info[6], [5]string{campus, "", "", "", ""})
	}

	// 学院
	if d.data[campus][college] == nil {
		d.data[campus][college] = map[string]map[string][]string{}
		if d.info[5] == nil {
			d.info[5] = make([][5]string, 0)
		}
		d.info[5] = append(d.info[5], [5]string{campus, college, "", "", ""})
	}

	// 年级
	if d.data[campus][college][grade] == nil {
		d.data[campus][college][grade] = map[string][]string{}
		if d.info[4] == nil {
			d.info[4] = make([][5]string, 0)
		}
		d.info[4] = append(d.info[4], [5]string{campus, college, grade, "", ""})
	}

	// 专业
	if d.data[campus][college][grade][major] == nil {
		d.data[campus][college][grade][major] = make([]string, 0)
		if d.info[3] == nil {
			d.info[3] = make([][5]string, 0)
		}
		d.info[3] = append(d.info[3], [5]string{campus, college, grade, major, ""})
	}

	// 如果已经添加过这个班级就返回
	for _, v := range d.data[campus][college][grade][major] {
		if v == class {
			return
		}
	}

	d.data[campus][college][grade][major] = append(d.data[campus][college][grade][major], class)
	if d.info[2] == nil {
		d.info[2] = make([][5]string, 0)
	}
	d.info[2] = append(d.info[2], [5]string{campus, college, grade, major, class})
}

// DataToJSON map转json字符串，用于高校人员结构持久化
func (d *Group) DataToJSON() (string, error) {
	res, err := json.Marshal(d.data)
	if err != nil {
		return "", err
	}

	return *(*string)(unsafe.Pointer(&res)), nil
}

// InfoToJSON map转json字符串，用于高校人员结构持久化
func (d *Group) InfoToJSON() (string, error) {
	res, err := json.Marshal(d.info)
	if err != nil {
		return "", err
	}

	return *(*string)(unsafe.Pointer(&res)), nil
}

