package theme

import (
	"github.com/charmbracelet/lipgloss"
)

// Theme defines the interface for a theme in the application.
type Theme interface {
	// Base colors
	Primary() lipgloss.AdaptiveColor
	Secondary() lipgloss.AdaptiveColor
	Accent() lipgloss.AdaptiveColor

	// Status colors
	Error() lipgloss.AdaptiveColor
	Info() lipgloss.AdaptiveColor
	Success() lipgloss.AdaptiveColor
	Warning() lipgloss.AdaptiveColor

	// Text colors
	Text() lipgloss.AdaptiveColor
	TextMuted() lipgloss.AdaptiveColor
	TextEmphasized() lipgloss.AdaptiveColor

	// Background colors
	Background() lipgloss.AdaptiveColor
	BackgroundSecondary() lipgloss.AdaptiveColor
	BackgroundDarker() lipgloss.AdaptiveColor

	// Border colors
	BorderNormal() lipgloss.AdaptiveColor
	BorderFocused() lipgloss.AdaptiveColor
	BorderDim() lipgloss.AdaptiveColor

	// Diff view colors
	DiffAdded() lipgloss.AdaptiveColor
	DiffRemoved() lipgloss.AdaptiveColor
	DiffContext() lipgloss.AdaptiveColor
	DiffHunkHeader() lipgloss.AdaptiveColor
	DiffHighlightAdded() lipgloss.AdaptiveColor
	DiffHighlightRemoved() lipgloss.AdaptiveColor
	DiffAddedBg() lipgloss.AdaptiveColor
	DiffRemovedBg() lipgloss.AdaptiveColor
	DiffContextBg() lipgloss.AdaptiveColor
	DiffLineNumber() lipgloss.AdaptiveColor
	DiffAddedLineNumberBg() lipgloss.AdaptiveColor
	DiffRemovedLineNumberBg() lipgloss.AdaptiveColor

	// Markdown colors
	MarkdownText() lipgloss.AdaptiveColor
	MarkdownHeading() lipgloss.AdaptiveColor
	MarkdownLink() lipgloss.AdaptiveColor
	MarkdownLinkText() lipgloss.AdaptiveColor
	MarkdownCode() lipgloss.AdaptiveColor
	MarkdownBlockQuote() lipgloss.AdaptiveColor
	MarkdownEmph() lipgloss.AdaptiveColor
	MarkdownStrong() lipgloss.AdaptiveColor
	MarkdownHorizontalRule() lipgloss.AdaptiveColor
	MarkdownListItem() lipgloss.AdaptiveColor
	MarkdownListEnumeration() lipgloss.AdaptiveColor
	MarkdownImage() lipgloss.AdaptiveColor
	MarkdownImageText() lipgloss.AdaptiveColor
	MarkdownCodeBlock() lipgloss.AdaptiveColor

	// Syntax highlighting colors
	SyntaxComment() lipgloss.AdaptiveColor
	SyntaxKeyword() lipgloss.AdaptiveColor
	SyntaxFunction() lipgloss.AdaptiveColor
	SyntaxVariable() lipgloss.AdaptiveColor
	SyntaxString() lipgloss.AdaptiveColor
	SyntaxNumber() lipgloss.AdaptiveColor
	SyntaxType() lipgloss.AdaptiveColor
	SyntaxOperator() lipgloss.AdaptiveColor
	SyntaxPunctuation() lipgloss.AdaptiveColor
}

