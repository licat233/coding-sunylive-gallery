package utils

import (
	"regexp"
	"strings"
)

//图片url校验
func ImgUrlVerify(url string) (bool, error) {
	//长度不能超过500
	if len(url) >= 500 {
		return false, nil
	}
	matched, err := regexp.MatchString("(https|http)://([\\w-]+\\.)+[\\w-]+(/[\\w-./?%&=]*)?", url)
	if err != nil {
		return matched, err
	}
	if !matched {
		return false, nil
	}
	index1 := strings.Index(url, "?")
	if index1 != -1 {
		url = url[:index1]
	}
	index2 := strings.LastIndex(url, ".")
	if index2 == -1 {
		return false, nil
	}
	fileType := url[index2+1:]
	return regexp.MatchString("^(jpg|png|jpeg|gif|ico|svg|webp|bmp)?$", fileType)
}

//ImgUrlPick 提取图片url
func ImgUrlPick(str string) (imgUrls []string, err error) {
	str = strings.Trim(str, " ")
	//匹配所有的url
	reg := regexp.MustCompile("(https|http)://([\\w-]+\\.)+[\\w-]+(/[\\w-./?%&=]*)?")
	urls := reg.FindAllString(str, -1)
	//校验每一条url
	for _, url := range urls {
		ok, e := ImgUrlVerify(url)
		if e != nil {
			err = e
			return
		}
		if !ok {
			// 此条url不符合要求，跳过
			continue
		}
		// 此条url符合要求，加入返回结果
		imgUrls = append(imgUrls, url)
	}
	return
}
