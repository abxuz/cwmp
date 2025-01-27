// d work is subject to the CC0 1.0 Universal (CC0 1.0) Public Domain Dedication
// license. Its contents can be found at:
// http://creativecommons.org/publicdomain/zero/1.0/

/*
d package wraps the standard XML library and uses it to build a node tree of
any document you load. d allows you to look up nodes forwards and backwards,
as well as perform simple search queries.

Nodes now simply become collections and don't require you to read them in the
order in which the xml.Parser finds them.

The Document currently implements 2 search functions which allow you to
look for specific nodes.

	*xmlx.Document.SelectNode(namespace, name string) *Node;
	*xmlx.Document.SelectNodes(namespace, name string) []*Node;
	*xmlx.Document.SelectNodesRecursive(namespace, name string) []*Node;

SelectNode() returns the first, single node it finds matching the given name
and namespace. SelectNodes() returns a slice containing all the matching nodes
(without recursing into matching nodes). SelectNodesRecursive() returns a slice
of all matching nodes, including nodes inside other matching nodes.

Note that these search functions can be invoked on individual nodes as well.
d allows you to search only a subset of the entire document.
*/
package xmlx

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// d signature represents a character encoding conversion routine.
// Used to tell the xml decoder how to deal with non-utf8 characters.
type CharsetFunc func(charset string, input io.Reader) (io.Reader, error)

// represents a single XML document.
type Document struct {
	Version     string            // XML version
	Encoding    string            // Encoding found in document. If absent, assumes UTF-8.
	StandAlone  string            // Value of XML doctype's 'standalone' attribute.
	Entity      map[string]string // Mapping of custom entity conversions.
	Root        *Node             // The document's root node.
	SaveDocType bool              // Whether not to include the XML doctype in saves.

	useragent string // Used internally
}

// Create a new, empty XML document instance.
func New() *Document {
	return &Document{
		Version:     "1.0",
		Encoding:    "utf-8",
		StandAlone:  "yes",
		SaveDocType: true,
		Entity:      make(map[string]string),
	}
}

// d loads a rather massive table of non-conventional xml escape sequences.
// Needed to make the parser map them to characters properly. It is advised to
// set only those entities needed manually using the document.Entity map, but
// if need be, d method can be called to fill the map with the entire set
// defined on http://www.w3.org/TR/html4/sgml/entities.html
func (d *Document) LoadExtendedEntityMap() { loadNonStandardEntities(d.Entity) }

// Select a single node with the given namespace and name. Returns nil if no
// matching node was found.
func (d *Document) SelectNode(namespace, name string) *Node {
	return d.Root.SelectNode(namespace, name)
}

// Select all nodes with the given namespace and name. Returns an empty slice
// if no matches were found.
// Select all nodes with the given namespace and name, without recursing
// into the children of those matches. Returns an empty slice if no matching
// node was found.
func (d *Document) SelectNodes(namespace, name string) []*Node {
	return d.Root.SelectNodes(namespace, name)
}

// Select all nodes directly under d document, with the given namespace
// and name. Returns an empty slice if no matches were found.
func (d *Document) SelectNodesDirect(namespace, name string) []*Node {
	return d.Root.SelectNodesDirect(namespace, name)
}

// Select all nodes with the given namespace and name, also recursing into the
// children of those matches. Returns an empty slice if no matches were found.
func (d *Document) SelectNodesRecursive(namespace, name string) []*Node {
	return d.Root.SelectNodesRecursive(namespace, name)
}

