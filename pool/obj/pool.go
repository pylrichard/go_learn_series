package obj

/*
	创建/销毁比较消耗资源的对象(比如TCP连接、数据库连接)时，通常需要将这些对象进行池化，以免多次重复创建消耗系统资源
	使用Buffer Channel构建一种对象池，Buffer channel有一个容量(capacity)，表示该Channel中最多存放消息的数量，可以作为对象池的大小
	当初始化对象池时，在Buffer Channel中初始化存放的对象，使用对象时从Buffer Channel中取出对象，使用完毕之后，再将对象放回到Channel中
 */
import (
	"errors"
	"fmt"
	"time"
)

type reusableObj struct {}

type objPool struct {
	bufChan chan *reusableObj
}

func NewObjPool(num int) *objPool {
	objPool := objPool{}
	objPool.bufChan = make(chan *reusableObj, num)
	/*
		可以创建一个空的Buffer Channel，在使用池中对象时再创建对象
	 */
	for i := 0; i < num; i++ {
		objPool.bufChan <-&reusableObj{}
	}

	return &objPool
}

func (p *objPool) GetObj(timeout time.Duration) (*reusableObj, error) {
	select {
	case obj := <-p.bufChan:
		return obj, nil
	//超时控制，符合高可用系统slow response危害大于fast failed的设计原则
	case <-time.After(timeout):
		return nil, errors.New("timeout error")
	}
}

func (p *objPool) ReleaseObj(obj *reusableObj) error {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("obj pool overflow")
	}
}

func TestObjPool() {
	pool := NewObjPool(8)
	for i := 0; i < 9; i++ {
		if v, err := pool.GetObj(time.Second * 2); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%T \n", v)
			if err := pool.ReleaseObj(v); err != nil {
				fmt.Println(err)
			}
		}
	}
}