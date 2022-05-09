package main

import (
	"fmt"
	"sunylive-gallery/utils"
	"testing"
)

func Test_reg(t *testing.T) {
	str := "s = https://alifei05.cfp.cn/cms/image/image/3fbff94a68e74fe4a5095c74b7aa5b0b.jpg https://tenfei04.cfp.cn/cms/image/image/3bea8a443a6043968a5bb489e5db5787.jpghttps://alifei02.cfp.cn/cms/image/image/e932041cbf8b45fca258ad6668a2c68f.jpghttps://tenfei01.cfp.cn/cms/image/image/1e79c817271840159097bef7dfcc3a71.jpghttps://alifei05.cfp.cn/cms/image/image/dab4eb5028244667bfb6e99f19ebbe6d.jpghttps://alifei03.cfp.cn/cms/image/image/743ad00bda224bb89d89e7157d7d4af7.jpghttps://alifei05.cfp.cn/cms/image/image/4fa4de9105e24ac2abe1ad36a8a6e33d.jpeghttps://tenfei01.cfp.cn/cms/image/image/b536a71be885467abdfc6f656a1a30d7.jpg https://alifei04.cfp.cn/cms/image/image/e9ee134a1bc24bd9859dc11bf73bc4aa.jpghttps://tenfei03.cfp.cn/cms/image/image/92751dedbb864b7ea1eaa89de07bca8a.jpg"
	matched, err := utils.ImgUrlPick(str)
	fmt.Println(matched, err)
}
