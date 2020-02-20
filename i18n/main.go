package main



func main() {
	Tr:=i18n.NewLocale()
	Tr.LoadPath("config/locales")

	fmt.Println(Tr.Translate("submit"))
//输出Submit
Tr.SetLocale("zh")
fmt.Println(Tr.Translate("submit"))
//输出“提交”
}
