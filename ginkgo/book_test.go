package ginkgo_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"utils/ginkgo"
)

var _ = Describe("Book", func() {
	var (
		longBook  ginkgo.Book
		shortBook ginkgo.Book
	)

	BeforeEach(func() {
		longBook = ginkgo.Book{
			Title:  "Les Miserables",
			Author: "Victor Hugo",
			Pages:  1488,
		}

		shortBook = ginkgo.Book{
			Title:  "Fox In Socks",
			Author: "Dr. Seuss",
			Pages:  24,
		}
	})

	Describe("Categorizing book length", func() {
		Context("With more than 300 pages", func() {
			It("should be a novel", func() {
				Expect(longBook.CategoryByLength()).To(Equal("NOVEL"))
			})
		})

		Context("With fewer than 300 pages", func() {
			It("should be a short story", func() {
				Expect(shortBook.CategoryByLength()).To(Equal("SHORT STORY"))
			})
		})
	})

})
