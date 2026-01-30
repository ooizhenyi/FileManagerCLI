package ui

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestUI(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "UI Suite")
}

var _ = Describe("UI Model", func() {
	var (
		m model
	)

	BeforeEach(func() {
	})

	It("should initialize with correct path", func() {
		m = InitialModel("/tmp")
		Expect(m.currentPath).To(Equal("/tmp"))
	})

	It("should navigate up when backspace is pressed", func() {
		m = InitialModel("/usr/bin")

		msg := tea.KeyMsg{Type: tea.KeyBackspace}
		updatedModel, _ := m.Update(msg)

		m = updatedModel.(model)
		Expect(m.currentPath).To(Equal("/usr"))
	})
})
