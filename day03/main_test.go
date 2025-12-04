package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("day02", func() {

	DescribeTable("testing joltage",
		func(l []int, r int) {
			rs := joltage(l)
			Expect(r).To(Equal(rs))
		},
		Entry("test", []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1}, 98),
		Entry("test", []int{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1}, 92),
	)

	DescribeTable("testing joltage 2",
		func(l []int, r int) {
			rs := joltageTwo(l)
			Expect(r).To(Equal(rs))
		},
		Entry("test", []int{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1}, 888911112111),
		Entry("test", []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1}, 987654321111),
		Entry("test", []int{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9}, 811111111119),
		Entry("test", []int{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8}, 434234234278),
	)

})
