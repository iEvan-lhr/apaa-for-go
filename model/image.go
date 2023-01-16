package model

import (
	"encoding/json"
	"github.com/iEvan-lhr/apaa-for-go/structs"
	tools "github.com/iEvan-lhr/exciting-tool"
	"github.com/iEvan-lhr/nihility-dust/anything"
)

type Image struct {
}

func (image Image) GetImages(mission chan *anything.Mission, data []any) {
	var images []structs.UserImage
	find := data[0].(*structs.UserImage)
	structs.Find(find, &images)
	if len(images) > 0 {
		mission <- &anything.Mission{Pursuit: []any{string(tools.Marshal(structs.ImageRes{Status: 200, Ans: "succ", Data: string(tools.ReturnValueByTwo(json.Marshal(images)).([]byte))}))}}
	} else {
		mission <- &anything.Mission{Pursuit: []any{string(tools.Marshal(structs.ImageRes{Status: 202, Ans: "Fail"}))}}
	}
}
