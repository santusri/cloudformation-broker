package cfbroker_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cf-platform-eng/cloudformation-broker/cfbroker"
)

var _ = Describe("Config", func() {
	var (
		config Config

		validConfig = Config{
			Region:               "cloudformation-region",
			CloudFormationPrefix: "cf",
			Catalog: Catalog{
				[]Service{
					Service{
						ID:          "service-1",
						Name:        "Service 1",
						Description: "Service 1 description",
					},
				},
			},
		}
	)

	Describe("Validate", func() {
		BeforeEach(func() {
			config = validConfig
		})

		It("does not return error if all sections are valid", func() {
			err := config.Validate()
			Expect(err).ToNot(HaveOccurred())
		})

		It("returns error if Region is not valid", func() {
			config.Region = ""

			err := config.Validate()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Must provide a non-empty Region"))
		})

		It("returns error if CloudFormationPrefix is not valid", func() {
			config.CloudFormationPrefix = ""

			err := config.Validate()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Must provide a non-empty CloudFormationPrefix"))
		})

		It("returns error if Catalog is not valid", func() {
			config.Catalog = Catalog{
				[]Service{
					Service{},
				},
			}

			err := config.Validate()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Validating Catalog configuration"))
		})
	})
})
