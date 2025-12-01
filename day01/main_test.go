package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("day01", func() {

	DescribeTable("testing rotations",
		func(startingPos int, operation string, endingPos int, zeroCrosses int) {
			result, zerocross, err := rotate(startingPos, operation)
			Expect(result).To(Equal(endingPos))
			Expect(zerocross).To(Equal(zeroCrosses))
			Expect(err).To(BeNil())
		},
		Entry("when rotating right but not past 99", 10, "R10", 20, 0),
		Entry("when rotating left but not past 0", 10, "L5", 5, 0),
		Entry("when rotating right past 99", 50, "R60", 10, 1),
		Entry("when rotating left past 0", 50, "L60", 90, 1),
		Entry("when rotating right past 0 many times", 50, "R400", 50, 4),
		Entry("when rotating left past 0 many times", 50, "L400", 50, 4),
		Entry("when rotating left past 0 many times", 50, "L150", 0, 2),
		Entry("when rotating right past 0 many times", 50, "R150", 0, 2),
	)
})
