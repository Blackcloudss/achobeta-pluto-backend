package snowflake

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
	"tgwp/log/zlog"
	"time"
)

var (
	// Epoch 是 Twitter 的 Snowflake 时间戳初始时间（毫秒），设为 2010 年 11 月 4 日 01:42:54 UTC
	// 您可以根据需要自定义初始时间。
	Epoch int64 = 1288834974657

	// NodeBits 表示用于节点（Node）的位数
	// 总共 22 位可以分配给节点和步数（Step）
	NodeBits uint8 = 10

	// StepBits 表示用于步数（Step）的位数
	// 总共 22 位可以分配给节点和步数
	StepBits uint8 = 12

	mu        sync.Mutex
	nodeMax   int64 = -1 ^ (-1 << NodeBits)
	nodeMask        = nodeMax << StepBits
	stepMask  int64 = -1 ^ (-1 << StepBits)
	timeShift       = NodeBits + StepBits
	nodeShift       = StepBits
)

// JSONSyntaxError 是在解析 JSON 时遇到无效 ID 时返回的错误类型。
type JSONSyntaxError struct{ original []byte }

func (j JSONSyntaxError) Error() string {
	return fmt.Sprintf("无效的 Snowflake ID %q", string(j.original))
}

// Node 结构体包含生成 Snowflake ID 所需的基本信息
type Node struct {
	mu    sync.Mutex
	epoch time.Time
	time  int64
	node  int64
	step  int64

	nodeMax   int64
	nodeMask  int64
	stepMask  int64
	timeShift uint8
	nodeShift uint8
}

// ID 是用于 Snowflake ID 的自定义类型，用于附加方法
type ID int64

// NewNode 返回一个可以用来生成 Snowflake ID 的新节点
func NewNode(node int64) (*Node, error) {

	mu.Lock()
	nodeMax = -1 ^ (-1 << NodeBits)
	nodeMask = nodeMax << StepBits
	stepMask = -1 ^ (-1 << StepBits)
	timeShift = NodeBits + StepBits
	nodeShift = StepBits
	mu.Unlock()

	n := Node{}
	n.node = node
	n.nodeMax = -1 ^ (-1 << NodeBits)
	n.nodeMask = n.nodeMax << StepBits
	n.stepMask = -1 ^ (-1 << StepBits)
	n.timeShift = NodeBits + StepBits
	n.nodeShift = StepBits

	if n.node < 0 || n.node > n.nodeMax {
		return nil, errors.New("节点号必须在 0 到 " + strconv.FormatInt(n.nodeMax, 10) + " 之间")
	}

	var curTime = time.Now()
	// 加入 time.Duration，以确保使用单调时钟（若可用）
	n.epoch = curTime.Add(time.Unix(Epoch/1000, (Epoch%1000)*1000000).Sub(curTime))

	return &n, nil
}

// Generate 创建并返回一个唯一的 Snowflake ID
// 确保唯一性：
// - 系统时间准确
// - 不会有多个节点使用相同的节点 ID
func (n *Node) Generate() ID {

	n.mu.Lock()

	now := time.Since(n.epoch).Nanoseconds() / 1000000

	if now == n.time {
		n.step = (n.step + 1) & n.stepMask

		if n.step == 0 {
			for now <= n.time {
				now = time.Since(n.epoch).Nanoseconds() / 1000000
			}
		}
	} else {
		n.step = 0
	}

	n.time = now

	r := ID((now)<<n.timeShift |
		(n.node << n.nodeShift) |
		(n.step),
	)

	n.mu.Unlock()
	return r
}

// Int64 返回 Snowflake ID 的 int64 表示
func (f ID) Int64() int64 {
	return int64(f)
}

// ParseInt64 将 int64 转换为 Snowflake ID
func ParseInt64(id int64) ID {
	return ID(id)
}

// String 返回 Snowflake ID 的字符串表示
func (f ID) String() string {
	return strconv.FormatInt(int64(f), 10)
}

// ParseString 将字符串转换为 Snowflake ID
func ParseString(id string) (ID, error) {
	i, err := strconv.ParseInt(id, 10, 64)
	return ID(i), err
}

// Node 返回 Snowflake ID 的节点号
// 以下函数将在未来版本中移除。
func (f ID) Node() int64 {
	return int64(f) & nodeMask >> nodeShift
}

// Step 返回 Snowflake ID 的步数（或序列号）
// 以下函数将在未来版本中移除。
func (f ID) Step() int64 {
	return int64(f) & stepMask
}

// MarshalJSON 返回 Snowflake ID 的 JSON 字节数组表示
func (f ID) MarshalJSON() ([]byte, error) {
	buff := make([]byte, 0, 22)
	buff = append(buff, '"')
	buff = strconv.AppendInt(buff, int64(f), 10)
	buff = append(buff, '"')
	return buff, nil
}
func GenId(point int) (id string, err error) {
	node, err := NewNode(int64(point))
	if err != nil {
		zlog.Errorf("生成 Node 出错")
	}
	id = node.Generate().String()
	return
}
