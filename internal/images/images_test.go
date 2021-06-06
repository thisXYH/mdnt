package images

import (
	"regexp"
	"testing"
)

// 测试正则的匹配
func TestRegexp(t *testing.T) {
	content := `
	![1](../.resources/2021-05-13-21-48-32.png)
	![2](/.resources/2021-05-13-21-48-32.png)
	![3](..\.resources\2021-05-13-21-48-32.png)
	![4](\.resources\2021-05-13-21-48-32.png)
	![5](\.resources\2021-05-13-21-48-32.png)
	![6](https://upload-images.jianshu.io/upload_images/24630328-df0561b8c1d131a1.image)
	![7](https://upload-images.jianshu.io/upload_images/24630328-df0561b8c1d131a1.png)
	`
	re := regexp.MustCompile(`!\[.*\]\((.+(?:\\|\/)(.+))\)`)
	result := re.FindAllStringSubmatch(content, -1)
	for _, v := range result {
		t.Log("--------------------\n")
		t.Log("orginal:", v[0]+"\n")
		t.Log("path:", v[1]+"\n")
		t.Log("name:", v[2]+"\n")
	}

	if len(result) != 7 {
		t.Fail()
	}
}
