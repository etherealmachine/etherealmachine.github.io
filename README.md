etherealmachine.github.io

Yet another static site generator for a github.io page...

Design:

* Markdown pages go in /pages.
* Pages have a optional YAML header that gets mapped to a Go struct (`Meta` in `page.go`) and used to control page generation.
* Generation consists pf
	* For each page:
  		* Extract the page metadata
  		* Run the Markdown parser to generate the page content
  * Process site metadata (e.g. recent pages)
  * Run the index template with site metadata
  * Run the page templates with page content and site metadata
