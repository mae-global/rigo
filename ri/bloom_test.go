package ri

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"fmt"
)

func Test_BloomFilter(t *testing.T) {

	Convey("Bloom Filter",t,func() {
		Convey("Small",func() {

			f := NewBloomFilter(1)
			So(f.Len(),ShouldEqual,MinimalBloomFilterSize)
		})
		
		Convey("Normal Usage",func() {

			f := NewBloomFilter(DefaultBloomFilterSize)
			So(f.Len(),ShouldEqual,DefaultBloomFilterSize)

			f = f.Append(fmt.Sprintf("Alice In Wonderland"))
			So(f,ShouldNotBeNil)

			//fmt.Printf("\n%s\n",f.Print())

			So(f.IsMember("Alice In Wonderland"),ShouldBeTrue)
			So(f.IsMember("Fred"),ShouldBeFalse)
			
			for i := 0; i < 100; i++ {
				f = f.Append(fmt.Sprintf("Alice_%03d",i))

				So(f.IsMember(fmt.Sprintf("Alice_%03d",i)),ShouldBeTrue)
				So(f.IsMember("Fred"),ShouldBeFalse)
			}

			//fmt.Printf("%s\n",f.Print())
		})
	})
}

func Test_RiBloomFilter(t *testing.T) {

	Convey("Ri Bloom Filter",t,func() {
		f := RiBloomFilter()
		So(f,ShouldNotBeNil)

		So(f.IsMember("Sphere"),ShouldBeTrue)
		So(f.IsMember("Foobar"),ShouldBeFalse)
	//	fmt.Printf("RiBloomFilter\n %s\n",f.Print())
	})
}



func Benchmark_BloomFilterHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hash(fmt.Sprintf("Alice_%03d",i))
	}
}

func Benchmark_BloomFilterAppendSingle(b *testing.B) {

	f := NewBloomFilter(DefaultBloomFilterSize)

	for i := 0; i < b.N; i++ {
		f = f.Append("Alice")
	}
}

func Benchmark_BloomFilterAppend(b *testing.B) {

	f := NewBloomFilter(DefaultBloomFilterSize)

	for i := 0; i < b.N; i++ {
		f = f.Append(fmt.Sprintf("Alice_%03d",i),fmt.Sprintf("Fred_%03d",i),fmt.Sprintf("Eve_%03d",i))
	}
}

func Benchmark_BloomFilterMembership(b *testing.B) {

	f := NewBloomFilter(DefaultBloomFilterSize)
	f = f.Append("Alice","Fred","Eve")

	for i := 0; i < b.N; i++ {
		f.IsMember("Alice","Fred","Eve")
	}
}

func Benchmark_BloomFilterMemberShipSingle(b *testing.B) {

	f := NewBloomFilter(DefaultBloomFilterSize)
	f = f.Append("Alice")

	for i := 0; i < b.N; i++ {
		f.IsMember("Alice")
	}
}

			
				
			
