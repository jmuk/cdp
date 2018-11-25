// Code generated by cdpgen. DO NOT EDIT.

package overlay

import (
	"encoding/json"

	"github.com/mafredri/cdp/protocol/dom"
	"github.com/mafredri/cdp/protocol/page"
	"github.com/mafredri/cdp/protocol/runtime"
)

// GetHighlightObjectForTestArgs represents the arguments for GetHighlightObjectForTest in the Overlay domain.
type GetHighlightObjectForTestArgs struct {
	NodeID dom.NodeID `json:"nodeId"` // Id of the node to get highlight object for.
}

// NewGetHighlightObjectForTestArgs initializes GetHighlightObjectForTestArgs with the required arguments.
func NewGetHighlightObjectForTestArgs(nodeID dom.NodeID) *GetHighlightObjectForTestArgs {
	args := new(GetHighlightObjectForTestArgs)
	args.NodeID = nodeID
	return args
}

// GetHighlightObjectForTestReply represents the return values for GetHighlightObjectForTest in the Overlay domain.
type GetHighlightObjectForTestReply struct {
	Highlight json.RawMessage `json:"highlight"` // Highlight data for the node.
}

// HighlightFrameArgs represents the arguments for HighlightFrame in the Overlay domain.
type HighlightFrameArgs struct {
	FrameID             page.FrameID `json:"frameId"`                       // Identifier of the frame to highlight.
	ContentColor        *dom.RGBA    `json:"contentColor,omitempty"`        // The content box highlight fill color (default: transparent).
	ContentOutlineColor *dom.RGBA    `json:"contentOutlineColor,omitempty"` // The content box highlight outline color (default: transparent).
}

// NewHighlightFrameArgs initializes HighlightFrameArgs with the required arguments.
func NewHighlightFrameArgs(frameID page.FrameID) *HighlightFrameArgs {
	args := new(HighlightFrameArgs)
	args.FrameID = frameID
	return args
}

// SetContentColor sets the ContentColor optional argument. The
// content box highlight fill color (default: transparent).
func (a *HighlightFrameArgs) SetContentColor(contentColor dom.RGBA) *HighlightFrameArgs {
	a.ContentColor = &contentColor
	return a
}

// SetContentOutlineColor sets the ContentOutlineColor optional argument.
// The content box highlight outline color (default: transparent).
func (a *HighlightFrameArgs) SetContentOutlineColor(contentOutlineColor dom.RGBA) *HighlightFrameArgs {
	a.ContentOutlineColor = &contentOutlineColor
	return a
}

// HighlightNodeArgs represents the arguments for HighlightNode in the Overlay domain.
type HighlightNodeArgs struct {
	HighlightConfig HighlightConfig         `json:"highlightConfig"`         // A descriptor for the highlight appearance.
	NodeID          *dom.NodeID             `json:"nodeId,omitempty"`        // Identifier of the node to highlight.
	BackendNodeID   *dom.BackendNodeID      `json:"backendNodeId,omitempty"` // Identifier of the backend node to highlight.
	ObjectID        *runtime.RemoteObjectID `json:"objectId,omitempty"`      // JavaScript object id of the node to be highlighted.
}

// NewHighlightNodeArgs initializes HighlightNodeArgs with the required arguments.
func NewHighlightNodeArgs(highlightConfig HighlightConfig) *HighlightNodeArgs {
	args := new(HighlightNodeArgs)
	args.HighlightConfig = highlightConfig
	return args
}

// SetNodeID sets the NodeID optional argument. Identifier of the node
// to highlight.
func (a *HighlightNodeArgs) SetNodeID(nodeID dom.NodeID) *HighlightNodeArgs {
	a.NodeID = &nodeID
	return a
}

// SetBackendNodeID sets the BackendNodeID optional argument.
// Identifier of the backend node to highlight.
func (a *HighlightNodeArgs) SetBackendNodeID(backendNodeID dom.BackendNodeID) *HighlightNodeArgs {
	a.BackendNodeID = &backendNodeID
	return a
}

