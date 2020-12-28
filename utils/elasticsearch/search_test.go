package elasticsearch

import (
	"encoding/json"
	"testing"
)

func SearchTest (t *testing.T) {

	var request = &SearchRequest{
		Start:   nil,
		End:     nil,
		AppName: "",
		Match:   nil,
		Sort:    false,
		From:    0,
		Size:    0,
	}

	manager := NewSearchManager()
	result , err := manager.Search(request)
	if err != nil {
		t.Errorf(err.Error())
	}
	b , err := json.Marshal(result)
	t.Log(string(b))
}
