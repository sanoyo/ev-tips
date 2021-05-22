package main

import "fmt"

// 抽象クラス
type Decorator interface {
	Decorate() string
}

// 既存クラス
// XXCircle
type Banner struct {
	str string
}

func (self *Banner) getString() string {
	return "*" + self.str + "*"
}

// クラス
// Circle
// 構造体の埋込による継承
type EmbeddedDecorateBanner struct {
	*Banner
}

func NewEmbeddedDecorateBanner(str string) *EmbeddedDecorateBanner {
	return &EmbeddedDecorateBanner{&Banner{str}}
}

// Circle
// インターフェースの実装と埋め込んだ構造体のメソッドによるアダプタ
func (self *EmbeddedDecorateBanner) Decorate() string {
	return self.getString()
}

func main() {
	var decorator Decorator
	decorator = NewEmbeddedDecorateBanner("AAAA")

	aa := decorator.Decorate()
	fmt.Println(aa)
}