// Default is the default theme used by the application.
type Default struct {
	// Base colors
	PrimaryColor   lipgloss.AdaptiveColor
	SecondaryColor lipgloss.AdaptiveColor
	AccentColor    lipgloss.AdaptiveColor

	// Status colors
	ErrorColor   lipgloss.AdaptiveColor
	WarningColor lipgloss.AdaptiveColor
	SuccessColor lipgloss.AdaptiveColor
	InfoColor    lipgloss.AdaptiveColor

	// Text colors
	TextColor           lipgloss.AdaptiveColor
	TextMutedColor      lipgloss.AdaptiveColor
	TextEmphasizedColor lipgloss.AdaptiveColor

	// Background colors
	BackgroundColor          lipgloss.AdaptiveColor
	BackgroundSecondaryColor lipgloss.AdaptiveColor
	BackgroundDarkerColor    lipgloss.AdaptiveColor

	// Border colors
	BorderNormalColor  lipgloss.AdaptiveColor
	BorderFocusedColor lipgloss.AdaptiveColor
	BorderDimColor     lipgloss.AdaptiveColor

	// Diff view colors
	DiffAddedColor               lipgloss.AdaptiveColor
	DiffRemovedColor             lipgloss.AdaptiveColor
	DiffContextColor             lipgloss.AdaptiveColor
	DiffHunkHeaderColor          lipgloss.AdaptiveColor
	DiffHighlightAddedColor      lipgloss.AdaptiveColor
	DiffHighlightRemovedColor    lipgloss.AdaptiveColor
	DiffAddedBgColor             lipgloss.AdaptiveColor
	DiffRemovedBgColor           lipgloss.AdaptiveColor
	DiffContextBgColor           lipgloss.AdaptiveColor
	DiffLineNumberColor          lipgloss.AdaptiveColor
	DiffAddedLineNumberBgColor   lipgloss.AdaptiveColor
	DiffRemovedLineNumberBgColor lipgloss.AdaptiveColor

	// Markdown colors
	MarkdownTextColor            lipgloss.AdaptiveColor
	MarkdownHeadingColor         lipgloss.AdaptiveColor
	MarkdownLinkColor            lipgloss.AdaptiveColor
	MarkdownLinkTextColor        lipgloss.AdaptiveColor
	MarkdownCodeColor            lipgloss.AdaptiveColor
	MarkdownBlockQuoteColor      lipgloss.AdaptiveColor
	MarkdownEmphColor            lipgloss.AdaptiveColor
	MarkdownStrongColor          lipgloss.AdaptiveColor
	MarkdownHorizontalRuleColor  lipgloss.AdaptiveColor
	MarkdownListItemColor        lipgloss.AdaptiveColor
	MarkdownListEnumerationColor lipgloss.AdaptiveColor
	MarkdownImageColor           lipgloss.AdaptiveColor
	MarkdownImageTextColor       lipgloss.AdaptiveColor
	MarkdownCodeBlockColor       lipgloss.AdaptiveColor

	// Syntax highlighting colors
	SyntaxCommentColor     lipgloss.AdaptiveColor
	SyntaxKeywordColor     lipgloss.AdaptiveColor
	SyntaxFunctionColor    lipgloss.AdaptiveColor
	SyntaxVariableColor    lipgloss.AdaptiveColor
	SyntaxStringColor      lipgloss.AdaptiveColor
	SyntaxNumberColor      lipgloss.AdaptiveColor
	SyntaxTypeColor        lipgloss.AdaptiveColor
	SyntaxOperatorColor    lipgloss.AdaptiveColor
	SyntaxPunctuationColor lipgloss.AdaptiveColor
}

// Implement the Theme interface for BaseTheme
func (d *Default) Primary() lipgloss.AdaptiveColor   { return d.PrimaryColor }
func (d *Default) Secondary() lipgloss.AdaptiveColor { return d.SecondaryColor }
func (d *Default) Accent() lipgloss.AdaptiveColor    { return d.AccentColor }

func (d *Default) Error() lipgloss.AdaptiveColor   { return d.ErrorColor }
func (d *Default) Warning() lipgloss.AdaptiveColor { return d.WarningColor }
func (d *Default) Success() lipgloss.AdaptiveColor { return d.SuccessColor }
func (d *Default) Info() lipgloss.AdaptiveColor    { return d.InfoColor }

func (d *Default) Text() lipgloss.AdaptiveColor           { return d.TextColor }
func (d *Default) TextMuted() lipgloss.AdaptiveColor      { return d.TextMutedColor }
func (d *Default) TextEmphasized() lipgloss.AdaptiveColor { return d.TextEmphasizedColor }

func (d *Default) Background() lipgloss.AdaptiveColor          { return d.BackgroundColor }
func (d *Default) BackgroundSecondary() lipgloss.AdaptiveColor { return d.BackgroundSecondaryColor }
func (d *Default) BackgroundDarker() lipgloss.AdaptiveColor    { return d.BackgroundDarkerColor }

func (d *Default) BorderNormal() lipgloss.AdaptiveColor  { return d.BorderNormalColor }
func (d *Default) BorderFocused() lipgloss.AdaptiveColor { return d.BorderFocusedColor }
func (d *Default) BorderDim() lipgloss.AdaptiveColor     { return d.BorderDimColor }

