# 目录结构
>Dialogue 对话</br>
>>Hello.yml</br>
>>Punish-1.yml</br>

>Missions 任务</br>
>>Easy	简单</br>
>>>XXX-Missions.yml</br>
>>>XXX-Missions.yml</br>

>>Normal	普通</br>
>>Hard	困难</br>

>Impossible	王の不可能完成の任务</br>
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
```
Contents: 你好啊我亲爱的奴隶 {{ .Name }} ，从今天开始我就是你的新主人了，你一定有很多问题，不过没关系我们相处的时间还长着呢，我会好好照顾你的 哼哼哼。如果不想被惩罚的话，就乖乖的听我说话，按照我说的做，否则哼哼。
Options: 
  Option-1: 
    Contents: 是的主人
    Action: 
      Go: Next
  Option-2: 
    Contents: 去一边去，我不认你做我的主人
    Action: 
      Go: Punish
  Option-3: 
    Contents: 虽然但是，你是谁
    Action:
      Go: Dialogue
      Contents: 我是你亲爱的主人啊，快点向我服从吧
      Options:
        Option-1:
          Contents: 是的，我会服从
          Action:
            Go: Next
        Option-2:
          Contents: 不，我不会服从
          Action:
            Go: Punish
```
</br>
其中 {{.Mission}} 为任务 {{.Name}} 为用户名</br>
文件为YML文件一条对话占用一个YML文件</br>
Go后可执行的操作详见GO功能表，以及我事先写的一些示例，相信聪明的~~M~~们一下就能看懂
各位笔下留情，不要写太过于暴露（参考已经被封了的Wearable Technology，谁也不希望这个仓库被封⑧）或者太难完成的任务（注意身体）</br>

# 变量表
待完善

# GO功能表
待完善
