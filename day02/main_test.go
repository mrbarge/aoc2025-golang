package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("day02", func() {

	DescribeTable("testing range",
		func(s string, l int, r int) {
			tl, tr := getRange(s)
			Expect(tl).To(Equal(l))
			Expect(tr).To(Equal(r))
		},
		Entry("test", "11-22", 11, 22),
		Entry("test", "824824821-824824827", 824824821, 824824827),
	)

	DescribeTable("testing isInvalid",
		func(i int, isinvalid bool) {
			e := isInvalid(i)
			Expect(e).To(Equal(isinvalid))
		},
		Entry("test", 11, true),
		Entry("test", 22, true),
		Entry("test", 13, false),
		Entry("test", 1188511880, false),
		Entry("test", 1188511885, true),
		Entry("test", 222222, true),
		Entry("test", 22222, false),
	)

	DescribeTable("testing isInvalidTwo",
		func(i int, isinvalid bool) {
			e := isInvalidTwo(i)
			Expect(e).To(Equal(isinvalid))
		},
		Entry("test 11", 11, true),
		Entry("test 22", 22, true),
		Entry("test 13", 13, false),
		Entry("test 999", 999, true),
		Entry("test 1010", 1010, true),
		Entry("test 1011", 1011, false),
		Entry("test 10110", 10110, false),
		Entry("test 2121212121", 2121212121, true),
	)

})
