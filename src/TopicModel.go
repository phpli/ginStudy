package src

type Topic struct {
	TopicId         int    `json:"id"`
	TopicTitle      string `json:"title" binding:"min=4,max=20"`
	TopicShortTitle string `json:"short_title" binding:"required,nefield=TopicTitle"`
	UserIp          string `json:"ip" binding:"ipv4"`
	TopicScore      int    `json:"score" binding:"omitempty,gt=5"`
	TopicUrl        string `json:"url" binding:"omitempty,topicurl"`
}

//func CreateTopic(id int, title string) Topic {
//	return Topic{id,title}
//}

type TopicQuery struct {
	UserName string `json:"user_name" form:"username"`
	Page     int    `json:"page" form:"page" binding:"required"`
	PageSize int    `json:"page_size" form:"page_size"`
}
