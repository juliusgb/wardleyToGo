package wardley

import (
	"image"

	svg "github.com/ajstarks/svgo"
	"github.com/owulveryck/wardleyToGo"
	"github.com/owulveryck/wardleyToGo/components"
	"github.com/owulveryck/wardleyToGo/internal/utils"
	"gonum.org/v1/gonum/graph"
)

const (
	RegularEdge wardleyToGo.EdgeType = iota | wardleyToGo.EdgeType(components.Wardley)
	EvolvedComponentEdge
	EvolvedEdge
)

type Collaboration struct {
	F, T  wardleyToGo.Component
	Label string
	Type  wardleyToGo.EdgeType
}

// From returns the from node of the edge.
func (c *Collaboration) From() graph.Node {
	return c.F
}

// To returns the to node of the edge.
func (c *Collaboration) To() graph.Node {
	return c.T
}

// ReversedEdge returns the edge reversal of the receiver
// if a reversal is valid for the data type.
// When a reversal is valid an edge of the same type as
// the receiver with nodes of the receiver swapped should
// be returned, otherwise the receiver should be returned
// unaltered.
func (c *Collaboration) ReversedEdge() graph.Edge {
	return &Collaboration{
		F:     c.T,
		T:     c.F,
		Label: c.Label,
		Type:  c.Type,
	}
}

func (c *Collaboration) GetType() wardleyToGo.EdgeType {
	return c.Type
}

func (c *Collaboration) SVGDraw(s *svg.SVG, r image.Rectangle) {
	fromCoord := c.F.(wardleyToGo.Component).GetPosition()
	toCoord := c.T.(wardleyToGo.Component).GetPosition()
	coordsF := utils.CalcCoords(fromCoord, r)
	coordsT := utils.CalcCoords(toCoord, r)
	switch c.Type {
	case RegularEdge:
		s.Line(coordsF.X, coordsF.Y,
			coordsT.X, coordsF.Y,
			`stroke="grey"`, `stroke-width="1"`)
	case EvolvedComponentEdge:
		s.Line(coordsF.X, coordsF.Y,
			coordsT.X, coordsF.Y,
			`stroke-dasharray="5 5"`, `stroke="red"`, `stroke-width="1"`, `marker-end="url(#arrow)"`)
	case EvolvedEdge:
		s.Line(coordsF.X, coordsF.Y,
			coordsT.X, coordsF.Y,
			`stroke="red"`, `stroke-width="1"`)
	default:
		s.Line(coordsF.X, coordsF.Y,
			coordsT.X, coordsF.Y,
			`stroke="grey"`, `stroke-width="1"`)
	}
}
