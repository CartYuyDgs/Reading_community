package models

import "time"

type Book struct {
	BookId         int       `orm:"pk;auto"json:"book_id"`
	BookName       string    `orm:"size(500)"json:"book_name"`
	Identify       string    `orm:"size(100);unique"json:"identify"`
	OrderIndex     int       `orm:"default(0)"json:"order_index"`
	Descripition   string    `orm:"size(1000)"json:"descripition"`
	Cover          string    `orm:"size(1000)" json:"cover"`             //封面地址
	Editor         string    `orm:"size(50)" json:"editor"`              //编辑器类型: "markdown"
	Status         int       `orm:"default(0)" json:"status"`            //状态:0 正常 ; 1 已删除
	PrivatelyOwned int       `orm:"default(0)" json:"privately_owned"`   // 是否私有: 0 公开 ; 1 私有
	PrivateToken   string    `orm:"size(500);null" json:"private_token"` // 私有图书访问Token
	MemberId       int       `orm:"size(100)" json:"member_id"`
	CreateTime     time.Time `orm:"type(datetime);auto_now_add" json:"create_time"` //创建时间
	ModifyTime     time.Time `orm:"type(datetime);auto_now_add" json:"modify_time"`
	ReleaseTime    time.Time `orm:"type(datetime);" json:"release_time"` //发布时间
	DocCount       int       `json:"doc_count"`                          //文档数量
	CommentCount   int       `orm:"type(int)" json:"comment_count"`
	Vcnt           int       `orm:"default(0)" json:"vcnt"`              //阅读次数
	Collection     int       `orm:"column(star);default(0)" json:"star"` //收藏次数
	Score          int       `orm:"default(40)" json:"score"`            //评分
	CntScore       int       //评分人数
	CntComment     int       //评论人数
	Author         string    `orm:"size(50)"`                      //来源
	AuthorURL      string    `orm:"column(author_url);size(1000)"` //来源链接
}

func (m *Book) TableName() string {
	return TNBook()
}

func (m *Book) NewBook() *Book {
	return &Book{}
}
