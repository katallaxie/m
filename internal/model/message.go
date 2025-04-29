package model

import "github.com/google/uuid"

// Role is a chat role.
type Role string

// String returns the string representation of the chat model.
func (r Role) String() string {
	return string(r)
}

const (
	// RoleUser indicates that a message was send by a user.
	RoleUser Role = "user"
	// RoleHuman indicates that a message was send by a human.
	RoleHuman Role = "human"
	// RoleAI indicates that a message was send by an AI.
	RoleAI Role = "ai"
	// RoleSystem indicates that a message was send by the system.
	RoleSystem Role = "system"
	// RoleAssistant indicates that a message was send by an assistant.
	RoleAssistant Role = "assistant"
	// RoleTool indicates that a message was send by a tool.
	RoleTool Role = "tool"
	// RoleFunction indicates that a message was send by a function.
	RoleFunction Role = "function"
	// RoleGeneric indicates that a message was send by a generic role.
	RoleGeneric Role = "generic"
)

// Message is an interface that represents a message in the application.
type Message interface {
	// ID returns the unique identifier of the message.
	ID() string
	// Role returns the role of the message.
	Role() Role
	// Content returns the content of the message.
	Content() string
}

var (
	_ Message = (*UserMessage)(nil)
	_ Message = (*HumanMessage)(nil)
	_ Message = (*AIMessage)(nil)
	_ Message = (*SystemMessage)(nil)
	_ Message = (*GenericMessage)(nil)
	_ Message = (*ToolMessage)(nil)
)

// Human chat message.
type HumanMessage struct {
	id      string
	content string
}

// ID returns the unique identifier of the message.
func (m *HumanMessage) ID() string {
	return m.id
}

// Content returns the content of the message.
func (m *HumanMessage) Content() string {
	return m.content
}

// Role returns the role of the message.
func (m *HumanMessage) Role() Role {
	return RoleHuman
}

// NewUserMessage returns a new user message.
func NewUserMessage(content string) *UserMessage {
	return &UserMessage{
		id:      uuid.NewString(),
		content: content,
	}
}

// UserMessage is a user chat message.
type UserMessage struct {
	id      string
	content string
}

// ID returns the unique identifier of the message.
func (m *UserMessage) ID() string {
	return m.id
}

// Content returns the content of the message.
func (m *UserMessage) Content() string {
	return m.content
}

// Role returns the role of the message.
func (m *UserMessage) Role() Role {
	return RoleUser
}

// AIMessage is an AI chat message.
type AIMessage struct {
	id      string
	content string
}

// ID returns the unique identifier of the message.
func (m *AIMessage) ID() string {
	return m.id
}

// Role returns the role of the message.
func (m *AIMessage) Role() Role {
	return RoleAI
}

// Content returns the content of the message.
func (m *AIMessage) Content() string {
	return m.content
}

// SystemMessage is a system chat message.
type SystemMessage struct {
	id      string
	content string
}

// ID returns the unique identifier of the message.
func (m *SystemMessage) ID() string {
	return m.id
}

// Content returns the content of the message.
func (m *SystemMessage) Content() string {
	return m.content
}

// Role returns the role of the message.
func (m *SystemMessage) Role() Role {
	return RoleAI
}

// GenericMessage is a generic chat message.
type GenericMessage struct {
	id      string
	content string
	name    string
}

// ID returns the unique identifier of the message.
func (m *GenericMessage) ID() string {
	return m.id
}

// Content returns the content of the message.
func (m *GenericMessage) Content() string {
	return m.content
}

// Role returns the role of the message.
func (m *GenericMessage) Role() Role {
	return RoleGeneric
}

// GetName returns the name of the message.
func (m *GenericMessage) GetName() string {
	return m.name
}

// ToolMessage is a tool chat message.
type ToolMessage struct {
	id      string
	content string
}

// Content returns the content of the message.
func (m *ToolMessage) Content() string {
	return m.content
}

// Role returns the role of the message.
func (m *ToolMessage) Role() Role {
	return RoleTool
}

// ID returns the ID of the message.
func (m *ToolMessage) ID() string {
	return m.id
}
