package repo

import (
	"blog/model"
	"blog/repo/db"
	"blog/tool"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// GetPropertyByKey get option by optionKey
func GetPropertyByKey(optionKey string) *model.Option {
	option := &model.Option{}
	db.GetMDB().Model(option).Where("option_key = ?", optionKey).First(option)
	return option
}

// SaveProperty save property
func SaveProperty(optionKey string, optionValue string) *model.Option {
	option := &model.Option{
		OptionKey:   optionKey,
		OptionValue: optionValue,
	}
	db.GetMDB().Model(option).Where("option_key = ?", optionKey).Create(option)
	return option
}

// UpdateProperty save property
func UpdateProperty(optionKey string, optionValue string) *model.Option {
	option := &model.Option{
		OptionValue: optionValue,
	}
	db.GetMDB().Model(option).Where("option_key = ?", optionKey).Update(option)
	return option
}

// GetBirthday return birthday
func GetBirthday() int64 {
	option := GetPropertyByKey("birthday")
	if option == nil {
		now := time.Now().Unix()
		SaveProperty("birthday", strconv.FormatInt(now, 10))
		return now
	}
	birthday, err := strconv.ParseInt(option.OptionValue, 10, 64)
	if err != nil {
		log.Printf("birthday convert failed:%v", err)
		return 0
	}
	return birthday
}

// GetBlogBaseURL get blog base url
func GetBlogBaseURL() string {
	var blogURL string
	option := GetPropertyByKey("blog_url")
	if option != nil && tool.IsNotBlank(option.OptionValue) {
		blogURL = option.OptionValue
	}

	if tool.IsNotBlank(blogURL) {
		strings.Replace(blogURL, "/", "", len(blogURL)-1)
	} else {
		blogURL = fmt.Sprintf("http://%s:%s", "127.0.0.1", "65530")
	}
	return blogURL
}