// SetObjectID sets the ObjectID optional argument. JavaScript object
// id of the node to be highlighted.
func (a *HighlightNodeArgs) SetObjectID(objectID runtime.RemoteObjectID) *HighlightNodeArgs {
	a.ObjectID = &objectID
	return a
}

// HighlightQuadArgs represents the arguments for HighlightQuad in the Overlay domain.
type HighlightQuadArgs struct {
	Quad         dom.Quad  `json:"quad"`                   // Quad to highlight
	Color        *dom.RGBA `json:"color,omitempty"`        // The highlight fill color (default: transparent).
	OutlineColor *dom.RGBA `json:"outlineColor,omitempty"` // The highlight outline color (default: transparent).
}

// NewHighlightQuadArgs initializes HighlightQuadArgs with the required arguments.
func NewHighlightQuadArgs(quad dom.Quad) *HighlightQuadArgs {
	args := new(HighlightQuadArgs)
	args.Quad = quad
	return args
}

// SetColor sets the Color optional argument. The highlight fill color
// (default: transparent).
func (a *HighlightQuadArgs) SetColor(color dom.RGBA) *HighlightQuadArgs {
	a.Color = &color
	return a
}

// SetOutlineColor sets the OutlineColor optional argument. The
// highlight outline color (default: transparent).
func (a *HighlightQuadArgs) SetOutlineColor(outlineColor dom.RGBA) *HighlightQuadArgs {
	a.OutlineColor = &outlineColor
	return a
}

// HighlightRectArgs represents the arguments for HighlightRect in the Overlay domain.
type HighlightRectArgs struct {
	X            int       `json:"x"`                      // X coordinate
	Y            int       `json:"y"`                      // Y coordinate
	Width        int       `json:"width"`                  // Rectangle width
	Height       int       `json:"height"`                 // Rectangle height
	Color        *dom.RGBA `json:"color,omitempty"`        // The highlight fill color (default: transparent).
	OutlineColor *dom.RGBA `json:"outlineColor,omitempty"` // The highlight outline color (default: transparent).
}

// NewHighlightRectArgs initializes HighlightRectArgs with the required arguments.
func NewHighlightRectArgs(x int, y int, width int, height int) *HighlightRectArgs {
	args := new(HighlightRectArgs)
	args.X = x
	args.Y = y
	args.Width = width
	args.Height = height
	return args
}

// SetColor sets the Color optional argument. The highlight fill color
// (default: transparent).
func (a *HighlightRectArgs) SetColor(color dom.RGBA) *HighlightRectArgs {
	a.Color = &color
	return a
}

// SetOutlineColor sets the OutlineColor optional argument. The
// highlight outline color (default: transparent).
func (a *HighlightRectArgs) SetOutlineColor(outlineColor dom.RGBA) *HighlightRectArgs {
	a.OutlineColor = &outlineColor
	return a
}

// SetInspectModeArgs represents the arguments for SetInspectMode in the Overlay domain.
type SetInspectModeArgs struct {
	Mode            InspectMode      `json:"mode"`                      // Set an inspection mode.
	HighlightConfig *HighlightConfig `json:"highlightConfig,omitempty"` // A descriptor for the highlight appearance of hovered-over nodes. May be omitted if `enabled == false`.
}

// NewSetInspectModeArgs initializes SetInspectModeArgs with the required arguments.
func NewSetInspectModeArgs(mode InspectMode) *SetInspectModeArgs {
	args := new(SetInspectModeArgs)
	args.Mode = mode
	return args
}

// SetHighlightConfig sets the HighlightConfig optional argument. A
// descriptor for the highlight appearance of hovered-over nodes. May
// be omitted if `enabled == false`.
func (a *SetInspectModeArgs) SetHighlightConfig(highlightConfig HighlightConfig) *SetInspectModeArgs {
	a.HighlightConfig = &highlightConfig
	return a
}

// SetPausedInDebuggerMessageArgs represents the arguments for SetPausedInDebuggerMessage in the Overlay domain.
type SetPausedInDebuggerMessageArgs struct {
	Message *string `json:"message,omitempty"` // The message to display, also triggers resume and step over controls.
}

