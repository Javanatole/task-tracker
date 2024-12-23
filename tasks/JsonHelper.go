package tasks

import "encoding/json"

type JsonHelper struct {
	fileHelper FileHelper
}

func (jsonHelper JsonHelper) writeJSONTasks(content JSONTasks) {
	contentAsJson, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}
	err = jsonHelper.fileHelper.writeContentIntoFile(string(contentAsJson))
	if err != nil {
		panic(err)
	}
}

func (jsonHelper JsonHelper) readContentFromFile() JSONTasks {
	// first we read the content of the file
	content, err := jsonHelper.fileHelper.readContentFromFile()
	if err != nil {
		// in case we discover an error, we re-write the content of the file
		err = jsonHelper.fileHelper.writeContentIntoFile(DefaultContent)
		if err != nil {
			// in case we can't re-write the file we launch a panic error
			panic(err)
		}
		// in case
		content = DefaultContent
	}

	contentTasks, err := ParseJSONTasks(content)
	if err != nil {
		contentTasks = JSONTasks{[]Task{}}
	}
	return contentTasks
}
