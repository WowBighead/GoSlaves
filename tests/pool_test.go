package slaves

import (
	"testing"

	"github.com/themester/GoSlaves"
)

func BenchmarkSlavePool(b *testing.B) {
	ch := make(chan int, b.N)
	done := make(chan struct{})

	sp := slaves.NewPool(func(obj interface{}) {
		ch <- obj.(int)
	})

	go func() {
		var i = 0
		for i < b.N {
			select {
			case <-ch:
				i++
			}
		}
		done <- struct{}{}
	}()

	for i := 0; i < b.N; i++ {
		sp.Serve(i)
	}
	<-done
	close(ch)
	close(done)
}
