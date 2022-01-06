package data

type DataBlob struct {
	Attributes []string
	Data       []map[string]interface{}
}

func (this *DataBlob) InitDataBlob(rows ...map[string]interface{}) {
	this.Data = rows
	this.Attributes = make([]string, 0, len(rows[0]))

	for attr := range rows[0] {
		this.Attributes = append(this.Attributes, attr)
	}
}
