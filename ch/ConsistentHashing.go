package ConsistentHashing

import (
    "hash/crc32"
    "sync"
    "os"
    "sort"
    "strconv"
    "fmt"
    "time"
)

type node struct {
    Ip string
}

type virtualnode struct {
	Node *node
    Id uint32
    VId uint32
}
type Obj2node struct {
    m_node_lock sync.RWMutex
    m_node map[uint32]*virtualnode
    m_is_active bool
    
    m_CacheNum uint32
    m_virtualNum uint32
}

func getCrc(key string) uint32 {  
    if len(key) < 64 {  
        var scratch [64]byte  
        copy(scratch[:], key)   
        return crc32.ChecksumIEEE(scratch[:len(key)])  
    }  
    return crc32.ChecksumIEEE([]byte(key))  
}
func (v *Obj2node) Init(cache,vnum uint32) {
    v.m_node_lock.Lock()
    v.m_node = make(map[uint32]*virtualnode)
    
    v.m_CacheNum = cache
    v.m_virtualNum = vnum
    
    v.m_node_lock.Unlock()
}

func (v *Obj2node) AddNode(ip string) {
    
    v.m_node_lock.Lock()
    v.m_is_active = false
    var tmpnode node
    tmpnode.Ip=ip
    for i:=uint32(0);i<v.m_virtualNum;i++ {
        var tmpvnode virtualnode
        tmpvnode.VId = i
        tmpvnode.Id = 0
        tmpvnode.Node = &tmpnode
        
        key := getCrc(ip+"#"+strconv.Itoa(int(i))) % v.m_CacheNum
        existNode, ok := v.m_node[key]    // = &tmpvnode
        if ok {
            if existNode.Id == 0 {
                if existNode.VId < tmpvnode.VId {
                    continue
                } else if existNode.VId == tmpvnode.VId {
                    if existNode.Node.Ip < ip {
                        continue
                    }
                }
            }
        }
        v.m_node[key] = &tmpvnode
    }
    v.m_node_lock.Unlock()
}

func (v *Obj2node)Active() bool{
    v.m_node_lock.Lock()

    sorted_keys := make([]int, 0)
    for i,_ := range v.m_node {
        sorted_keys = append(sorted_keys, int(i))
    }
    sort.Ints(sorted_keys)
    
    if len(sorted_keys) == 0 {
        return false
    }
    
    var pre_key uint32 = v.m_CacheNum+1
    var now_key uint32
    var first_key uint32 = uint32(sorted_keys[0])
    for _, k := range sorted_keys {
        var kk = uint32(k)
        if v.m_node[kk].Id == 0 {
            now_key = kk
            if pre_key == v.m_CacheNum+1 {
                pre_key = now_key
                continue
            }
            
            var j uint32 = 1
            for i:=pre_key+1; i<now_key; i++ {
                var tmpvnode virtualnode
                tmpvnode.VId = v.m_node[pre_key].VId
                tmpvnode.Id = j
                j++
                tmpvnode.Node = v.m_node[pre_key].Node
                v.m_node[i] = &tmpvnode
            }
            pre_key = now_key
        }
    }
    
    var j uint32 = 1
    for i:=pre_key+1;i<v.m_CacheNum;i++ {
        var tmpvnode virtualnode
        tmpvnode.VId = v.m_node[pre_key].VId
        tmpvnode.Id = j
        j++
        tmpvnode.Node = v.m_node[pre_key].Node
        v.m_node[i] = &tmpvnode
    }

    for i:=uint32(0);i<first_key;i++ {
        var tmpvnode virtualnode
        tmpvnode.VId = v.m_node[pre_key].VId
        tmpvnode.Id = j
        j++
        tmpvnode.Node = v.m_node[pre_key].Node
        v.m_node[i] = &tmpvnode
    }
    
    v.m_is_active = true
    v.m_node_lock.Unlock()
    
    return true
}


func (v *Obj2node)Get(obj string) string{

    var res string 
    v.m_node_lock.RLock()
    if v.m_is_active {
        res = v.m_node[getCrc(obj) % v.m_CacheNum].Node.Ip
    } else {
        fmt.Fprintln(os.Stderr, time.Now(), "Obj2node is not actived")
        panic("Obj2node is not actived")
        os.Exit(0)
    }
    v.m_node_lock.RUnlock()
    return res
}

func (v *Obj2node)Delete(ip string){
    v.m_node_lock.Lock()
    v.m_is_active = false
    for i:=uint32(0);i<v.m_virtualNum;i++ {
        key := getCrc(ip+"#"+strconv.Itoa(int(i))) % v.m_CacheNum
        existNode, ok := v.m_node[key]
        if !ok {
            continue
        }
        if existNode.Node.Ip != ip {
            continue
        }
        delete(v.m_node, key)
    }
    v.m_node_lock.Unlock()
}

func (v *Obj2node)Debug() {
    v.m_node_lock.RLock()
    fmt.Println("CacheNum:",v.m_CacheNum, "virtualNum:", v.m_virtualNum)
    fmt.Println("real cache num:",len(v.m_node))
    fmt.Println("is actived ?",v.m_is_active)
    
    
    count := make(map[string]int)
    for i:=uint32(0);i<v.m_CacheNum;i++ {
        count[v.m_node[i].Node.Ip]++
        fmt.Println(i,v.m_node[i].Node.Ip, v.m_node[i].Id, v.m_node[i].VId)
    }
    fmt.Println(count)
    v.m_node_lock.RUnlock()
}


















