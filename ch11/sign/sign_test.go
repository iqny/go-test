package sign

import (
	"testing"
)

func TestCreateSign(t *testing.T) {
	var param = make(map[string]string)
	param["timestamp"] = "1638866931"
	param["service"] = "douyin_order_list"
	param["access_token"] = "DATA_CENTER_APP_SECRET"
	got:=CreateSign(param)
	want := "5f3ef9dbef366cefe3a7247630eb8059"
	if want!=got {
		t.Errorf("got %s want %s", got, want)
	}
}