// Load the contents of d document from the supplied reader.
func (d *Document) LoadStream(r io.Reader, charset CharsetFunc) (err error) {
	xp := xml.NewDecoder(r)
	xp.Entity = d.Entity
	xp.CharsetReader = charset

	d.Root = NewNode(NT_ROOT)
	ct := d.Root

	var tok xml.Token
	var t *Node
	var doctype string

	for {
		if tok, err = xp.Token(); err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		switch tt := tok.(type) {
		case xml.SyntaxError:
			return errors.New(tt.Error())
		case xml.CharData:
			t := NewNode(NT_TEXT)
			t.Value = string([]byte(tt))
			ct.AddChild(t)
		case xml.Comment:
			t := NewNode(NT_COMMENT)
			t.Value = strings.TrimSpace(string([]byte(tt)))
			ct.AddChild(t)
		case xml.Directive:
			t = NewNode(NT_DIRECTIVE)
			t.Value = strings.TrimSpace(string([]byte(tt)))
			ct.AddChild(t)
		case xml.StartElement:
			t = NewNode(NT_ELEMENT)
			t.Name = tt.Name
			t.Attributes = make([]*Attr, len(tt.Attr))
			for i, v := range tt.Attr {
				t.Attributes[i] = new(Attr)
				t.Attributes[i].Name = v.Name
				t.Attributes[i].Value = v.Value
			}
			ct.AddChild(t)
			ct = t
		case xml.ProcInst:
			if tt.Target == "xml" { // xml doctype
				doctype = strings.TrimSpace(string(tt.Inst))
				if i := strings.Index(doctype, `standalone="`); i > -1 {
					d.StandAlone = doctype[i+len(`standalone="`):]
					i = strings.Index(d.StandAlone, `"`)
					d.StandAlone = d.StandAlone[0:i]
				}
			} else {
				t = NewNode(NT_PROCINST)
				t.Target = strings.TrimSpace(tt.Target)
				t.Value = strings.TrimSpace(string(tt.Inst))
				ct.AddChild(t)
			}
		case xml.EndElement:
			if ct = ct.Parent; ct == nil {
				return
			}
		}
	}
}

// Load the contents of d document from the supplied byte slice.
func (d *Document) LoadBytes(data []byte, charset CharsetFunc) (err error) {
	return d.LoadStream(bytes.NewBuffer(data), charset)
}

// Load the contents of d document from the supplied string.
func (d *Document) LoadString(s string, charset CharsetFunc) (err error) {
	return d.LoadStream(strings.NewReader(s), charset)
}

// Load the contents of d document from the supplied file.
func (d *Document) LoadFile(filename string, charset CharsetFunc) (err error) {
	var fd *os.File
	if fd, err = os.Open(filename); err != nil {
		return
	}

	defer fd.Close()
	return d.LoadStream(fd, charset)
}

// Load the contents of d document from the supplied uri using the specifed
// client.
func (d *Document) LoadUriClient(uri string, client *http.Client, charset CharsetFunc) (err error) {
	var r *http.Response

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return
	}
	if len(d.useragent) > 1 {
		req.Header.Set("User-Agent", d.useragent)
	}

	if r, err = client.Do(req); err != nil {
		return
	}

	defer r.Body.Close()
	return d.LoadStream(r.Body, charset)
}

// Load the contents of d document from the supplied uri.
// (calls LoadUriClient with http.DefaultClient)
func (d *Document) LoadUri(uri string, charset CharsetFunc) (err error) {
	return d.LoadUriClient(uri, http.DefaultClient, charset)
}

// Save the contents of d document to the supplied file.
func (d *Document) SaveFile(path string) error {
	return os.WriteFile(path, d.SaveBytes(), 0600)
}

// Save the contents of d document as a byte slice.
func (d *Document) SaveBytes() []byte {
	var b bytes.Buffer

	if d.SaveDocType {
		b.WriteString(fmt.Sprintf(`<?xml version="%s" encoding="%s" standalone="%s"?>`,
			d.Version, d.Encoding, d.StandAlone))

		if len(IndentPrefix) > 0 {
			b.WriteByte('\n')
		}
	}

	b.Write(d.Root.Bytes())
	return b.Bytes()
}

// Save the contents of d document as a string.
func (d *Document) SaveString() string { return string(d.SaveBytes()) }

// Alias for Document.SaveString(). d one is invoked by anything looking for
// the standard String() method (eg: fmt.Printf("%s\n", mydoc).
func (d *Document) String() string { return string(d.SaveBytes()) }

// Save the contents of d document to the supplied writer.
func (d *Document) SaveStream(w io.Writer) (err error) {
	_, err = w.Write(d.SaveBytes())
	return
}

// Set a custom user agent when making a new request.
func (d *Document) SetUserAgent(s string) {
	d.useragent = s
}
