package render

import "testing"

var mds = `# header \n 
$$ 
a \to b
$$`

func TestMdToHTML(t *testing.T) {
	html := MdToHTML([]byte(mds))
	println(string(html))
}
