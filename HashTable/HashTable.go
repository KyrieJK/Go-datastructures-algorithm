package HashTable

import "hash/fnv"

type HashTable struct{
	data [256]*TableItem
}

type TableItem struct{
	key string
	data int
	next *TableItem
}

func (table *HashTable) Add(key string,i int){
	position := generateHash(key)
	if table.data[position] == nil{
		table.data[position] = &TableItem{key:key,data:i}
		return
	}
	current := table.data[position]
	for current.next != nil{
		current = current.next
	}
	current.next = &TableItem{key:key,data:i}
}

func (table *HashTable) Get(key string) (int,bool){
	position := generateHash(key)
	current := table.data[position]
	for current !=nil{
		if current.key == key{
			return current.data,true
		}
		current = current.next
	}
	return -1,false
}

func (table *HashTable) Set(key string,value int) bool{
	position:=generateHash(key)
	current := table.data[position]
	for current != nil{
		if current.key == key{
			current.data = value
			return true
		}
		current = current.next
	}
	return false
}

func (table *HashTable) Remove(key string) bool{
	position := generateHash(key)
	current := table.data[position]
	if current == nil{
		return false
	}

	if current.key == key{
		current = current.next
		return true
	}

	for current.next != nil{
		if current.next.key == key{
			current.next = current.next.next
			return true
		}
		current = current.next
	}
	return false
}

/* Hash生成算法 FNV-1a */
func generateHash(s string) uint8 {
	hash := fnv.New32a()
	hash.Write([]byte(s))
	return uint8(hash.Sum32() % 256)
}

