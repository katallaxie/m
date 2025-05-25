package styles

import (
	"github.com/katallaxie/m/internal/ui/theme"

	"github.com/charmbracelet/lipgloss"
)

var (
	ImageBakcground = "#212121"
)

// Style generation functions that use the current theme

// BaseStyle returns the base style with background and foreground colors
func BaseStyle() lipgloss.Style {
	t := theme.Current()

	return lipgloss.NewStyle().
		Background(t.Background()).
		Foreground(t.Text())
}

// Regular returns a basic unstyled lipgloss.Style
func Regular() lipgloss.Style {
	return lipgloss.NewStyle()
}

// Bold returns a bold style
func Bold() lipgloss.Style {
	return Regular().Bold(true)
}

// Padded returns a style with horizontal padding
func Padded() lipgloss.Style {
	return Regular().Padding(0, 1)
}

// Border returns a style with a normal border
func Border() lipgloss.Style {
	t := theme.Current()

	return Regular().
		Border(lipgloss.NormalBorder()).
		BorderForeground(t.BorderNormal())
}

// ThickBorder returns a style with a thick border
func ThickBorder() lipgloss.Style {
	t := theme.Current()
	return Regular().
		Border(lipgloss.ThickBorder()).
		BorderForeground(t.BorderNormal())
}

// DoubleBorder returns a style with a double border
func DoubleBorder() lipgloss.Style {
	t := theme.Current()
	return Regular().
		Border(lipgloss.DoubleBorder()).
		BorderForeground(t.BorderNormal())
}

// FocusedBorder returns a style with a border using the focused border color
func FocusedBorder() lipgloss.Style {
	t := theme.Current()
	return Regular().
		Border(lipgloss.NormalBorder()).
		BorderForeground(t.BorderFocused())
}

// DimBorder returns a style with a border using the dim border color
func DimBorder() lipgloss.Style {
	t := theme.Current()
	return Regular().
		Border(lipgloss.NormalBorder()).
		BorderForeground(t.BorderDim())
}

// PrimaryColor returns the primary color from the current theme
func PrimaryColor() lipgloss.AdaptiveColor {
	return theme.Current().Primary()
}

// SecondaryColor returns the secondary color from the current theme
func SecondaryColor() lipgloss.AdaptiveColor {
	return theme.Current().Secondary()
}

// AccentColor returns the accent color from the current theme
func AccentColor() lipgloss.AdaptiveColor {
	return theme.Current().Accent()
}

// ErrorColor returns the error color from the current theme
func ErrorColor() lipgloss.AdaptiveColor {
	return theme.Current().Error()
}

// WarningColor returns the warning color from the current theme
func WarningColor() lipgloss.AdaptiveColor {
	return theme.Current().Warning()
}

// SuccessColor returns the success color from the current theme
func SuccessColor() lipgloss.AdaptiveColor {
	return theme.Current().Success()
}

// InfoColor returns the info color from the current theme
func InfoColor() lipgloss.AdaptiveColor {
	return theme.Current().Info()
}

// TextColor returns the text color from the current theme
func TextColor() lipgloss.AdaptiveColor {
	return theme.Current().Text()
}

// TextMutedColor returns the muted text color from the current theme
func TextMutedColor() lipgloss.AdaptiveColor {
	return theme.Current().TextMuted()
}

// TextEmphasizedColor returns the emphasized text color from the current theme
func TextEmphasizedColor() lipgloss.AdaptiveColor {
	return theme.Current().TextEmphasized()
}

// BackgroundColor returns the background color from the current theme
func BackgroundColor() lipgloss.AdaptiveColor {
	return theme.Current().Background()
}

// BackgroundSecondaryColor returns the secondary background color from the current theme
func BackgroundSecondaryColor() lipgloss.AdaptiveColor {
	return theme.Current().BackgroundSecondary()
}

// BackgroundDarkerColor returns the darker background color from the current theme
func BackgroundDarkerColor() lipgloss.AdaptiveColor {
	return theme.Current().BackgroundDarker()
}

// BorderNormalColor returns the normal border color from the current theme
func BorderNormalColor() lipgloss.AdaptiveColor {
	return theme.Current().BorderNormal()
}

// BorderFocusedColor returns the focused border color from the current theme
func BorderFocusedColor() lipgloss.AdaptiveColor {
	return theme.Current().BorderFocused()
}

// BorderDimColor returns the dim border color from the current theme
func BorderDimColor() lipgloss.AdaptiveColor {
	return theme.Current().BorderDim()
}
