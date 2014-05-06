package apps

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"

	"github.com/cloudfoundry/cf-acceptance-tests/helpers"
	"github.com/pivotal-cf-experimental/cf-test-helpers/cf"
	"github.com/pivotal-cf-experimental/cf-test-helpers/generator"
)

var _ = Describe("Buildpacks", func() {
	var appName string

	BeforeEach(func() {
		appName = generator.RandomName()
	})

	AfterEach(func() {
		Expect(cf.Cf("delete", appName, "-f").Wait(DEFAULT_TIMEOUT)).To(Exit(0))
	})

	Describe("node", func() {
		It("makes the app reachable via its bound route", func() {
			Expect(cf.Cf("push", appName, "-p", helpers.NewAssets().Node, "-c", "node app.js").Wait(CF_PUSH_TIMEOUT)).To(Exit(0))

			Expect(helpers.CurlAppRoot(appName)).To(ContainSubstring("Hello from a node app!"))
		})
	})

	Describe("java", func() {
		It("makes the app reachable via its bound route", func() {
			Expect(cf.Cf("push", appName, "-p", helpers.NewAssets().Java).Wait(CF_PUSH_TIMEOUT)).To(Exit(0))

			Expect(helpers.CurlAppRoot(appName)).To(ContainSubstring("Hello, from your friendly neighborhood Java JSP!"))
		})
	})

	Describe("go", func() {
		It("makes the app reachable via its bound route", func() {
			Expect(cf.Cf("push", appName, "-p", helpers.NewAssets().Go).Wait(CF_PUSH_TIMEOUT)).To(Exit(0))

			Expect(helpers.CurlAppRoot(appName)).To(ContainSubstring("go, world"))
		})
	})
})
