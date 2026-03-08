package ui

import (
	"strings"
	"testing"

	"github.com/bytewiredev/bytewire/pkg/dom"
)

func TestProgressBar_DefaultWidth(t *testing.T) {
	html := dom.RenderHTML(ProgressBar(50))
	if !strings.Contains(html, "width:50%") {
		t.Errorf("expected %q in output, got: %s", "width:50%", html)
	}
}

func TestProgressBar_DefaultColor(t *testing.T) {
	html := dom.RenderHTML(ProgressBar(50))
	if !strings.Contains(html, "background-color:#22c55e") {
		t.Errorf("expected %q in output, got: %s", "background-color:#22c55e", html)
	}
}

func TestProgressBar_CustomColor(t *testing.T) {
	html := dom.RenderHTML(ProgressBar(75, "#ff0000"))
	if !strings.Contains(html, "background-color:#ff0000") {
		t.Errorf("expected %q in output, got: %s", "background-color:#ff0000", html)
	}
}

func TestProgressBar_ClampNegative(t *testing.T) {
	html := dom.RenderHTML(ProgressBar(-10))
	if !strings.Contains(html, "width:0%") {
		t.Errorf("expected %q in output, got: %s", "width:0%", html)
	}
}

func TestProgressBar_ClampOver100(t *testing.T) {
	html := dom.RenderHTML(ProgressBar(150))
	if !strings.Contains(html, "width:100%") {
		t.Errorf("expected %q in output, got: %s", "width:100%", html)
	}
}

func TestProgressBarLabeled_LabelText(t *testing.T) {
	html := dom.RenderHTML(ProgressBarLabeled("Upload", 75))
	if !strings.Contains(html, "Upload") {
		t.Errorf("expected %q in output, got: %s", "Upload", html)
	}
}

func TestProgressBarLabeled_PercentageText(t *testing.T) {
	html := dom.RenderHTML(ProgressBarLabeled("Upload", 75))
	if !strings.Contains(html, "75%") {
		t.Errorf("expected %q in output, got: %s", "75%", html)
	}
}

func TestSpinner_BorderTopColor(t *testing.T) {
	html := dom.RenderHTML(Spinner())
	if !strings.Contains(html, "border-top-color:#22c55e") {
		t.Errorf("expected %q in output, got: %s", "border-top-color:#22c55e", html)
	}
}

func TestSpinner_Animation(t *testing.T) {
	html := dom.RenderHTML(Spinner())
	if !strings.Contains(html, "animation:spin") {
		t.Errorf("expected %q in output, got: %s", "animation:spin", html)
	}
}

func TestSpinnerCSS_KeyframesSpin(t *testing.T) {
	if !strings.Contains(SpinnerCSS, "@keyframes spin") {
		t.Errorf("expected %q in SpinnerCSS, got: %s", "@keyframes spin", SpinnerCSS)
	}
}
