package models

type SyllabusResponse struct {
	Syllabus []*Section
}

type Section struct {
	ID          string      `json:"sectionID"`
	Name        string      `json:"sectionName"`
	Subsections []*Material `json:"subSections"`
}

type Material struct {
	ID   string `json:"materialID"`
	Name string `json:"materialName"`
	Type string `json:"materialType"`
}

type SectionContentResponse struct {
	ID 					string							`json:"sectionID"`
	Subsections []*MaterialContent	`json:"subSections"`
}

type MaterialContent struct {
	ID 					string `json:"materialID"`
	Name 				string `json:"materialName"`
	Type 				string `json:"materialType"`
	Content 		string `json:"materialContent"`
	ContentText string `json:"materialContentText"`
}