package rectangle_test

import (
	. "skeleton-go-test/internal/rectangle"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Rectangle", func() {
	var subject *Rectangle

	BeforeEach(func() {
		subject = NewRectangle(2.0, 3.0)
	})

	Context("Area", func() {
		It("Returns six for sides of two and three", func() {
			Expect(subject.Area()).To(Equal(6.0))
		})

		It("Returns six for sides of one and six", func() {
			subject = NewRectangle(1.0, 6.0)

			Expect(subject.Area()).To(Equal(6.0))
		})
	})

})
