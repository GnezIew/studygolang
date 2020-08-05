package main

type Player struct {
	currPos Vec2	//当前位置
	targetPos Vec2 	//目标位置
	speed	float32		//移动速度
}

//s设置玩家的移动目标
func (p *Player) MoveTo(v Vec2){
	p.targetPos = v
}

//获取当前位置
func (p *Player) Pos() Vec2{
	return p.currPos
}


//判断是否到达目的地
func (p *Player) IsArrived() bool {
	//通过计算当前玩家位置与目标位置的距离不超过移动的步长， 判断已经到达目标点
	return p.currPos.DistanceTo(p.targetPos) < p.speed
}

//更新玩家的位置
func (p *Player) Update(){
	if !p.IsArrived(){
		//计算当前位置指向目标的朝向
		dir := p.targetPos.Sub(p.currPos).Normalize()
		//添加速度矢量生成行的位置
		newPos := p.currPos.Add(dir.Scale(p.speed))
		//移动完成后更新当前位置
		p.currPos = newPos
	}
}

//创建新玩家
func NewPlayer(speed float32) *Player{
	return &Player{
		speed:     speed,
	}
}