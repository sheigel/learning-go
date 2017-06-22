package kata_test
import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "codewarrior/kata"
)

func testequal(h, bounce, window float64, exp int, message string) {
	var ans = BouncingBall(h, bounce, window)
	Expect(ans).To(Equal(exp), message)
}

var _ = Describe("Test Example", func() {

	It("should handle basic cases", func() {
		testequal(3, 0.66, 1.5, 3,"h=3")
		testequal(2, 0.5, 1, 2,"h=2")
		testequal(40, 0.4, 10, 3,"h=40")
		testequal(10, 0.6, 10, -1,"h=10")
		testequal(40, 1, 10, -1,"h=10, b=1")
		testequal(5, -1, 1.5, -1,"h=5")
		testequal(100, 0.1, 1, 4,"h=100")
	})
})

