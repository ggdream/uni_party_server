package cmap

// CMap 分片的RWMutex+Map
type CMap struct {
	// 分区切片
	area []*cmapSubarea
	// 分区个数
	length int
}

// New 实例化一个CMap
func New(areaCount int) *CMap {
	area := make([]*cmapSubarea, areaCount)
	for i := 0; i < areaCount; i++ {
		area[i] = &cmapSubarea{
			data: make(map[string]interface{}),
		}
	}

	return &CMap{
		area:   area,
		length: areaCount,
	}
}

// hash 哈希函数
func (c *CMap) hash(key string) uint32 {
	return fnv32(key) / uint32(c.length)
}

// Set 添加或更新值
func (c *CMap) Set(key string, value interface{}) {
	c.area[c.hash(key)].Set(key, value)
}

// Get 获取值
func (c *CMap) Get(key string) (value interface{}, ok bool) {
	return c.area[c.hash(key)].Get(key)
}

// Del 删除值
func (c *CMap) Del(key string) {
	c.area[c.hash(key)].Del(key)
}
