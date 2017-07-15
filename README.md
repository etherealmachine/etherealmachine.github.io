etherealmachine.github.io

Yet another static site generator for a github.io page..

* Markdown pages go in /pages.
* Pages have a optional YAML header that gets mapped to a Go struct (`Meta` in `page.go`) and used to control page generation.
* Generation consists of
	* For each page:
  		* Extract the page metadata
  		* Run the Markdown parser to generate the page content
  * Process site metadata (e.g. recent pages)
  * Run the index template with site metadata
  * Run the page templates with page content and site metadata

# Markdown Library

The big difference is that I wanted to mix HTML and Markdown in a single document. I like the expressivity of HTML but am not willing to commit to the syntax required just to write a paragraph. I couldn't find a Markdown library that let you mix Markdown into HTML. Once you drop into HTML, you're done. So I wrote one myself - https://github.com/etherealmachine/markdown. Not recommended for actual use, but it works for this site, and it gave me an excuse to write a recursive descent parser, which is one of the great joys in life.