# 目录结构
>Dialogue 对话</br>
>>Hello.html</br>
>>Punish-1.html</br>

>Missions 任务</br>
>>Easy	简单</br>
>>>XXX-Mission.yml</br>
>>>XXX-Mission.yml</br>

>>Normal	普通</br>
>>Hard	困难</br>
>>Impossible	王の不可能完成の任务</br>
>>Punish	惩罚任务（如果完不成主人给的任务，~~就撅好屁股挨罚吧~~）</br>
>>>XXX-Punish.yml

>Pictures 图片资源</br>
>>Illustrations 二次元插图（虽然不可能有，不要直接传PNG这样的图，真的会被封，请编码成Base64之后传TXT格式，十分感谢各位Hero）</br>
>>Static 固定UI图片
# 特别注意
目录下一切文件名称，分类都为纯英文字符（或拼音）。命名格式遵从XXX（您给起的名字）-Mission(Punish)
# 任务编写
遵从YAML语法格式，详细格式如下</br>
```
Mission:
  Categories: Language			 #任务分类
  Describe: 我是一个任务                  #任务的简介
  Contents: 我要你叫我一声"主人"          #任务的内容
  Keywords: 叫"主人"		      #任务的实质操作，关键词
RestTime:
  FixedTime: 0				 #固定休息时间，单位（s），设定为非0后系统会固定休息一段时间再进行下次任务
  DynamicTime: 0			 #可变休息时间，单位（s），系统会随机一个在 0-设置时间 范围内的时间与固定时间进行相加
Punish:
  Level: 1				 #惩罚等级
  Categories: Language			 #惩罚分类，未指定则随机
  Contents: Call-Punish			 #指定惩罚名称，未指定则随机
```
# 对话编写
请您使用如下格式编写剧情对话</br>
恭喜你抽到了 {{.Mission}} 的任务，我亲爱的奴隶 {{.Name}}</br>
其中 {{.Mission}} 为任务 {{.Name}} 为用户名</br>
文件为Html文件，可以适当引入Html标签和Css美化（但我不保证被解析出来会不会炸，所以最好别写太华丽吧）一条对话占用一个Html文件</br>
各位笔下留情，不要写太过于暴露（参考已经被封了的Wearable Technology，谁也不希望这个仓库被封⑧）或者太难完成的任务（注意身体）</br>
#变量表
待完善