// NewSetPausedInDebuggerMessageArgs initializes SetPausedInDebuggerMessageArgs with the required arguments.
func NewSetPausedInDebuggerMessageArgs() *SetPausedInDebuggerMessageArgs {
	args := new(SetPausedInDebuggerMessageArgs)

	return args
}

// SetMessage sets the Message optional argument. The message to
// display, also triggers resume and step over controls.
func (a *SetPausedInDebuggerMessageArgs) SetMessage(message string) *SetPausedInDebuggerMessageArgs {
	a.Message = &message
	return a
}

// SetShowDebugBordersArgs represents the arguments for SetShowDebugBorders in the Overlay domain.
type SetShowDebugBordersArgs struct {
	Show bool `json:"show"` // True for showing debug borders
}

// NewSetShowDebugBordersArgs initializes SetShowDebugBordersArgs with the required arguments.
func NewSetShowDebugBordersArgs(show bool) *SetShowDebugBordersArgs {
	args := new(SetShowDebugBordersArgs)
	args.Show = show
	return args
}

// SetShowFPSCounterArgs represents the arguments for SetShowFPSCounter in the Overlay domain.
type SetShowFPSCounterArgs struct {
	Show bool `json:"show"` // True for showing the FPS counter
}

// NewSetShowFPSCounterArgs initializes SetShowFPSCounterArgs with the required arguments.
func NewSetShowFPSCounterArgs(show bool) *SetShowFPSCounterArgs {
	args := new(SetShowFPSCounterArgs)
	args.Show = show
	return args
}

// SetShowPaintRectsArgs represents the arguments for SetShowPaintRects in the Overlay domain.
type SetShowPaintRectsArgs struct {
	Result bool `json:"result"` // True for showing paint rectangles
}

// NewSetShowPaintRectsArgs initializes SetShowPaintRectsArgs with the required arguments.
func NewSetShowPaintRectsArgs(result bool) *SetShowPaintRectsArgs {
	args := new(SetShowPaintRectsArgs)
	args.Result = result
	return args
}

// SetShowScrollBottleneckRectsArgs represents the arguments for SetShowScrollBottleneckRects in the Overlay domain.
type SetShowScrollBottleneckRectsArgs struct {
	Show bool `json:"show"` // True for showing scroll bottleneck rects
}

// NewSetShowScrollBottleneckRectsArgs initializes SetShowScrollBottleneckRectsArgs with the required arguments.
func NewSetShowScrollBottleneckRectsArgs(show bool) *SetShowScrollBottleneckRectsArgs {
	args := new(SetShowScrollBottleneckRectsArgs)
	args.Show = show
	return args
}

// SetShowHitTestBordersArgs represents the arguments for SetShowHitTestBorders in the Overlay domain.
type SetShowHitTestBordersArgs struct {
	Show bool `json:"show"` // True for showing hit-test borders
}

// NewSetShowHitTestBordersArgs initializes SetShowHitTestBordersArgs with the required arguments.
func NewSetShowHitTestBordersArgs(show bool) *SetShowHitTestBordersArgs {
	args := new(SetShowHitTestBordersArgs)
	args.Show = show
	return args
}

// SetShowViewportSizeOnResizeArgs represents the arguments for SetShowViewportSizeOnResize in the Overlay domain.
type SetShowViewportSizeOnResizeArgs struct {
	Show bool `json:"show"` // Whether to paint size or not.
}

// NewSetShowViewportSizeOnResizeArgs initializes SetShowViewportSizeOnResizeArgs with the required arguments.
func NewSetShowViewportSizeOnResizeArgs(show bool) *SetShowViewportSizeOnResizeArgs {
	args := new(SetShowViewportSizeOnResizeArgs)
	args.Show = show
	return args
}

// SetSuspendedArgs represents the arguments for SetSuspended in the Overlay domain.
type SetSuspendedArgs struct {
	Suspended bool `json:"suspended"` // Whether overlay should be suspended and not consume any resources until resumed.
}

// NewSetSuspendedArgs initializes SetSuspendedArgs with the required arguments.
func NewSetSuspendedArgs(suspended bool) *SetSuspendedArgs {
	args := new(SetSuspendedArgs)
	args.Suspended = suspended
	return args
}
