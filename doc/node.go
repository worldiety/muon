package doc

type NodeType string

const (
	TText    NodeType = "text"
	TPreface          = "preface"
	TChapter          = "chapter"
)

/*
Node structure is as follows

		 {
		"type": "document",
		"title": "my book",
		"authors": [
			{
				"type":"author REDUNDANT!!!",
				"name": "Torben Schinke",
				"mail": "trash@worldiety.de"
			}
		],
		"segments": [
			{
				"type": "section",
				"chapters": [
					{
						"type": "chapter REDUNDANT???",
						"title": "my first chapter",
						"sections": [
							{
								"type": "section REDUNDANT???",
								"title": "my section",
								"body": [],
								"subsections": [
									{
										"type": "subsection REDUNDANT???",
										"title": "my subsection",
										"body": [
											{
												"type": "text POLYMORPH, REQUIRED",
												"value": "hello"
											},
											{
												"type": "it POLYMORPH, REQUIRED",
												"body": [
													{
														"type": "text POLYMORPH, REQUIRED",
														"value": "hello world"
													}
												]
											},
											{
												"type": "br POLYMORPH, REQUIRED"
											}
										]
									}
								]
							}
						]
					}
				]
			}
		]
	}
*/
type Node struct {
	Type       string
	Value      string
	Body       []*Node
	Properties []*Node
}

