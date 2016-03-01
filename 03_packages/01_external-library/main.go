package main

import (
	"fmt"

	"github.com/russross/blackfriday"
)

/*

This library blackfriday helps to convert markdown to html

*/

func main() {
	markdown := []byte(`
    # Title
    * topic 1
    * topic 2
    * topic 3
    `)

	output := blackfriday.MarkdownBasic(markdown)

	fmt.Println(string(output))

	/* output
	<pre><code># Title
	* topic 1
	* topic 2
	* topic 3
	</code></pre>
	*/
}
