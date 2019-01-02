package main

import "testing"

var r *Robot

func init() {
	r = NewRobot("xxxxxxxxxxxxxxxxxxxxxxx")
}

func TestTextMsg(t *testing.T) {
	textMsg := NewTextMsg("测试文本信息")
	_, err := r.SendMsg(textMsg)
	if err != nil /*|| resp.Errcode != 0*/ {
		t.Errorf("test failed,err:%v", err)
	}
}

func TestTextMsg1(t *testing.T) {
	textMsg := NewTextMsg("测试文本信息@158XXXXXXXXX")
	textMsg.At = At{
		AtMobiles: []string{"158XXXXXXXXX"},
	}
	_, err := r.SendMsg(textMsg)
	if err != nil /*|| resp.Errcode != 0*/ {
		t.Errorf("test failed,err:%v", err)
	}
}

func TestTextMsg2(t *testing.T) {
	textMsg := NewTextMsg("测试文本信息")
	textMsg.At = At{
		IsAtAll: true,
	}
	_, err := r.SendMsg(textMsg)
	if err != nil /*|| resp.Errcode != 0*/ {
		t.Errorf("test failed,err:%v", err)
	}
}

func TestLinkMsg(t *testing.T) {
	linkMsg := NewLinkMsg("标题", "内容", "http://www.baidu.com")
	_, err := r.SendMsg(linkMsg)
	if err != nil /*|| resp.Errcode != 0*/ {
		t.Errorf("test failed,err:%v", err)
	}
}

func TestLinkMsg1(t *testing.T) {
	linkMsg := NewLinkMsg("标题", "内容", "http://www.baidu.com")
	linkMsg.Link.PicUrl = "https://note.youdao.com/favicon.ico"
	_, err := r.SendMsg(linkMsg)
	if err != nil /*|| resp.Errcode != 0*/ {
		t.Errorf("test failed,err:%v", err)
	}
}

func TestMarkdownMsg(t *testing.T) {
	markdownMsg := NewMarkdownMsg("杭州天气", "#### 杭州天气  \n > 9度，西北风1级，空气良89，相对温度73%\n\n > ![screenshot](http://i01.lw.aliimg.com/media/lALPBbCc1ZhJGIvNAkzNBLA_1200_588.png)\n  > ###### 10点20分发布 [天气](http://www.thinkpage.cn/)")
	_, err := r.SendMsg(markdownMsg)
	if err != nil /*|| resp.Errcode != 0*/ {
		t.Errorf("test failed,err:%v", err)
	}
}

func TestMarkdownMsg1(t *testing.T) {
	markdownMsg := NewMarkdownMsg("杭州天气", "#### 杭州天气  \n > @158XXXXXXXXX 9度，西北风1级，空气良89，相对温度73%\n\n > ![screenshot](http://i01.lw.aliimg.com/media/lALPBbCc1ZhJGIvNAkzNBLA_1200_588.png)\n  > ###### 10点20分发布 [天气](http://www.thinkpage.cn/)")
	markdownMsg.At = At{
		AtMobiles: []string{"158XXXXXXXXX"},
	}
	_, err := r.SendMsg(markdownMsg)
	if err != nil /*|| resp.Errcode != 0*/ {
		t.Errorf("test failed,err:%v", err)
	}
}

func TestMarkdownMsg2(t *testing.T) {
	markdownMsg := NewMarkdownMsg("杭州天气", "#### 杭州天气  \n > 9度，西北风1级，空气良89，相对温度73%\n\n > ![screenshot](http://i01.lw.aliimg.com/media/lALPBbCc1ZhJGIvNAkzNBLA_1200_588.png)\n  > ###### 10点20分发布 [天气](http://www.thinkpage.cn/)")
	markdownMsg.At = At{
		IsAtAll: true,
	}
	_, err := r.SendMsg(markdownMsg)
	if err != nil /*|| resp.Errcode != 0*/ {
		t.Errorf("test failed,err:%v", err)
	}
}

func TestActionCardMsg(t *testing.T) {
	actionCardMsg := NewWholeActionCardMsg("乔布斯 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身", "![screenshot](@lADOpwk3K80C0M0FoA) \n #### 乔布斯 20 年前想打造的苹果咖啡厅 \n\n Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划", "阅读全文", "https://www.dingtalk.com/")
	_, err := r.SendMsg(actionCardMsg)
	if err != nil /*|| resp.Errcode != 0*/ {
		t.Errorf("test failed,err:%v", err)
	}
}

func TestActionCardMsg1(t *testing.T) {
	actionCardMsg := NewWholeActionCardMsg("乔布斯 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身", "![screenshot](@lADOpwk3K80C0M0FoA) \n #### 乔布斯 20 年前想打造的苹果咖啡厅 \n\n Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划", "阅读全文", "https://www.dingtalk.com/")
	btns := []Button{
		{
			Title:     "百度",
			ActionURL: "http://www.baidu.com",
		},
		{
			Title:     "新浪",
			ActionURL: "http://www.sina.com",
		},
	}
	actionCardMsg = NewIndependentCardMsg("选择网站", "年度最佳网站", btns)
	actionCardMsg.ActionCard.BtnOrientation = "1"
	actionCardMsg.ActionCard.HideAvatar = "1"
	_, err := r.SendMsg(actionCardMsg)
	if err != nil /*|| resp.Errcode != 0*/ {
		t.Errorf("test failed,err:%v", err)
	}
}

func TestFeedCardMsg(t *testing.T) {
	feedCardMsg := NewFeedCardMsg([]FeedCardLink{
		{
			Title:      "时代的火车向前开",
			MessageURL: "https://mp.weixin.qq.com/s?__biz=MzA4NjMwMTA2Ng==&mid=2650316842&idx=1&sn=60da3ea2b29f1dcc43a7c8e4a7c97a16&scene=2&srcid=09189AnRJEdIiWVaKltFzNTw&from=timeline&isappinstalled=0&key=&ascene=2&uin=&devicetype=android-23&version=26031933&nettype=WIFI",
			PicURL:     "http://i01.lw.aliimg.com/media/lALPBbCc1ZhJGIvNAkzNBLA_1200_588.png",
		},
		{
			Title:      "时代的火车向前开2",
			MessageURL: "https://mp.weixin.qq.com/s?__biz=MzA4NjMwMTA2Ng==&mid=2650316842&idx=1&sn=60da3ea2b29f1dcc43a7c8e4a7c97a16&scene=2&srcid=09189AnRJEdIiWVaKltFzNTw&from=timeline&isappinstalled=0&key=&ascene=2&uin=&devicetype=android-23&version=26031933&nettype=WIFI",
			PicURL:     "http://i01.lw.aliimg.com/media/lALPBbCc1ZhJGIvNAkzNBLA_1200_588.png",
		},
		{
			Title:      "时代的火车向前开3",
			MessageURL: "https://mp.weixin.qq.com/s?__biz=MzA4NjMwMTA2Ng==&mid=2650316842&idx=1&sn=60da3ea2b29f1dcc43a7c8e4a7c97a16&scene=2&srcid=09189AnRJEdIiWVaKltFzNTw&from=timeline&isappinstalled=0&key=&ascene=2&uin=&devicetype=android-23&version=26031933&nettype=WIFI",
			PicURL:     "http://i01.lw.aliimg.com/media/lALPBbCc1ZhJGIvNAkzNBLA_1200_588.png",
		},
	})
	_, err := r.SendMsg(feedCardMsg)
	if err != nil /*|| resp.Errcode != 0*/ {
		t.Errorf("test failed,err:%v", err)
	}
}
