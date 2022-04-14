package singleton

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLazyInstance(t *testing.T) {
	assert.Equal(t, GetInstance("njdsnj"), GetInstance("sdkjb"))
}

func TestName(t *testing.T) {
	jenny := GetInstance("jenny")
	jenny.ShowMyHeart()
	kimi := GetInstance("kimi")
	kimi.ShowMyHeart()
	asf := GetInstance("sdas")
	asf.ShowMyHeart()
	sd := GetInstance("sdad")
	sd.ShowMyHeart()

}

func BenchmarkGetLazyInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetInstance("sd") != GetInstance("saknjudsa") {
				b.Errorf("test fail")
			}
		}
	})
}