func (d *Default) DiffAdded() lipgloss.AdaptiveColor            { return d.DiffAddedColor }
func (d *Default) DiffRemoved() lipgloss.AdaptiveColor          { return d.DiffRemovedColor }
func (d *Default) DiffContext() lipgloss.AdaptiveColor          { return d.DiffContextColor }
func (d *Default) DiffHunkHeader() lipgloss.AdaptiveColor       { return d.DiffHunkHeaderColor }
func (d *Default) DiffHighlightAdded() lipgloss.AdaptiveColor   { return d.DiffHighlightAddedColor }
func (d *Default) DiffHighlightRemoved() lipgloss.AdaptiveColor { return d.DiffHighlightRemovedColor }
func (d *Default) DiffAddedBg() lipgloss.AdaptiveColor          { return d.DiffAddedBgColor }
func (d *Default) DiffRemovedBg() lipgloss.AdaptiveColor        { return d.DiffRemovedBgColor }
func (d *Default) DiffContextBg() lipgloss.AdaptiveColor        { return d.DiffContextBgColor }
func (d *Default) DiffLineNumber() lipgloss.AdaptiveColor       { return d.DiffLineNumberColor }
func (d *Default) DiffAddedLineNumberBg() lipgloss.AdaptiveColor {
	return d.DiffAddedLineNumberBgColor
}
func (d *Default) DiffRemovedLineNumberBg() lipgloss.AdaptiveColor {
	return d.DiffRemovedLineNumberBgColor
}

func (d *Default) MarkdownText() lipgloss.AdaptiveColor       { return d.MarkdownTextColor }
func (d *Default) MarkdownHeading() lipgloss.AdaptiveColor    { return d.MarkdownHeadingColor }
func (d *Default) MarkdownLink() lipgloss.AdaptiveColor       { return d.MarkdownLinkColor }
func (d *Default) MarkdownLinkText() lipgloss.AdaptiveColor   { return d.MarkdownLinkTextColor }
func (d *Default) MarkdownCode() lipgloss.AdaptiveColor       { return d.MarkdownCodeColor }
func (d *Default) MarkdownBlockQuote() lipgloss.AdaptiveColor { return d.MarkdownBlockQuoteColor }
func (d *Default) MarkdownEmph() lipgloss.AdaptiveColor       { return d.MarkdownEmphColor }
func (d *Default) MarkdownStrong() lipgloss.AdaptiveColor     { return d.MarkdownStrongColor }
func (d *Default) MarkdownHorizontalRule() lipgloss.AdaptiveColor {
	return d.MarkdownHorizontalRuleColor
}
func (d *Default) MarkdownListItem() lipgloss.AdaptiveColor { return d.MarkdownListItemColor }
func (d *Default) MarkdownListEnumeration() lipgloss.AdaptiveColor {
	return d.MarkdownListEnumerationColor
}
func (d *Default) MarkdownImage() lipgloss.AdaptiveColor     { return d.MarkdownImageColor }
func (d *Default) MarkdownImageText() lipgloss.AdaptiveColor { return d.MarkdownImageTextColor }
func (d *Default) MarkdownCodeBlock() lipgloss.AdaptiveColor { return d.MarkdownCodeBlockColor }

func (d *Default) SyntaxComment() lipgloss.AdaptiveColor     { return d.SyntaxCommentColor }
func (d *Default) SyntaxKeyword() lipgloss.AdaptiveColor     { return d.SyntaxKeywordColor }
func (d *Default) SyntaxFunction() lipgloss.AdaptiveColor    { return d.SyntaxFunctionColor }
func (d *Default) SyntaxVariable() lipgloss.AdaptiveColor    { return d.SyntaxVariableColor }
func (d *Default) SyntaxString() lipgloss.AdaptiveColor      { return d.SyntaxStringColor }
func (d *Default) SyntaxNumber() lipgloss.AdaptiveColor      { return d.SyntaxNumberColor }
func (d *Default) SyntaxType() lipgloss.AdaptiveColor        { return d.SyntaxTypeColor }
func (d *Default) SyntaxOperator() lipgloss.AdaptiveColor    { return d.SyntaxOperatorColor }
func (d *Default) SyntaxPunctuation() lipgloss.AdaptiveColor { return d.SyntaxPunctuationColor }
