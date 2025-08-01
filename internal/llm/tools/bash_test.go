package tools

import (
	"strings"
	"testing"

	"github.com/charmbracelet/crush/internal/config"
)

func boolPtr(b bool) *bool {
	return &b
}

func TestBashToolGitCoAuthoring(t *testing.T) {
	// Test with co-authoring enabled
	t.Run("CoAuthoringEnabled", func(t *testing.T) {
		cfg := &config.Config{}
		cfg.Options = &config.Options{EnableGitCoAuthoring: boolPtr(true)}
		
		bashTool := NewBashTool(nil, "/tmp", cfg).(*bashTool)
		description := bashTool.bashDescription()
		
		if !strings.Contains(description, "Co-Authored-By: Crush") {
			t.Error("Expected co-authoring instructions when enabled")
		}
		if !strings.Contains(description, "💘 Generated with Crush") {
			t.Error("Expected Crush signature when co-authoring enabled")
		}
	})

	// Test with co-authoring disabled
	t.Run("CoAuthoringDisabled", func(t *testing.T) {
		cfg := &config.Config{}
		cfg.Options = &config.Options{EnableGitCoAuthoring: boolPtr(false)}
		
		bashTool := NewBashTool(nil, "/tmp", cfg).(*bashTool)
		description := bashTool.bashDescription()
		
		if strings.Contains(description, "Co-Authored-By: Crush") {
			t.Error("Expected no co-authoring instructions when disabled")
		}
		if strings.Contains(description, "💘 Generated with Crush") {
			t.Error("Expected no Crush signature when co-authoring disabled")
		}
	})

	// Test with nil config (should not crash)
	t.Run("NilConfig", func(t *testing.T) {
		bashTool := NewBashTool(nil, "/tmp", nil).(*bashTool)
		description := bashTool.bashDescription()
		
		// Should not crash and should not include co-authoring
		if strings.Contains(description, "Co-Authored-By: Crush") {
			t.Error("Expected no co-authoring instructions with nil config")
		}
	})

	// Test with default (nil pointer) - should be enabled
	t.Run("DefaultEnabled", func(t *testing.T) {
		cfg := &config.Config{}
		cfg.Options = &config.Options{} // EnableGitCoAuthoring is nil, should default to true
		
		bashTool := NewBashTool(nil, "/tmp", cfg).(*bashTool)
		description := bashTool.bashDescription()
		
		if !strings.Contains(description, "Co-Authored-By: Crush") {
			t.Error("Expected co-authoring instructions when using default (should be enabled)")
		}
		if !strings.Contains(description, "💘 Generated with Crush") {
			t.Error("Expected Crush signature when using default (should be enabled)")
		}
	})
}