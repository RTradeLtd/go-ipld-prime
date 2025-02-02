package typed

import (
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/schema"
)

// typed.Node is a superset of the ipld.Node interface, and has additional behaviors.
//
// A typed.Node can be inspected for its schema.Type and schema.Kind,
// which conveys much more and richer information than the Data Model layer
// ipld.ReprKind.
//
// There are many different implementations of typed.Node.
// One implementation can wrap any other existing ipld.Node (i.e., it's zero-copy)
// and promises that it has *already* been validated to match the typesystem.Type;
// another implementation similarly wraps any other existing ipld.Node, but
// defers to the typesystem validation checking to fields that are accessed;
// and when using code generation tools, all of the generated native Golang
// types produced by the codegen will each individually implement typed.Node.
//
// Note that typed.Node can wrap *other* typed.Node instances.
// Imagine you have two parts of a very large code base which have codegen'd
// components which are from different versions of a schema.  Smooth migrations
// and zero-copy type-safe data sharing between them: We can accommodate that!
type Node interface {
	ipld.Node

	Type() schema.Type
}

// unboxing is... ugh, we probably should codegen an unbox method per concrete type.
//  (or, attach them to the non-pointer type, which would namespace in an alloc-free way, but i don't know if that's anything but confusing.)
//  there are notes about this from way back at 2019.01; reread to see if any remain relevant and valid.
// main important point is: it's not gonna be casting.
//  if casting was sufficient to unbox, it'd mean every method on the Node interface would be difficult to use as a field name on a struct type.  undesirable.
//   okay, or, alternative, we flip this to `superapi.Footype{}.Fields().FrobFieldName()`.  that strikes me as unlikely to be pleasing, though.
//    istm we can safely expect direct use of field names much, much more often that flipping back and forth to hypergeneric node; so we should optimize syntax for that accordingly.
