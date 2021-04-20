package models

type Graphic struct {
	Tech struct {
		TechKnowledge Custom
		BestPractices Custom
		CodingStyle   Custom
	}
	Team struct {
		TeamPlayer    Custom
		Commited      Custom
		Communication Custom
	}
	Perfo struct {
		WorkQuality    Custom
		ClientOriented Custom
	}
}

type Custom struct {
	Value int
	Count int
}

