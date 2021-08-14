package main

//使用状态机解字符串分词逻辑
//规则空格 #代表分词标记遇到空格和#代表一个单词结束
//&和.为连接符,如果&和. 前面或者后面有单词,则为单词,如果 &和.连续或者单独存在不包括单词则不算单词

import (
	"fmt"
)

func main() {
	//words 里包含的单词为 1&2 1.2 x.y 1.1 j i x. .y  一共8个单词
	words := "1&2 1.2 ... &&& .&.&  x.y 1.1 j#i x. .y"
	fmt.Println(WordsStatistics(words))
}

//WordsStatistics 单词数统计方法
func WordsStatistics(s string) int {
	words := NewWords()
	Num := 0
	for k, v := range s {
		NowChar := fmt.Sprintf("%c", v)
		if words.IsEnd(NowChar) {
			if words.WordStatus() {
				Num++
			}
			words.init()
		} else {
			words.IsWord(NowChar)
		}
		//todo 这里应该还可以优化下将几个if合并下
		if k == (len(s) - 1) {
			if words.WordStatus() {
				Num++
			}
		}

	}
	return Num
}

type Words struct {
	isword bool
}

func NewWords() *Words {
	return &Words{
		isword: false,
	}
}

//判断当前输入字符串是否为单词
func (w *Words) IsWord(s string) {
	//如果状态已经标记为是单词则不需要判断
	if w.isword {
		return
	}
	if s != "&" && s != "." {
		w.isword = true
	}
}

//返回单词状态是否为单词
func (w *Words) WordStatus() bool {
	return w.isword
}

//初始化
func (w *Words) init() {
	w.isword = false
}

//判断当前单词是否为结束标记
func (w *Words) IsEnd(s string) bool {
	if s == " " || s == "#" {
		return true
	}
	return false
}
