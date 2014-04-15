package emitter_test

import (
	"github.com/cloudfoundry-incubator/dropsonde/emitter"
	"github.com/cloudfoundry-incubator/dropsonde/events"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

type fakeDataSource struct {
}

func (fds *fakeDataSource) GetHeartbeatEvent() events.Event {
	return events.NewTestEvent(42)
}

var _ = Describe("HeartbeatGenerator", func() {
	Describe("BeginGeneration", func() {
		var (
			fakeEmitter          *emitter.FakeEmitter
			heartbeatEventSource = &fakeDataSource{}
		)

		BeforeEach(func() {
			origin := "testHeartbeatEmitter/0"
			fakeEmitter = emitter.NewFake(origin)

			emitter.HeartbeatInterval = 10 * time.Millisecond
		})

		It("periodically emits heartbeats, and the emitter can be closed properly", func() {
			stopChannel, _ := emitter.BeginGeneration(heartbeatEventSource, fakeEmitter)

			Eventually(func() int { return len(fakeEmitter.GetMessages()) }).Should(BeNumerically(">=", 2))

			close(stopChannel)
			Eventually(fakeEmitter.IsClosed).Should(BeTrue())
		})
	})
})
