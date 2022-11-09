package src

//如下代码如果这样写，好像调用位置就会有编码问题，会将特殊字符转换为ASCII码，而且还转不回来那种，比方说 " 转换为 &#34;
//所以就把这几行迁移回main.go了，如果有大佬能解决，那鄙人十分感激

/*
func process(t *template.Template, vars interface{}) string {
	var tmplBytes bytes.Buffer

	err := t.Execute(&tmplBytes, vars)
	if err != nil {
		fmt.Println(err)
	}
	return tmplBytes.String()
}
func ProcessString(str string, vars interface{}) string {
	tmpl, err := template.New("tmpl").Parse(str)

	if err != nil {
		fmt.Println(err)
	}
	return process(tmpl, vars)
}
*/
