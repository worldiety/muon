package html

import (
	"github.com/worldiety/muon/document"
	"html/template"
	"io"
)

func Write(doc *document.Model, writer io.Writer) error {
	tpl, err := template.New("render.go->tmpl").Parse(tmpl)
	if err != nil {
		return err
	}
	return tpl.Execute(writer, doc)
}

const tmpl = `
<html>
	<header>
	<title>{{.Title}}</title>
	<style>
		.content {
			max-width:900px;
			margin-left:auto;
			margin-right:auto;
			font-family: roboto-thin,open-sans,sans-serif;
			color: black;
		}

		.title{
			text-align: center;
			font-weight: 100;
		}
		
		.subtitle{
			text-align: center;
			font-weight: 100;
		    font-size: x-large;
		}

		.preface{
			font-style: italic;
			page-break-before: always;
		}

		h2{
			font-weight: 100;
			page-break-before: always;
		}

		h3{
			font-weight: 100;
		}

		a{
			text-decoration: none;
			color: black;
		}

		@page {
		  size: A4;
		  margin: 20mm 20mm 20mm 20mm;
counter-increment: page ;
		}
		.pageNumber { content: counter(page); }

		footer,header{
			display: none;
		}

		@media print {

		  	footer {
				position: fixed;
				bottom: 0;
			}

			header {
				position: fixed;
				top: 0;
			}

			.content {
				font-size: 12pt;
			}
			p {
				page-break-inside: avoid;
			}
			footer,header{
				display: initial;
			}
		}
	</style>
	</header>
	<body>
		<div class="content">
			<h1 class="title">{{.Title}}</h1>
			<h1 class="subtitle">{{.Subtitle}}</h1>

			<span class=".page-break"/>

			<p class="preface">
			Preface
			<br>
			{{.Preface}}
			</p>
			<h2>Table of contents</h2>
			<ul class="toc">
			{{range $i,$chap := .Elements}}
				<li class="level-1"><a href="#chapter-{{$i}}">{{$chap.Title}}</a></li>
				{{if .Elements}}
					<ul class="toc-subchapter">
					{{range $i,$chap := .Elements}}
						<li class="level-2"><a href="#subchapter-{{$i}}">{{$chap.Title}}</a></li>
					{{end}}
					</ul>
				{{end}}
			{{end}}
			</ul>
		
			{{range $i,$chap := .Elements}}
				<h2><a name="chapter-{{$i}}">{{$chap.Title}}</a></h2>
				{{range .Body}}
					<p>{{.}}</p>
				{{end}}

				{{range $i,$chap := .Elements}}
					<h3><a name="subchapter-{{$i}}">{{$chap.Title}}</a></h3>
					{{range .Body}}
						<p>{{.}}</p>
					{{end}}
				{{end}}
			{{end}}


			<header>
			  text in the docs header
			   <hr/>	
			</header>

			<footer>
 			  <hr/>
			  This is the text that goes at the bottom of every page.<span class="pageNumber"></span>
			</footer>
		</div>

	</body>
</html>
`
